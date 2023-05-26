package trapri

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/degendata"
	"github.com/benleb/gloomberg/internal/external"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/collectionsource"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/osmodels"
	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/benleb/gloomberg/internal/nemo/standard"
	"github.com/benleb/gloomberg/internal/nemo/tokencollections"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/benleb/gloomberg/internal/notify"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/ticker"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/benleb/gloomberg/internal/utils/wwatcher"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

func TokenTransactionFormatter(gb *gloomberg.Gloomberg, queueTokenTransactions chan *totra.TokenTransaction, queueWsOutTokenTransactions chan *totra.TokenTransaction, queueWsInTokenTransactions chan *totra.TokenTransaction, terminalPrinterQueue chan string) {
	gbl.Log.Debugf("🧱 starting ttx formatter worker")

	if viper.GetBool("websockets.client.enabled") {
		go func() {
			for ttx := range queueWsInTokenTransactions {
				go formatTokenTransaction(gb, ttx, terminalPrinterQueue)
			}
		}()
	}

	// ! critical path !
	// this is the main loop of the formatter
	// blocking/delaying here will block/delay the whole stream
	// when adding additional calls here, prefer goroutines with conditional select
	for ttx := range queueTokenTransactions {
		go formatTokenTransaction(gb, ttx, terminalPrinterQueue)

		// send to ws if webserver enabled & the queue is not congested
		if viper.GetBool("web.enabled") {
			if len(queueWsOutTokenTransactions) < cap(queueWsOutTokenTransactions)-10 {
				queueWsOutTokenTransactions <- ttx
			} else {
				gbl.Log.Warnf("🧱 ws out queue is congested")
			}
		}

		// send to bluechip ticker
		if viper.GetBool("notifications.bluechip.enabled") {
			ticker.BlueChips.CheckForBlueChipInvolvment(ttx)
		}

		if viper.GetBool("notifications.smart_wallets.enabled") {
			ticker.AlphaCaller.AddEvent(ttx)
		}
	}
}

func formatTokenTransaction(gb *gloomberg.Gloomberg, ttx *totra.TokenTransaction, terminalPrinterQueue chan string) {
	// // check if we already know the transaction the log belongs to
	// output.AlreadyPrintedMu.Lock()
	// known, ok := output.AlreadyPrinted[txHash]
	// output.AlreadyPrintedMu.Unlock()

	// if known && ok {
	// 	// we already know this transaction
	// 	gbl.Log.Infof("already printed tx: %s", style.Bold(txHash.String()))

	// 	return
	// }

	ctx := context.Background()

	// fake a txHash for listings
	txHash := common.Hash{}
	if ttx.Action != totra.Listing || ttx.Tx != nil {
		txHash = ttx.Tx.Hash()
	}

	// is a collections from configured collections + own wallets
	isOwnCollection := false

	// a watched wallet is involved
	isOwnWallet := gb.OwnWallets.ContainsAddressFromSlice(ttx.GetNFTSenderAndReceiverAddresses()) != internal.ZeroAddress
	isWatchUsersWallet := gb.Watcher.ContainsAddressFromSlice(ttx.GetNFTSenderAndReceiverAddresses()) != internal.ZeroAddress

	// is this an intentional purchase or a dump into bids?
	// isBidDump := false

	// telegram notification
	if viper.GetBool("notifications.telegram.enabled") && (isOwnWallet || isWatchUsersWallet) && ttx.Action != totra.Transfer {
		gbl.Log.Infof("🧱 sending telegram notification | isOwnWallet: %+v | isWatchUsersWallet: %+v", isOwnWallet, isWatchUsersWallet)

		go notify.SendNotification(gb, ttx)
	}

	// if its a single-collection transaction we set the collection as the currentCollection
	// from here on already, otherwise we set it to nil and fill it later in the loop
	// over the collections/transfers
	var currentCollection *collections.Collection

	if len(ttx.GetTransfersByContract()) >= 1 && currentCollection == nil {
		currentCollection = tokencollections.GetCollection(gb, ttx.Transfers[0].Token.Address, ttx.Transfers[0].Token.ID.Int64())
	}

	// defaults
	priceStyle := style.DarkWhiteStyle
	priceArrowColor := style.DarkGray

	if ttx.GetPrice().Ether() >= 0.01 {
		priceArrowColor = style.GetPriceShadeColor(ttx.GetPrice().Ether())
	}

	switch ttx.Action {
	case totra.ReBurn, totra.Burn, totra.Airdrop, totra.Transfer:
		priceStyle = style.DarkerGrayStyle
	}

	// build the line to be displayed
	out := strings.Builder{}

	var divider string

	var priceCurrencyStyle lipgloss.Style
	if currentCollection != nil {
		priceCurrencyStyle = currentCollection.Style().Copy()
		divider = style.Sharrow.Copy().Foreground(priceArrowColor).Faint(true).String()
	} else {
		priceCurrencyStyle = priceStyle.Copy().Foreground(style.DarkGray)
	}

	formattedCurrencySymbol := priceCurrencyStyle.Render("Ξ")
	formattedFaintCurrencySymbol := priceCurrencyStyle.Copy().Faint(true).Render("Ξ")

	out.WriteString(ttx.Marketplace.RenderFaintTag())

	// timestamp styling
	// WEN...??
	now := time.Now()
	currentTime := now.Format("15:04:05")
	timeNow := style.Gray5Style.Render(currentTime)

	// prepare links
	etherscanURL, _, blurURL := utils.GetLinks(txHash, ttx.Transfers[0].Token.Address, ttx.Transfers[0].Token.ID.Int64())

	// print collection name and token id
	fmtTokensTransferred := make([]string, 0)
	fmtTokensHistory := make([]string, 0)
	ttxCollections := make(map[common.Address]*collections.Collection, 0)

	// contract addresses of the burned token(s)
	// used in reburnes for nicer formatting
	burnedTokenTransferIndex := -1

	for contractAddress, transfers := range ttx.GetTransfersByContract() {
		if transfers[0].Standard == standard.ERC20 {
			continue
		}

		switch {
		// erc20
		case transfers[0].Standard == standard.ERC20:
			continue

		// NFTfi and so on...
		case internal.LoanContracts[contractAddress] != "":
			continue

		// Uniswap V3: Positions NFT
		case contractAddress == common.HexToAddress("0xc36442b4a4522e871399cd717abdd847ab11fe88"):
			return
		}

		fmtTokenIds := make(map[common.Address][]string, 0)
		fmtHistoryTokenIds := make(map[common.Address][]string, 0)

		var name string

		collection := tokencollections.GetCollection(gb, contractAddress, transfers[0].Token.ID.Int64())

		var ensMetadata *external.ENSMetadata

		if collection.IgnorePrinting {
			continue
		}

		ttxCollections[contractAddress] = collection

		numCollectionTokens := int64(0)

		for _, transfer := range transfers {
			fmtTokenID := strings.Builder{}

			// ignore transfers of more than 9999 tokens
			if transfer.AmountTokens.Cmp(big.NewInt(9999)) > 0 {
				gbl.Log.Debugf("♾️ amountTokens > 9999 for token %s", transfer.Token.ShortID())

				continue
			}

			// add rank if available
			if degendata.Metadatas[transfer.Token.Address] != nil {
				if _, ok := degendata.Metadatas[transfer.Token.Address][transfer.Token.ID.Int64()]; ok {
					rank := degendata.Metadatas[transfer.Token.Address][transfer.Token.ID.Int64()].Score.Rank
					topX := float64(rank) / float64(collection.Metadata.TotalSupply)

					var rankSymbol string

					switch {
					case topX <= 0.01:
						rankSymbol = "🥇"
					case topX <= 0.1:
						rankSymbol = "🥈"
					case topX <= 0.25:
						rankSymbol = "🥉"
					default:
						rankSymbol = "|"
					}

					fmtTokenID.WriteString(style.TrendLightGreenStyle.Copy().Bold(false).Render(fmt.Sprintf(" ・%d %s・ ", rank, rankSymbol)))
					log.Debug(degendata.Metadatas[transfer.Token.Address][transfer.Token.ID.Int64()])
				}
			}

			// add number of tokens transferred
			if transfer.AmountTokens.Cmp(big.NewInt(1)) > 0 {
				fmtTokenID.WriteString(style.DarkGrayStyle.Render(transfer.AmountTokens.String() + "x"))
			}

			fmtTokenID.WriteString(formatTokenID(collection, transfer.Token.ID))

			// add a marker for burned tokens
			if transfer.To == internal.ZeroAddress {
				fmtTokenID.WriteString("🔥")

				if ttx.Action == totra.ReBurn {
					burnedTokenTransferIndex = len(fmtTokensTransferred)
				}
			}

			// if it is an ENS we use the resolved domain name as "token id"
			if transfer.Token.Address == internal.ENSContractAddress || transfer.Token.Address == internal.ENSNameWrapperContractAddress {
				// set custom collection name
				collection.Name = "ENS"

				// get ens token metadata
				metadata, err := external.GetENSMetadataForTokenID(transfer.Token.ID)
				if err == nil && metadata != nil {
					ensMetadata = metadata

					fmtTokenID.Reset()
					fmtTokenID.WriteString(style.TerminalLink(metadata.URL, collection.Render(metadata.Name)))
				} else {
					gbl.Log.Debugf("getting ens metadata failed: %s | %v", fmt.Sprint(transfer.Token.ID), err)
				}
			}

			isOwnCollection = collection.Source == collectionsource.FromWallet || collection.Source == collectionsource.FromConfiguration

			// link each token id to opensea
			_, openseaURL, _ := utils.GetLinks(txHash, transfer.Token.Address, transfer.Token.ID.Int64())

			if collection == nil && transfer.Standard == standard.ERC1155 {
				collection = tokencollections.GetCollection(gb, contractAddress, transfer.Token.ID.Int64())
			}

			// for erc1155 we get the current total supply of the token
			// useful for mints, burns, etc., especially if they happen over a long* period of time
			var fmtTotalSupply string

			if transfer.Standard == standard.ERC1155 && (isOwnWallet || isOwnCollection || isWatchUsersWallet) {
				// if supply, err := gb.Nodes.TotalSupplyERC1155(ctx, transfer.Token.Address, transfer.Token.ID); err == nil {
				if supply, err := gb.ProviderPool.ERC1155TotalSupply(ctx, transfer.Token.Address, transfer.Token.ID); err == nil {
					// fmtTokenID.WriteString(style.DarkGrayStyle.Render(" /") + collection.StyleSecondary().Copy().Faint(true).Render(supply.String()))
					fmtTotalSupply = style.DarkGrayStyle.Render(" /") + collection.StyleSecondary().Copy().Faint(true).Render(supply.String())
				}
			}

			fmtTokenIds[transfer.Token.Address] = append(fmtTokenIds[transfer.Token.Address], style.TerminalLink(openseaURL, fmtTokenID.String())+fmtTotalSupply)
			fmtHistoryTokenIds[transfer.Token.Address] = append(fmtHistoryTokenIds[transfer.Token.Address], formatTokenID(collection, transfer.Token.ID)+fmtTotalSupply)

			if isOwnCollection {
				currentCollection = collection
			}

			numCollectionTokens += transfer.AmountTokens.Int64()
		}

		if ttx.IsReBurn() {
			numCollectionTokens /= 2
		}

		fmtEvent := strings.Builder{}
		fmtHistoryEvent := strings.Builder{}

		if numCollectionTokens > 1 {
			numberStyle, _ := getNumberStyles(int(numCollectionTokens))

			//
			// check if this was sweep or someone just dumped a lot of tokens into (blur) bids
			if tfFrom := ttx.GetNonZeroNFTSenders(); len(tfFrom) > 0 {
				for sender, transfers := range tfFrom {
					if numCollectionTokens == int64(len(transfers)) && ttx.From == sender {
						// all tokens sold by the same address -> dumped into bids
						// isBidDump = true
						if numCollectionTokens > 5 || ttx.GetPrice().Ether() > 3.0 {
							// if its a significant amount of tokens or ether we use a reddish style
							numberStyle = style.TrendRedStyle
						} else {
							// otherwise we use a light red style
							numberStyle = style.TrendLightRedStyle
						}
					}
				}
			}

			fmtEvent.WriteString(numberStyle.Render(fmt.Sprintf("%d", numCollectionTokens)) + "x ")
			fmtHistoryEvent.WriteString(numberStyle.Render(fmt.Sprintf("%d", numCollectionTokens)) + "x ")
		}

		// handle special cases
		switch {
		// if the collection name is empty we use a placeholder for now
		case collection.Name == "":
			name = "⸘Unknown‽"

		// if its an ENS nft, we use the resolved domain name as token id and slightly modify the collection name
		case (contractAddress == internal.ENSContractAddress || contractAddress == internal.ENSNameWrapperContractAddress) && ensMetadata != nil:
			name = collection.Style().Copy().Faint(true).Render(collection.Name + collection.StyleSecondary().Render(":"))

		// default collection name
		default:
			name = collection.Render(collection.Name)
		}

		if ttx.IsListing() {
			// print token name for listing
			name = collection.Render(ttx.Transfers[0].Token.Name)
		}

		maxShown := 5
		idsShown := int(math.Min(float64(len(fmtTokenIds[contractAddress])), float64(maxShown)))

		// use a variant without a link for the history
		// needed due to a bug causing unnecessary line breaks
		fmtEvent.WriteString(name + " " + strings.Join(fmtTokenIds[contractAddress][:idsShown], collection.StyleSecondary().Copy().Faint(true).Render(", ")))
		fmtHistoryEvent.WriteString(name + " " + strings.Join(fmtHistoryTokenIds[contractAddress][:idsShown], collection.StyleSecondary().Copy().Faint(true).Render(", ")))

		if len(fmtTokenIds[contractAddress]) > maxShown && collection != nil {
			fmtEvent.WriteString(collection.StyleSecondary().Render("…"))
			fmtHistoryEvent.WriteString(collection.StyleSecondary().Render("…"))
		}

		fmtTokensTransferred = append(fmtTokensTransferred, fmtEvent.String())

		if collection.Show.History {
			fmtTokensHistory = append(fmtTokensHistory, fmtHistoryEvent.String())
		}

		// counting
		if !ttx.IsListing() {
			ttx.TotalTokens += numCollectionTokens
		}

		// collection counting
		switch ttx.Action {
		case totra.Sale, totra.Purchase:
			collection.AddSale(ttx.AmountPaid, uint64(numCollectionTokens))
		case totra.Mint:
			collection.AddMintVolume(ttx.AmountPaid, uint64(numCollectionTokens))
		}
	}

	// total counting
	if gb.Stats != nil {
		switch ttx.Action {
		case totra.Sale, totra.Purchase:
			gb.Stats.AddSale(ttx.TotalTokens, ttx.AmountPaid)
		case totra.Mint:
			gb.Stats.AddMint(ttx.TotalTokens)
		}
	}

	if ttx.IsListing() {
		timeNow = style.Gray7Style.Render(currentTime)
	} else if isOwnCollection {
		timeNow = currentCollection.Style().Copy().Bold(true).Render(currentTime)
	}

	// highlight line if the seller or buyer is a wallet from the configured wallets
	if isOwnWallet {
		timeNow = lipgloss.NewStyle().Foreground(style.Pink).Bold(true).Render(currentTime)
	}

	// is own wallet or collection
	isOwn := isOwnWallet || isOwnCollection

	// time & type
	out.WriteString(timeNow)
	out.WriteString(" " + ttx.Action.Icon())
	out.WriteString(" " + divider)

	var fixWidthPrice string
	if ttx.GetPrice().Ether() < 100.0 {
		fixWidthPrice = fmt.Sprintf("%6.3f", ttx.GetPrice().Ether())
	} else {
		fixWidthPrice = fmt.Sprintf("%6.2f", ttx.GetPrice().Ether())
	}

	if before, after, found := strings.Cut(fixWidthPrice, "."); found {
		beforeSepStyle := style.DarkWhiteStyle
		sepStyle := style.GrayStyle

		switch {
		case ttx.Action == totra.Burn || ttx.Action == totra.ReBurn || ttx.Action == totra.Airdrop || ttx.Action == totra.Transfer:
			beforeSepStyle = style.DarkerGrayStyle
			sepStyle = style.LightGrayStyle.Copy().Faint(true)

			priceStyle = style.DarkerGrayStyle

		case ttx.GetPrice().Ether() < 1.0:
			beforeSepStyle = style.GrayStyle
			sepStyle = priceStyle

		case ttx.GetPrice().Ether() > 0.0:
			// get a color with saturation depending on the tx price
			beforeSepStyle = style.DarkWhiteStyle
			sepStyle = priceStyle.Copy().Foreground(priceArrowColor).Faint(true)

			priceStyle = style.DarkWhiteStyle
		}

		fixWidthPrice = beforeSepStyle.Render(before) + sepStyle.Render(".") + priceStyle.Render(after)
	}

	if len(fmtTokensTransferred) == 0 {
		gbl.Log.Debugf("🧐 no tokens transferred: %s | %+v", style.TerminalLink(utils.GetEtherscanTxURL(txHash.String()), txHash.String()), ttx.Transfers)

		for _, transfer := range ttx.Transfers {
			gbl.Log.Debugf(
				"  transfer of %dx %s | %+v",
				transfer.AmountTokens,
				style.TerminalLink(utils.GetEtherscanTokenURL(transfer.Token.Address.String()), style.ShortenAddress(&transfer.Token.Address)),
				transfer.Standard,
			)
		}

		return
	}

	// price
	fmtPrice := fixWidthPrice
	out.WriteString(" " + fmtPrice + formattedCurrencySymbol)

	// if all collections in a tx have the IgnorePrinting flag set, don't print the tx
	for _, collection := range ttxCollections {
		if !collection.IgnorePrinting {
			ttx.DoNotPrint = false

			break
		}

		ttx.DoNotPrint = true
	}

	// average price (makes no sense for multi-collections tx)
	averagePrice := ttx.GetPrice()
	if ttx.TotalTokens > 1 {
		averagePrice = price.NewPrice(big.NewInt(0).Div(ttx.AmountPaid, big.NewInt(ttx.TotalTokens)))
	}

	formattedAveragePriceEther := fmt.Sprintf("%6.3f", averagePrice.Ether())

	//
	// min value
	// belowAvgPriceMultiplier := 3.0
	totalValueBelowMinValue := ttx.GetPrice().Ether() < viper.GetFloat64("show.min_value")
	// avgValueBelowMultiMinValue := averagePrice.Ether()*belowAvgPriceMultiplier < viper.GetFloat64("show.min_value")
	if ttx.GetPrice().Ether() > 0.0 && totalValueBelowMinValue { // && avgValueBelowMultiMinValue {
		gbl.Log.Debugf("price is below min_value, not showing")

		ttx.DoNotPrint = true
	}

	// average price per item
	pricePerItemStyle := style.DarkerGrayStyle
	if averagePrice.Wei().Cmp(ttx.GetPrice().Wei()) < 0 {
		_, pricePerItemStyle = getNumberStyles(int(ttx.TotalTokens))
	}

	// print average per-item price (does not make sense anymore in multi-collection tx)
	out.WriteString(" " + pricePerItemStyle.Render(formattedAveragePriceEther))
	out.WriteString(formattedFaintCurrencySymbol)

	// // floor price TODO fix this
	// var trendIndicatorStyle lipgloss.Style
	// if currentCollection != nil {
	// 	trendIndicatorStyle = style.CreateTrendIndicator(currentCollection.PreviousFloorPrice, currentFloorPrice)
	// } else {
	// 	trendIndicatorStyle = style.CreateTrendIndicator(0.0, currentFloorPrice)
	// }

	currentFloorPriceStyle := style.DarkerGrayStyle

	// if currentFloorPrice > 0.0 {
	// 	currentFloorPriceStyle = style.GrayStyle.Copy().Faint(true)
	// }

	// trendIndicatorFaintStyle := trendIndicatorStyle.Copy().Faint(true)
	// out.WriteString(" " + currentFloorPriceStyle.Render(fmt.Sprintf("%6.3f", currentFloorPrice)) + trendIndicatorFaintStyle.Render("Ξ"))

	// print sales for collection
	if viper.GetBool("show.sales") {
		out.WriteString(" | " + fmt.Sprintf("%dx", currentCollection.Counters.Sales) + style.BoldStyle.Render(""))

		if currentCollection.Counters.Sales < 10 {
			out.WriteString(" ")
		}
		if currentCollection.Counters.Sales < 100 {
			out.WriteString(" ")
		}
		// print bluechip collection sales
		if ticker.BlueChips != nil && ticker.BlueChips.GetStats(currentCollection.ContractAddress) != nil {
			out.WriteString("/" + lipgloss.NewStyle().Foreground(style.OpenseaToneBlue).Faint(true).Render(fmt.Sprintf("%d", ticker.BlueChips.GetStats(currentCollection.ContractAddress).Sales)))
		} else {
			out.WriteString(" ")
		}
		// print collection volume
		volume := utils.WeiToEther(currentCollection.Counters.SalesVolume)
		out.WriteString(" | " + currentFloorPriceStyle.Render(fmt.Sprintf("Σ%6.1fΞ", volume)))

		if ttx.Action == totra.Mint {
			out.WriteString(" | " + fmt.Sprintf("%dx", currentCollection.Counters.Mints))
		}
	}

	//
	// show the burned token(s) on the same line on the right side
	if ttx.Action == totra.ReBurn && len(fmtTokensTransferred) == 2 && burnedTokenTransferIndex >= 0 && burnedTokenTransferIndex < len(fmtTokensTransferred) {
		// flips between 0 and 1 depending on the burnedTokenTransferIndex
		redeemedTokenTransferIndex := 1 - burnedTokenTransferIndex

		fmtTokensTransferred = []string{fmt.Sprint(
			fmtTokensTransferred[redeemedTokenTransferIndex],
			style.TrendRedStyle.Render("  ⇄  "),
			fmtTokensTransferred[burnedTokenTransferIndex],
		)}
	}

	// show the first collection/token on the same line
	// and further collections/tokens on the next lines
	out.WriteString("  " + fmtTokensTransferred[0] + " ")

	// links blur
	if ttx.TotalTokens == 1 {
		if ttx.Transfers[0].Standard == standard.ERC721 {
			out.WriteString(" | " + style.GrayBoldStyle.Copy().Foreground(style.BlurOrange).Faint(true).Render(style.TerminalLink(blurURL, "BL")))
		}
	}

	// link etherscan
	if ttx.Action != totra.Listing {
		out.WriteString(" | " + style.GrayStyle.Render(style.TerminalLink(etherscanURL, "ES")))
	}

	// // for burns the line ends after the etherscan link and we do not need a trailing pipe
	// if !ttx.IsBurn() {
	// 	out.WriteString(" | ")
	// }

	out.WriteString(" | ")

	var transferFrom common.Address

	// show "from" if its not a listing
	if !ttx.IsMint() && !ttx.IsListing() {
		var fmtFrom string

		if tfFrom := ttx.GetNonZeroNFTSenders(); len(tfFrom) > 0 {
			for fromAddr := range tfFrom {
				transferFrom = fromAddr

				break
			}
		} else {
			transferFrom = ttx.From
		}

		fromStyle := lipgloss.NewStyle().Foreground(style.GenerateColorWithSeed(transferFrom.Hash().Big().Int64()))

		if fromENS, err := gb.ProviderPool.ReverseResolveAddressToENS(ctx, transferFrom); err == nil {
			gbl.Log.Debugf("🤷 from address %s has ENS %s", transferFrom.Hex(), fromENS)
			fmtFrom = fromStyle.Render(fromENS)
		} else {
			gbl.Log.Debugf("🤷‍♀️ from address %s has NO ENS", transferFrom.Hex())
			fmtFrom = style.ShortenAddressStyled(&transferFrom, fromStyle)
		}

		out.WriteString(fmtFrom)
	}

	// buyer
	var fmtBuyer string

	buyer := ttx.Transfers[0].To

	if ttx.IsReBurn() && len(ttx.Transfers) > 1 && buyer == internal.ZeroAddress {
		buyer = ttx.Transfers[1].To
	} else if ttx.IsBurn() {
		for _, fromAddresses := range ttx.GetNFTSenderAddresses() {
			if fromAddresses != internal.ZeroAddress {
				buyer = fromAddresses

				break
			}
		}
	}

	buyerStyle := lipgloss.NewStyle().Foreground(style.GenerateColorWithSeed(buyer.Hash().Big().Int64()))

	if buyerENS, err := gb.ProviderPool.ReverseResolveAddressToENS(ctx, buyer); err == nil {
		gbl.Log.Debugf("✅ resolved ENS name for %s: %s", buyer.Hex(), buyerENS)

		fmtBuyer = buyerStyle.Render(buyerENS)
	} else {
		gbl.Log.Debugf("❌ failed to resolve ENS name for %s: %s", buyer.Hex(), err)

		fmtBuyer = style.ShortenAddressStyled(&buyer, buyerStyle)
	}

	arrow := style.DividerArrowRight
	if ttx.IsListing() || ttx.IsBurn() {
		arrow = style.DividerArrowLeft
	}

	out.WriteString(arrow.String() + fmtBuyer)

	// 'maybe important wallet' indicator
	if wwatcher.MIWC.MIWs.Contains(buyer) {
		level := strings.Repeat(" 👀", int(math.Min(3.0, float64(wwatcher.MIWC.WeightedMIWs[buyer]))))
		out.WriteString(" " + level)

		// out.WriteString("   " + style.PinkBoldStyle.Render(level))
	}

	// sales/listings count & salira | think about how to do this for multi-collection tx?!
	if currentCollection.Counters.Sales+currentCollection.Counters.Listings > 0 {
		var salesAndListings string

		if currentCollection.Counters.Listings > 0 {
			salesAndListings = fmt.Sprint(
				style.TrendLightGreenStyle.Render(fmt.Sprint(currentCollection.Counters.Sales)),
				currentCollection.Render("/"),
				style.TrendLightRedStyle.Render(fmt.Sprint(currentCollection.Counters.Listings)),
			)
		} else {
			salesAndListings = fmt.Sprint(style.TrendLightGreenStyle.Render(fmt.Sprint(currentCollection.Counters.Sales)))
		}

		out.WriteString(" | " + salesAndListings)

		if previousMASaLiRa, currentMASaLiRa := currentCollection.CalculateSaLiRa(currentCollection.ContractAddress, gb.Rueidi); currentMASaLiRa > 0 {
			// coloring moving average salira
			saLiRaStyle := style.TrendGreenStyle

			if previousMASaLiRa > currentMASaLiRa {
				saLiRaStyle = style.TrendRedStyle
			}

			salira := fmt.Sprint(
				style.CreateTrendIndicator(previousMASaLiRa, currentMASaLiRa),
				saLiRaStyle.Render(fmt.Sprintf("%4.2f", currentMASaLiRa)),
				// currentCollection.Render("slr"),
			)

			out.WriteString(style.GrayStyle.Render(" ~ ") + salira)
		} else if cachedSalira, err := gb.Rueidi.GetCachedSalira(ctx, currentCollection.ContractAddress); cachedSalira > 0 && err == nil {
			salira := fmt.Sprint(
				style.GrayStyle.Render(" ~ "),
				style.GrayStyle.Render(fmt.Sprintf("%4.2f", cachedSalira)),
				currentCollection.Render("*"),
			)

			out.WriteString(salira)
		}
	}

	// multi-line output for multi-collection events
	if len(fmtTokensTransferred) > 1 {
		for _, fmtTokenCollection := range fmtTokensTransferred[1:] {
			out.WriteString("\n" + strings.Repeat(" ", 31))
			out.WriteString(style.DarkGrayStyle.Render("+") + fmtTokenCollection)
		}
	}

	// add blue chip icons
	if viper.GetBool("notifications.bluechip.enabled") {
		if ticker.BlueChips.ContainsWallet(buyer) && ttx.Action != totra.Burn {
			if ticker.BlueChips.CollectionStats[currentCollection.ContractAddress] != nil {
				out.WriteString(" | " + fmt.Sprintf("%d", ticker.BlueChips.CollectionStats[currentCollection.ContractAddress].Sales) + style.BoldStyle.Render("🔵"))
			}

			for i, blueChipTypes := range ticker.BlueChips.WalletMap[buyer].Holder {
				if i == 0 {
					out.WriteString("·")
				}

				out.WriteString(style.BoldStyle.Render(ticker.GetEmojiMapping(blueChipTypes)))
			}
		}
	}

	// add manifold event to manifold ticker
	if viper.GetBool("notifications.manifold.enabled") && (!viper.GetBool("notifications.disabled")) {
		if ttx.IsMovingNFTs() && ttx.Tx.To() != nil && ticker.Manifold.IsManifoldContractAddress(*ttx.Tx.To()) {
			if viper.GetBool("notifications.manifold.enabled") {
				gbl.Log.Debugf("tx %s is a tx to the manifold (lazy claim) contract", ttx.TxReceipt.TxHash.Hex())
				ticker.Manifold.AppendManifoldEvent(ttx)
			}
		}
	}

	//
	// building output string done
	//

	// dont apply excludes to "own" events
	if !(isOwnWallet || isWatchUsersWallet) {
		// DoNotPrint can be set by the "pipeline" the tx is going through (e.g. when a collection has the IgnorePrinting flag set)
		if ttx.DoNotPrint {
			gbl.Log.Debugf("skipping tx %s | doNotPrint flaf: %v | %+v", style.Bold(txHash.String()), ttx.DoNotPrint, ttx)

			return
		}

		if !currentCollection.Show.Mints && (ttx.Action == totra.Mint || ttx.Action == totra.Airdrop) && !viper.GetBool("show.mints") {
			gbl.Log.Debugf("skipping mint %s | viper.GetBool(show.mints): %v | %+v", style.Bold(txHash.String()), viper.GetBool("show.mints"), ttx)

			return
		}

		if (ttx.Action == totra.Burn) && !viper.GetBool("show.burns") {
			gbl.Log.Debugf("skipping burn/airdrop %s | viper.GetBool(show.burns): %v | %+v", style.Bold(txHash.String()), viper.GetBool("show.burns"), ttx)

			return
		}

		if (ttx.Action == totra.ReBurn) && !viper.GetBool("show.reburns") {
			gbl.Log.Debugf("skipping re-burn %s | viper.GetBool(show.burns): %v | %+v", style.Bold(txHash.String()), viper.GetBool("show.reburns"), ttx)

			return
		}

		if (ttx.Action == totra.Transfer) && !viper.GetBool("show.transfers") {
			gbl.Log.Debugf("skipping transfer %s | viper.GetBool(show.transfers): %v | %+v", style.Bold(txHash.String()), viper.GetBool("show.transfers"), ttx)

			return
		}

		if (ttx.Action == totra.Unknown) && !viper.GetBool("show.unknown") {
			gbl.Log.Debugf("skipping unknown %s | viper.GetBool(show.unknown): %v | %+v", style.Bold(txHash.String()), viper.GetBool("show.unknown"), ttx)

			return
		}
	}

	//
	// 🌈 finally print the sale/listing/whatever 🌈
	if !viper.GetBool("ui.headless") {
		terminalPrinterQueue <- out.String()
	}

	// add to history
	if isOwn && !ttx.IsLoan() {
		if (!ttx.IsListing() || (ttx.IsListing() && isOwnWallet)) && gb.Stats != nil {
			gb.Stats.EventHistory = append(gb.Stats.EventHistory, ttx.AsHistoryTokenTransaction(currentCollection, fmtTokensHistory))
		}
	}

	// outputLine := OutputLine{
	// 	PlatformSymbol: ttx.Marketplace.Tag,
	// 	DateTime:       currentTime,
	// 	PriceSymbol:    "→",
	// 	ActionSymhol:   ttx.Action.Icon(),
	// 	FromToSymbol:   "⇄",
	// 	Price:          ttx.GetPrice().Ether(),
	// 	PricePerItem:   averagePrice.Ether(),
	// 	CollectionName: currentCollection.Name,
	// 	TokenIDs:       []int64{0, 2, 4}, // TODO: ttx.TokenIDsByContract()
	// 	TxHash:         txHash.String(),
	// 	Buyer:          osmodels.Account{Address: buyer.String(), User: ""},
	// 	Seller:         osmodels.Account{Address: transferFrom.String(), User: ""},
	// 	NumSales:       int(currentCollection.Counters.Sales),
	// 	NumListings:    int(currentCollection.Counters.Listings),
	// 	SaLiRa:         currentCollection.SaLiRa.Value(),

	// 	Colors: OutputColors{
	// 		Platform:            ttx.Marketplace.Color,
	// 		DateTime:            style.DarkGray,
	// 		PriceSymbol:         priceArrowColor,
	// 		FromToSymbol:        style.DarkGray,
	// 		Collection:          currentCollection.Colors.Primary,
	// 		CollectionSecondary: currentCollection.Colors.Secondary,
	// 		Buyer:               style.GenerateColorWithSeed(buyer.Hash().Big().Int64()),
	// 		Seller:              style.GenerateColorWithSeed(transferFrom.Hash().Big().Int64()),
	// 		SaLiRa:              style.DarkGray,
	// 	},
	// }

	// // terminalPrinterQueue <- fmt.Sprint(outputLine)

	// gbl.Log.Debugf("outputLine: %+v", outputLine)
}

func getNumberStyles(numEvents int) (lipgloss.Style, lipgloss.Style) {
	var numberStyle, pricePerItemStyle lipgloss.Style

	switch {
	case numEvents > 7:
		numberStyle = style.AlmostWhiteStyle
		pricePerItemStyle = style.DarkWhiteStyle
	case numEvents > 4:
		numberStyle = style.DarkWhiteStyle
		pricePerItemStyle = style.LightGrayStyle
	case numEvents > 1:
		numberStyle = style.LightGrayStyle
		pricePerItemStyle = style.GrayStyle
	default:
		numberStyle = style.GrayStyle
		pricePerItemStyle = style.DarkGrayStyle
	}

	return numberStyle, pricePerItemStyle
}

func formatTokenID(collection *collections.Collection, tokenID *big.Int) string {
	shortened := false

	// shorten token id if it's too long
	if tokenID.Cmp(big.NewInt(999_999)) > 0 {
		tokenID = big.NewInt(tokenID.Int64() % 10000)
		shortened = true
	}

	// token id
	prefix := collection.StyleSecondary().Render("#")
	id := collection.Style().Render(fmt.Sprint(tokenID))

	if shortened {
		id += collection.StyleSecondary().Render("…")
	}

	return prefix + id
}

type OutputColors struct {
	Platform            lipgloss.Color
	DateTime            lipgloss.Color
	PriceSymbol         lipgloss.Color
	FromToSymbol        lipgloss.Color
	Collection          lipgloss.Color
	CollectionSecondary lipgloss.Color
	Buyer               lipgloss.Color
	Seller              lipgloss.Color
	SaLiRa              lipgloss.Color
}

type OutputLine struct {
	PlatformSymbol string
	DateTime       string
	PriceSymbol    string
	ActionSymhol   string
	FromToSymbol   string
	Price          float64
	PricePerItem   float64
	CollectionName string
	TokenIDs       []int64
	TxHash         string
	Buyer          osmodels.Account
	Seller         osmodels.Account
	NumSales       int
	NumListings    int
	SaLiRa         float64

	Colors OutputColors
}

package trapri

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/cache"
	"github.com/benleb/gloomberg/internal/collections"
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

	for ttx := range queueTokenTransactions {
		go formatTokenTransaction(gb, ttx, terminalPrinterQueue)

		// send to ws
		queueWsOutTokenTransactions <- ttx
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

	// fake a txHash for listings
	txHash := common.Hash{}
	if ttx.Action != totra.Listing || ttx.Tx != nil {
		txHash = ttx.Tx.Hash()
	}

	// // print the eventTx
	// originNode := gb.Nodes.GetRandomNode()

	// isMintOrSaleOrPurchase := collections.EventType(ttx.Action) == collections.Mint || collections.EventType(ttx.Action) == collections.Sale || collections.EventType(ttx.Action) == collections.Purchase

	// is a collections from configured collections + own wallets
	isOwnCollection := false

	// for _, event := range ttx.Events {
	// 	if event.Collection != nil {
	// 		event.Collection = GetCollection(gb, event.ContractAddress, event.TokenID)

	// 		if event.Collection != nil && event.Collection.IgnorePrinting {
	// 			gbl.Log.Debugf("event/collection %s is ignored for printing", style.Bold(event.Collection.Name))

	// 			return
	// 		}

	// 		if event.Collection.Name == "" && event.ContractAddress == internal.ENSContractAddress {
	// 			event.Collection.Name = "ENS"
	// 		}
	// 	}
	// }

	// // unify events
	// unifyEventsMap := make(map[string]*trotra.TokenEvent, 0)

	// for _, ev := range ttx.Events {
	// 	if ev.Collection != nil {
	// 		unifyEventsMap[utils.GetNFTID(ev.ContractAddress, ev.TokenID)] = ev
	// 	}
	// }

	// uniqueNFTEvents := make([]*trotra.TokenEvent, 0)

	// for _, ev := range unifyEventsMap {
	// 	uniqueNFTEvents = append(uniqueNFTEvents, ev)
	// }

	// ttx.Events = uniqueNFTEvents

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

	// price-dependent styling
	// var currentFloorPrice float64

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

	// migration helper
	marker := " "

	// // origin node marker
	// if viper.GetBool("log.debug") {
	// 	out.WriteString(originNode.GetStyledMarker())
	// }

	// out.WriteString(style.GrayStyle.Render("·"))
	out.WriteString(ttx.Marketplace.RenderFaintTag())

	// timestamp styling
	// WEN...??
	now := time.Now()
	currentTime := now.Format("15:04:05")
	timeNow := style.GrayStyle.Copy().Faint(true).Render(currentTime)

	// // explicitly configured colors
	// // TODO implement multi-collection handling
	// switch {
	// case collections.EventType(ttx.Action) == collections.Sale && ttx.Events[0].Collection.Highlight.Sales != "":
	// 	timeNow = lipgloss.NewStyle().Foreground(ttx.Events[0].Collection.Highlight.Sales).Render(currentTime)
	// case ttx.Events[0].Collection.Highlight.Color != "":
	// 	timeNow = lipgloss.NewStyle().Foreground(ttx.Events[0].Collection.Highlight.Color).Render(currentTime)
	// }

	// prepare links
	etherscanURL, _, blurURL := utils.GetLinks(txHash, ttx.Transfers[0].Token.Address, ttx.Transfers[0].Token.ID.Int64())
	//
	// price per item
	// numEvents := len(ttx.EventsByContract[ttx.Events[0].ContractAddress])
	// numEventsPerContract := make(map[common.Address]int, 0)

	// numEvents := 0
	// for _, events := range ttx.EventsByContract {
	// 	numEvents += len(events)
	// }

	// if numEvents == 0 {
	// 	gbl.Log.Warnf("numEvents == 0 for %s", txHash)

	// 	return
	// }

	// if ttx.Events[0].Standard == standard.ERC1155 {
	// 	numEvents = 0
	// 	for _, event := range ttx.EventsByContract {
	// 		for _, event := range event {
	// 			numEvents += int(event.Amount.Int64())
	// 		}
	// 	}
	// }

	// print collection name and token id
	// TODO implement multi-collection handling
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
		case totra.LoanContracts[contractAddress] != "":
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
				// if supply, err := gb.Nodes.TotalSupplyERC1155(context.Background(), transfer.Token.Address, transfer.Token.ID); err == nil {
				if supply, err := gb.ProviderPool.ERC1155TotalSupply(context.Background(), transfer.Token.Address, transfer.Token.ID); err == nil {
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

		// counting
		switch ttx.Action {
		case totra.Sale, totra.Purchase:
			// numItems := uint64(numCollectionTokens)
			// amountPaid := ttx.AmountPaid

			// // if its a bid dump we just count half of the items
			// if isBidDump {
			// 	numItems /= 2
			// 	amountPaid = amountPaid.Div(amountPaid, big.NewInt(2))
			// }

			// collection.AddSale(amountPaid, numItems)

			collection.AddSale(ttx.AmountPaid, uint64(numCollectionTokens))
		case totra.Mint:
			collection.AddMint()
		}
	}

	// TODO implement multi-collection handling
	if isOwnCollection && ttx.IsListing() {
		coloredColon := currentCollection.Render(":")
		timeStyle := style.Gray7Style.Render
		timeNow = fmt.Sprint(timeStyle(fmt.Sprintf("%02d", now.Hour())), coloredColon, timeStyle(fmt.Sprintf("%02d", now.Minute())), coloredColon, timeStyle(fmt.Sprintf("%02d", now.Second())))
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
	out.WriteString(marker + ttx.Action.Icon())
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

	// TODO: disable for multi-collection tx?
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

	// currentFloorPriceStyle := style.DarkerGrayStyle

	// if currentFloorPrice > 0.0 {
	// 	currentFloorPriceStyle = style.GrayStyle.Copy().Faint(true)
	// }

	// trendIndicatorFaintStyle := trendIndicatorStyle.Copy().Faint(true)
	// out.WriteString(" " + currentFloorPriceStyle.Render(fmt.Sprintf("%6.3f", currentFloorPrice)) + trendIndicatorFaintStyle.Render("Ξ"))

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
	if !ttx.IsListing() {
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

		if fromENS, err := gb.ProviderPool.ResolveENSForAddress(context.TODO(), transferFrom); err == nil {
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

	if buyerENS, err := gb.ProviderPool.ResolveENSForAddress(context.TODO(), buyer); err == nil {
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

	// maybe importan wallet indicator
	if wwatcher.MIWC.MIWs.Contains(buyer) {
		var miwLevel string
		if wwatcher.MIWC.WeightedMIWs[buyer] > 1 {
			miwLevel = "⭐ " + strconv.Itoa(wwatcher.MIWC.WeightedMIWs[buyer]) + " ⭐"
		} else {
			miwLevel = strconv.Itoa(wwatcher.MIWC.WeightedMIWs[buyer])
		}

		out.WriteString("   " + style.PinkBoldStyle.Render(fmt.Sprintf("👀 MIW! %s 👀", miwLevel)))
	}

	// TODO think about how to do this for multi-collection tx
	// sales/listings count & salira
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

		if previousMASaLiRa, currentMASaLiRa := currentCollection.CalculateSaLiRa(currentCollection.ContractAddress); currentMASaLiRa > 0 {
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
		} else if cachedSalira, err := cache.GetSalira(context.TODO(), currentCollection.ContractAddress); cachedSalira > 0 && err == nil {
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

	if !viper.GetBool("ui.headless") {
		terminalPrinterQueue <- out.String()
	}

	if isOwn && ticker.StatsTicker != nil && !ttx.IsLoan() {
		if !ttx.IsListing() || (ttx.IsListing() && isOwnWallet) {
			ticker.StatsTicker.EventHistory = append(ticker.StatsTicker.EventHistory, ttx.AsHistoryTokenTransaction(currentCollection, fmtTokensHistory))
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

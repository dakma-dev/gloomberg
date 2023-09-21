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
	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/external"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/jobs"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg/gbgrpc"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg/gbgrpc/gen"
	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/benleb/gloomberg/internal/nemo/standard"
	"github.com/benleb/gloomberg/internal/nemo/tokencollections"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/benleb/gloomberg/internal/notify"
	"github.com/benleb/gloomberg/internal/opensea"
	seawatcher "github.com/benleb/gloomberg/internal/seawa"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/ticker"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/benleb/gloomberg/internal/utils/wwatcher"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/kr/pretty"
	"github.com/spf13/viper"
)

// func TokenTransactionFormatter(gb *gloomberg.Gloomberg, seawa *seawatcher.SeaWatcher, queueWsOutTokenTransactions chan *totra.TokenTransaction, queueWsInTokenTransactions chan *totra.TokenTransaction) {.
func TokenTransactionFormatter(gb *gloomberg.Gloomberg, seawa *seawatcher.SeaWatcher) {
	gbl.Log.Debugf("🧱 starting ttx formatter worker")

	// if viper.GetBool("websockets.client.enabled") {
	// 	go func() {
	// 		for ttx := range queueWsInTokenTransactions {
	// 			go formatTokenTransaction(gb, seawa, ttx)
	// 		}
	// 	}()
	// }

	// ! critical path !
	// this is the main loop of the formatter
	// blocking/delaying here will block/delay the whole stream
	// when adding additional calls here, prefer goroutines with conditional select

	tokenTransactionsChannel := gb.SubscribeTokenTransactions()

	for workerID := 1; workerID <= viper.GetInt("server.workers.ttxFormatter"); workerID++ {
		go func() {
			for ttx := range tokenTransactionsChannel {
				go formatTokenTransaction(gb, seawa, ttx)
			}
		}()
	}
}

func formatTokenTransaction(gb *gloomberg.Gloomberg, seawa *seawatcher.SeaWatcher, ttx *totra.TokenTransaction) {
	ctx := context.Background()

	// parsed event to be used for the web-ui
	parsedEvent := degendb.ParsedEvent{Other: make(map[string]interface{})}

	// fake a txHash for listings
	txHash := common.Hash{}
	if ttx.Tx != nil && ttx.TxHash != (common.Hash{}) {
		txHash = ttx.TxHash
	} else {
		// generate random Hash
		txHash = common.BytesToHash(ttx.From.Bytes()) //  *ttx.Tx.To().Bytes())
	}

	parsedEvent.TxHash = txHash

	// is a collections from configured collections + own wallets
	isOwnCollection := false

	// a watched wallet is involved
	nftTransactors := ttx.GetNFTSenderAndReceiverAddresses()
	isOwnWallet := gb.OwnWallets.ContainsAddressFromSlice(nftTransactors.ToSlice()) != internal.ZeroAddress
	isWatchUsersWallet := gb.Watcher.ContainsAddressFromSlice(nftTransactors.ToSlice()) != internal.ZeroAddress

	// is this an intentional purchase or a dump into bids?
	// isBidDump := false

	// telegram notification
	if viper.GetBool("notifications.telegram.enabled") && (isOwnWallet || isWatchUsersWallet) { //  && ttx.Action != degendb.Transfer {
		gbl.Log.Infof("🧱 sending telegram notification | isOwnWallet: %+v | isWatchUsersWallet: %+v", isOwnWallet, isWatchUsersWallet)

		go notify.SendNotification(gb, ttx)
	}

	// if it's a single-collection transaction we set the collection as the currentCollection
	// from here on already, otherwise we set it to nil and fill it later in the loop
	// over the collections/transfers
	var currentCollection *collections.Collection

	if len(ttx.GetTransfersByContract()) >= 1 && currentCollection == nil || ttx.Action == degendb.CollectionOffer {
		currentCollection = tokencollections.GetCollection(gb, ttx.Transfers[0].Token.Address, ttx.Transfers[0].Token.ID.Int64())
	}

	// defaults
	priceStyle := style.DarkWhiteStyle
	priceArrowColor := style.DarkGray

	if ttx.GetPrice() != nil && ttx.GetPrice().Ether() >= 0.01 {
		priceArrowColor = style.GetPriceShadeColor(ttx.GetPrice().Ether())
	}

	parsedEvent.Colors.PriceArrow = priceArrowColor

	switch ttx.Action {
	case degendb.BurnRedeem, degendb.Burn, degendb.Airdrop, degendb.Transfer:
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

	var ok bool
	parsedEvent.Colors.PriceCurrency, ok = priceCurrencyStyle.GetForeground().(lipgloss.Color)
	if !ok {
		parsedEvent.Colors.PriceCurrency = style.DarkGray
	}

	formattedCurrencySymbol := priceCurrencyStyle.Render("Ξ")
	formattedFaintCurrencySymbol := priceCurrencyStyle.Copy().Faint(true).Render("Ξ")

	out.WriteString(ttx.Marketplace.RenderFaintTag())

	// timestamp styling
	// WEN...??
	now := time.Now()
	currentTime := now.Format("15:04:05")
	timeNow := style.Gray5Style.Render(currentTime)

	parsedEvent.ReceivedAt = now

	// prepare links
	etherscanURL, openSeaURL, blurURL := utils.GetLinks(txHash, ttx.Transfers[0].Token.Address, ttx.Transfers[0].Token.ID.Int64())

	parsedEvent.BlurURL = blurURL
	parsedEvent.EtherscanURL = etherscanURL
	parsedEvent.OpenSeaURL = openSeaURL

	// print collection name and token id
	fmtTokensTransferred := make([]string, 0)
	fmtTokensHistory := make([]string, 0)
	ttxCollections := make(map[common.Address]*collections.Collection)

	// contract addresses of the burned token(s)
	// used in reburnes for nicer formatting
	burnedTokenTransferIndex := -1

	transferredCollections := make([]degendb.TransferredCollection, 0)

	for contractAddress, transfers := range ttx.GetTransfersByContract() {
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

		fmtTokenIds := make(map[common.Address][]string)
		fmtHistoryTokenIds := make(map[common.Address][]string)

		var name string

		collection := tokencollections.GetCollection(gb, contractAddress, transfers[0].Token.ID.Int64())

		var ensMetadata *external.ENSMetadata

		if collection.IgnorePrinting {
			continue
		}

		//
		// experiment: fetch & store the first txs for this contract

		if viper.GetBool("experiments.firsttxs") {
			collectionFileNamePrefix := collection.Name
			if collection.OpenseaSlug != "" {
				collectionFileNamePrefix = collection.OpenseaSlug
			}

			if collectionFileNamePrefix != "" && ttx.GetPrice().Ether() >= viper.GetFloat64("gloomberg.firstTxs.min_value") {
				jobs.AddJob("firsttxs", "etherscan", gloomberg.JobFirstTxsForContract, collectionFileNamePrefix, contractAddress)
			}
		}

		ttxCollections[contractAddress] = collection

		numCollectionTokens := int64(0)

		transferredTokens := make([]degendb.TransferredToken, 0)

		for _, transfer := range transfers {
			fmtTokenID := strings.Builder{}

			transferredToken := degendb.TransferredToken{}

			// ignore transfers of more than 9999 tokens
			if transfer.AmountTokens.Cmp(big.NewInt(9999)) > 0 {
				gbl.Log.Debugf("♾️ amountTokens > 9999 for token %s", transfer.Token.ShortID())

				continue
			}

			// highlight grifters
			if collection.ContractAddress == internal.GrifterContractAddress {
				if transfer.Token.ID.Uint64() < 666 {
					collection.Name = "Grifter"
				} else {
					gbl.Log.Infof("🧱 grifter token id out of range: %d", transfer.Token.ID.Int64())

					continue
				}
			}

			transferredToken.ID = transfer.Token.ID.Int64()

			// add rank if available
			var fmtRank string
			if gb.Ranks[transfer.Token.Address] != nil {
				if rank := gb.Ranks[transfer.Token.Address][transfer.Token.ID.Int64()].Rank; rank > 0 {
					rankSymbol := gb.Ranks[transfer.Token.Address][transfer.Token.ID.Int64()].GetRankSymbol(collection.Metadata.TotalSupply)

					transferredToken.Rank = rank
					transferredToken.RankSymbol = rankSymbol

					// fmtRank := style.TrendLightGreenStyle.Copy().Bold(false).Render(fmt.Sprintf("%d%s", rank, rankSymbol))
					fmtRank := lipgloss.NewStyle().Foreground(style.OpenseaToneBlue).Render(fmt.Sprintf("%d%s", rank, rankSymbol))

					fmtTokenID.WriteString(fmtRank)
				}
			}

			// add number of tokens transferred
			if transfer.AmountTokens.Cmp(big.NewInt(1)) > 0 {
				fmtTokenID.WriteString(style.DarkGrayStyle.Render(transfer.AmountTokens.String() + "x"))

				transferredToken.Amount = transfer.AmountTokens.Int64()
			}

			transferredTokens = append(transferredTokens, transferredToken)

			fmtTokenID.WriteString(formatTokenID(collection, transfer.Token.ID))

			// add a marker for burned tokens
			if transfer.To == internal.ZeroAddress {
				fmtTokenID.WriteString("🔥")

				if ttx.Action == degendb.BurnRedeem {
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

			isOwnCollection = collection.Source == collections.FromWallet || collection.Source == collections.FromConfiguration

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
			fmtHistoryTokenIds[transfer.Token.Address] = append(fmtHistoryTokenIds[transfer.Token.Address], fmtRank+formatTokenID(collection, transfer.Token.ID)+fmtTotalSupply)

			if isOwnCollection {
				currentCollection = collection
			}

			numCollectionTokens += transfer.AmountTokens.Int64()

			if viper.GetBool("experiments.eip6551") {
				jobs.AddJob("eip6551", "node", gloomberg.JobCheckEIP6551TokenAccount, gb, transfer.Token.Address, transfer.Token.ID)
			}
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
							// if it's a significant amount of tokens or ether we use a reddish style
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

		transferredCollection := degendb.TransferredCollection{
			CollectionName: collection.Name,
			From:           ttx.From.Hex(),

			TransferredTokens: transferredTokens,

			Colors: degendb.CollectionColors{
				Primary:   collection.Colors.Primary,
				Secondary: collection.Colors.Secondary,
			},
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

		// total supply
		if currentCollection.Metadata.TotalSupply > 0 && currentCollection.Metadata.TotalSupply < 99999 {
			fmtTotalSupply := fmt.Sprint(currentCollection.Metadata.TotalSupply)

			if currentCollection.Metadata.TotalSupply > 999 {
				shortTotalSupply := int(currentCollection.Metadata.TotalSupply / 1000)
				fmtTotalSupply = fmt.Sprint(shortTotalSupply) + "k"
			}

			fmtEvent.WriteString(style.DarkGrayStyle.Render(" /") + collection.StyleSecondary().Copy().Faint(true).Render(fmtTotalSupply))
		}

		fmtTokensTransferred = append(fmtTokensTransferred, fmtEvent.String())

		if collection.Show.History {
			fmtTokensHistory = append(fmtTokensHistory, fmtHistoryEvent.String())
		}

		// count sales/purchases
		if degendb.SaleTypes.Contains(ttx.Action) {
			ttx.TotalTokens += numCollectionTokens

			collection.AddSales(ttx.AmountPaid, uint64(numCollectionTokens))
		}

		// count mints
		if ttx.Action == degendb.Mint {
			collection.AddMintVolume(ttx.AmountPaid, uint64(numCollectionTokens))
		}

		transferredCollections = append(transferredCollections, transferredCollection)
	}

	// // add number of tokens transferred
	// if transfer.AmountTokens.Cmp(big.NewInt(1)) > 0 {
	// 	fmtTokenID.WriteString(style.DarkGrayStyle.Render(transfer.AmountTokens.String() + "x"))

	// 	transferredToken.Amount = transfer.AmountTokens.Int64()
	// }

	// log.Printf("transfer.AmountTokens: %+v", transfer.AmountTokens)

	parsedEvent.TransferredCollections = transferredCollections

	// total counting
	if gb.Stats != nil {
		var eventType degendb.EventType
		switch ttx.Action {
		case degendb.Sale, degendb.Purchase:
			eventType = degendb.Sale
		case degendb.Mint:
			eventType = degendb.Mint
		case degendb.Listing:
			eventType = degendb.Listing
		}

		gb.Stats.AddEvent(eventType, ttx.TotalTokens, ttx.AmountPaid)
	}

	parsedEvent.Colors.Time = style.DarkGray

	if ttx.IsListing() {
		timeNow = style.Gray7Style.Render(currentTime)
		parsedEvent.Colors.Time = style.Gray7
	} else if isOwnCollection {
		timeNow = currentCollection.Style().Copy().Bold(true).Render(currentTime)
		parsedEvent.Colors.Time = currentCollection.Colors.Primary
	}

	// highlight line if the seller or buyer is a wallet from the configured wallets
	if isOwnWallet {
		timeNow = lipgloss.NewStyle().Foreground(style.Pink).Bold(true).Render(currentTime)
		parsedEvent.Colors.Time = lipgloss.Color(style.Pink.Dark)
	}

	// is our own wallet or collection
	isOwn := isOwnWallet || isOwnCollection

	// time & type
	out.WriteString(timeNow)
	out.WriteString(" " + ttx.Action.Icon())
	out.WriteString(" " + divider)

	parsedEvent.Action = ttx.Action.String()
	parsedEvent.Typemoji = ttx.Action.Icon()

	var fixWidthPrice string
	if ttx.GetPrice() != nil && ttx.GetPrice().Ether() < 100.0 {
		fixWidthPrice = fmt.Sprintf("%6.3f", ttx.GetPrice().Ether())
	} else {
		fixWidthPrice = fmt.Sprintf("%6.2f", ttx.GetPrice().Ether())
	}

	if before, after, found := strings.Cut(fixWidthPrice, "."); found {
		beforeSepStyle := style.DarkWhiteStyle
		sepStyle := style.GrayStyle

		switch {
		case ttx.Action == degendb.Burn || ttx.Action == degendb.BurnRedeem || ttx.Action == degendb.Airdrop || ttx.Action == degendb.Transfer:
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

	parsedEvent.Colors.Price, ok = priceStyle.GetForeground().(lipgloss.Color)
	if !ok {
		parsedEvent.Colors.Price = style.DarkGray
	}

	if len(fmtTokensTransferred) == 0 && ttx.Tx != nil {
		gbl.Log.Debugf("🧐 no tokens transferred: %s | %+v", style.TerminalLink(utils.GetEtherscanTxURL(txHash.String()), txHash.String()), ttx.Transfers)
		gbl.Log.Debugf("no tokens transferred: %+v", fmt.Sprintf("%+v", pretty.Formatter(ttx)))

		for _, transfer := range ttx.Transfers {
			gbl.Log.Debugf(
				"  transfer of %dx %s | %+v",
				transfer.AmountTokens,
				style.TerminalLink(utils.GetEtherscanTokenURLForAddress(transfer.Token.Address), style.ShortenAddress(transfer.Token.Address)),
				transfer.Standard,
			)
		}

		return
	}

	// price
	fmtPrice := fixWidthPrice
	out.WriteString(" " + fmtPrice + formattedCurrencySymbol)

	parsedEvent.Price = ttx.GetPrice() // fmt.Sprintf("%6.3f", ttx.GetPrice().Ether())

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

	priceWidth := "%6.3f"
	if averagePrice.Ether() >= 100.0 {
		priceWidth = "%6.2f"
	}

	formattedAveragePriceEther := fmt.Sprintf(priceWidth, averagePrice.Ether())

	//
	// min value
	// if the average price is below the min_value and the total price is below
	// the min_value * min_value_multiplier, don't show the tx in the stream
	if !isOwnCollection || (!ttx.IsListing() && !ttx.IsItemBid() && !ttx.IsCollectionOffer()) {
		if ttx.GetPrice().Ether() > 0.0 && averagePrice.Ether() > 0.0 {
			minValue := viper.GetFloat64("show.min_value")

			averageBelowMinValue := averagePrice.Ether() < minValue
			totalBelowMultiMinValue := ttx.GetPrice().Ether() < minValue*viper.GetFloat64("show.min_value_multiplier")

			gbl.Log.Debugf("total: %f | avg: %f | averageBelowMinValue: %+v | totalBelowMultiMinValue: %+v", ttx.GetPrice().Ether(), averagePrice.Ether(), averageBelowMinValue, totalBelowMultiMinValue)

			if averageBelowMinValue && totalBelowMultiMinValue {
				gbl.Log.Debugf("price is below min_value, not showing")

				ttx.DoNotPrint = true
			}
		}
	}

	out.WriteString(ttx.GetPurchaseOrBidIndicator())

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
		numLastSales, _ := currentCollection.GetSaLiCount()

		out.WriteString(" | " + fmt.Sprintf("%dx", numLastSales) + style.BoldStyle.Render(""))

		if numLastSales < 10 {
			out.WriteString(" ")
		}
		if numLastSales < 100 {
			out.WriteString(" ")
		}
		// print bluechip collection sales
		if ticker.BlueChips != nil && ticker.BlueChips.GetStats(currentCollection.ContractAddress) != nil {
			out.WriteString("/" + lipgloss.NewStyle().Foreground(style.OpenseaToneBlue).Faint(true).Render(fmt.Sprintf("%d", ticker.BlueChips.GetStats(currentCollection.ContractAddress).GetTXCount())))
		} else {
			out.WriteString(" ")
		}
		// print collection volume
		volume := utils.WeiToEther(currentCollection.Counters.SalesVolume)
		out.WriteString(" | " + currentFloorPriceStyle.Render(fmt.Sprintf("Σ%6.1fΞ", volume)))

		if ttx.Action == degendb.Mint {
			out.WriteString(" | " + fmt.Sprintf("%dx", currentCollection.Counters.Mints))
		}
	}

	//
	// show the burned token(s) on the same line on the right side
	if ttx.Action == degendb.BurnRedeem && len(fmtTokensTransferred) == 2 && burnedTokenTransferIndex >= 0 && burnedTokenTransferIndex < len(fmtTokensTransferred) {
		// flips between 0 and 1 depending on the burnedTokenTransferIndex
		redeemedTokenTransferIndex := 1 - burnedTokenTransferIndex

		fmtTokensTransferred = []string{fmt.Sprint(
			fmtTokensTransferred[redeemedTokenTransferIndex],
			style.TrendRedStyle.Render("  ⇄  "),
			fmtTokensTransferred[burnedTokenTransferIndex],
		)}
	}

	if len(fmtTokensTransferred) == 0 {
		gloomberg.PrWarn(fmt.Sprintf("no tokens transferred in tx %s", style.TerminalLink(etherscanURL)))

		gbl.Log.Warnf("no tokens transferred: %+v", fmt.Sprintf("%+v", pretty.Formatter(ttx)))

		return
	}
	// show the first collection/token on the same line
	// and further collections/tokens on the next lines
	out.WriteString("  " + fmtTokensTransferred[0] + " ") //nolint:gosec

	// links blur
	if ttx.TotalTokens == 1 {
		if ttx.Transfers[0].Standard == standard.ERC721 {
			out.WriteString(" | " + style.GrayBoldStyle.Copy().Foreground(style.BlurOrange).Faint(true).Render(style.TerminalLink(blurURL, "BL")))
		}
	}

	// link etherscan
	if ttx.Action != degendb.Listing {
		out.WriteString(" | " + style.GrayStyle.Render(style.TerminalLink(etherscanURL, "ES")))
	}

	// // for burns the line ends after the etherscan link, and we do not need a trailing pipe
	// if !ttx.IsBurn() {
	// 	out.WriteString(" | ")
	// }

	out.WriteString(" | ")

	var transferFrom common.Address

	// show "from" if it's not a listing
	if !ttx.IsMint() && !ttx.IsListing() && !ttx.IsItemBid() && !ttx.IsCollectionOffer() {
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
		parsedEvent.Colors.From, ok = fromStyle.GetForeground().(lipgloss.Color)
		if !ok {
			parsedEvent.Colors.From = style.DarkGray
		}

		if fromENS, err := gb.ProviderPool.ReverseResolveAddressToENS(ctx, transferFrom); err == nil {
			gbl.Log.Debugf("🤷 from address %s has ENS %s", transferFrom.Hex(), fromENS)
			fmtFrom = fromStyle.Render(fromENS)
			parsedEvent.From = gb.DegenDB.NewDegen(fromENS, []common.Address{transferFrom}, "", "", 0, []degendb.Tag{})

			parsedEvent.FromAddress = transferFrom
		} else {
			gbl.Log.Debugf("🤷‍♀️ from address %s has NO ENS", transferFrom.Hex())
			fmtFrom = style.ShortenAddressStyled(&transferFrom, fromStyle)
			// shortName := style.ShortenAddress(&transferFrom)
			parsedEvent.From = gb.DegenDB.NewDegen(fromENS, []common.Address{transferFrom}, "", "", 0, []degendb.Tag{})
			parsedEvent.FromAddress = transferFrom
		}

		out.WriteString(fmtFrom)
	}

	// buyer
	var fmtBuyer string

	buyer := ttx.Transfers[0].To

	if ttx.IsReBurn() && len(ttx.Transfers) > 1 && buyer == internal.ZeroAddress {
		buyer = ttx.Transfers[1].To
	} else if ttx.IsBurn() {
		for _, fromAddresses := range ttx.GetNFTSenderAddresses().ToSlice() {
			if fromAddresses != internal.ZeroAddress {
				buyer = fromAddresses

				break
			}
		}
	}

	buyerStyle := lipgloss.NewStyle().Foreground(style.GenerateColorWithSeed(buyer.Hash().Big().Int64()))

	parsedEvent.Colors.To, ok = buyerStyle.GetForeground().(lipgloss.Color)
	if !ok {
		parsedEvent.Colors.To = style.DarkGray
	}

	if buyerENS, err := gb.ProviderPool.ReverseResolveAddressToENS(ctx, buyer); err == nil {
		gbl.Log.Debugf("✅ resolved ENS name for %s: %s", buyer.Hex(), buyerENS)

		fmtBuyer = buyerStyle.Render(buyerENS)

		parsedEvent.To = gb.DegenDB.NewDegen(buyerENS, []common.Address{buyer}, "", "", 0, []degendb.Tag{})
		parsedEvent.ToAddress = buyer
	} else {
		gbl.Log.Debugf("❌ failed to resolve ENS name for %s: %s", buyer.Hex(), err)

		fmtBuyer = style.ShortenAddressStyled(&buyer, buyerStyle)

		parsedEvent.To = gb.DegenDB.NewDegen(buyerENS, []common.Address{buyer}, "", "", 0, []degendb.Tag{})
		parsedEvent.ToAddress = buyer
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

	// don't apply excludes to "own" events
	if !(isOwnWallet || isWatchUsersWallet) {
		// DoNotPrint can be set by the "pipeline" the tx is going through (e.g. when a collection has the IgnorePrinting flag set)
		if ttx.DoNotPrint {
			log.Debugf("skipping tx %s | doNotPrint flaf: %v | %+v", style.Bold(txHash.String()), ttx.DoNotPrint, ttx)

			return
		}

		if !isOwnCollection && !currentCollection.Show.Mints && (ttx.Action == degendb.Mint || ttx.Action == degendb.Airdrop) && !viper.GetBool("show.mints") {
			log.Debugf("skipping mint %s | viper.GetBool(show.mints): %v | %+v", style.Bold(txHash.String()), viper.GetBool("show.mints"), ttx)

			return
		}

		if (ttx.Action == degendb.Burn) && !viper.GetBool("show.burns") {
			log.Debugf("skipping burn/airdrop %s | viper.GetBool(show.burns): %v | %+v", style.Bold(txHash.String()), viper.GetBool("show.burns"), ttx)

			return
		}

		if (ttx.Action == degendb.BurnRedeem) && !viper.GetBool("show.reburns") {
			log.Debugf("skipping re-burn %s | viper.GetBool(show.burns): %v | %+v", style.Bold(txHash.String()), viper.GetBool("show.reburns"), ttx)

			return
		}

		if (ttx.Action == degendb.Transfer) && !viper.GetBool("show.transfers") {
			log.Debugf("skipping transfer %s | viper.GetBool(show.transfers): %v | %+v", style.Bold(txHash.String()), viper.GetBool("show.transfers"), ttx)

			return
		}

		if (ttx.Action == degendb.Unknown) && !viper.GetBool("show.unknown") {
			log.Debugf("skipping unknown %s | viper.GetBool(show.unknown): %v | %+v", style.TerminalLink(txHash.String(), style.ShortenHashStyled(txHash)), viper.GetBool("show.unknown"), ttx)

			return
		}
	}

	// sales/listings count & salira | think about how to do this for multi-collection tx?!
	numLastSales, numLastListings := currentCollection.GetSaLiCount()

	if numLastSales+numLastListings > 0 {
		var salesAndListings string

		if numLastListings > 0 {
			salesAndListings = fmt.Sprint(
				style.TrendLightGreenStyle.Render(fmt.Sprint(numLastSales)),
				currentCollection.Render("/"),
				style.TrendLightRedStyle.Render(fmt.Sprint(numLastListings)),
			)
		} else {
			salesAndListings = fmt.Sprint(style.TrendLightGreenStyle.Render(fmt.Sprint(numLastSales)))

			//
			// auto-subscribe to opensea events after X sales
			if autoSubscribeAfterSales := viper.GetUint64("seawatcher.auto_subscribe_after_sales"); uint64(numLastSales) >= autoSubscribeAfterSales {
				if currentCollection.OpenseaSlug == "" {
					currentCollection.OpenseaSlug = opensea.GetCollectionSlug(currentCollection.ContractAddress)
				}

				if !seawa.IsSubscribed(currentCollection.OpenseaSlug) {
					// if seawa.SubscribeForSlug(currentCollection.OpenseaSlug) {
					// gprcClient := gbgrpc.NewClient("")

					if gbgrpc.GRPCClient != nil {
						subsriptionRequest := &gen.SubscriptionRequest{EventTypes: []gen.EventType{gen.EventType_ITEM_LISTED}, Collections: []string{currentCollection.OpenseaSlug}} //nolint:nosnakecase

						_, err := gbgrpc.GRPCClient.Subscribe(context.Background(), subsriptionRequest, nil)
						if err != nil {
							log.Errorf("failed to subscribe to events for %s: %s", currentCollection.OpenseaSlug, err)
						}

						log.Printf("subsubsubbbyyyyyyyyyyyyyyy: %+v", subsriptionRequest)
					}

					if seawa.SubscribeForSlug(currentCollection.OpenseaSlug, []gen.EventType{gen.EventType_ITEM_LISTED, gen.EventType_ITEM_RECEIVED_BID, gen.EventType_COLLECTION_OFFER}) > 0 { //nolint:nosnakecase
						seawa.Pr(fmt.Sprintf("auto-subscribed to events for %s (after %d sales) | stats resetted", style.AlmostWhiteStyle.Render(currentCollection.OpenseaSlug), autoSubscribeAfterSales))
					}
				}
			}
		}

		out.WriteString(" | " + salesAndListings)

		//
		// SaLiRas
		if timeframedSaLiRas := currentCollection.GetPrettySaLiRas(); len(timeframedSaLiRas) > 0 {
			out.WriteString(style.DarkGrayStyle.Render(" ~ ") + strings.Join(timeframedSaLiRas, " "))

			// add collection symbol ad the end for easier matching between salira and collection
			if currentCollection.Metadata != nil && currentCollection.Metadata.Symbol != "" {
				out.WriteString(style.DarkGrayStyle.Render(" | ") + currentCollection.Style().Copy().Faint(true).Render(currentCollection.Metadata.Symbol))
			}
		}
	}

	// multi-line output for multi-collection events
	if len(fmtTokensTransferred) > 1 {
		for _, fmtTokenCollection := range fmtTokensTransferred[1:] { //nolint:gosec
			out.WriteString("\n" + strings.Repeat(" ", 32))
			out.WriteString(style.DarkGrayStyle.Render("+") + fmtTokenCollection)
		}
	}

	// add blue chip icons
	if viper.GetBool("notifications.bluechip.enabled") {
		if ticker.BlueChips.ContainsWallet(buyer) && ttx.Action != degendb.Burn {
			if ticker.BlueChips.CollectionStats[currentCollection.ContractAddress] != nil {
				out.WriteString(" | " + fmt.Sprintf("%d", ticker.BlueChips.CollectionStats[currentCollection.ContractAddress].Sales) + style.BoldStyle.Render("🔵"))
			}

			for i, blueChipTypes := range ticker.BlueChips.WalletMap[buyer].Types {
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

	// // don't apply excludes to "own" events
	// if !(isOwnWallet || isWatchUsersWallet) {
	// 	// DoNotPrint can be set by the "pipeline" the tx is going through (e.g. when a collection has the IgnorePrinting flag set)
	// 	if ttx.DoNotPrint {
	// 		gbl.Log.Debugf("skipping tx %s | doNotPrint flaf: %v | %+v", style.Bold(txHash.String()), ttx.DoNotPrint, ttx)

	// 		return
	// 	}

	// 	if !currentCollection.Show.Mints && (ttx.Action == degendb.Mint || ttx.Action == degendb.Airdrop) && !viper.GetBool("show.mints") {
	// 		gbl.Log.Debugf("skipping mint %s | viper.GetBool(show.mints): %v | %+v", style.Bold(txHash.String()), viper.GetBool("show.mints"), ttx)

	// 		return
	// 	}

	// 	if (ttx.Action == degendb.Burn) && !viper.GetBool("show.burns") {
	// 		gbl.Log.Debugf("skipping burn/airdrop %s | viper.GetBool(show.burns): %v | %+v", style.Bold(txHash.String()), viper.GetBool("show.burns"), ttx)

	// 		return
	// 	}

	// 	if (ttx.Action == degendb.BurnRedeem) && !viper.GetBool("show.reburns") {
	// 		gbl.Log.Debugf("skipping re-burn %s | viper.GetBool(show.burns): %v | %+v", style.Bold(txHash.String()), viper.GetBool("show.reburns"), ttx)

	// 		return
	// 	}

	// 	if (ttx.Action == degendb.Transfer) && !viper.GetBool("show.transfers") {
	// 		gbl.Log.Debugf("skipping transfer %s | viper.GetBool(show.transfers): %v | %+v", style.Bold(txHash.String()), viper.GetBool("show.transfers"), ttx)

	// 		return
	// 	}

	// 	if (ttx.Action == degendb.Unknown) && !viper.GetBool("show.unknown") {
	// 		gbl.Log.Debugf("skipping unknown %s | viper.GetBool(show.unknown): %v | %+v", style.TerminalLink(txHash.String(), style.ShortenHashStyled(txHash)), viper.GetBool("show.unknown"), ttx)

	// 		return
	// 	}
	// }

	//
	// 🌈 finally print the sale/listing/whatever 🌈
	if !viper.GetBool("ui.headless") {
		// terminalPrinterQueue <- out.String()
		if ttx.IsListing() && !isOwnCollection {
			return
		}

		// highlight special events with newlines above and below
		printLine := out.String()
		if ttx.Highlight {
			printLine = "\n" + printLine + "\n"
		}

		// print to terminal
		// gb.In.PrintToTerminal <- printLine
		gloomberg.TerminalPrinterQueue <- printLine

		gb.In.ParsedEvents <- &parsedEvent
	}

	// add to history
	if isOwn && (!ttx.IsLoan() && !ttx.IsItemBid()) { // && ttx.Action != degendb.ItemBid && ttx.Action != degendb.CollectionOffer {
		if (!ttx.IsListing() || (ttx.IsListing() && isOwnWallet)) && currentCollection.Source != collections.FromConfiguration && gb.Stats != nil {
			// gb.Stats.EventHistory = append(gb.Stats.EventHistory, ttx.AsHistoryTokenTransaction(currentCollection, fmtTokensHistory))

			// TODO: fix/remove this...
			parsedEvent.Other["fmtTokensHistory"] = fmtTokensHistory

			parsedEvent.IsOwnWallet = isOwnWallet
			parsedEvent.IsOwnCollection = isOwnCollection
			parsedEvent.IsWatchUsersWallet = isWatchUsersWallet
			parsedEvent.PurchaseOrBidIndicator = ttx.GetPurchaseOrBidIndicator()

			// ...and actually use this!!
			gb.RecentOwnEvents.Add(&parsedEvent)

			// new event added to own recent events, send the whole slice to the ui
			if gb.RecentOwnEvents.Cardinality() > 0 {
				gb.In.RecentOwnEvents <- gb.RecentOwnEvents.ToSlice()
			}

			gbl.Log.Debugf("trapri added event to history: %+v", gb.RecentOwnEvents.Cardinality())
		}
	}
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

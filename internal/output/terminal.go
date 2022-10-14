package output

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/cache"
	"github.com/benleb/gloomberg/internal/chainwatcher/wwatcher"
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/external"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/models/gloomberg"
	"github.com/benleb/gloomberg/internal/models/topic"
	"github.com/benleb/gloomberg/internal/nodes"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/ticker"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/benleb/gloomberg/internal/utils/notifications"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

// var rarities = openrarity.LoadRaritiesFromJSONs()
func FormatEvent(gb *gloomberg.Gloomberg, event *collections.Event, queueOutput chan<- string) {
	gbl.Log.Debugf("FormatEvent | event: %+v", event)

	var collection *collections.GbCollection

	if event.Collection == nil {
		gbl.Log.Warnf("FormatEvent | event.Collection is nil")

		//
		// collection information
		gb.CollectionDB.RWMu.RLock()
		collection = gb.CollectionDB.Collections[event.ContractAddress]
		gb.CollectionDB.RWMu.RUnlock()

		if collection == nil && event.ContractAddress != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			name := ""

			if topic.Topic(event.Topic) == topic.TransferSingle && gb.Nodes != nil {
				if tokenName, err := gb.Nodes.GetRandomNode().GetERC1155TokenName(event.ContractAddress, event.TokenID); err == nil && tokenName != "" {
					name = tokenName
					gbl.Log.Debugf("found token name: %s | %s", name, event.ContractAddress.String())
				} else if err != nil {
					gbl.Log.Debugf("failed to get collection name: %s", err)
				}
			}

			collection = collections.NewCollection(event.ContractAddress, name, gb.Nodes, models.FromStream)

			if collection != nil {
				gb.CollectionDB.RWMu.Lock()
				gb.CollectionDB.Collections[event.ContractAddress] = collection
				gb.CollectionDB.RWMu.Unlock()
			} else {
				// atomic.AddUint64(&StatsBTV.DiscardedUnknownCollection, 1)
				gbl.Log.Warnf("🗑️ collection is nil | cw.CollectionDB.UserCollections[subLog.Address] -> %v | %v | TxHash: %v / %d", gb.CollectionDB.Collections[event.ContractAddress], event.ContractAddress.String(), event.TxHash, event.TxItemCount)
				return
			}
		}

		event.Collection = collection
	}

	var (
		priceStyle      lipgloss.Style
		priceArrowColor lipgloss.Color
	)

	eventWithStyle := &collections.EventWithStyle{}

	isMint := event.EventType == collections.Mint
	isMintOrTransfer := event.EventType == collections.Mint || event.EventType == collections.Transfer
	isSaleOrMint := event.EventType == collections.Sale || event.EventType == collections.Purchase || event.EventType == collections.Mint
	isMultiItemTx := event.TxItemCount > 1

	// var (
	// 	priceStyle      lipgloss.Style
	// 	priceArrowColor lipgloss.Color
	// )

	var pricePerItem *big.Int
	if event.EventType == collections.Sale && isMultiItemTx {
		// pricePerItem = event.PriceWei.Div() / big.NewInt(event.TxItemCount)
		pricePerItem = big.NewInt(0).Div(event.PriceWei, big.NewInt(int64(event.TxItemCount)))
	} else {
		// pricePerItem = event.PriceWei.Uint64()
		pricePerItem = event.PriceWei
	}

	event.PricePerItem = pricePerItem
	event.CollectionColor = event.Collection.Colors.Primary

	// format price in ether
	priceEther, _ := nodes.WeiToEther(event.PriceWei).Float64()
	priceEtherPerItem, _ := nodes.WeiToEther(pricePerItem).Float64()

	var previousMovingAverage, currentMovingAverage float64

	if event.EventType == collections.Sale {
		if gb.OwnWallets.Contains(event.To.Address) {
			event.EventType = collections.Purchase
		}

		// recalculate moving average
		itemPrice, _ := nodes.WeiToEther(pricePerItem).Float64()
		previousMovingAverage, currentMovingAverage = event.Collection.CalculateArtificialFloor(itemPrice)

		// get a color with saturation depending on the tx price
		priceStyle = style.DarkWhiteStyle
		priceArrowColor = style.GetPriceShadeColor(priceEther)

		eventWithStyle.PriceEtherColor = "#DDDDDD"
		eventWithStyle.PriceArrowColor = style.GetPriceShadeColor(priceEther)
	} else {
		// if this is a mint/transfer/listing, we don't touch the moving average
		currentMovingAverage = event.Collection.ArtificialFloor.Value()
		previousMovingAverage = currentMovingAverage

		priceStyle = style.GrayStyle
		priceArrowColor = "#333333"

		eventWithStyle.PriceEtherColor = "#666666"
		eventWithStyle.PriceArrowColor = "#333333"
	}

	// item number style
	var numberStyle, pricePerItemStyle lipgloss.Style

	switch {
	case event.TxItemCount > 7:
		numberStyle = style.AlmostWhiteStyle
		pricePerItemStyle = style.DarkWhiteStyle
	case event.TxItemCount > 4:
		numberStyle = style.DarkWhiteStyle
		pricePerItemStyle = style.LightGrayStyle
	case event.TxItemCount > 1:
		numberStyle = style.LightGrayStyle
		pricePerItemStyle = style.GrayStyle
	default:
		numberStyle = style.GrayStyle
		pricePerItemStyle = style.DarkGrayStyle
	}

	priceCurrencyStyle := event.Collection.Style().Copy().Faint(isMintOrTransfer)
	formattedCurrencySymbol := priceCurrencyStyle.Render("Ξ")
	currentMovingAverageStyle := style.GrayStyle.Copy().Faint(isMintOrTransfer)

	trendIndicator := style.CreateTrendIndicator(previousMovingAverage, currentMovingAverage)
	divider := style.Sharrow.Copy().Foreground(priceArrowColor).String()

	isOwnCollection := event.Collection.Source == models.FromWallet || event.Collection.Source == models.FromConfiguration
	isOwnWallet := gb.OwnWallets.Contains(event.From.Address) || gb.OwnWallets.Contains(event.To.Address)
	isWatchUsersWallet := gb.WatchUsers.Contains(event.From.Address) || gb.WatchUsers.Contains(event.To.Address)
	listingBelowPrice := event.Collection.Highlight.ListingsBelowPrice > 0.0 && event.Collection.Highlight.ListingsBelowPrice <= priceEther

	// buyer
	toStyle := lipgloss.NewStyle().Foreground(style.GenerateColorWithSeed(event.To.Address.Hash().Big().Int64()))
	to := style.ShortenAddressStyled(&event.To.Address, toStyle)

	var ensName string

	// check if the ENS name is already in the cache
	if name, err := cache.GetENSName(event.To.Address); err == nil && name != "" {
		gbl.Log.Debugf("cache | cached ENS name: %s", name)
		ensName = name
	} else {
		gbl.Log.Debugf("cache | no cached ENS name for %s | trying to resolve...", event.To.Address.String())

		// if not, try to resolve it
		if viper.IsSet("api_keys.etherscan") && gb.Nodes != nil {
			gbl.Log.Debugf("cache | ENS name not cached, trying to resolve...")

			if name, err := gb.Nodes.GetENSForAddress(event.To.Address); err == nil && name != "" {
				gbl.Log.Debugf("cache | resolved ENS name: %s", name)
				ensName = name
			}
		}
	}

	if ensName != "" {
		to = toStyle.Render(ensName)
		event.ToENS = ensName
		cache.StoreENSName(event.To.Address, ensName)
	} else if event.ToENS != "" {
		to = toStyle.Render(event.ToENS)
	}

	// if ensName == "" && viper.IsSet("api_keys.etherscan") && gb.Nodes != nil {
	// 	gbl.Log.Infof("cache | ENS name not cached, trying to resolve...")

	// 	if name, err := gb.Nodes.GetENSForAddress(event.To.Address); err == nil && name != "" {
	// 		gbl.Log.Infof("cache | resolved ENS name: %s", name)

	// 		ensName = name
	// 		to = toStyle.Render(ensName)
	// 		event.ToENS = ensName

	// 		cache.StoreENSName(event.To.Address, ensName)
	// 	} else if event.ToENS != "" {
	// 		to = toStyle.Render(event.ToENS)
	// 	}
	// }

	// WEN...??
	now := time.Now()
	currentTime := now.Format("15:04:05")
	timeNow := style.GrayStyle.Copy().Faint(true).Render(currentTime)
	eventWithStyle.Time = now
	eventWithStyle.TimeColor = "#999999"

	eventWithStyle.CollectionName = event.Collection.Name

	// WHAT...??
	var tokenInfo string
	if isMultiItemTx {
		tokenInfo = fmt.Sprintf("%s %s", numberStyle.Render(fmt.Sprintf("%dx", event.TxItemCount)), event.Collection.Style().Faint(isMint).Render(event.Collection.Name))

		eventWithStyle.TxItemCount = event.TxItemCount
	} else if event.Collection.ContractAddress == external.ENSContract {
		ensName := "Ethereum Name Service"
		if event.ENSMetadata != nil && event.ENSMetadata.Name != "" {
			ensName = event.ENSMetadata.Name
		}

		tokenInfo = fmt.Sprintf(
			"%s %s",
			event.Collection.Style().Copy().Faint(true).Render(event.Collection.Name+":"),
			event.Collection.Style().Copy().Faint(false).Render(ensName),
		)

		eventWithStyle.TxItemCount = event.TxItemCount
	} else {
		tokenInfo = internal.FormatTokenInfo(event.TokenID, event.Collection, isMint, true)

		eventWithStyle.TokenID = event.TokenID
	}

	// PRETTY...??
	collectionStyle := lipgloss.NewStyle().Foreground(event.Collection.Colors.Primary)
	eventWithStyle.CollectionColor = event.Collection.Colors.Primary

	if event.EventType == collections.Sale && isOwnCollection {
		timeNow = collectionStyle.Render(currentTime)
		eventWithStyle.TimeColor = eventWithStyle.CollectionColor

		notifications.SendNotification(event.Collection.Name, tokenInfo)
		gbl.Log.Debugf("SendNotification | collection: %s, tokenInfo: %s", event.Collection.Name, tokenInfo)
	}

	// highlight line if the seller or buyer is a wallet from the configured wallets
	if isOwnWallet {
		if event.EventType == collections.Listing {
			timeNow = lipgloss.NewStyle().Foreground(style.OpenseaToneBlue).Bold(true).Render(currentTime)
		} else {
			timeNow = lipgloss.NewStyle().Foreground(style.Pink).Bold(true).Render(currentTime)
		}

		eventWithStyle.TimeColor = lipgloss.Color(style.OwnerGreen.Dark)
	}

	// check if listing is below configured max. price
	if listingBelowPrice {
		var timeStyle lipgloss.Style

		if event.EventType == collections.Listing {
			timeStyle = style.PinkBoldStyle
			priceStyle = style.BoldStyle
			eventWithStyle.TimeColor = lipgloss.Color(style.Pink.Dark)
		} else {
			timeStyle = lipgloss.NewStyle().Foreground(style.ShadesPink[3])
			eventWithStyle.TimeColor = style.ShadesPink[3]
		}

		timeNow = timeStyle.Render(currentTime)
	}

	switch {
	case event.EventType == collections.Sale && event.Collection.Highlight.Sales != "":
		timeNow = lipgloss.NewStyle().Foreground(event.Collection.Highlight.Sales).Render(currentTime)
		eventWithStyle.TimeColor = event.Collection.Highlight.Sales
	case event.EventType == collections.Listing && event.Collection.Highlight.Listings != "":
		timeNow = lipgloss.NewStyle().Foreground(event.Collection.Highlight.Listings).Render(currentTime)
		eventWithStyle.TimeColor = event.Collection.Highlight.Listings
	case event.Collection.Highlight.Color != "":
		timeNow = lipgloss.NewStyle().Foreground(event.Collection.Highlight.Color).Render(currentTime)
		eventWithStyle.TimeColor = event.Collection.Highlight.Color
	}

	arrow := style.DividerArrowRight
	if event.EventType == collections.Listing {
		arrow = style.DividerArrowLeft
	}

	var openseaURL string
	if event.Permalink != "" {
		openseaURL = event.Permalink
	} else {
		openseaURL = fmt.Sprintf("https://opensea.io/assets/%s/%d", event.Collection.ContractAddress, event.TokenID)
	}

	etherscanURL := fmt.Sprintf("https://etherscan.io/tx/%s", event.TxHash)

	marker := " "

	if listingBelowPrice {
		marker = style.PinkBoldStyle.Render("*")
	} else if isOwnCollection && event.EventType == collections.Sale {
		// if itemPrice, _ := priceEtherPerItem.Float64(); itemPrice >= viper.GetFloat64("show.min_value") {
		if priceEtherPerItem >= viper.GetFloat64("show.min_value") {
			if isOwnWallet {
				marker = style.OwnerGreenBoldStyle.Render("*")
				eventWithStyle.Marker = "*"
				eventWithStyle.MarkerColor = "#73F59F"
			}
		}
	}

	// add to event history
	if isOwnCollection && event.EventType == collections.Sale { // && itemPrice >= viper.GetFloat64("show.min_value") {
		ticker.StatsTicker.EventHistory = append(ticker.StatsTicker.EventHistory, event)
	} else if isOwnWallet {
		ticker.StatsTicker.EventHistory = append(ticker.StatsTicker.EventHistory, event)
	}

	eventWithStyle.Verbose = viper.GetBool("log.verbose")

	// build the line to be displayed
	out := strings.Builder{}

	if viper.GetBool("log.verbose") {
		if event.EventType == collections.Listing {
			out.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#20293d")).Render("OS"))

			eventWithStyle.Source = "OS"
			eventWithStyle.SourceColor = "#20293d"
		} else if gb.Nodes != nil && len(*gb.Nodes) > 0 {
			out.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#1A1A1A")).Render(fmt.Sprint(gb.Nodes.GetNodeByID(event.NodeID).Marker)))

			eventWithStyle.Source = utils.StripANSI(fmt.Sprint(gb.Nodes.GetNodeByID(event.NodeID).Marker))
			eventWithStyle.SourceColor = "#1A1A1A"
		} else {
			eventWithStyle.Source = lipgloss.NewStyle().Foreground(lipgloss.Color("#1A1A1A")).Render(fmt.Sprint(event.NodeID))
			eventWithStyle.SourceColor = "#1A1A1A"
		}

		out.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#111111")).Render("|"))
	}

	out.WriteString(timeNow)
	out.WriteString(marker + event.EventType.Icon())
	out.WriteString(" " + divider)

	eventWithStyle.EventType = event.EventType
	eventWithStyle.EventEmoji = event.EventType.Icon()
	eventWithStyle.Marker = marker

	eventWithStyle.PriceEther = fmt.Sprintf("%6.3f", priceEther)

	eventWithStyle.OpenseaURL = openseaURL
	eventWithStyle.EtherscanURL = etherscanURL

	// price
	if event.EventType == collections.Listing {
		out.WriteString(" " + priceStyle.Render(style.TerminalLink(openseaURL, fmt.Sprintf("%6.3f", priceEther))))
	} else {
		out.WriteString(" " + priceStyle.Render(fmt.Sprintf("%6.3f", priceEther)))
	}

	out.WriteString(formattedCurrencySymbol)

	// moving average (artificial) floor price
	out.WriteString("  " + trendIndicator)
	out.WriteString(currentMovingAverageStyle.Render(fmt.Sprintf("%6.3f", currentMovingAverage)))

	// price per item
	//if false && event.EventType == collections.Listing && rarities[event.Collection.OpenseaSlug] != nil {
	//	rank := rarities[event.Collection.OpenseaSlug][int(event.TokenID.Int64())].Rank
	//	out.WriteString("   " + style.PinkBoldStyle.Render(fmt.Sprintf("%4d", rank)) + " ")
	//} else {
	out.WriteString(" " + pricePerItemStyle.Render(fmt.Sprintf("%6.3f", nodes.WeiToEther(pricePerItem))))
	out.WriteString(priceCurrencyStyle.Copy().Faint(true).Render("Ξ"))

	// collection/token info
	out.WriteString("  " + tokenInfo)

	// total supply
	if totalSupply := event.Collection.Metadata.TotalSupply; totalSupply > 0 && !isMultiItemTx {
		outputTotalSupply := fmt.Sprintf("%d", totalSupply)
		totalSupplyStyle := collectionStyle

		// make it less noticeable for big collections
		if totalSupply > 1000 {
			outputTotalSupply = fmt.Sprintf("%dK", uint(totalSupply/1000))
			totalSupplyStyle = collectionStyle.Copy().Faint(true)
		}

		eventWithStyle.CollectionTotalSupply = totalSupply

		out.WriteString(style.DarkGrayStyle.Render(" /") + totalSupplyStyle.Render(outputTotalSupply))
	}

	// link opensea
	out.WriteString(" | " + style.GrayBoldStyle.Copy().Foreground(style.OpenseaToneBlue).Render(style.TerminalLink(openseaURL, "OpenSea")))

	// link etherscan
	if event.EventType != collections.Listing {
		out.WriteString(" | " + style.GrayStyle.Render(style.TerminalLink(etherscanURL, "Etherscan")))
	}

	// buyer
	out.WriteString(" | " + arrow.String())
	out.WriteString(" " + to)

	// maybe importan wallet indicator
	if wwatcher.MIWC.MIWs.Contains(event.To.Address) {
		var miwLevel string
		if wwatcher.MIWC.WeightedMIWs[event.To.Address] > 1 {
			miwLevel = "⭐ " + strconv.Itoa(wwatcher.MIWC.WeightedMIWs[event.To.Address]) + " ⭐"
		} else {
			miwLevel = strconv.Itoa(wwatcher.MIWC.WeightedMIWs[event.To.Address])
		}

		out.WriteString("   " + style.PinkBoldStyle.Render(fmt.Sprintf("👀 MIW! %s 👀", miwLevel)))
	}

	// log topic (for debugging)
	if viper.GetBool("log.debug") {
		out.WriteString(" | " + topic.Topic(event.Topic).String())
	}

	// // automatically fetch listings for collections with more than opensea.auto_list_min_sales sales
	// if event.Collection.Counters.Sales == viper.GetUint64("opensea.auto_list_min_sales") {
	// 	slug := opensea.GetCollectionSlug(event.Collection.ContractAddress)
	// 	ossw.SubscribeToListingsForCollectionSlug(nil, slug, nil)
	// 	event.Collection.ResetStats()

	// 	gbl.Log.Infof("auto-subscribed to %s after %d sales", event.Collection.Name, event.Collection.Counters.Sales)

	// 	queueOutput <- fmt.Sprintf(
	// 		" %s auto-subscribed to %s after %d sales",
	// 		style.PinkBoldStyle.Render(">"),
	// 		event.Collection.Name,
	// 		event.Collection.Counters.Sales,
	// 	)
	// }

	// counting
	if event.EventType == collections.Sale || event.EventType == collections.Purchase {
		event.Collection.AddSale(event.PriceWei, event.TxItemCount)
	} else if event.EventType == collections.Mint {
		event.Collection.AddMint()
	}

	// sales/listings count & salira
	eventWithStyle.ListingsCount = event.Collection.Counters.Listings
	eventWithStyle.SalesCount = event.Collection.Counters.Sales

	if event.Collection.Counters.Sales+event.Collection.Counters.Listings > 0 {
		var salesAndListings string

		if event.Collection.Counters.Listings > 0 {
			salesAndListings = fmt.Sprint(
				style.TrendLightGreenStyle.Render(fmt.Sprint(event.Collection.Counters.Sales)),
				collectionStyle.Render("/"),
				style.TrendLightRedStyle.Render(fmt.Sprint(event.Collection.Counters.Listings)),
			)
		} else {
			salesAndListings = fmt.Sprint(style.TrendLightGreenStyle.Render(fmt.Sprint(event.Collection.Counters.Sales)))
		}

		out.WriteString(" | " + salesAndListings)

		// coloring  moving average salira
		saLiRaStyle := style.TrendGreenStyle

		if previousMASaLiRa, currentMASaLiRa := event.Collection.CalculateSaLiRa(); currentMASaLiRa > 0 {
			if previousMASaLiRa > currentMASaLiRa {
				saLiRaStyle = style.TrendRedStyle
			}

			salira := fmt.Sprint(
				style.CreateTrendIndicator(previousMASaLiRa, currentMASaLiRa),
				saLiRaStyle.Render(fmt.Sprintf("%4.2f", currentMASaLiRa)),
				event.Collection.Style().Copy().Faint(true).Render("slr"),
				// style.DarkGrayStyle.Render("slr"),
			)
			out.WriteString(style.GrayStyle.Render(" ~ ") + salira)
		}
	}

	// mark the line if the seller or buyer is a wallet from the configured wallets
	if isOwnWallet {
		out.WriteString(" " + style.PinkBoldStyle.Render("*"))

		if event.EventType != collections.Listing {
			outputLine := "\n" + out.String() + "\n"
			out.Reset()
			out.WriteString(outputLine)
		}
	}

	// mark the line if the listing is below configured max. price
	if listingBelowPrice && event.EventType == collections.Listing {
		outputLine := "\n" + out.String() + "\n"
		out.Reset()
		out.WriteString(outputLine)
	}

	// print to terminal
	queueOutput <- out.String()

	//
	// telegram notification
	if isSaleOrMint && isWatchUsersWallet && viper.GetBool("notifications.telegram.enabled") {
		go func() {
			// did someone buy or sell something?
			var triggerAddress common.Address

			if gb.WatchUsers.Contains(event.To.Address) {
				triggerAddress = event.To.Address
			} else {
				triggerAddress = event.From.Address
			}

			// get the username of the wallet that triggered the notification
			var userName string

			if user := (*gb.WatchUsers)[triggerAddress]; user != nil {
				if user.TgUsername != "" {
					userName = "@" + user.TgUsername
				} else {
					userName = (*gb.WatchUsers)[triggerAddress].Name
				}
			} else {
				userName = "⸘Unknown‽"

				gbl.Log.Warnf("could not find user for address %s", triggerAddress.Hex())
				fmt.Printf("could not find user for address %s\n", triggerAddress.Hex())
			}

			// build the message to send
			msgTelegram := strings.Builder{}
			msgTelegram.WriteString(event.EventType.Icon())
			msgTelegram.WriteString(" " + strings.ReplaceAll(userName, "_", "\\_"))
			msgTelegram.WriteString(" " + event.EventType.ActionName())
			msgTelegram.WriteString(" " + internal.FormatTokenInfo(event.TokenID, event.Collection, false, false))
			msgTelegram.WriteString(" for **" + fmt.Sprintf("%.3f", priceEther) + "Ξ**")
			msgTelegram.WriteString("\n[Etherscan](" + etherscanURL + ")")
			msgTelegram.WriteString(" · [Opensea](" + openseaURL + ")")

			// try to get the token image url from its metadata
			var imageURI string

			if uri, err := gb.Nodes.GetRandomNode().GetTokenImageURI(event.Collection.ContractAddress, event.TokenID); err != nil {
				gbl.Log.Warnf("error getting token image (uri): %v", err)
			} else if strings.HasSuffix(uri, ".gif") {
				gbl.Log.Infof("token image uri is a gif -> not usable in tg msg: %s", uri)
			} else {
				imageURI = utils.ReplaceSchemeWithGateway(uri)
			}

			// send telegram message
			if msg, err := notifications.SendTelegramMessage(viper.GetInt64("telegram_chat_id"), msgTelegram.String(), imageURI); err != nil {
				gbl.Log.Warnf("failed to send telegram message | imageURI: '%s' | msgTelegram: '%s' | err: %s", imageURI, msgTelegram.String(), err)
			} else {
				rawMsg := msgTelegram.String()
				if msg.Text != "" {
					rawMsg = msg.Text
				}
				gbl.Log.Infof("sent telegram message | %s", strings.Replace(rawMsg, "\n", " | ", -1))
			}
		}()
	}
}

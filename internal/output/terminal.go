package output

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

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

		if collection == nil && event.ContractAddress != common.HexToAddress("0xbF230AEbBd288C69CcEA4ffe45d0F004261a895c") {
			name := ""

			if topic.Topic(event.Topic) == topic.TransferSingle && gb.Nodes != nil {
				if tokenName, err := gb.Nodes.GetERC1155TokenName(event.ContractAddress, event.TokenID); err == nil && tokenName != "" {
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
				gbl.Log.Warnf("🗑️ collection is nil | cw.CollectionDB.UserCollections[subLog.Address] -> %v | %v | TxHash: %v / %d", gb.CollectionDB.Collections[event.ContractAddress], event.ContractAddress.String(), event.TxHash, event.TxLogCount)
				return
			}
		}

		event.Collection = collection
	}

	var (
		priceStyle      lipgloss.Style
		priceArrowColor lipgloss.Color
	)

	// priceEtherPerItem, _ := nodes.WeiToEther(big.NewInt(int64(event.PriceWei.Uint64() / event.TxLogCount))).Float64()
	pricePerItem := big.NewInt(0).Div(event.PriceWei, big.NewInt(int64(event.TxLogCount)))

	priceEther, _ := nodes.WeiToEther(event.PriceWei).Float64()
	priceEtherPerItem, _ := nodes.WeiToEther(pricePerItem).Float64()

	//
	// conditions (review needed Oo)
	isMultiItemTx := event.TxLogCount > 1

	isMint := event.EventType == collections.Mint
	isMintOrTransfer := event.EventType == collections.Mint || event.EventType == collections.Transfer
	isMintOrSale := event.EventType == collections.Sale || event.EventType == collections.Purchase || event.EventType == collections.Mint

	isOwnCollection := event.Collection.Source == models.FromWallet || event.Collection.Source == models.FromConfiguration

	isOwnWallet := false
	if isMultiItemTx {
		isOwnWallet = gb.OwnWallets.ContainsOneOf(event.ToAddresses) != utils.ZeroAddress || gb.Watcher.ContainsOneOf(event.ToAddresses) != utils.ZeroAddress
	} else {
		isOwnWallet = gb.OwnWallets.Contains(event.To.Address) || gb.Watcher.Contains(event.To.Address)
	}

	isWatchUsersWallet := gb.Watcher.ContainsOneOf(event.FromAddresses) != utils.ZeroAddress || gb.Watcher.ContainsOneOf(event.ToAddresses) != utils.ZeroAddress

	// set type to purchase if "we" are on the buyer side
	if event.EventType == collections.Sale && (gb.OwnWallets.ContainsOneOf(event.ToAddresses) != utils.ZeroAddress || gb.Watcher.ContainsOneOf(event.ToAddresses) != utils.ZeroAddress) {
		event.EventType = collections.Purchase
	}

	var currentFloorPrice float64

	//
	// price-dependent styling
	if event.EventType == collections.Sale {
		// recalculate moving average
		event.Collection.PreviousFloorPrice, currentFloorPrice = event.Collection.CalculateFloorPrice(priceEtherPerItem)

		priceStyle = style.DarkWhiteStyle

		// get a color with saturation depending on the tx price
		priceArrowColor = style.GetPriceShadeColor(priceEther)
	} else {
		// if this is a mint/transfer/listing, we don't touch the moving average
		currentFloorPrice = (*event.Collection.FloorPrice).Value()
		event.Collection.PreviousFloorPrice = currentFloorPrice

		priceStyle = style.GrayStyle
		priceArrowColor = "#333333"
	}

	event.PriceArrowColor = priceArrowColor

	priceCurrencyStyle := event.Collection.Style().Copy().Faint(isMintOrTransfer)
	formattedCurrencySymbol := priceCurrencyStyle.Render("Ξ")
	divider := style.Sharrow.Copy().Foreground(priceArrowColor).Faint(true).String()

	currentFloorPriceStyle := style.GrayStyle.Copy().Faint(true)

	trendIndicator := style.CreateTrendIndicator(event.Collection.PreviousFloorPrice, currentFloorPrice)

	var numberStyle, pricePerItemStyle lipgloss.Style

	switch {
	case event.TxLogCount > 7:
		numberStyle = style.AlmostWhiteStyle
		pricePerItemStyle = style.DarkWhiteStyle
	case event.TxLogCount > 4:
		numberStyle = style.DarkWhiteStyle
		pricePerItemStyle = style.LightGrayStyle
	case event.TxLogCount > 1:
		numberStyle = style.LightGrayStyle
		pricePerItemStyle = style.GrayStyle
	default:
		numberStyle = style.GrayStyle
		pricePerItemStyle = style.DarkGrayStyle
	}

	// buyer styling
	event.ToColor = style.GenerateColorWithSeed(event.To.Address.Hash().Big().Int64())
	toStyle := lipgloss.NewStyle().Foreground(event.ToColor)
	to := style.ShortenAddressStyled(&event.To.Address, toStyle)

	//
	// ens
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

	// WEN...??
	now := time.Now()
	currentTime := now.Format("15:04:05")
	timeNow := style.GrayStyle.Copy().Faint(true).Render(currentTime)

	// WHAT...??
	var tokenInfo string
	if isMultiItemTx {
		tokenInfo = fmt.Sprintf("%s %s", numberStyle.Render(fmt.Sprintf("%dx", event.TxLogCount)), event.Collection.Style().Faint(isMint).Render(event.Collection.Name))
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
	} else {
		tokenInfo = style.FormatTokenInfo(event.TokenID, event.Collection.Name, event.Collection.Style(), event.Collection.StyleSecondary(), isMint, true)
	}

	// PRETTY...??
	collectionStyle := lipgloss.NewStyle().Foreground(event.Collection.Colors.Primary)

	if event.EventType == collections.Sale && isOwnCollection {
		timeNow = collectionStyle.Render(currentTime)

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
	}

	// // check if listing is below configured max. price
	// if isListingBelowPrice {
	// 	var timeStyle lipgloss.Style

	// 	if event.EventType == collections.Listing {
	// 		timeStyle = style.PinkBoldStyle
	// 		priceStyle = style.BoldStyle
	// 	} else {
	// 		timeStyle = lipgloss.NewStyle().Foreground(style.ShadesPink[3])
	// 	}

	// 	timeNow = timeStyle.Render(currentTime)
	// }

	switch {
	case event.EventType == collections.Sale && event.Collection.Highlight.Sales != "":
		timeNow = lipgloss.NewStyle().Foreground(event.Collection.Highlight.Sales).Render(currentTime)
	case event.EventType == collections.Listing && event.Collection.Highlight.Listings != "":
		timeNow = lipgloss.NewStyle().Foreground(event.Collection.Highlight.Listings).Render(currentTime)
	case event.Collection.Highlight.Color != "":
		timeNow = lipgloss.NewStyle().Foreground(event.Collection.Highlight.Color).Render(currentTime)
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

	// blur.io link
	var blurURL string

	// if slug, err := cache.GetBlurSlug(event.Collection.ContractAddress); err == nil && slug != "" {
	// 	gbl.Log.Infof("cache | cached blur slug: %s", slug)
	// 	blurURL = fmt.Sprintf("https://blur.io/collection/%s", strings.Replace(strings.ToLower(slug), " ", "", -1))
	// } else if event.Collection.OpenseaSlug != "" {
	// 	gbl.Log.Infof("cache | no cached blur slug for %s | trying to resolve...", event.Collection.ContractAddress)
	// 	blurURL = fmt.Sprintf("https://blur.io/collection/%s", strings.Replace(strings.ToLower(event.Collection.OpenseaSlug), " ", "", -1))
	// } else {
	// 	gbl.Log.Infof("cache | no cached blur slug for %s | queuing to fetch...", event.Collection.ContractAddress)
	// 	gb.QueueSlugs <- event.Collection.ContractAddress
	// }

	etherscanURL := fmt.Sprintf("https://etherscan.io/tx/%s", event.TxHash)

	marker := " "

	// if isListingBelowPrice {
	// 	marker = style.PinkBoldStyle.Render("*")
	// } else if isOwnCollection && event.EventType == collections.Sale {
	// 	if priceEtherPerItem >= viper.GetFloat64("show.min_value") {
	// 		if isOwnWallet {
	// 			marker = style.OwnerGreenBoldStyle.Render("*")
	// 		}
	// 	}
	// }

	// add to event history
	isSaleOrPurchase := event.EventType == collections.Sale || event.EventType == collections.Purchase

	if event.Discarded == nil {
		if isOwnWallet || (isOwnCollection && isSaleOrPurchase) {
			ticker.StatsTicker.EventHistory = append(ticker.StatsTicker.EventHistory, event)
		} else if gb.OwnWallets.Contains(event.To.Address) {
			ticker.StatsTicker.EventHistory = append(ticker.StatsTicker.EventHistory, event)
		}
	}

	// build the line to be displayed
	out := strings.Builder{}

	if viper.GetBool("log.verbose") {
		if event.EventType == collections.Listing {
			out.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#20293d")).Render("OS"))
		} else if gb.Nodes != nil && len(*gb.Nodes) > 0 {
			out.WriteString(gb.Nodes.GetNodeByID(event.NodeID).GetStyledMarker())
		}

		out.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#111111")).Render("|"))
	}

	// time & type
	out.WriteString(timeNow)
	out.WriteString(marker + event.EventType.Icon())
	out.WriteString(" " + divider)

	// price
	if event.EventType == collections.Listing {
		out.WriteString(" " + priceStyle.Render(style.TerminalLink(openseaURL, fmt.Sprintf("%6.3f", priceEther))))
	} else {
		out.WriteString(" " + priceStyle.Render(fmt.Sprintf("%6.3f", priceEther)))
	}

	out.WriteString(formattedCurrencySymbol)

	// price per item
	out.WriteString(" " + pricePerItemStyle.Render(fmt.Sprintf("%6.3f", priceEtherPerItem)))
	out.WriteString(priceCurrencyStyle.Copy().Faint(true).Render("Ξ"))

	// (artificial) floor price
	out.WriteString("  " + trendIndicator)
	out.WriteString(currentFloorPriceStyle.Render(fmt.Sprintf("%6.3f", currentFloorPrice)))

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

		out.WriteString(style.DarkGrayStyle.Render(" /") + totalSupplyStyle.Render(outputTotalSupply))
	}

	// listing price in relation to the collection floor
	if (event.EventType == collections.Listing || event.EventType == collections.Purchase || isOwnCollection) && currentFloorPrice > 0 {
		if fpRatio := (priceEtherPerItem / currentFloorPrice) * 100; fpRatio > 0 {
			var fpStyle lipgloss.Style

			fpRatioDifference := int(fpRatio - 100)

			if fpRatioDifference > 0 {
				fpStyle = style.TrendRedStyle.Copy().Faint(true)
			} else if fpRatioDifference < 0 {
				fpStyle = style.TrendGreenStyle.Copy()
			} else {
				fpStyle = style.GrayStyle.Copy()
			}

			// out.WriteString(" " + style.PinkBoldStyle.Render("·") + " ")
			// out.WriteString(" " + style.DarkGrayStyle.Render("·") + " ")
			out.WriteString("  ")
			out.WriteString(fpStyle.Bold(false).Render(fmt.Sprintf("%+d%%", fpRatioDifference)))
		}
	}

	// marker for collections which are contained in the buy rules
	if gb.BuyRules.Rules[event.Collection.ContractAddress] != nil {
		out.WriteString(" " + style.PinkBoldStyle.Render("·"))
	}

	// link opensea
	out.WriteString(" | " + style.GrayBoldStyle.Copy().Foreground(style.OpenseaToneBlue).Render(style.TerminalLink(openseaURL, "OpenSea")))

	// link blur
	if blurURL != "" {
		out.WriteString(" | " + style.GrayBoldStyle.Copy().Foreground(style.BlurOrange).Faint(true).Render(style.TerminalLink(blurURL, "blur")))
	}

	// link etherscan
	if event.EventType != collections.Listing {
		out.WriteString(" | " + style.GrayStyle.Render(style.TerminalLink(etherscanURL, "ES")))
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
		event.Collection.AddSale(event.PriceWei, event.TxLogCount)
	} else if event.EventType == collections.Mint {
		event.Collection.AddMint()
	}

	// sales/listings count & salira
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

		if previousMASaLiRa, currentMASaLiRa := event.Collection.CalculateSaLiRa(event.Collection.ContractAddress); currentMASaLiRa > 0 {
			// coloring moving average salira
			saLiRaStyle := style.TrendGreenStyle

			if previousMASaLiRa > currentMASaLiRa {
				saLiRaStyle = style.TrendRedStyle
			}

			salira := fmt.Sprint(
				style.CreateTrendIndicator(previousMASaLiRa, currentMASaLiRa),
				saLiRaStyle.Render(fmt.Sprintf("%4.2f", currentMASaLiRa)),
				event.Collection.Render("slr"),
			)

			out.WriteString(style.GrayStyle.Render(" ~ ") + salira)
		} else if cachedSalira, err := cache.GetSalira(event.Collection.ContractAddress); cachedSalira > 0 && err == nil {
			salira := fmt.Sprint(
				style.GrayStyle.Render(" ~ "),
				style.GrayStyle.Render(fmt.Sprintf("%4.2f", cachedSalira)),
				event.Collection.Render("slr*"),
			)

			out.WriteString(salira)
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

	// // mark the line if the listing is below configured max. price
	// if isListingBelowPrice && event.EventType == collections.Listing {
	// 	outputLine := "\n" + out.String() + "\n"
	// 	out.Reset()
	// 	out.WriteString(outputLine)
	// }

	// print to terminal
	if event.Discarded == nil || event.Discarded.PrintInStream {
		queueOutput <- out.String()
	}

	// set price of listing as the new fp
	if event.EventType == collections.Listing {
		if currentFloorPrice == 0.0 || priceEther < currentFloorPrice {
			(*event.Collection.FloorPrice).Set(priceEther)
			cache.StoreFloor(event.Collection.ContractAddress, priceEther)
		}
	}

	go cache.StoreFloor(event.Collection.ContractAddress, currentFloorPrice)

	//
	// telegram notification
	if isMintOrSale && isWatchUsersWallet && viper.GetBool("notifications.telegram.enabled") {
		// try to acquire the lock
		notificationLock, err := cache.NotificationLock(event.TxHash)

		if !notificationLock || err != nil {
			gbl.Log.Debugf("notification lock for %s already exists", event.TxHash)
			return
		}

		gbl.Log.Infof("notification lock for %s acquired, trying to send...", event.TxHash)

		go func() {
			// did someone buy or sell something?
			var triggerAddress common.Address

			if trigger := gb.Watcher.ContainsOneOf(event.ToAddresses); trigger != utils.ZeroAddress {
				triggerAddress = trigger
			} else if trigger := gb.Watcher.ContainsOneOf(event.FromAddresses); trigger != utils.ZeroAddress {
				triggerAddress = trigger
			} else {
				return
			}

			// get correct erc721 sale
			for _, transfer := range event.ERC721Transfers {
				watchUser := ((*gb.Watcher).WatchUsers)[transfer.To]
				if watchUser == nil {
					watchUser = ((*gb.Watcher).WatchUsers)[transfer.From]
				}

				if watchUser == nil {
					continue
				}
				sendTelegramNotification(gb, triggerAddress, watchUser, event, transfer.TokenId, priceEtherPerItem, etherscanURL, openseaURL)
			}
			for _, transfer := range event.ERC1155Transfers {
				watchUser := ((*gb.Watcher).WatchUsers)[transfer.To]
				if watchUser == nil {
					watchUser = ((*gb.Watcher).WatchUsers)[transfer.From]
				}

				if watchUser == nil {
					continue
				}
				sendTelegramNotification(gb, triggerAddress, watchUser, event, transfer.Id, priceEtherPerItem, etherscanURL, openseaURL)
			}
		}()
	}
}

func sendTelegramNotification(gb *gloomberg.Gloomberg, triggerAddress common.Address, watchUser *models.WatchUser, event *collections.Event, tokenId *big.Int, priceEtherPerItem float64, etherscanURL string, openseaURL string) {
	// get the username of the wallet that triggered the notification
	var userName string

	user := ((*gb.Watcher).UserAddresses)[triggerAddress]
	// watchUser := ((*gb.Watcher).WatchUsers)[triggerAddress]

	if watchUser != nil {
		if watchUser.TelegramUsername != "" {
			userName = "@" + watchUser.TelegramUsername
		} else {
			userName = watchUser.Name
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
	msgTelegram.WriteString(" (" + style.ShortenAddress(&triggerAddress) + ")")
	msgTelegram.WriteString(" " + event.EventType.ActionName())
	msgTelegram.WriteString(" " + style.FormatTokenInfo(tokenId, event.Collection.Name, event.Collection.Style(), event.Collection.StyleSecondary(), false, false))
	msgTelegram.WriteString(" for **" + fmt.Sprintf("%.3f", priceEtherPerItem) + "Ξ**")
	msgTelegram.WriteString("\n[Etherscan](" + etherscanURL + ")")
	msgTelegram.WriteString(" · [Opensea](" + openseaURL + ")")

	// try to get the token image url from its metadata
	var imageURI string

	if uri, err := gb.Nodes.GetTokenImageURI(event.Collection.ContractAddress, tokenId); err != nil {
		gbl.Log.Warnf("error getting token image (uri): %v", err)
	} else if strings.HasSuffix(uri, ".gif") {
		gbl.Log.Infof("token image uri is a gif -> not usable in tg msg: %s", uri)
	} else {
		imageURI = utils.ReplaceSchemeWithGateway(uri)
	}

	// send telegram message
	if msg, err := notifications.SendTelegramMessage(user.TelegramChatID, msgTelegram.String(), imageURI); err != nil {
		gbl.Log.Warnf("failed to send telegram message | imageURI: '%s' | msgTelegram: '%s' | err: %s", imageURI, msgTelegram.String(), err)
	} else {
		rawMsg := msgTelegram.String()
		if msg.Text != "" {
			rawMsg = msg.Text
		}
		gbl.Log.Infof("sent telegram message | %s", strings.Replace(rawMsg, "\n", " | ", -1))
	}
}

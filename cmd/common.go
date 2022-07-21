package cmd

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/awesome-gocui/gocui"
	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/cache"
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/gbnode"
	"github.com/benleb/gloomberg/internal/glicker"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/notifications"
	"github.com/benleb/gloomberg/internal/opensea"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/subscriptions"
	"github.com/benleb/gloomberg/internal/wwatcher"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/viper"
)

var (
	Version   string
	Commit    string
	BuildDate string
	BuiltBy   string

	queueEvents   = make(chan *collections.Event, 1024)
	queueLogs     = make(chan types.Log, 1024)
	queueListings = make(chan *models.ItemListedEvent, 1024)
	queueOutput   = make(chan string, 1024)
)

func getNodes() *gbnode.NodeCollection {
	nodesSpinner := style.GetSpinner("setting up node connections...")
	_ = nodesSpinner.Start()

	nodes := gbnode.GetNodesFromConfig(viper.GetStringSlice("endpoints"))
	numNodes := len(nodes.GetNodes())

	// stop spinner
	nodesSpinner.StopMessage(fmt.Sprint(style.BoldStyle.Render(fmt.Sprint(numNodes)), " nodes connected\n"))
	_ = nodesSpinner.Stop()

	if numNodes == 0 {
		gbl.Log.Fatal("No node providers found")
	}

	return nodes
}

func getWallets(nodes *gbnode.NodeCollection) *models.Wallets {
	// set up spinner
	walletSpinner := style.GetSpinner("setting up wallets...")
	_ = walletSpinner.Start()

	wallets := wwatcher.GetWalletsFromConfig(viper.GetStringSlice("wallets"), nodes)

	// stop spinner
	if len(*wallets) > 0 {
		walletSpinner.StopMessage(fmt.Sprint(style.BoldStyle.Render(fmt.Sprint(len(*wallets))), " wallets: ", strings.Join(wallets.FormattedNames(), ", ")) + "\n")
		_ = walletSpinner.Stop()
	} else {
		_ = walletSpinner.StopFail()
	}

	return wallets
}

func formatEvent(ctx context.Context, g *gocui.Gui, event *collections.Event, nodes *gbnode.NodeCollection, wallets *models.Wallets, outputLines *chan string) {
	namesCache := &wwatcher.NamesCache{
		Names: make(map[common.Address]string),
		RWMu:  &sync.RWMutex{},
	}

	gbl.Log.Debugf("FormatEvent | event: %+v", event)

	isMint := event.EventType == collections.Mint
	isMintOrTransfer := event.EventType == collections.Mint || event.EventType == collections.Transfer
	isSaleOrMint := event.EventType == collections.Sale || event.EventType == collections.Purchase || event.EventType == collections.Mint
	isMultiItemTx := event.TxItemCount > 1

	var (
		priceStyle      lipgloss.Style
		priceArrowColor lipgloss.Color
	)

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
	priceEther := subscriptions.WeiToEther(event.PriceWei)
	priceEtherPerItem := subscriptions.WeiToEther(pricePerItem)

	var previousMovingAverage, currentMovingAverage float64

	if event.EventType == collections.Sale {
		if wallets.Contains(event.To.Address) {
			event.EventType = collections.Purchase
		}

		// recalculate moving average
		itemPrice, _ := subscriptions.WeiToEther(pricePerItem).Float64()
		previousMovingAverage, currentMovingAverage = event.Collection.CalculateArtificialFloor(itemPrice)

		// get a color with saturation depending on the tx price
		priceStyle = style.DarkWhiteStyle
		priceArrowColor = style.GetPriceShadeColor(priceEther)
	} else {
		// if this is a mint/transfer/listing, we don't touch the moving average
		currentMovingAverage = event.Collection.ArtificialFloor.Value()
		previousMovingAverage = currentMovingAverage

		priceStyle = style.GrayStyle
		priceArrowColor = "#333333"
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

	isOwnCollection := event.Collection.Source == collections.Wallet || event.Collection.Source == collections.Configuration
	ownWalletInvolved := wallets.Contains(event.From.Address) || wallets.Contains(event.To.Address)
	wwatcherWalletInvolved := wwatcher.Recipients.Contains(event.From.Address) || wwatcher.Recipients.Contains(event.To.Address)
	listingBelowPrice := event.Collection.Highlight.ListingsBelowPrice > 0.0 && big.NewFloat(event.Collection.Highlight.ListingsBelowPrice).Cmp(priceEther) > 0

	// buyer
	toStyle := lipgloss.NewStyle().Foreground(style.GenerateColorWithSeed(event.To.Address.Hash().Big().Int64()))
	to := style.ShortenAddressStyled(&event.To.Address, toStyle)

	var ensName string

	if viper.GetBool("redis.enabled") {
		if rdb := cache.GetRedisClient(); rdb != nil {
			if cachedName, err := rdb.Get(ctx, cache.KeyENS(event.To.Address)).Result(); err == nil && cachedName != "" {
				ensName = cachedName

				gbl.Log.Infof("ensName from redis: %s", ensName)
				to = toStyle.Render(ensName)
			}
		}
	}

	if ensName == "" && viper.IsSet("api_keys.etherscan") {
		if ensName := wwatcher.GetENSForAddress(nodes, event.To.Address, namesCache); ensName != "" {
			gbl.Log.Debugf("ensName from etherscan: %s", ensName)
			to = toStyle.Render(ensName)
			event.ToENS = ensName
		} else if event.ToENS != "" {
			to = toStyle.Render(event.ToENS)
		}
	}

	// WEN...??
	currentTime := time.Now().Format("15:04:05")
	timeNow := style.GrayStyle.Copy().Faint(true).Render(currentTime)

	// WHAT...??
	var tokenInfo string
	if isMultiItemTx {
		tokenInfo = fmt.Sprintf("%s %s", numberStyle.Render(fmt.Sprintf("%dx", event.TxItemCount)), event.Collection.Style().Faint(isMint).Render(event.Collection.Name))
	} else {
		tokenInfo = internal.FormatTokenInfo(event.TokenID, event.Collection, isMint, true)
	}

	// PRETTY...??
	collectionStyle := lipgloss.NewStyle().Foreground(event.Collection.Colors.Primary)

	if event.EventType == collections.Sale && isOwnCollection {
		timeNow = collectionStyle.Render(currentTime)

		notifications.SendNotification(event.Collection.Name, tokenInfo)
		gbl.Log.Debugf("SendNotification | collection: %s, tokenInfo: %s", event.Collection.Name, tokenInfo)
	}

	// highlight line if the seller or buyer is a wallet from the configured wallets
	if ownWalletInvolved {
		timeNow = lipgloss.NewStyle().Foreground(style.OwnerGreen).Bold(true).Render(currentTime)
	}

	// check if listing is below configured max. price
	if listingBelowPrice {
		var timeStyle lipgloss.Style

		if event.EventType == collections.Listing {
			timeStyle = style.PinkBoldStyle
			priceStyle = style.BoldStyle
		} else {
			timeStyle = lipgloss.NewStyle().Foreground(style.ShadesPink[3])
		}

		timeNow = timeStyle.Render(currentTime)
	}

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

	etherscanURL := fmt.Sprintf("https://etherscan.io/tx/%s", event.TxHash)

	marker := " "
	if listingBelowPrice {
		marker = style.PinkBoldStyle.Render("*")
	} else if isOwnCollection && event.EventType == collections.Sale {
		if itemPrice, _ := priceEtherPerItem.Float64(); itemPrice >= viper.GetFloat64("show.min_price") {
			if ownWalletInvolved {
				marker = style.OwnerGreenBoldStyle.Render("*")
			}
		}
	}

	// add to event history
	if isOwnCollection && event.EventType == collections.Sale {
		glicker.StatsTicker.EventHistory = append(glicker.StatsTicker.EventHistory, event)
	}

	// build the line to be displayed
	out := strings.Builder{}
	out.WriteString(timeNow)
	out.WriteString(marker + event.EventType.Icon())
	out.WriteString(" " + divider)

	// price
	if event.EventType == collections.Listing {
		out.WriteString(" " + priceStyle.Render(TerminalLink(openseaURL, fmt.Sprintf("%6.3f", priceEther))))
	} else {
		out.WriteString(" " + priceStyle.Render(fmt.Sprintf("%6.3f", priceEther)))
	}

	out.WriteString(formattedCurrencySymbol)

	// moving average (artificial) floor price
	out.WriteString("  " + trendIndicator)
	out.WriteString(currentMovingAverageStyle.Render(fmt.Sprintf("%6.3f", currentMovingAverage)))

	// price per item
	out.WriteString(" " + pricePerItemStyle.Render(fmt.Sprintf("%6.3f", subscriptions.WeiToEther(pricePerItem))))
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

		out.WriteString(style.DarkGrayStyle.Render(" /") + totalSupplyStyle.Render(outputTotalSupply))
	}

	// link opensea
	out.WriteString(" | " + style.GrayBoldStyle.Copy().Foreground(style.OpenseaToneBlue).Render(TerminalLink(openseaURL, "OpenSea")))

	// link etherscan
	if event.EventType != collections.Listing {
		out.WriteString(" | " + style.GrayStyle.Render(TerminalLink(etherscanURL, "Etherscan")))
	}

	// buyer
	out.WriteString(" | " + arrow.String())
	out.WriteString(" " + to)

	// maybe importan wallet indicator
	if wwatcher.MIWs.Contains(event.To.Address) {
		out.WriteString("   " + style.PinkBoldStyle.Render("👀 MIW! 👀"))
	}

	// log topic (for debugging)
	if viper.GetBool("log.debug") {
		out.WriteString(" | " + gbnode.Topic(event.Topic).String())
	}

	// automatically fetch listings for collections with more than opensea.auto_list_min_sales sales
	if event.Collection.Counters.Sales == viper.GetUint64("opensea.auto_list_min_sales") {
		slug := opensea.GetCollectionSlug(event.Collection.ContractAddress)
		opensea.SubscribeToCollectionSlug(nil, slug, nil)
		event.Collection.ResetStats()

		gbl.Log.Infof("auto-subscribed to %s after %d sales", event.Collection.Name, event.Collection.Counters.Sales)

		*outputLines <- fmt.Sprintf(
			" %s auto-subscribed to %s after %d sales",
			style.PinkBoldStyle.Render(">"),
			event.Collection.Name,
			event.Collection.Counters.Sales,
		)
	}

	// counting
	if event.EventType == collections.Sale || event.EventType == collections.Purchase {
		go event.Collection.AddSale(event.PriceWei, uint64(event.TxItemCount))
	} else if event.EventType == collections.Mint {
		go event.Collection.AddMint()
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

		// coloring  moving average salira
		saLiRaStyle := style.TrendGreenStyle

		if previousMASaLiRa, currentMASaLiRa := event.Collection.CalculateSaLiRa(); currentMASaLiRa > 0 {
			if previousMASaLiRa > currentMASaLiRa {
				saLiRaStyle = style.TrendRedStyle
			}

			salira := fmt.Sprint(
				style.CreateTrendIndicator(previousMASaLiRa, currentMASaLiRa),
				saLiRaStyle.Render(fmt.Sprintf("%5.3f", currentMASaLiRa)),
				style.DarkGrayStyle.Render("slr"),
			)
			out.WriteString(style.GrayStyle.Render(" ~ ") + salira)
		}
	}

	// mark the line if the seller or buyer is a wallet from the configured wallets
	if ownWalletInvolved {
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

	if g != nil {
		// print to tui
		gbl.Log.Debugf("updating gui now... %p", g)

		g.UpdateAsync(func(g *gocui.Gui) error {
			streamView, err := g.View("main")
			if err != nil {
				gbl.Log.Errorf("error getting streamView: %+v", err.Error())

				return err
			}

			gbl.Log.Warnf("streamView: %p", streamView)

			if _, err = streamView.Write([]byte(out.String() + "\n")); err != nil {
				gbl.Log.Errorf("error writing streamView: %+v", err.Error())

				return err
			}

			return nil
		})

		gbl.Log.Debugf("...done %p\n", g)
	} else {
		// print to terminal
		*outputLines <- out.String()
	}

	// telegram notification
	gbl.Log.Debug("telegram notify: %v | address match: %v | tgBot: %v", viper.GetBool("telegram.enabled"), wwatcherWalletInvolved, notifications.TgBot != nil)

	if isSaleOrMint && wwatcherWalletInvolved && viper.GetBool("telegram.enabled") && notifications.TgBot != nil {
		gbl.Log.Warn("sending telegram notification...")

		go func() {
			gbl.Log.Warnf("tg wwatcher address found: %s | %s\n", event.From.Address, event.To.Address)

			var triggerAddress common.Address

			// linkURL := etherscanURL

			if wwatcher.Recipients.Contains(event.To.Address) {
				triggerAddress = event.To.Address

				if event.EventType == collections.Sale {
					event.EventType = collections.Purchase
				} else if event.EventType == collections.Listing {
					// currently not reachable as Listing events are filtered before here
					// linkURL = openseaURL
				}
			} else {
				triggerAddress = event.From.Address
			}

			gbl.Log.Warnf("tg wwatcher address: %s\n", triggerAddress.Hex())

			var userName string
			if name := wwatcher.Recipients[triggerAddress].TgUsername; name != "" {
				userName = "@" + name
			} else {
				userName = wwatcher.Recipients[triggerAddress].Name
			}

			msgTelegram := strings.Builder{}
			msgTelegram.WriteString(event.EventType.Icon())
			msgTelegram.WriteString(" " + strings.ReplaceAll(userName, "_", "\\_"))
			msgTelegram.WriteString(" " + event.EventType.ActionName())
			msgTelegram.WriteString(" " + strings.ReplaceAll(internal.FormatTokenInfo(event.TokenID, event.Collection, isMint, false), "#", "\\#"))
			msgTelegram.WriteString(" for **" + fmt.Sprintf("%.3f", priceEther) + "Ξ**")
			// msgTelegram.WriteString("\n" + linkURL)
			msgTelegram.WriteString("\n[Etherscan](" + etherscanURL + ")")
			msgTelegram.WriteString(" · [Opensea](" + openseaURL + ")")

			chatID := viper.GetInt64("telegram_chat_id")

			var imageURI string

			if uri, err := nodes.GetRandomNode().GetTokenImageURI(event.Collection.ContractAddress, event.TokenID); err != nil || strings.HasSuffix(uri, ".gif") {
				gbl.Log.Errorf("error getting token image uri: %v", err)
			} else {
				imageURI = strings.Replace(uri, "ipfs://", "https://ipfs.io/ipfs/", 1)
			}

			if msg, err := notifications.SendTelegramMessage(chatID, msgTelegram.String(), imageURI); err != nil {
				gbl.Log.Errorf("failed to send telegram message: %s | imageURI: %s | err: %s\n", msg, imageURI, err)
			} else {
				gbl.Log.Warnf("sent telegram message: %s\n", msg.Text)
				gbl.Log.Infof("sent telegram message: %s\n", msg)
			}
		}()
	}
}

// TerminalLink formats a link for the terminal using ANSI codes.
func TerminalLink(params ...string) string {
	var text string

	url := params[0]

	if len(params) >= 2 {
		text = params[1]
	} else {
		text = url
	}

	return fmt.Sprintf("\033]8;;%s\033\\%s\033]8;;\033\\", url, text)
}

package trapri

import (
	"fmt"
	"strings"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/benleb/gloomberg/internal/pusu"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/kr/pretty"
	"github.com/spf13/viper"
)

// SeaWatcherEventsHandler handles all incoming & decoded events from OpenSea and dispatches them to the appropriate handler.
func SeaWatcherEventsHandler(gb *gloomberg.Gloomberg) {
	chanItemListed := gb.SubscribeItemListed()
	chanItemReceivedBid := gb.SubscribeItemReceivedBid()
	chanCollectionOffer := gb.SubscribeCollectionOffer()
	chanMetadataUpdated := gb.SubscribeItemMetadataUpdated()

	for i := 0; i < viper.GetInt("trapri.numOpenSeaEventhandlers"); i++ {
		go func(i int) {
			log.Debugf("  👨‍🔧 OpenseaEventsHandler %d started", i)

			// gloomberg.Prf("⚓️ OpenseaEventsHandler %d started", i)

			for {
				select {
				case event := <-chanItemListed:
					gbl.Log.Debugf("  📢 item listed: %+v", pretty.Sprintf("%#v", event))

					// gloomberg.Prf("⚓️ handling %s event...", event.EventType)
					go HandleItemListed(gb, event)

					if viper.GetBool("pubsub.server.enabled") || viper.GetBool("grpc.server.enabled") {
						// publish via pubsub
						publishChannel := internal.PubSubSeaWatcherListings + "/" + event.Payload.Item.NftID.ContractAddress().Hex() + "/" + event.Payload.Item.NftID.TokenID().String()
						pusu.Publish(gb, publishChannel, event)
						gloomberg.PrModf("seawa", "⚓️ published %s event: %+v", event.EventType, event.Payload.Item.Name)

						// output to terminal
						collectionPrimaryStyle := lipgloss.NewStyle().Foreground(style.GenerateColorWithSeed(event.Payload.Item.NftID.ContractAddress().Hash().Big().Int64()))
						price := price.NewPrice(event.Payload.BasePrice)
						fmtCurrencySymbol := collectionPrimaryStyle.Bold(false).Render("Ξ")
						fmtPrice := style.BoldAlmostWhite(fmt.Sprintf("%5.2f", price.Ether())) + fmtCurrencySymbol
						fmtItemName := collectionPrimaryStyle.Bold(true).Render(event.Payload.Item.Name)
						fmtItemLink := style.TerminalLink(event.Payload.Item.Permalink, fmtItemName)
						eventType := degendb.EventType(degendb.GetEventType(event.EventType))

						gloomberg.Prf("%s %s %s", eventType.Icon(), fmtPrice, fmtItemLink)
					}

				case event := <-chanItemReceivedBid:
					gbl.Log.Debugf("  💦 item received bid: %+v", pretty.Sprintf("%#v", event))

					// log.Print("  🎭 💦 item received bid")
					// pretty.Println(event)

					go HandleItemReceivedBid(gb, event)

				case event := <-chanCollectionOffer:
					gbl.Log.Debugf("  🦕 collection offer: %+v", pretty.Sprintf("%#v", event))

					// go HandleCollectionOffer(gb, event)

				case event := <-chanMetadataUpdated:
					// filter lawless cloaknet transponders due to spam
					if event.Payload.ContractAddress() == common.HexToAddress("0xd3a0b315023243632a15fd623d6f33314193df4e") {
						continue
					}

					if len(event.Payload.Traits) > 0 {
						fmtTraits := make([]string, 0)
						for _, trait := range event.Payload.Traits {
							fmtTraits = append(fmtTraits, trait.StringBold())
						}

						log.Printf("  🎭 | %s #%s", event.Payload.Name, event.Payload.NftID.TokenID().String())
						log.Printf("  🎭 → %v", strings.Join(fmtTraits, style.DarkGrayStyle.Render(" | ")))

						gbl.Log.Info("  🎭 metadata updated:")
						gbl.Log.Info(pretty.Sprint(event))
					}

					// go HandleMetadataUpdated(gb, event)
				}
			}
		}(i)
	}
}

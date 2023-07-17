package trapri

import (
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

// OpenseaEventsHandler handles all incoming & decoded events from OpenSea and dispatches them to the appropriate handler.
func OpenseaEventsHandler(gb *gloomberg.Gloomberg) {
	chanItemListed := gb.SubscribeItemListed()
	chanItemReceivedBid := gb.SubscribeItemReceivedBid()
	chanCollectionOffer := gb.SubscribeCollectionOffer()
	chanMetadataUpdated := gb.SubscribeItemMetadataUpdated()

	for i := 0; i < viper.GetInt("trapri.numOpenSeaEventhandlers"); i++ {
		go func() {
			for {
				select {
				case event := <-chanItemListed:
					gbl.Log.Debugf("%s event received at trapri.eventWorker: %#v", event.Event, event)

					go HandleItemListed(gb, event)

				case event := <-chanItemReceivedBid:
					gbl.Log.Debugf("%s event received at trapri.eventWorker: %#v", event.Event, event)

					go HandleItemReceivedBid(gb, event)

				case event := <-chanCollectionOffer:
					gbl.Log.Debugf("%s event received at trapri.eventWorker: %#v", event.Event, event)

					// go HandleCollectionOffer(gb, event)

				case event := <-chanMetadataUpdated:
					gbl.Log.Debugf("%s event received at trapri.eventWorker: %#v", event.Event, event)

					if event.Payload.Item.NftID.ContractAddress() == common.HexToAddress("0xd3a0b315023243632a15fd623d6f33314193df4e") {
						continue
					}

					log.Printf("  🦄 metadata updated: %#v", event)
					// go HandleMetadataUpdated(gb, event)
				}
			}
		}()
	}
}

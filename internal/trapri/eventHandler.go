package trapri

import (
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/charmbracelet/log"
)

// OpenseaEventsHandler handles all incoming & decoded events from OpenSea and dispatches them to the appropriate handler.
func OpenseaEventsHandler(gb *gloomberg.Gloomberg) {
	chanItemListed := gb.SubscribeItemListed()
	chanItemReceivedBid := gb.SubscribeItemReceivedBid()
	chanCollectionOffer := gb.SubscribeCollectionOffer()

	for {
		select {
		case event := <-chanItemListed:
			log.Debugf("%s event received at trapri.eventWorker: %#v", event.Event, event)

			go HandleItemListed(gb, event)

		case event := <-chanItemReceivedBid:
			log.Debugf("%s event received at trapri.eventWorker: %#v", event.Event, event)

			go HandleItemReceivedBid(gb, event)

		case event := <-chanCollectionOffer:
			log.Debugf("%s event received at trapri.eventWorker: %#v", event.Event, event)

			go HandleCollectionOffer(gb, event)
		}
	}
}

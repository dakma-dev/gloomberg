package pusu

import (
	"context"
	"encoding/json"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/osmodels"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/benleb/gloomberg/internal/trapri"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/redis/rueidis"
)

func SubscribeToSales(gb *gloomberg.Gloomberg, channel string, queueTokenTransactions chan *totra.TokenTransaction) {
	err := gb.Rdb.Receive(context.Background(), gb.Rdb.B().Subscribe().Channel(channel).Build(), func(msg rueidis.PubSubMessage) {
		// validate json
		if !json.Valid([]byte(msg.Message)) {
			gbl.Log.Warnf("❗️ invalid json: %s", msg.Message)

			return
		}

		// create the event transaction
		var ttx totra.TokenTransaction

		// unmarshal event transaction from json
		err := json.Unmarshal([]byte(msg.Message), &ttx)
		if err != nil {
			gbl.Log.Warnf("❗️ error unmarshalling event Tx: %+v | %s", msg.Message, err)

			return
		}

		queueTokenTransactions <- &ttx
	})
	if err != nil {
		gbl.Log.Errorf("❌ error subscribing to redis channel %s: %s", channel, err.Error())

		return
	}
}

// SubscribeToListings subscribes to all collections for which we have a slug.
func SubscribeToListings(gb *gloomberg.Gloomberg, queueTokenTransactions chan *totra.TokenTransaction) {
	slugAddresses := gb.CollectionDB.OpenseaSlugAddresses()
	if len(slugAddresses) == 0 {
		gbl.Log.Warn("❌ no slugs to send to gloomberg server")

		return
	}

	// create a list of channels to subscribe to
	channels := make([]string, 0)

	for _, collectionAddress := range slugAddresses {
		channelPattern := internal.TopicSeaWatcher + "/" + collectionAddress.String() + "/*"

		channels = append(channels, channelPattern)
	}

	err := gb.Rdb.Receive(context.Background(), gb.Rdb.B().Psubscribe().Pattern(channels...).Build(), func(msg rueidis.PubSubMessage) {
		gbl.Log.Debugf("🚇 received msg on channel %s: %s", msg.Channel, msg.Message)

		var itemListedEvent osmodels.ItemListedEvent

		// validate json
		if !json.Valid([]byte(msg.Message)) {
			gbl.Log.Warnf("❗️ invalid json: %s", msg.Message)

			return
		}

		// unmarshal
		if err := json.Unmarshal([]byte(msg.Message), &itemListedEvent); err != nil {
			gbl.Log.Errorf("❌ error json.Unmarshal: %+v\n", err.Error())

			return
		}

		// nftID is a string in the format <chain>/<contract>/<tokenID>
		nftID := itemListedEvent.GetNftID()

		log.Printf("itemListedEvent: %+v", itemListedEvent)
		log.Printf("nftID: %+v", nftID)

		//
		// discard listings for ignored collections
		if collection, ok := gb.CollectionDB.Collections[common.HexToAddress(nftID[1])]; ok && collection.IgnorePrinting {
			gbl.Log.Debugf("🗑️ ignoring printing for collection %s", collection.Name)

			return
		}

		// print
		trapri.FormatListing(gb, &itemListedEvent, queueTokenTransactions)
	})
	if err != nil {
		gbl.Log.Errorf("❌ error subscribing to redis channels %s: %s", channels, err.Error())

		return
	}
}

func Publish(gb *gloomberg.Gloomberg, channel string, event any) {
	// marshal event to json
	marshalledEvent, err := json.Marshal(event)
	if err != nil {
		gbl.Log.Warnf("❗️ error marshalling event: %s", err)

		return
	}

	// publish event to redis pubsub channel
	if gb.Rdb.Do(context.Background(), gb.Rdb.B().Publish().Channel(channel).Message(string(marshalledEvent)).Build()).Error() != nil {
		gbl.Log.Warnf("❗️ error publishing event to redis: %s", err.Error())
	} else {
		gbl.Log.Debug("published event to redis")
	}
}

package pusu

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/osmodels"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/benleb/gloomberg/internal/trapri"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func SubscribeToSales(gb *gloomberg.Gloomberg, channel string, queueTokenTransactions chan *totra.TokenTransaction) {
	pubsub := gb.Rdb.Subscribe(context.Background(), channel)

	ch := pubsub.Channel(redis.WithChannelSize(1024))

	for msg := range ch {
		gbl.Log.Infof("🚇 subscribe channel %s (%d)", msg.Channel, len(ch))

		// validate json
		if !json.Valid([]byte(msg.Payload)) {
			gbl.Log.Warnf("❗️ invalid json: %s", msg.Payload)

			continue
		}

		// create the event transaction
		var ttx totra.TokenTransaction

		// unmarshal event transaction from json
		err := json.Unmarshal([]byte(msg.Payload), &ttx)
		if err != nil {
			gbl.Log.Warnf("❗️ error unmarshalling event Tx: %+v | %s", msg.Payload, err)

			continue
		}

		queueTokenTransactions <- &ttx
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

	pubsub := gb.Rdb.PSubscribe(context.Background(), channels...)

	ch := pubsub.Channel(redis.WithChannelSize(1024))

	for i := 0; i < viper.GetInt("server.pubsub.listings"); i++ {
		go func(i int) {
			gbl.Log.Infof("🚇 starting pusu listings receiver #%d | subscriptions: %d", i, len(channels))

			for msg := range ch {
				gbl.Log.Debugf("🚇 received msg on channel %s (%d): %s", msg.Channel, len(ch), msg.Payload)

				var itemListedEvent osmodels.ItemListedEvent

				// validate json
				if !json.Valid([]byte(msg.Payload)) {
					gbl.Log.Warnf("❗️ invalid json: %s", msg.Payload)

					continue
				}

				// unmarshal
				if err := json.Unmarshal([]byte(msg.Payload), &itemListedEvent); err != nil {
					gbl.Log.Errorf("❌ error json.Unmarshal: %+v\n", err.Error())

					continue
				}

				// nftID is a string in the format <chain>/<contract>/<tokenID>
				nftID := strings.Split(itemListedEvent.Payload.Item.NftID, "/")
				if len(nftID) != 3 {
					gbl.Log.Warnf("🤷‍♀️ error parsing nftID: %s | %+v", itemListedEvent.Payload.Item.NftID, nftID)

					continue
				}

				//
				// discard listings for ignored collections
				if collection, ok := gb.CollectionDB.Collections[common.HexToAddress(nftID[1])]; ok && collection.IgnorePrinting {
					gbl.Log.Debugf("🗑️ ignoring printing for collection %s", collection.Name)

					continue
				}

				// print
				trapri.FormatListing(gb, &itemListedEvent, queueTokenTransactions)
			}
		}(i)
	}
}

func Publish(gb *gloomberg.Gloomberg, channel string, event any) {
	// marshal event to json
	marshalledEvent, err := json.Marshal(event)
	if err != nil {
		gbl.Log.Warnf("❗️ error marshalling event: %s", err)

		return
	}

	// publish event to redis pubsub
	err = gb.Rdb.Publish(context.Background(), channel, marshalledEvent).Err()
	if err != nil {
		gbl.Log.Warnf("❗️ error publishing event to redis: %s", err)
	} else {
		gbl.Log.Debug("published event to redis")
	}
}

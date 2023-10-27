package pusu

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/apex/log"
	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/benleb/gloomberg/internal/seawa/models"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/mitchellh/mapstructure"
	"github.com/redis/rueidis"
	"github.com/spf13/viper"
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

// SubscribeToListingsViaRedis subscribes to all collections for which we have a slug.
func SubscribeToListingsViaRedis(gb *gloomberg.Gloomberg) {
	slugAddresses := gb.CollectionDB.OpenseaSlugAddresses()
	if len(slugAddresses) == 0 {
		gbl.Log.Warn("❌ no slugs to send to gloomberg server")

		return
	}

	// create a list of channels to subscribe to
	channels := make([]string, 0)

	// TODO investigate why this seems not to work in go (but in the redis cli) 🤨
	// for _, collectionAddress := range slugAddresses {
	// 	// channelPattern := internal.PubSubSeaWatcherListings + "/" + collectionAddress.Hex() + "/*"
	// 	// channelPattern := internal.PubSubSeaWatcher + "/*/" + collectionAddress.Hex()
	// 	// channelPattern := internal.PubSubSeaWatcher + "/" + collectionAddress.Hex() + "/*"
	// 	channels = append(channels, channelPattern)
	// }

	channels = append(channels, internal.PubSubSeaWatcher+"/*/*")

	gbl.Log.Infof("🚇 subscribing to redis channels %s", channels)

	eventMessages := make(chan rueidis.PubSubMessage, viper.GetInt("gloomberg.eventhub.inQueuesSize"))

	for i := 0; i < viper.GetInt("gloomberg.eventhub.numHandler"); i++ {
		go eventHandler(gb, &eventMessages)
	}

	err := gb.Rdb.Receive(context.Background(), gb.Rdb.B().Psubscribe().Pattern(channels...).Build(), func(msg rueidis.PubSubMessage) {
		eventMessages <- msg

		gbl.Log.Debugf("🚇 received msg on channel %s", msg.Channel)

		// handle event
		// go handleEvent(gb, msg)
	})
	if err != nil {
		gbl.Log.Errorf("❌ error subscribing to redis channels %s: %s", channels, err.Error())

		return
	}
}

func eventHandler(gb *gloomberg.Gloomberg, eventMessages *chan rueidis.PubSubMessage) {
	for msg := range *eventMessages {
		handleEvent(gb, msg)
	}
}

func handleEvent(gb *gloomberg.Gloomberg, msg rueidis.PubSubMessage) {
	var rawEvent map[string]interface{}

	// validate json
	if !json.Valid([]byte(msg.Message)) {
		gbl.Log.Warnf("❗️ invalid json: %s", msg.Message)

		return
	}

	// unmarshal json
	if err := json.Unmarshal([]byte(msg.Message), &rawEvent); err != nil {
		gbl.Log.Errorf("❌ error json.Unmarshal: %+v\n", err.Error())

		return
	}

	// decode event to general event
	var generalEvent models.GeneralEvent

	// decode event
	rawDecoderConfig := models.GetEventDecoderConfig()
	rawDecoderConfig.Result = &generalEvent
	decoder, _ := mapstructure.NewDecoder(&rawDecoderConfig)

	err := decoder.Decode(rawEvent)
	if err != nil {
		log.Infof("⚓️❌ decoding incoming event failed: %+v | %+v", msg.Message, err)

		return
	}

	// decoder config
	decoderConfig := models.GetEventDecoderConfig()

	switch degendb.GetEventType(generalEvent.EventType) {
	case degendb.Listing:
		var itemListed models.ItemListed

		decoderConfig.Result = &itemListed
		decoder, _ := mapstructure.NewDecoder(&decoderConfig)

		err := decoder.Decode(rawEvent)
		if err != nil {
			log.Infof("⚓️❌ decoding incoming %s event failed: %s", generalEvent, err)

			return
		}

		// push to event hub
		gb.In.ItemListed <- &itemListed

	case degendb.Bid:
		var itemReceivedBid models.ItemReceivedBid

		decoderConfig.Result = &itemReceivedBid
		decoder, _ := mapstructure.NewDecoder(&decoderConfig)

		err := decoder.Decode(rawEvent)
		if err != nil {
			log.Infof("⚓️❌ decoding incoming %s event failed: %s", generalEvent, err)

			return
		}

		// push to event hub
		gb.In.ItemReceivedBid <- &itemReceivedBid

	case degendb.CollectionOffer:
		var collectionOffer models.CollectionOffer

		// decoderConfig := models.GetEventDecoderConfig()
		decoderConfig.Result = &collectionOffer
		decoder, _ := mapstructure.NewDecoder(&decoderConfig)

		err := decoder.Decode(rawEvent)
		if err != nil {
			log.Infof("⚓️❌ decoding incoming event failed: %+v %+v", collectionOffer, err)

			return
		}

		// push to event hub
		gb.In.CollectionOffer <- &collectionOffer

	case degendb.MetadataUpdated:
		var itemMetadataUpdated models.ItemMetadataUpdated

		decoderConfig.Result = &itemMetadataUpdated
		decoder, _ := mapstructure.NewDecoder(&decoderConfig)

		err := decoder.Decode(rawEvent)
		if err != nil {
			log.Infof("⚓️❌ decoding incoming event failed: %+v %+v", itemMetadataUpdated, err)

			return
		}

		// push to event hub
		gb.In.ItemMetadataUpdated <- &itemMetadataUpdated

	default:
		gbl.Log.Warnf("❗️ unknown event type: %s", generalEvent.EventType)
		gbl.Log.Warnf("❗️         %#v", generalEvent)
	}

	logEvent(generalEvent)
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

func logEvent(generalEvent models.GeneralEvent) {
	primaryStyle, _ := style.GenerateAddressStyles(generalEvent.ContractAddress())

	fmtCurrencySymbol := primaryStyle.Bold(false).Render("Ξ")
	fmtPrice := style.BoldAlmostWhite(fmt.Sprintf("%7.4f", generalEvent.BasePrice().Ether())) + fmtCurrencySymbol

	fmtItem := primaryStyle.Bold(true).Render(generalEvent.ItemName())

	fmtFrom := style.FormatAddress(&generalEvent.Payload.Maker.Address)

	out := strings.Builder{}
	out.WriteString(degendb.GetEventType(generalEvent.EventType).Icon())
	out.WriteString(" " + fmtPrice)
	out.WriteString(" " + fmtItem)
	out.WriteString(" " + style.DividerArrowLeft.String())
	out.WriteString(fmtFrom)

	gbl.Log.Debugf(out.String())
}

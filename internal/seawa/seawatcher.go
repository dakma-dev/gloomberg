package seawa

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/osmodels"
	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/benleb/gloomberg/internal/pusu"
	"github.com/benleb/gloomberg/internal/seawa/models"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils/hooks"
	"github.com/charmbracelet/log"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mitchellh/mapstructure"
	"github.com/nshafer/phx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/redis/rueidis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type SeaWatcher struct {
	// UnimplementedSeaWatcherServer

	// channel for events received from the opensea stream
	receivedEvents chan map[string]interface{}

	// phoenix channels client
	phoenixSocket *phx.Socket

	// subscribed phoenix channels
	channels map[string]*phx.Channel

	// subscribed slugs/events
	subscriptions map[string]map[degendb.EventType]func()

	runLocal        bool
	runPubsubServer bool

	// redis client
	rdb rueidis.Client

	gb *gloomberg.Gloomberg

	mu *sync.RWMutex
}

var availableEventTypes = mapset.NewSet[degendb.EventType](degendb.Listing, degendb.CollectionOffer, degendb.Bid)

var eventsReceivedCounter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "gloomberg_seawatcher_events_received_count_total",
	Help: "The number of received events from OpenSea API.",
})

// func NewSeaWatcher(apiToken string, rdb rueidis.Client) *SeaWatcher {.
func NewSeaWatcher(apiToken string, gb *gloomberg.Gloomberg) *SeaWatcher {
	// we might not connect to the stream api locally if we use other ways to get the events
	runLocalAPIClient := viper.GetBool("seawatcher.local")

	if runLocalAPIClient && apiToken == "" {
		log.Info("no OpenSea api token provided, skipping OpenSea stream api")

		return nil
	}

	endpointURL := fmt.Sprint(osmodels.StreamAPIEndpoint, "?token=", apiToken)

	endpoint, err := url.Parse(endpointURL)
	if err != nil {
		log.Info(err)

		return nil
	}

	sw := &SeaWatcher{
		receivedEvents: make(chan map[string]interface{}, 1024),
		subscriptions:  make(map[string]map[degendb.EventType]func()),

		channels: make(map[string]*phx.Channel),

		runLocal:        runLocalAPIClient,
		runPubsubServer: viper.GetBool("seawatcher.pubsub"),

		gb:  gb,
		rdb: gb.Rdb,

		mu: &sync.RWMutex{},
	}

	// create phoenix socket
	if runLocalAPIClient {
		sw.phoenixSocket = phx.NewSocket(endpoint)
		sw.phoenixSocket.Logger = phx.NewCustomLogger(phx.LoggerLevel(phx.LogWarning), zap.NewStdLog(gbl.Log.Desugar()))

		// exponential backoff for reconnects
		sw.phoenixSocket.ReconnectAfterFunc = func(attempt int) time.Duration {
			// max waitTime is 2^7 = 128sec
			waitTime := time.Second * time.Duration(math.Pow(2.0, float64(int(math.Min(float64(attempt), 5)))))

			sw.Prf("❌ reconnecting to OpenSea failed (#%d) 😩 trying again in %dsec..", attempt, int(waitTime.Seconds()))

			return waitTime
		}

		// error function
		sw.phoenixSocket.OnError(func(err error) { gbl.Log.Errorf("❌ seawa socket error: %+v", err) })

		// called on successful connection to the socket/OpenSea
		sw.phoenixSocket.OnOpen(func() {
			sw.Pr("✅ connected to the OpenSea stream")
		})

		// called on disconnect/connection breaks to the socket/OpenSea
		sw.phoenixSocket.OnClose(func() {
			sw.Pr("❕ connection to OpenSea closed, trying to reconnect...")

			err := sw.phoenixSocket.Reconnect()
			if err != nil {
				sw.Prf("❌ reconnecting to OpenSea stream failed: %s", err)
			}
		})

		// initial connection to the socket/OpenSea
		if sw.phoenixSocket != nil {
			sw.Pr("connecting to OpenSea...")

			if err := sw.phoenixSocket.Connect(); err != nil {
				socketError := errors.New("OpenSea stream socket connection failed: " + err.Error())
				sw.Prf("❌ socket error: %s", socketError.Error())

				return nil
			}
		}
	}

	// // subscribe to mgmt channel
	// if viper.GetBool("pubsub.server.enabled") {
	// 	go sw.subscribeToMgmt()
	// }

	return sw
}

// Pr prints messages from seawatcher to the terminal.
func (sw *SeaWatcher) Pr(message string) {
	gloomberg.PrWithKeywordAndIcon("🌊", style.OpenSea.Render("seawa"), message)
}

// Prf formats and prints messages from seawatcher to the terminal.
func (sw *SeaWatcher) Prf(format string, a ...interface{}) {
	sw.Pr(fmt.Sprintf(format, a...))
}

// eventHandler handles incoming stream api events and forwards them as map.
func (sw *SeaWatcher) eventHandler(response any) {
	eventsReceivedCounter.Inc()

	rawEvent, ok := response.(map[string]interface{})
	if !ok {
		log.Errorf("⚓️❌ error in type assertion of received event: %+v", response)

		return
	}

	itemEventType, ok := rawEvent["event_type"].(string)
	if !ok {
		log.Warnf("⚓️ 🤷‍♀️ unknown event type: %s", rawEvent["event_type"])

		return
	}

	// decode event
	var generalEvent models.GeneralEvent

	// decoder config
	decoderConfig := models.GetEventDecoderConfig()
	decoderConfig.Result = &generalEvent
	decoder, _ := mapstructure.NewDecoder(&decoderConfig)

	err := decoder.Decode(rawEvent)
	if err != nil {
		log.Info("⚓️❌ decoding incoming %s event failed: %s", style.Bold(itemEventType), err)

		return
	}

	contractAddress := generalEvent.ContractAddress()
	collectionStyle, _ := style.GenerateAddressStyles(contractAddress)
	collectionStyle = collectionStyle.Bold(true)
	fmtItemName := collectionStyle.Render(generalEvent.ItemNameLink())

	// sw.Prf("⚓️ mapstructure generalEvent %s event: %+v", generalEvent.EventType, collectionStyle.Render(contractAddress.Hex()))

	// switch osmodels.OpenSeaEventType(generalEvent.EventType) {
	switch degendb.GetEventType(generalEvent.EventType) {
	// item listed
	case degendb.Listing:
		var itemListed *models.ItemListed

		decoderConfig.Result = &itemListed
		decoder, _ := mapstructure.NewDecoder(&decoderConfig)

		err := decoder.Decode(rawEvent)
		if err != nil {
			log.Info("⚓️❌ decoding incoming %s event failed: %s", style.Bold(itemEventType), err)

			return
		}

		// push to eventHub for further processing
		sw.gb.In.ItemListed <- itemListed

	case degendb.Bid:
		var itemReceivedBid *models.ItemReceivedBid

		decoderConfig.Result = &itemReceivedBid
		decoder, _ := mapstructure.NewDecoder(&decoderConfig)

		err := decoder.Decode(rawEvent)
		if err != nil {
			log.Info("⚓️❌ decoding incoming %s event failed: %s", style.Bold(itemEventType), err)

			return
		}

		// push to eventHub for further processing
		sw.gb.In.ItemReceivedBid <- itemReceivedBid

	case degendb.CollectionOffer:
		var collectionOffer *models.CollectionOffer

		decoderConfig.Result = &collectionOffer
		decoder, _ := mapstructure.NewDecoder(&decoderConfig)

		err := decoder.Decode(rawEvent)
		if err != nil {
			log.Infof("⚓️❌ decoding incoming %s event failed: %s", style.Bold(itemEventType), err)

			return
		}

		// push to eventHub for further processing
		sw.gb.In.CollectionOffer <- collectionOffer

		// pretty.Println(collectionOffer)

	case degendb.MetadataUpdated:
		var itemMetadataUpdated *models.ItemMetadataUpdated

		decoderConfig.Result = &itemMetadataUpdated
		decoder, _ := mapstructure.NewDecoder(&decoderConfig)

		err := decoder.Decode(rawEvent)
		if err != nil {
			log.Infof("⚓️❌ decoding incoming %s event failed: %s", style.Bold(itemEventType), err)

			return
		}

		// push to eventHub for further processing
		sw.gb.In.ItemMetadataUpdated <- itemMetadataUpdated
	}

	if viper.GetBool("pubsub.server.enabled") {
		publishChannel := internal.PubSubSeaWatcher + "/" + generalEvent.EventType + "/" + contractAddress.Hex()
		pusu.Publish(sw.gb, publishChannel, rawEvent)
	}

	// 💄 styled log
	perItemPrice := price.NewPrice(big.NewInt(0).Div(generalEvent.Payload.GetPrice().Wei(), big.NewInt(int64(generalEvent.Payload.Quantity))))

	// go logEvent(sw, degendb.GetEventType(generalEvent.EventType), contractAddress, generalEvent.BasePrice(), fmtItemName, &generalEvent.Payload.Maker.Address)
	go logEvent(sw, degendb.GetEventType(generalEvent.EventType), contractAddress, perItemPrice, fmtItemName, &generalEvent.Payload.Maker.Address)
}

func (sw *SeaWatcher) Subscribe(subscriptions degendb.SlugSubscriptions) uint64 {
	if !viper.GetBool("pubsub.server.enabled") && !viper.GetBool("seawatcher.pubsub") && !viper.GetBool("seawatcher.local") {
		// runs on the pubsub client side
		gbl.Log.Infof("⚓️ subscribing to: %+v", subscriptions)

		if len(subscriptions) > 0 {
			sw.gb.PublishSlubSubscriptions(subscriptions)
		} else {
			sw.gb.PublishOwnSlubSubscription()
		}

		return uint64(len(subscriptions))
	}

	// runs on the pubsub server side
	newCollectionSubscriptions := uint64(0)

	for _, slugSubscription := range subscriptions {
		if sw.IsSubscribedToEvents(slugSubscription.Slug, slugSubscription.Events) {
			log.Debugf("⚓️ ☕️ already subscribed events for %s", style.BoldAlmostWhite(slugSubscription.Slug))

			continue
		}

		subscribedEvents := make(map[degendb.EventType]func())
		slugEventSubscriptions := mapset.NewSet[degendb.EventType]()

		for _, eventType := range slugSubscription.Events {
			// if subscribedEvents[eventType] != nil {
			// 	log.Debugf("⚓️ ☕️ already subscribed to %s events for %s", eventType, slugSubscription.Slug)

			// 	continue
			// }

			subscribedEvents[eventType] = sw.on(eventType, slugSubscription.Slug, sw.eventHandler)

			// newEventSubscriptions++
			slugEventSubscriptions.Add(eventType)

			time.Sleep(time.Millisecond * 37)
		}

		newCollectionSubscriptions++

		sw.mu.Lock()
		sw.subscriptions[slugSubscription.Slug] = subscribedEvents
		sw.mu.Unlock()

		fmtSlug := style.BoldAlmostWhite(slugSubscription.Slug)
		fmtDivider := style.GrayStyle.Render("|")

		// no collection db -> no addresses -> no colors 😢 fix me!
		if collection := sw.gb.CollectionDB.GetCollectionForSlug(slugSubscription.Slug); collection != nil {
			log.Debugf("⏮️ resetting stats for %s", slugSubscription.Slug)

			collection.ResetStats()

			// use collection colors
			fmtSlug = collection.Render(slugSubscription.Slug)
			fmtDivider = collection.Style().Copy().Faint(true).Render("|")
		}

		sw.Prf("%s: %s", fmtSlug, strings.Join(slugSubscription.ToStringSlice(), fmtDivider))

		time.Sleep(time.Millisecond * 137)
	}

	return newCollectionSubscriptions // newEventSubscriptions
}

// func (sw *SeaWatcher) UnubscribeForSlug(slug string, _ mapset.Set[degendb.EventType]) uint64 {
// 	return sw.Unsubscribe([]string{slug}, nil)
// }

func (sw *SeaWatcher) Unsubscribe(subscriptions degendb.SlugSubscriptions) uint64 {
	numUnsubscribed := uint64(0)

	for _, slugSubscription := range subscriptions {
		if sw.IsSubscribedToEvents(slugSubscription.Slug, slugSubscription.Events) {
			log.Debugf("⚓️ ☕️ not subscribed to events for %s", slugSubscription)

			return 0
		}

		sw.mu.Lock()
		slugSubscriptions := sw.subscriptions[slugSubscription.Slug]
		sw.mu.Unlock()

		if slugSubscriptions != nil {
			// unsubscribe
			for _, unsubscribe := range slugSubscriptions {
				unsubscribe()
			}

			// remove slug
			sw.mu.Lock()
			sw.subscriptions[slugSubscription.Slug] = nil
			sw.mu.Unlock()
		}

		numUnsubscribed++
	}

	log.Debugf("unsubscribed from %d collections/slugs", numUnsubscribed)

	return numUnsubscribed
}

// func (sw *SeaWatcher) UnubscribeForSlugs(subscriptions degendb.SlugSubscriptions) uint64 {
// 	numUnsubscribed := uint64(0)

// 	for _, slug := range slugs {
// 		if sw.IsSubscribedToAllEvents(slug) {
// 			log.Debugf("⚓️ ☕️ not subscribed to events for %s", slug)

// 			return 0
// 		}

// 		sw.mu.Lock()
// 		slugSubscriptions := sw.subscriptions[slug]
// 		sw.mu.Unlock()

// 		if slugSubscriptions != nil {
// 			// unsubscribe
// 			for _, unsubscribe := range slugSubscriptions {
// 				unsubscribe()
// 			}

// 			// remove slug
// 			sw.mu.Lock()
// 			sw.subscriptions[slug] = nil
// 			sw.mu.Unlock()
// 		}

// 		numUnsubscribed++
// 	}

// 	log.Debugf("unsubscribed %d from opensea events", numUnsubscribed)

// 	return numUnsubscribed
// }

func (sw *SeaWatcher) IsSubscribedToAllEvents(slug string) bool {
	sw.mu.Lock()
	slugSubscriptions, ok := sw.subscriptions[slug]
	sw.mu.Unlock()

	if !ok { // || len(slugSubscriptions) < availableEventTypes.Cardinality() {
		return false
	}

	for _, eventType := range availableEventTypes.ToSlice() {
		sw.mu.Lock()
		cancelSubscriptionFunc, ok := slugSubscriptions[eventType]
		sw.mu.Unlock()

		if !ok {
			log.Errorf("⚓️❌ error while checking existing eventtype subscriptions for %s / %+v", slug, eventType)

			return false
		}

		if cancelSubscriptionFunc == nil {
			return false
		}
	}

	return true
}

func (sw *SeaWatcher) IsSubscribedToEvents(slug string, subscriptionEventTypes []degendb.EventType) bool {
	sw.mu.Lock()
	slugSubscriptions, ok := sw.subscriptions[slug]
	sw.mu.Unlock()

	if !ok { // || len(slugSubscriptions) < availableEventTypes.Cardinality() {
		return false
	}

	for _, eventType := range subscriptionEventTypes {
		if _, ok := slugSubscriptions[eventType]; !ok {
			return false
		}
	}

	return true
}

func (sw *SeaWatcher) createChannel(topic string) *phx.Channel {
	channel := sw.phoenixSocket.Channel(topic, nil)

	join, err := channel.Join()
	if err != nil {
		sw.Prf("⚓️❌ failed to join channel: %s", err)

		return nil
	}

	join.Receive("ok", func(_ any) {
		log.Debugf("👋 joined channel: %s", channel.Topic())
	})

	join.Receive("error", func(response any) {
		log.Warn("failed to joined channel:", channel.Topic(), response)
	})

	sw.mu.Lock()
	sw.channels[topic] = channel
	sw.mu.Unlock()

	return channel
}

func (sw *SeaWatcher) getChannel(topic string) *phx.Channel {
	channel, ok := sw.channels[topic]
	if !ok {
		channel = sw.createChannel(topic)
	}

	return channel
}

func (sw *SeaWatcher) on(eventType degendb.EventType, collectionSlug string, eventHandler func(response any)) func() {
	topic := fmt.Sprintf("collection:%s", collectionSlug)
	openseaEvent := strings.ToLower(eventType.OpenseaEventName())

	log.Debugf("Fetching channel %s", topic)
	channel := sw.getChannel(topic)

	log.Debugf("subscribing to %s events on %s", openseaEvent, topic)
	channel.On(openseaEvent, eventHandler)

	channel.OnClose(func(payload any) {
		sw.Prf("⚠️ Channel %s closed: %s", topic, payload)
	})

	log.Debugf("␚ subscribed to %s for %s", openseaEvent, collectionSlug)

	return func() {
		sw.Prf("Unsubscribing from %s events on %s", openseaEvent, topic)

		leave, err := channel.Leave()
		if err != nil {
			sw.Prf("channel.Leave err: %s", err)
		}

		leave.Receive("ok", func(_ any) {
			delete(sw.channels, collectionSlug)
			sw.Prf("Successfully left channel %s listening for %s", topic, openseaEvent)
		})
	}
}

// func (sw *SeaWatcher) subscribeToMgmt() {
// 	log.Debugf("subscribing to mgmt channel %s", internal.PubSubSeaWatcherMgmt)

// 	subscriptionsChannel := sw.gb.SubscribeSeawatcherSubscriptions()

// 	for {
// 		subscriptionEvent := <-subscriptionsChannel
// 		sw.serverHandleMgmtEvent(subscriptionEvent)
// 	}
// }

// ServerSubscribeToPubsubMgmt starts the seawatcher by subscribing to the mgmt channel and listening for new slugs to subscribe to.
func (sw *SeaWatcher) ServerSubscribeToPubsubMgmt() {
	sw.Prf("👔 subscribing to mgmt channel %s", style.AlmostWhiteStyle.Render(internal.PubSubSeaWatcherMgmt))

	err := sw.rdb.Receive(context.Background(), sw.rdb.B().Subscribe().Channel(internal.PubSubSeaWatcherMgmt).Build(), func(msg rueidis.PubSubMessage) {
		log.Debugf("👔 received msg on channel %s: %s", msg.Channel, msg.Message)

		// validate json
		if !json.Valid([]byte(msg.Message)) {
			gbl.Log.Warnf("❗️ invalid json: %s", msg.Message)

			return
		}

		// unmarshal json to map
		var rawEvent map[string]interface{}
		if err := json.Unmarshal([]byte(msg.Message), &rawEvent); err != nil {
			log.Errorf("⚓️❌ error json.Unmarshal to map: %+v", err)

			return
		}

		// decode to event
		var subscriptionEvent *models.SubscriptionEvent
		decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			DecodeHook: mapstructure.ComposeDecodeHookFunc(hooks.StringToEventTypeHookFunc()),
			Result:     &subscriptionEvent,
		})

		err := decoder.Decode(rawEvent)
		if err != nil {
			log.Infof("⚓️❌ decoding incoming event failed: %+v | %+v", msg.Message, err)

			return
		}

		go sw.serverHandleMgmtEvent(subscriptionEvent)
	})
	if err != nil {
		log.Errorf("❌ error subscribing to redis channels %s: %s", internal.PubSubSeaWatcherMgmt, err.Error())

		return
	}
}

func (sw *SeaWatcher) serverHandleMgmtEvent(subscriptionEvent *models.SubscriptionEvent) {
	switch subscriptionEvent.Action {
	case models.SendSlugs:
		// SendSlugs can be ignored on server side for now
		return

	case models.Subscribe, models.Unsubscribe:
		if !viper.GetBool("seawatcher.local") {
			sw.Prf("⚓️❌ can't subscribe to mgmt events, not running local OpenSea client")

			return
		}

		sw.Prf("👔 received %s for %s collections/slugs...", style.AlmostWhiteStyle.Render(subscriptionEvent.Action.String()), style.AlmostWhiteStyle.Render(strconv.Itoa(len(subscriptionEvent.Collections))))
		if len(subscriptionEvent.Collections) == 0 {
			log.Error("⚓️❌ incoming collection slugs msg is empty")

			return
		}

		if viper.GetString("api_keys.opensea") == "" {
			log.Error("⚓️❌ OpenSea api key is not set, can't subscribe to listings")

			return
		}

		var action func(subscriptions degendb.SlugSubscriptions) uint64

		switch subscriptionEvent.Action {
		case models.Subscribe:
			action = sw.Subscribe
		case models.Unsubscribe:
			action = sw.Unsubscribe
		}

		newEventSubscriptions := action(subscriptionEvent.Collections)

		sw.Prf(
			"👔 subscribed for %s new collections/slugs | total subscribed collections: %s",
			style.AlmostWhiteStyle.Render(strconv.FormatUint(newEventSubscriptions, 10)),
			style.AlmostWhiteStyle.Render(strconv.Itoa(len(sw.subscriptions))),
		)

	default:
		sw.Prf("👔 👀 received unknown mgmt event: %s", subscriptionEvent.Action.String())

		return
	}
}

func (sw *SeaWatcher) ServerRequestSlugSubscriptions() {
	// build "SendSlugs" event
	requestSlugsEvent := &models.MgmtEvent{
		Action: models.SendSlugs,
	}

	// marshal event
	jsonMgmtEvent, err := json.Marshal(requestSlugsEvent)
	if err != nil {
		log.Error("⚓️❌ marshal failed for SendSlugs action: %s | %v", err, requestSlugsEvent)

		return
	}

	if sw.rdb.Do(context.Background(), sw.rdb.B().Publish().Channel(internal.PubSubSeaWatcherMgmt).Message(string(jsonMgmtEvent)).Build()).Error() != nil {
		log.Errorf("⚓️❌ error publishing %s to redis: %s", requestSlugsEvent.Action.String(), err.Error())
	} else {
		sw.Prf("👔 published %s event to %s", style.AlmostWhiteStyle.Render(requestSlugsEvent.Action.String()), style.AlmostWhiteStyle.Render(internal.PubSubSeaWatcherMgmt))
	}
}

func logEvent(sw *SeaWatcher, eventType *degendb.GBEventType, address *common.Address, price *price.Price, fmtItem string, from *common.Address) {
	primaryStyle, _ := style.GenerateAddressStyles(address)

	fmtCurrencySymbol := primaryStyle.Bold(false).Render("Ξ")
	fmtPrice := style.BoldAlmostWhite(fmt.Sprintf("%7.4f", price.Ether())) + fmtCurrencySymbol

	fmtFrom := style.FormatAddress(from)

	out := strings.Builder{}
	out.WriteString(eventType.Icon())
	out.WriteString(" " + fmtPrice)
	out.WriteString(" " + fmtItem)
	out.WriteString(" " + style.DividerArrowLeft.String())
	out.WriteString(fmtFrom)

	sw.Pr(out.String())
}

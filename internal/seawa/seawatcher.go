package seawa

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/url"
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

	// start worker for managing subscriptions
	if viper.GetBool("pubsub.client.enabled") || viper.GetBool("seawatcher.pubsub") {
		go sw.WorkerMgmtChannel()
	}

	return sw
}

// Pr prints messages from seawatcher to the terminal.
func (sw *SeaWatcher) Pr(message string) {
	gloomberg.PrWithKeywordAndIcon("🌊", style.OpenSea("seawa"), message)
}

// Prf formats and prints messages from seawatcher to the terminal.
func (sw *SeaWatcher) Prf(format string, a ...interface{}) {
	sw.Pr(fmt.Sprintf(format, a...))
}

func (sw *SeaWatcher) EventChannel() chan map[string]interface{} {
	return sw.receivedEvents
}

// func (sw *SeaWatcher) ActiveSubscriptions() map[string]map[osmodels.EventType]func() {.
func (sw *SeaWatcher) ActiveSubscriptions() map[string]map[degendb.EventType]func() {
	totalSubscriptions := 0

	sw.mu.Lock()
	slugSubscriptions := sw.subscriptions
	sw.mu.Unlock()

	for _, eventSubscriptions := range slugSubscriptions {
		totalSubscriptions += len(eventSubscriptions)
	}

	return slugSubscriptions
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
	fmtItemName := collectionStyle.Render(generalEvent.ItemName())

	// sw.Prf("⚓️ mapstructure generalEvent %s event: %+v", generalEvent.EventType, collectionStyle.Render(contractAddress.Hex()))

	switch osmodels.EventType(generalEvent.EventType) {
	// item listed
	case osmodels.ItemListed:
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

	case osmodels.ItemReceivedBid:
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

	case osmodels.CollectionOffer:
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

	case osmodels.ItemMetadataUpdated:
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
	go logEvent(sw, degendb.GetEventType(generalEvent.EventType), contractAddress, generalEvent.BasePrice(), fmtItemName, &generalEvent.Payload.Maker.Address)
}

func (sw *SeaWatcher) DecodeItemReceivedBidEvent(itemEvent map[string]interface{}) (osmodels.ItemReceivedBidEvent, error) {
	var collectionOfferEvent osmodels.ItemReceivedBidEvent

	decodeHooks := mapstructure.ComposeDecodeHookFunc(
		hooks.StringToAddressHookFunc(),
	)

	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: decodeHooks,
		Result:     &collectionOfferEvent,
	})

	err := decoder.Decode(itemEvent)
	if err != nil {
		sw.Prf("⚓️❌ decoding incoming OpenSea stream api ItemReceivedBidEvent failed: %s", err)

		return osmodels.ItemReceivedBidEvent{}, err
	}

	return collectionOfferEvent, err
}

func (sw *SeaWatcher) DecodeCollectionOfferEvent(itemEvent map[string]interface{}) (osmodels.CollectionOfferEvent, error) {
	var collectionOfferEvent osmodels.CollectionOfferEvent

	decodeHooks := mapstructure.ComposeDecodeHookFunc(
		hooks.StringToAddressHookFunc(),
	)

	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: decodeHooks,
		Result:     &collectionOfferEvent,
	})

	err := decoder.Decode(itemEvent)
	if err != nil {
		sw.Prf("⚓️❌ decoding incoming OpenSea stream api collection offer event failed: %s", err)

		return osmodels.CollectionOfferEvent{}, err
	}

	return collectionOfferEvent, err
}

func (sw *SeaWatcher) SubscribeForSlug(slug string, eventTypes []degendb.EventType) uint64 {
	eventTypesSet := mapset.NewSet[degendb.EventType]()
	for _, eventType := range eventTypes {
		eventTypesSet.Add(eventType)
	}

	return sw.SubscribeForSlugs([]string{slug}, eventTypesSet)
}

func (sw *SeaWatcher) SubscribeForSlugs(slugs []string, eventTypes mapset.Set[degendb.EventType]) uint64 {
	if !viper.GetBool("pubsub.server.enabled") && !viper.GetBool("seawatcher.pubsub") && !viper.GetBool("seawatcher.local") {
		gbl.Log.Infof("⚓️ subscribeing to: %+v", style.BoldAlmostWhite(strings.Join(slugs, ", ")))
		sw.gb.PublishCollectionsSlugs(slugs)

		return uint64(len(slugs))
	}

	newEventSubscriptions := uint64(0)

	for _, slug := range slugs {
		if sw.IsSubscribedToAllEvents(slug) {
			log.Debugf("⚓️ ☕️ already subscribed to OpenSea events for %s", slug)

			return 0
		}

		slugSubscriptions := make(map[degendb.EventType]func())
		slugEventSubscriptions := mapset.NewSet[string]()

		for _, eventType := range eventTypes.ToSlice() {
			if slugSubscriptions[eventType] != nil {
				log.Debugf("⚓️ ☕️ already subscribed to %s events for %s", eventType, slug)

				continue
			}

			slugSubscriptions[eventType] = sw.on(eventType, slug, sw.eventHandler)

			newEventSubscriptions++
			slugEventSubscriptions.Add(eventType.String())

			time.Sleep(time.Millisecond * 37)
		}

		sw.mu.Lock()
		sw.subscriptions[slug] = slugSubscriptions
		sw.mu.Unlock()

		fmtSlug := style.BoldAlmostWhite(slug)
		fmtDivider := style.GrayStyle.Render("|")

		// no collection db -> no addresses -> no colors 😢 fix me!
		if collection := sw.gb.CollectionDB.GetCollectionForSlug(slug); collection != nil {
			log.Debugf("⏮️ resetting stats for %s", slug)

			collection.ResetStats()

			// use collection colors
			fmtSlug = collection.Render(slug)
			fmtDivider = collection.Style().Copy().Faint(true).Render("|")
		}

		sw.Prf("%s: %s", fmtSlug, strings.Join(slugEventSubscriptions.ToSlice(), fmtDivider))

		time.Sleep(time.Millisecond * 137)
	}

	return newEventSubscriptions
}

func (sw *SeaWatcher) UnubscribeForSlug(slug string, _ mapset.Set[degendb.EventType]) uint64 {
	return sw.UnubscribeForSlugs([]string{slug}, nil)
}

func (sw *SeaWatcher) UnubscribeForSlugs(slugs []string, _ mapset.Set[degendb.EventType]) uint64 {
	numUnsubscribed := uint64(0)

	for _, slug := range slugs {
		if sw.IsSubscribedToAllEvents(slug) {
			log.Debugf("⚓️ ☕️ not subscribed to events for %s", slug)

			return 0
		}

		sw.mu.Lock()
		slugSubscriptions := sw.subscriptions[slug]
		sw.mu.Unlock()

		if slugSubscriptions != nil {
			// unsubscribe
			for _, unsubscribe := range slugSubscriptions {
				unsubscribe()
			}

			// remove slug
			sw.mu.Lock()
			sw.subscriptions[slug] = nil
			sw.mu.Unlock()
		}

		numUnsubscribed++
	}

	log.Debugf("unsubscribed %d from opensea events", numUnsubscribed)

	return numUnsubscribed
}

func (sw *SeaWatcher) IsSubscribedToAllEvents(slug string) bool {
	sw.mu.Lock()
	slugSubscriptions, ok := sw.subscriptions[slug]
	sw.mu.Unlock()

	if !ok || len(slugSubscriptions) < availableEventTypes.Cardinality() {
		return false
	}

	for _, eventType := range availableEventTypes.ToSlice() {
		cancelSubscriptionFunc, ok := slugSubscriptions[eventType]

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

	sw.channels[topic] = channel

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

func (sw *SeaWatcher) WorkerMgmtChannel() {
	log.Debugf("subscribing to mgmt channel %s", internal.PubSubSeaWatcherMgmt)

	mgmtChannel := sw.gb.SubscribeSeawatcherMgmt()

	for {
		mgmtEvent := <-mgmtChannel

		sw.handleMgmtEvent(mgmtEvent)
	}
}

// SubscribeToPubsubMgmt starts the seawatcher by subscribing to the mgmt channel and listening for new slugs to subscribe to.
func (sw *SeaWatcher) SubscribeToPubsubMgmt() {
	sw.Prf("👔 subscribing to mgmt channel %s", style.AlmostWhiteStyle.Render(internal.PubSubSeaWatcherMgmt))

	err := sw.rdb.Receive(context.Background(), sw.rdb.B().Subscribe().Channel(internal.PubSubSeaWatcherMgmt).Build(), func(msg rueidis.PubSubMessage) {
		log.Infof("👔 received msg on channel %s: %s", msg.Channel, msg.Message)

		var mgmtEvent *models.MgmtEvent

		if err := json.Unmarshal([]byte(msg.Message), &mgmtEvent); err != nil {
			log.Errorf("⚓️❌ error json.Unmarshal: %+v", err)

			return
		}

		sw.handleMgmtEvent(mgmtEvent)
	})
	if err != nil {
		log.Errorf("❌ error subscribing to redis channels %s: %s", internal.PubSubSeaWatcherMgmt, err.Error())

		return
	}

	sw.Prf("👔 subscribed to %s", style.AlmostWhiteStyle.Render(internal.PubSubSeaWatcherMgmt))
}

func (sw *SeaWatcher) handleMgmtEvent(mgmtEvent *models.MgmtEvent) {
	switch mgmtEvent.Action {
	case models.SendSlugs:
		// SendSlugs can be ignored on server side for now
		return

	case models.Subscribe, models.Unsubscribe:
		if !viper.GetBool("seawatcher.local") {
			sw.Prf("⚓️❌ can't subscribe to mgmt events, not running local OpenSea client")

			return
		}

		sw.Prf("received %s for %s collections/slugs...", style.AlmostWhiteStyle.Render(mgmtEvent.Action.String()), style.AlmostWhiteStyle.Render(fmt.Sprint(len(mgmtEvent.Slugs))))
		if len(mgmtEvent.Slugs) == 0 {
			log.Error("⚓️❌ incoming collection slugs msg is empty")

			return
		}

		if viper.GetString("api_keys.opensea") == "" {
			log.Error("⚓️❌ OpenSea api key is not set, can't subscribe to listings")

			return
		}

		var action func(slug []string, eventTypes mapset.Set[degendb.EventType]) uint64

		switch mgmtEvent.Action {
		case models.Subscribe:
			action = sw.SubscribeForSlugs
		case models.Unsubscribe:
			action = sw.UnubscribeForSlugs
		}

		newEventSubscriptions := action(mgmtEvent.Slugs, availableEventTypes)

		sw.Prf(
			"👔 successfully subscribed to %s new collections/slugs | total subscribed collections: %s",
			style.AlmostWhiteStyle.Render(fmt.Sprint(newEventSubscriptions)),
			style.AlmostWhiteStyle.Render(fmt.Sprint(len(sw.ActiveSubscriptions()))),
		)

	default:
		sw.Prf("👔 👀 received unknown mgmt event: %s", mgmtEvent.Action.String())

		return
	}
}

func (sw *SeaWatcher) PublishSendSlugs() {
	// build "SendSlugs" event
	sendSlugsEvent := &models.MgmtEvent{
		Action: models.SendSlugs,
	}

	// marshal event
	jsonMgmtEvent, err := json.Marshal(sendSlugsEvent)
	if err != nil {
		log.Error("⚓️❌ marshal failed for SendSlugs action: %s | %v", err, sendSlugsEvent)

		return
	}

	if sw.rdb.Do(context.Background(), sw.rdb.B().Publish().Channel(internal.PubSubSeaWatcherMgmt).Message(string(jsonMgmtEvent)).Build()).Error() != nil {
		log.Errorf("⚓️❌ error publishing %s to redis: %s", sendSlugsEvent.Action.String(), err.Error())
	} else {
		sw.Prf("👔 published %s event to %s", style.AlmostWhiteStyle.Render(sendSlugsEvent.Action.String()), style.AlmostWhiteStyle.Render(internal.PubSubSeaWatcherMgmt))
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

package seawa

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/url"
	"sync"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/osmodels"
	"github.com/benleb/gloomberg/internal/seawa/models"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils/hooks"
	"github.com/charmbracelet/log"
	"github.com/mitchellh/mapstructure"
	"github.com/nshafer/phx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/redis/rueidis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type SeaWatcher struct {
	// channel for events received from the opensea stream
	receivedEvents chan map[string]interface{}

	// phoenix channels client
	phoenixSocket *phx.Socket

	// subscribed phoenix channels
	channels map[string]*phx.Channel

	// subscribed slugs/events
	// subscriptions map[osmodels.EventType]map[string]func()
	subscriptions map[string]map[osmodels.EventType]func()

	// redis client
	rdb rueidis.Client

	gb *gloomberg.Gloomberg

	mu *sync.Mutex
}

var (
	availableEventTypes = []osmodels.EventType{osmodels.ItemListed, osmodels.ItemReceivedBid, osmodels.ItemMetadataUpdated} // , osmodels.CollectionOffer} //, osmodels.ItemMetadataUpdated} // ItemMetadataUpdated, ItemCancelled

	eventsReceivedTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "gloomberg_oswatcher_events_received_total",
		Help: "The total number of received events from the OpenSea api/stream",
	})
)

// func NewSeaWatcher(apiToken string, rdb rueidis.Client) *SeaWatcher {.
func NewSeaWatcher(apiToken string, gb *gloomberg.Gloomberg) *SeaWatcher {
	if apiToken == "" {
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
		subscriptions:  make(map[string]map[osmodels.EventType]func(), 0),

		channels: make(map[string]*phx.Channel),

		gb:  gb,
		rdb: gb.Rdb,

		mu: &sync.Mutex{},
	}

	// create phoenix socket
	sw.phoenixSocket = phx.NewSocket(endpoint)
	sw.phoenixSocket.Logger = phx.NewCustomLogger(phx.LoggerLevel(phx.LogWarning), zap.NewStdLog(gbl.Log.Desugar()))

	// exponential backoff for reconnects
	sw.phoenixSocket.ReconnectAfterFunc = func(attempt int) time.Duration {
		// max waitTime is 2^7 = 128sec
		waitTime := time.Second * time.Duration(math.Pow(2.0, float64(int(math.Min(float64(attempt), 7)))))

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

	// start worker for managing subscriptions
	go sw.WorkerMgmtChannel()

	return sw
}

// Pr prints messages from seawatcher to the terminal.
func (sw *SeaWatcher) Pr(message string) {
	sw.gb.PrWithKeywordAndIcon("🌊", style.OpenSea("seawa"), message)
}

// Prf formats and prints messages from seawatcher to the terminal.
func (sw *SeaWatcher) Prf(format string, a ...interface{}) {
	sw.Pr(fmt.Sprintf(format, a...))
}

func (sw *SeaWatcher) EventChannel() chan map[string]interface{} {
	return sw.receivedEvents
}

func (sw *SeaWatcher) ActiveSubscriptions() map[string]map[osmodels.EventType]func() {
	return sw.subscriptions
}

// eventHandler handles incoming stream api events and forwards them as map.
func (sw *SeaWatcher) eventHandler(response any) {
	eventsReceivedTotal.Inc()

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
	var metadata *mapstructure.Metadata

	// decoder config
	decoderConfig := &mapstructure.DecoderConfig{
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			hooks.StringToAddressHookFunc(),
			hooks.StringToHashHookFunc(),
			hooks.StringToBigIntHookFunc(),
			models.StringToNftIDHookFunc(),
			mapstructure.OrComposeDecodeHookFunc(
				hooks.StringToUnixTimeHookFunc(),
				mapstructure.StringToTimeHookFunc(time.RFC3339),
			),
		),
		Metadata: metadata,
	}

	switch osmodels.EventType(itemEventType) {
	// item listed
	case osmodels.ItemListed:
		var event *models.ItemListed

		decoderConfig.Result = &event
		decoder, _ := mapstructure.NewDecoder(decoderConfig)

		err := decoder.Decode(rawEvent)
		if err != nil {
			log.Info("⚓️❌ decoding incoming %s event failed: %s", style.Bold(itemEventType), err)

			return
		}

		// push to eventHub for further processing
		sw.gb.In.ItemListed <- event

	case osmodels.ItemReceivedBid:
		var event *models.ItemReceivedBid

		decoderConfig.Result = &event
		decoder, _ := mapstructure.NewDecoder(decoderConfig)

		err := decoder.Decode(rawEvent)
		if err != nil {
			log.Info("⚓️❌ decoding incoming %s event failed: %s", style.Bold(itemEventType), err)

			return
		}

		// push to eventHub for further processing
		sw.gb.In.ItemReceivedBid <- event

	case osmodels.CollectionOffer:
		var event *models.CollectionOffer

		decoderConfig.Result = &event
		decoder, _ := mapstructure.NewDecoder(decoderConfig)

		err := decoder.Decode(rawEvent)
		if err != nil {
			log.Infof("⚓️❌ decoding incoming %s event failed: %s", style.Bold(itemEventType), err)

			return
		}

		// push to eventHub for further processing
		sw.gb.In.CollectionOffer <- event
	}
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

func (sw *SeaWatcher) SubscribeForSlug(slug string) bool {
	sw.mu.Lock()
	alreadySubscribed := sw.subscriptions[slug]

	if alreadySubscribed != nil {
		sw.mu.Unlock()

		log.Debugf("⚓️ ☕️ already subscribed to OpenSea events for %s", slug)

		return false
	}

	sw.subscriptions[slug] = make(map[osmodels.EventType]func())

	for _, eventType := range availableEventTypes {
		sw.subscriptions[slug][eventType] = sw.on(eventType, slug, sw.eventHandler)
	}
	sw.mu.Unlock()

	if collection := sw.gb.CollectionDB.GetCollectionForSlug(slug); collection != nil {
		log.Debugf("⏮️ resetting stats for %s", slug)

		collection.ResetStats()
	}

	return true
}

func (sw *SeaWatcher) UnubscribeForSlug(slug string) bool {
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

		return true
	}

	log.Debugf("unsubscribed %s from opense events", slug)

	return false
}

func (sw *SeaWatcher) IsSubscribed(slug string) bool {
	sw.mu.Lock()
	alreadySubscribed, ok := sw.subscriptions[slug]
	sw.mu.Unlock()

	if ok && alreadySubscribed != nil {
		return true
	}

	return false
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

func (sw *SeaWatcher) on(eventType osmodels.EventType, collectionSlug string, eventHandler func(response any)) func() {
	topic := fmt.Sprintf("collection:%s", collectionSlug)

	log.Debugf("Fetching channel %s", topic)
	channel := sw.getChannel(topic)

	log.Debugf("Subscribing to %s events on %s", eventType, topic)
	channel.On(string(eventType), eventHandler)

	channel.OnClose(func(payload any) {
		sw.Prf("⚠️ Channel %s closed: %s", topic, payload)
	})

	log.Debugf("␚ 🔔 subscribed to %s for %s", string(eventType), collectionSlug)

	return func() {
		sw.Prf("Unsubscribing from %s events on %s", eventType, topic)

		leave, err := channel.Leave()
		if err != nil {
			sw.Prf("channel.Leave err: %s", err)
		}

		leave.Receive("ok", func(_ any) {
			delete(sw.channels, collectionSlug)
			sw.Prf("Successfully left channel %s listening for %s", topic, eventType)
		})
	}
}

func (sw *SeaWatcher) WorkerMgmtChannel() {
	log.Debugf("subscribing to mgmt channel %s", internal.TopicSeaWatcherMgmt)

	mgmtChannel := sw.gb.SubscribeSeawatcherMgmt()

	for {
		mgmtEvent := <-mgmtChannel

		sw.handleMgmtEvent(mgmtEvent)
	}
}

// SubscribeToPubsubMgmt starts the seawatcher by subscribing to the mgmt channel and listening for new slugs to subscribe to.
func (sw *SeaWatcher) SubscribeToPubsubMgmt() {
	sw.Prf("subscribing to mgmt channel %s", style.AlmostWhiteStyle.Render(internal.TopicSeaWatcherMgmt))

	err := sw.rdb.Receive(context.Background(), sw.rdb.B().Subscribe().Channel(internal.TopicSeaWatcherMgmt).Build(), func(msg rueidis.PubSubMessage) {
		log.Debugf("⚓️ received msg on channel %s: %s", msg.Channel, msg.Message)

		var mgmtEvent *models.MgmtEvent

		if err := json.Unmarshal([]byte(msg.Message), &mgmtEvent); err != nil {
			log.Errorf("⚓️❌ error json.Unmarshal: %+v", err)

			return
		}

		sw.handleMgmtEvent(mgmtEvent)
	})
	if err != nil {
		log.Errorf("❌ error subscribing to redis channels %s: %s", internal.TopicSeaWatcherMgmt, err.Error())

		return
	}
}

func (sw *SeaWatcher) handleMgmtEvent(mgmtEvent *models.MgmtEvent) {
	switch mgmtEvent.Action {
	case models.SendSlugs:
		// SendSlugs can be ignored on server side for now
		return

	case models.Subscribe, models.Unsubscribe:
		sw.Prf("received %s for %s collections/slugs...", style.AlmostWhiteStyle.Render(mgmtEvent.Action.String()), style.AlmostWhiteStyle.Render(fmt.Sprint(len(mgmtEvent.Slugs))))
		if len(mgmtEvent.Slugs) == 0 {
			log.Error("⚓️❌ incoming collection slugs msg is empty")

			return
		}

		if viper.GetString("api_keys.opensea") == "" {
			log.Error("⚓️❌ OpenSea api key is not set, can't subscribe to listings")

			return
		}

		var action func(slug string) bool

		switch mgmtEvent.Action {
		case models.Subscribe:
			action = sw.SubscribeForSlug
		case models.Unsubscribe:
			action = sw.UnubscribeForSlug
		}

		newEventSubscriptions := 0

		for _, slug := range mgmtEvent.Slugs {
			if slug == "ens" {
				sw.Pr("⚓️ ␚ skipping ens for now")

				continue
			}

			if action(slug) {
				newEventSubscriptions++

				time.Sleep(337 * time.Millisecond)
			}
		}

		sw.Prf(
			"successfully subscribed to %s new collections/slugs | total subscribed collections: %s",
			style.AlmostWhiteStyle.Render(fmt.Sprint(newEventSubscriptions)),
			style.AlmostWhiteStyle.Render(fmt.Sprint(len(sw.ActiveSubscriptions()))),
		)

	default:
		sw.Prf("⚓️ 👀 received unknown mgmt event: %s", mgmtEvent.Action.String())

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

	if sw.rdb.Do(context.Background(), sw.rdb.B().Publish().Channel(internal.TopicSeaWatcherMgmt).Message(string(jsonMgmtEvent)).Build()).Error() != nil {
		log.Errorf("⚓️❌ error publishing %s to redis: %s", sendSlugsEvent.Action.String(), err.Error())
	} else {
		sw.Prf("📣 published %s event to %s", style.AlmostWhiteStyle.Render(sendSlugsEvent.Action.String()), style.AlmostWhiteStyle.Render(internal.TopicSeaWatcherMgmt))
	}
}

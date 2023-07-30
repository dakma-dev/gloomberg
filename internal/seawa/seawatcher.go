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
	gbgrpc "github.com/benleb/gloomberg/internal/nemo/gloomberg/gbgrpc/gen"
	"github.com/benleb/gloomberg/internal/nemo/osmodels"
	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/benleb/gloomberg/internal/seawa/models"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils/hooks"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/mitchellh/mapstructure"
	"github.com/nshafer/phx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/redis/rueidis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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
	// subscriptions map[osmodels.EventType]map[string]func()
	// subscriptions map[string]map[osmodels.EventType]func()
	subscriptions map[string]map[gbgrpc.EventType]func()

	runLocal        bool
	runPubsubServer bool
	runGRPCServer   bool

	// redis client
	rdb rueidis.Client

	gb *gloomberg.Gloomberg

	mu *sync.Mutex
}

// availableEventTypes = []osmodels.EventType{osmodels.ItemListed, osmodels.ItemMetadataUpdated, osmodels.ItemReceivedBid, osmodels.CollectionOffer} // , osmodels.ItemMetadataUpdated} // ItemMetadataUpdated, ItemCancelled.
var availableEventTypes = []gbgrpc.EventType{gbgrpc.EventType_ITEM_LISTED, gbgrpc.EventType_METADATA_UPDATED, gbgrpc.EventType_ITEM_RECEIVED_BID, gbgrpc.EventType_COLLECTION_OFFER} //nolint:nosnakecase // ItemMetadataUpdated} // ItemMetadataUpdated, ItemCancelled

var eventsReceivedCounter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "gloomberg_seawatcher_events_received_count_total",
	Help: "The number of received events from OpenSea API.",
})

// func NewSeaWatcher(apiToken string, rdb rueidis.Client) *SeaWatcher {.
func NewSeaWatcher(apiToken string, gb *gloomberg.Gloomberg) *SeaWatcher {
	// we might not connect to the stream api locally if we use grpc or other ways to get the events
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
		// subscriptions:  make(map[string]map[osmodels.EventType]func()),
		subscriptions: make(map[string]map[gbgrpc.EventType]func()),

		channels: make(map[string]*phx.Channel),

		runLocal:        runLocalAPIClient,
		runPubsubServer: viper.GetBool("seawatcher.pubsub"),
		runGRPCServer:   viper.IsSet("seawatcher.grpc.listen"),

		gb:  gb,
		rdb: gb.Rdb,

		mu: &sync.Mutex{},
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
	if viper.GetBool("seawatcher.grpc.pubsub") {
		go sw.WorkerMgmtChannel()
	}

	// if viper.GetBool("seawatcher.grpc.server.enabled") {
	// 	sw.Prf("starting grpc server...")

	// 	listenHost := viper.GetString("seawatcher.grpc.listen")
	// 	port := viper.GetUint16("seawatcher.grpc.port")
	// 	serverAddress := fmt.Sprintf("%s:%d", listenHost, port)

	// 	// configure grpc server
	// 	go func() {
	// 		grpcListener, err := net.Listen("tcp", serverAddress)
	// 		if err != nil {
	// 			log.Errorf("failed to listen: %v", err)
	// 		}

	// 		var opts []grpc.ServerOption
	// 		if creds, err := gloomberg.GetTLSCredentialsWithoutClientAuth(); err == nil {
	// 			opts = []grpc.ServerOption{grpc.Creds(creds)}
	// 		}

	// 		// start grpc server
	// 		grpcServer := grpc.NewServer(opts...)
	// 		RegisterSeaWatcherServer(grpcServer, sw)

	// 		go log.Error(grpcServer.Serve(grpcListener))
	// 	}()

	// 	sw.Prf("grpc server started on %+v", style.BoldAlmostWhite(serverAddress))
	// }

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

// func (sw *SeaWatcher) ActiveSubscriptions() map[string]map[osmodels.EventType]func() {.
func (sw *SeaWatcher) ActiveSubscriptions() map[string]map[gbgrpc.EventType]func() {
	return sw.subscriptions
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

	// log.Printf("⚓️ received %s event: %+v", itemEventType, rawEvent)

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

		// sw.Prf("📢 %s #%s listed for %sΞ", style.BoldAlmostWhite(event.Payload.Item.Name), style.BoldAlmostWhite(event.Payload.Item.NftID.TokenID().String()), event.Payload.EventPayload.GetPrice())

		// sw.gb.GloomHub.Publish(channel.ItemListed, event)

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

		// sw.Prf("💦 %s #%s received bid: %sΞ", style.BoldAlmostWhite(event.Payload.Item.Metadata.Name), style.BoldAlmostWhite(event.Payload.Item.NftID.TokenID().String()), style.BoldAlmostWhite(fmt.Sprint(event.Payload.EventPayload.GetPrice())))

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

		// name := event.Payload.CollectionCriteria.Slug
		// if collection := sw.gb.CollectionDB.GetCollectionForSlug(event.Payload.CollectionCriteria.Slug); collection != nil {
		// 	name = collection.Name
		// }

		// // parse tokenPrice
		// var tokenPrice *price.Price
		// if event.Payload.BasePrice != nil {
		// 	tokenPrice = price.NewPrice(event.Payload.BasePrice)
		// } else {
		// 	tokenPrice = price.NewPrice(big.NewInt(0))

		// 	gbl.Log.Warnf("🤷‍♀️ error parsing tokenPrice: %+v", event.Payload.BasePrice)
		// }

		// sw.Prf("🦕 %s collection offer: %sΞ", style.BoldAlmostWhite(name), style.BoldAlmostWhite(fmt.Sprintf("%5.3f", tokenPrice.Ether())))

	case osmodels.ItemMetadataUpdated:
		var event *models.ItemMetadataUpdated

		decoderConfig.Result = &event
		decoder, _ := mapstructure.NewDecoder(decoderConfig)

		// gbl.Log.Info("\n\n")
		// gbl.Log.Infof(fmt.Sprintf("raw: %# v", pretty.Formatter(rawEvent)))
		// gbl.Log.Info("")

		err := decoder.Decode(rawEvent)
		if err != nil {
			log.Infof("⚓️❌ decoding incoming %s event failed: %s", style.Bold(itemEventType), err)

			return
		}

		// gbl.Log.Infof(fmt.Sprintf("event: %# v", pretty.Formatter(event)))
		// gbl.Log.Info("\n\n")

		// push to eventHub for further processing
		sw.gb.In.ItemMetadataUpdated <- event
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

func (sw *SeaWatcher) SubscribeForSlug(slug string, eventTypes []gbgrpc.EventType) bool {
	return sw.SubscribeForSlugs([]string{slug}, eventTypes)
}

func (sw *SeaWatcher) SubscribeForSlugs(slug []string, eventTypes []gbgrpc.EventType) bool {
	if !viper.GetBool("seawatcher.local") {
		log.Warn("⚓️ subscribe discarded - no local OpenSea clients")
		log.Warn("⚓️ TODO implement subscribe via grpc (and maybe pubsub)")

		return false
	}

	for _, slug := range slug {
		if sw.IsSubscribed(slug) {
			log.Debugf("⚓️ ☕️ already subscribed to OpenSea events for %s", slug)

			return false
		}

		sw.mu.Lock()

		sw.subscriptions[slug] = make(map[gbgrpc.EventType]func())

		for _, eventType := range eventTypes {
			sw.subscriptions[slug][eventType] = sw.on(eventType, slug, sw.eventHandler)

			time.Sleep(time.Millisecond * 37)
		}

		sw.mu.Unlock()

		if collection := sw.gb.CollectionDB.GetCollectionForSlug(slug); collection != nil {
			log.Debugf("⏮️ resetting stats for %s", slug)

			collection.ResetStats()
		}

		time.Sleep(time.Millisecond * 137)
	}

	return true
}

func (sw *SeaWatcher) UnubscribeForSlug(slug string, _ []gbgrpc.EventType) bool {
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

func (sw *SeaWatcher) on(eventType gbgrpc.EventType, collectionSlug string, eventHandler func(response any)) func() {
	topic := fmt.Sprintf("collection:%s", collectionSlug)
	evType := strings.ToLower(eventType.String())

	log.Debugf("Fetching channel %s", topic)
	channel := sw.getChannel(topic)

	log.Debugf("subscribing to %s events on %s", evType, topic)
	channel.On(evType, eventHandler)

	channel.OnClose(func(payload any) {
		sw.Prf("⚠️ Channel %s closed: %s", topic, payload)
	})

	log.Debugf("␚ subscribed to %s for %s", evType, collectionSlug)

	return func() {
		sw.Prf("Unsubscribing from %s events on %s", evType, topic)

		leave, err := channel.Leave()
		if err != nil {
			sw.Prf("channel.Leave err: %s", err)
		}

		leave.Receive("ok", func(_ any) {
			delete(sw.channels, collectionSlug)
			sw.Prf("Successfully left channel %s listening for %s", topic, evType)
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
		log.Infof("⚓️ received msg on channel %s: %s", msg.Channel, msg.Message)

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

	sw.Prf("📣 subscribed to %s", style.AlmostWhiteStyle.Render(internal.TopicSeaWatcherMgmt))
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

		var action func(slug string, eventTypes []gbgrpc.EventType) bool

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

			if action(slug, availableEventTypes) {
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

func (sw *SeaWatcher) Subscribe(context.Context, *gbgrpc.SubscriptionRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (sw *SeaWatcher) GetItemListedEvents(req *gbgrpc.SubscriptionRequest, stream *gbgrpc.Gloomberg_GetEventsServer) error { //nolint:nosnakecase
	// TODO: remove when handler for other event types are implemented
	req.EventTypes = availableEventTypes

	// sw.Prf("received subscription request for %s collections/slugs (%s types each)...", style.BoldAlmostWhite(fmt.Sprint(len(req.Collections))), style.BoldAlmostWhite(fmt.Sprint(len(req.EventTypes))))

	newEventSubscriptions := 0

	go func() {
		for _, slug := range req.Collections {
			sw.Prf("subscribing to %s...", slug)

			if sw.SubscribeForSlug(slug, req.EventTypes) {
				newEventSubscriptions++

				time.Sleep(337 * time.Millisecond)
			}
		}

		sw.Prf(
			"successfully subscribed to %s new collections/slugs | total subscribed collections: %s",
			style.AlmostWhiteStyle.Render(fmt.Sprint(newEventSubscriptions)),
			style.AlmostWhiteStyle.Render(fmt.Sprint(len(sw.ActiveSubscriptions()))),
		)
	}()

	for event := range sw.gb.SubscribeItemListed() {
		// transform *models.ItemListed event to ItemListed grpc message
		itemListed := &gbgrpc.ItemListed{
			EventType: gbgrpc.EventType(gbgrpc.EventType_value[event.EventType]), //nolint:nosnakecase
			SentAt:    &timestamppb.Timestamp{Seconds: event.SentAt.Unix()},

			Payload: &gbgrpc.ItemListed_ItemListedPayload{ //nolint:nosnakecase
				Item: &gbgrpc.ItemListed_Item{ //nolint:nosnakecase
					Chain:     &gbgrpc.ItemListed_Chain{Name: "ethereum"}, //nolint:nosnakecase
					NftId:     event.Payload.Item.String(),
					Permalink: event.Payload.Item.Permalink,
					Metadata: &gbgrpc.ItemListed_Metadata{ //nolint:nosnakecase
						Name:         event.Payload.Item.Name,
						ImageUrl:     event.Payload.Item.ImageURL,
						AnimationUrl: event.Payload.Item.AnimationURL,
						MetadataUrl:  event.Payload.Item.MetadataURL,
					},
				},
				BasePrice:      event.Payload.BasePrice.String(),
				Collection:     &gbgrpc.ItemListed_Collection{Slug: event.Payload.Slug}, //nolint:nosnakecase
				IsPrivate:      event.Payload.IsPrivate,
				ListingDate:    &timestamppb.Timestamp{Seconds: event.Payload.ListingDate.Unix()},
				EventTimestamp: &timestamppb.Timestamp{Seconds: event.Payload.EventTimestamp.Unix()},
				Quantity:       uint32(event.Payload.Quantity),
				Maker:          &gbgrpc.ItemListed_Account{Address: event.Payload.Maker.Address.String()}, //nolint:nosnakecase
				Taker:          &gbgrpc.ItemListed_Account{Address: event.Payload.Taker.Address.String()}, //nolint:nosnakecase
				ExpirationDate: &timestamppb.Timestamp{Seconds: event.Payload.ExpirationDate.Unix()},
				OrderHash:      event.Payload.OrderHash.String(),
				PaymentToken: &gbgrpc.ItemListed_PaymentToken{ //nolint:nosnakecase
					Address:  event.Payload.Address.String(),
					Symbol:   event.Payload.Symbol,
					Name:     event.Payload.Name,
					Decimals: uint32(event.Payload.Decimals),
					UsdPrice: event.Payload.UsdPrice,
				},
			},
		}

		ev := &gbgrpc.Event{
			EventType: gbgrpc.EventType_ITEM_LISTED, //nolint:nosnakecase
			Payload: &gbgrpc.EventPayload{
				Kind: &gbgrpc.EventPayload_ItemListed{ //nolint:nosnakecase
					ItemListed: itemListed,
				},
			},
		}

		if err := (*stream).Send(ev); err != nil {
			log.Printf("❌ error sending event to grpc client: %s", err)

			return err
			// continue
		}

		// output to terminal
		price := price.NewPrice(event.Payload.BasePrice)
		eventType := degendb.EventType(degendb.GetEventType(event.EventType))
		collectionPrimaryStyle := lipgloss.NewStyle().Foreground(style.GenerateColorWithSeed(event.Payload.Item.NftID.ContractAddress().Hash().Big().Int64()))
		collectionSecondaryStyle := lipgloss.NewStyle().Foreground(style.GenerateColorWithSeed(event.Payload.Item.NftID.ContractAddress().Big().Int64() ^ 2))
		currencySymbol := collectionPrimaryStyle.Bold(false).Render("Ξ")

		fmtPrice := style.BoldAlmostWhite(fmt.Sprintf("%5.2f", price.Ether())) + currencySymbol
		fmtItemName := strings.ReplaceAll(collectionPrimaryStyle.Bold(true).Render(event.Payload.Item.Name), "#", collectionSecondaryStyle.Render("#"))

		fmtItemLink := style.TerminalLink(event.Payload.Item.Permalink, fmtItemName)
		// fmtCollectionLink := style.TerminalLink(utils.GetOpenseaCollectionLink(event.Payload.Slug), style.LightGrayStyle.Render(fmt.Sprint(event.Payload.Slug)))

		sw.Prf("%s %s %s", eventType.Icon(), fmtPrice, fmtItemLink)
	}

	return nil
}

package seawa

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/osmodels"
	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/benleb/gloomberg/internal/seawa/models"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils/hooks"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mitchellh/mapstructure"
	"github.com/nshafer/phx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/redis/rueidis"
	"github.com/spf13/viper"
)

type SeaWatcher struct {
	// channel for events received from the opensea stream
	receivedEvents chan map[string]interface{}

	// phoenix channels client
	phoenixSocket *phx.Socket

	// subscribed phoenix channels
	channels map[string]*phx.Channel

	// subscribed slugs/events
	subscriptions map[osmodels.EventType]map[string]func()

	// redis client
	rdb rueidis.Client

	gb *gloomberg.Gloomberg

	mu *sync.Mutex
}

var (
	AvailableEventTypes = []osmodels.EventType{osmodels.ItemListed, osmodels.ItemReceivedBid, osmodels.CollectionOffer} // , osmodels.ItemMetadataUpdated} // ItemMetadataUpdated, ItemCancelled

	eventsReceivedTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "gloomberg_oswatcher_events_received_total",
		Help: "The total number of received events from the opensea api/stream",
	})
	// eventsProcessedTotal = promauto.NewCounter(prometheus.CounterOpts{
	// 	Name: "gloomberg_oswatcher_events_processed_total",
	// 	Help: "The total number of processed events from the opensea api/stream",
	// })
	// eventsIgnoredTotal = promauto.NewCounter(prometheus.CounterOpts{
	// 	Name: "gloomberg_oswatcher_events_ignored_total",
	// 	Help: "The total number of ignored events from the opensea api/stream",
	// }).
)

// func NewSeaWatcher(apiToken string, rdb rueidis.Client) *SeaWatcher {.
func NewSeaWatcher(apiToken string, gb *gloomberg.Gloomberg) *SeaWatcher {
	var socket *phx.Socket

	if apiToken != "" {
		endpointURL := fmt.Sprint(osmodels.StreamAPIEndpoint, "?token=", apiToken)

		endpoint, err := url.Parse(endpointURL)
		if err != nil {
			log.Info(err)

			return nil
		}

		// create phoenix socket
		socket = phx.NewSocket(endpoint)
		socket.Logger = phx.NewSimpleLogger(phx.LoggerLevel(phx.LogError))

		socket.ReconnectAfterFunc = func(attempt int) time.Duration {
			log.Warnf("⚓️❕ opensea stream socket retry after %v..", time.Duration(attempt)*2*time.Second)

			return time.Duration(attempt) * 2 * time.Second
		}

		// error function
		onError := func(err error) { log.Info(err) }
		socket.OnError(onError)

		socket.OnClose(func() {
			log.Info("⚓️❕ opensea stream socket closed, reconnecting...")

			err := socket.Reconnect()
			if err != nil {
				onError(errors.New("opensea stream socket reconnecting failed: " + err.Error()))
			}
		})
	}

	client := &SeaWatcher{
		receivedEvents: make(chan map[string]interface{}, 1024),
		subscriptions:  make(map[osmodels.EventType]map[string]func(), 0),

		phoenixSocket: socket,
		channels:      make(map[string]*phx.Channel),

		gb: gb,
		// rdb: rdb,
		rdb: gb.Rdb,

		mu: &sync.Mutex{},
	}

	// create subscriptions map/registry
	for _, event := range AvailableEventTypes {
		client.subscriptions[event] = make(map[string]func(), 0)
	}

	if client.phoenixSocket != nil {
		if err := client.connect(); err != nil {
			socketError := errors.New("opensea stream socket connection failed: " + err.Error())
			log.Error("⚓️❌ " + socketError.Error())

			return nil
		}
	}

	return client
}

func (sw *SeaWatcher) EventChannel() chan map[string]interface{} {
	return sw.receivedEvents
}

func (sw *SeaWatcher) ActiveSubscriptions() map[osmodels.EventType]map[string]func() {
	return sw.subscriptions
}

func (sw *SeaWatcher) connect() error {
	return sw.phoenixSocket.Connect()
}

// eventHandler handles incoming stream api events and forwards them as map.
func (sw *SeaWatcher) eventHandler(response any) {
	eventsReceivedTotal.Inc()

	itemEvent, ok := response.(map[string]interface{})
	if !ok {
		log.Errorf("⚓️❌ error in type assertion of received event: %+v", response)

		return
	}

	log.Debugf("⚓️ received event: %+v", itemEvent)

	itemEventType, ok := itemEvent["event_type"].(string)
	if !ok {
		log.Warnf("⚓️ 🤷‍♀️ unknown event type: %s", itemEvent["event_type"])

		return
	}

	switch osmodels.EventType(itemEventType) {
	case osmodels.ItemSold:
		log.Debugf("⚓️ received SOLD %s: %+v", itemEventType, itemEvent)

	case osmodels.ItemReceivedBid:
		log.Debugf("⚓️ received BID %s: %+v", itemEventType, itemEvent)
		// sw.publishEvent(itemEvent)

	case osmodels.ItemMetadataUpdated:
		log.Debugf("⚓️ received METADATA %s: %+v", itemEventType, itemEvent)
		// sw.publishEvent(itemEvent)

	case osmodels.CollectionOffer:
		log.Debugf("⚓️ received COLLECTIONOFFER %s: %+v", itemEventType, itemEvent)

		collectionOfferEvent, err := sw.DecodeCollectionOfferEvent(itemEvent)
		if err != nil {
			return
		}

		// print collectionOfferEvent
		log.Debugf("⚓️ received COLLECTIONOFFER %s: %+v", itemEventType, collectionOfferEvent)

		// sw.publishEvent(collectionOfferEvent)

		// priceWeiRaw, _, err := big.ParseFloat(collectionOfferEvent.Payload.BasePrice, 10, 64, big.ToNearestEven)
		// if err != nil {
		//	log.Infof("⚓️❌ error parsing price: %s", err.Error())
	//		return
	//		}
	// priceWei, _ := priceWeiRaw.Int(nil)

	// eventType := osmodels.TxType[osmodels.EventType(itemEventType)]

	// collectionSlug := collectionOfferEvent.Payload.Collection.Slug

	// paymentTokenSymbol := collectionOfferEvent.Payload.PaymentToken.Symbol

	// quantity := collectionOfferEvent.Payload.Quantity

	// pricePerToken := priceWei.Div(priceWei, big.NewInt(int64(quantity)))

	// log.Infof("⚓️ 🔸 %s | %dx %s %s for %s", eventType.Icon(), quantity, style.TrendRedStyle.Render(fmt.Sprintf("%5.3f", price.NewPrice(pricePerToken).Ether())), paymentTokenSymbol, style.BoldStyle.Render(collectionSlug))

	case osmodels.ItemReceivedOffer:
		log.Debugf("⚓️ offer received %s: %+v", itemEventType, itemEvent)

		var itemReceivedOfferEvent osmodels.ItemReceivedOfferEvent

		err := mapstructure.Decode(itemEvent, &itemReceivedOfferEvent)
		if err != nil {
			log.Info("⚓️❌ decoding incoming opensea stream api ItemReceivedOffer event failed:", err)

			return
		}

		// sw.publishEvent(itemReceivedOfferEvent)

		// print itemReceivedOfferEvent
		log.Debugf("⚓️ received %s: %+v", itemEventType, itemReceivedOfferEvent)
		printItemReceivedOfferEvent(itemReceivedOfferEvent)

	case osmodels.ItemListed:
		var itemListedEvent osmodels.ItemListedEvent

		err := mapstructure.Decode(itemEvent, &itemListedEvent)
		if err != nil {
			log.Info("⚓️❌ decoding incoming opensea stream api ItemListed event failed:", err)

			return
		}

		// sw.publishItemListedToRedis(itemListedEvent)
		sw.gb.In.ItemListedEvents <- &itemListedEvent
	}

	sw.receivedEvents <- itemEvent
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
		log.Info("⚓️❌ decoding incoming opensea stream api ItemReceivedBidEvent failed:", err)

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
		log.Info("⚓️❌ decoding incoming opensea stream api collection offer event failed:", err)

		return osmodels.CollectionOfferEvent{}, err
	}

	return collectionOfferEvent, err
}

func printItemReceivedOfferEvent(itemOfferEvent osmodels.ItemReceivedOfferEvent) {
	priceWeiRaw, _, err := big.ParseFloat(itemOfferEvent.Payload.BasePrice, 10, 64, big.ToNearestEven)
	if err != nil {
		log.Infof("⚓️❌ error parsing price: %s", err.Error())

		return
	}
	priceWei, _ := priceWeiRaw.Int(nil)

	// nftID is a identification string in the format <chain>/<contract>/<tokenID>
	nftID := strings.Split(itemOfferEvent.Payload.Item.NftID, "/")
	if len(nftID) != 3 {
		log.Infof("⚓️ 🤷‍♀️ error parsing nftID: %s", itemOfferEvent.Payload.Item.NftID)

		return
	}
	eventType := osmodels.TxType[itemOfferEvent.StreamEvent]
	collectionPrimaryStyle := lipgloss.NewStyle().Foreground(style.GenerateColorWithSeed(common.HexToAddress(nftID[1]).Hash().Big().Int64())).Bold(true)
	collectionSecondaryStyle := lipgloss.NewStyle().Foreground(style.GenerateColorWithSeed(common.HexToAddress(nftID[1]).Big().Int64() ^ 2)).Bold(true)
	// get tokenID
	tID, _, _ := big.ParseFloat(nftID[2], 10, 64, big.ToNearestEven)
	tokenID, _ := tID.Int(nil)
	fmtTokenID := style.ShortenedTokenIDStyled(tokenID, collectionPrimaryStyle, collectionSecondaryStyle)
	// for erc1155 tokens itemOfferEvent.Payload.Item.Metadata.Name is the item name
	collectionName := strings.Split(itemOfferEvent.Payload.Item.Metadata.Name, " #")[0]
	fmtToken := style.BoldStyle.Render(fmt.Sprintf("%s %s", collectionPrimaryStyle.Render(collectionName), fmtTokenID))
	fmt.Println(itemOfferEvent.StreamEvent)
	log.Infof("⚓️ %s | %sΞ  %s ", eventType.Icon(), style.BoldStyle.Render(fmt.Sprintf("%5.3f", price.NewPrice(priceWei).Ether())), style.TerminalLink(itemOfferEvent.Payload.Item.Permalink, fmtToken))
}

func (sw *SeaWatcher) SubscribeForSlug(eventType osmodels.EventType, slug string) bool {
	sw.mu.Lock()
	alreadySubscribed := sw.subscriptions[eventType][slug]

	if alreadySubscribed != nil {
		sw.mu.Unlock()

		log.Debugf("⚓️ ☕️ already subscribed to %s for %s", eventType, slug)

		return false
	}

	sw.subscriptions[eventType][slug] = sw.on(eventType, slug, sw.eventHandler)
	sw.mu.Unlock()

	return true
}

func (sw *SeaWatcher) UnubscribeForSlug(eventType osmodels.EventType, slug string) bool {
	sw.mu.Lock()
	unsubscribe := sw.subscriptions[eventType][slug]
	sw.mu.Unlock()

	if unsubscribe != nil {
		// unsubscribe
		unsubscribe()

		// remove slug
		sw.mu.Lock()
		sw.subscriptions[eventType][slug] = nil
		sw.mu.Unlock()

		return true
	}

	log.Debugf("☕️ not subscribed to %s for %s (anymore)", eventType, slug)

	return false
}

func (sw *SeaWatcher) createChannel(topic string) *phx.Channel {
	channel := sw.phoenixSocket.Channel(topic, nil)

	join, err := channel.Join()
	if err != nil {
		log.Info(err)

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

	log.Debugf("␚ 🔔 subscribed to %s for %s", string(eventType), collectionSlug)

	return func() {
		log.Infof("Unsubscribing from %s events on %s", eventType, topic)

		leave, err := channel.Leave()
		if err != nil {
			log.Info("channel.Leave err:", err)
		}

		leave.Receive("ok", func(_ any) {
			delete(sw.channels, collectionSlug)
			log.Infof("Successfully left channel %s listening for %s", topic, eventType)
		})
	}
}

// Run starts the seawatcher by subscribing to the mgmt channel and listening for new slugs to subscribe to.
func (sw *SeaWatcher) Run() {
	// subscribe to mgmt channel
	log.Infof("⚓️ ␚ subscribing to mgmt channel %s", internal.TopicSeaWatcherMgmt)

	err := sw.rdb.Receive(context.Background(), sw.rdb.B().Subscribe().Channel(internal.TopicSeaWatcherMgmt).Build(), func(msg rueidis.PubSubMessage) {
		log.Debugf("⚓️ received msg on channel %s: %s", msg.Channel, msg.Message)

		var mgmtEvent *models.MgmtEvent

		if err := json.Unmarshal([]byte(msg.Message), &mgmtEvent); err != nil {
			log.Errorf("⚓️❌ error json.Unmarshal: %+v", err)

			return
		}

		switch mgmtEvent.Action {
		case models.SendSlugs:
			// SendSlugs can be ignored on server side for now
			return

		case models.Subscribe, models.Unsubscribe:
			log.Infof("⚓️ ␚ received %s for %s collections/slugs on %s, subscribing...", style.BoldStyle.Render(mgmtEvent.Action.String()), style.BoldStyle.Render(fmt.Sprint(len(mgmtEvent.Slugs))), internal.TopicSeaWatcherMgmt)
			if len(mgmtEvent.Slugs) == 0 {
				log.Error("⚓️❌ incoming collection slugs msg is empty")

				return
			}

			if viper.GetString("api_keys.opensea") == "" {
				log.Error("⚓️❌ opensea api key is not set, can't subscribe to listings")

				return
			}

			var action func(event osmodels.EventType, slug string) bool

			switch mgmtEvent.Action {
			case models.Subscribe:
				action = sw.SubscribeForSlug
			case models.Unsubscribe:
				action = sw.UnubscribeForSlug
			}

			// transform to string
			var events []string
			for _, event := range AvailableEventTypes {
				events = append(events, string(event))
			}

			// subscribe to which events?
			if len(mgmtEvent.Events) == 0 {
				// subscribe to all available events if none are specified
				log.Infof("⚓️ ␚ no events specified, subscribing to all available events (%+v)", strings.Join(events, ", "))
				// style.BoldStyle.Render(fmt.Sprint(len(mgmtEvent.Slugs)))
				mgmtEvent.Events = AvailableEventTypes
			}

			newSubscriptions := make(map[string][]osmodels.EventType, 0)
			newEventSubscriptions := 0

			for _, slug := range mgmtEvent.Slugs {
				if slug == "ens" {
					log.Info("⚓️ ␚ skipping ens for now")

					continue
				}

				for _, event := range mgmtEvent.Events {
					if action(event, slug) {
						newEventSubscriptions++

						if _, ok := newSubscriptions[slug]; !ok {
							newSubscriptions[slug] = make([]osmodels.EventType, 0)
						}

						newSubscriptions[slug] = append(newSubscriptions[slug], event)

						time.Sleep(137 * time.Millisecond)
					}
				}
			}

			log.Infof(
				"⚓️ ␚ successfully subscribed to %s new collections/slugs (%d events in total) | total subscriptions: %s",
				style.BoldStyle.Render(fmt.Sprint(len(newSubscriptions))),
				newEventSubscriptions,
				style.BoldStyle.Render(fmt.Sprint(len(sw.ActiveSubscriptions()[osmodels.ItemListed]))),
			)

		default:
			log.Infof("⚓️ 👀 received unknown mgmt event: %s", mgmtEvent.Action.String())

			return
		}
	})
	if err != nil {
		log.Errorf("❌ error subscribing to redis channels %s: %s", internal.TopicSeaWatcherMgmt, err.Error())

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
		log.Infof("⚓️ 📣 published %s event to %s", style.Bold(sendSlugsEvent.Action.String()), style.Bold(internal.TopicSeaWatcherMgmt))
	}
}

// func printItemListed(itemListedEvent osmodels.ItemListedEvent) {
// 	// parse price
// 	priceWeiRaw, _, err := big.ParseFloat(itemListedEvent.Payload.BasePrice, 10, 64, big.ToNearestEven)
// 	if err != nil {
// 		log.Infof("⚓️❌ error parsing price: %s", err.Error())

// 		return
// 	}
// 	priceWei, _ := priceWeiRaw.Int(nil)

// 	var listedBy string

// 	listedByAddress := common.HexToAddress(itemListedEvent.Payload.Maker.Address)
// 	listedByStyle := lipgloss.NewStyle().Foreground(style.GenerateColorWithSeed(listedByAddress.Big().Int64()))

// 	if itemListedEvent.Payload.Maker.User != "" {
// 		listedBy = listedByStyle.Render(itemListedEvent.Payload.Maker.User) + " (" + style.ShortenAddressStyled(&listedByAddress, listedByStyle) + ")"
// 	} else {
// 		listedBy = style.ShortenAddressStyled(&listedByAddress, listedByStyle)
// 	}

// 	// nftID is a identification string in the format <chain>/<contract>/<tokenID>
// 	nftID := strings.Split(itemListedEvent.Payload.Item.NftID, "/")
// 	if len(nftID) != 3 {
// 		log.Infof("⚓️ 🤷‍♀️ error parsing nftID: %s", itemListedEvent.Payload.Item.NftID)

// 		return
// 	}

// 	// collection styles
// 	collectionPrimaryStyle := lipgloss.NewStyle().Foreground(style.GenerateColorWithSeed(common.HexToAddress(nftID[1]).Hash().Big().Int64())).Bold(true)
// 	collectionSecondaryStyle := lipgloss.NewStyle().Foreground(style.GenerateColorWithSeed(common.HexToAddress(nftID[1]).Big().Int64() ^ 2)).Bold(true)

// 	// strip tokenID from collection name
// 	collectionName := strings.Split(itemListedEvent.Payload.Item.Metadata.Name, " #")[0]

// 	// get tokenID
// 	tID, _, _ := big.ParseFloat(nftID[2], 10, 64, big.ToNearestEven)
// 	tokenID, _ := tID.Int(nil)

// 	// format tokenID
// 	fmtTokenID := style.ShortenedTokenIDStyled(tokenID, collectionPrimaryStyle, collectionSecondaryStyle)
// 	eventType := osmodels.TxType[itemListedEvent.StreamEvent]

// 	fmtToken := style.BoldStyle.Render(fmt.Sprintf("%s %s", collectionPrimaryStyle.Render(collectionName), fmtTokenID))

// 	log.Infof("⚓️ %s | %sΞ  %s  | %s", eventType.Icon(), style.BoldStyle.Render(fmt.Sprintf("%5.3f", price.NewPrice(priceWei).Ether())), style.TerminalLink(itemListedEvent.Payload.Item.Permalink, fmtToken), listedBy)
// }

// func (sw *SeaWatcher) publishItemListedToRedis(itemListedEvent osmodels.ItemListedEvent) {
// 	// marshal event to json
// 	jsonEvent, err := json.Marshal(itemListedEvent)
// 	if err != nil {
// 		log.Info("⚓️❌ json.Marshal failed for incoming stream api event", err)

// 		return
// 	}

// 	// generate the redis pubsub channel name
// 	channel := internal.TopicSeaWatcher + "/" + itemListedEvent.ContractAddress().String() + "/" + string(itemListedEvent.StreamEvent)

// 	// publish event to redis
// 	if sw.rdb.Do(context.Background(), sw.rdb.B().Publish().Channel(channel).Message(string(jsonEvent)).Build()).Error() != nil {
// 		log.Warnf("⚓️❕ error publishing event to redis: %s", err.Error())
// 	}

// 	// this is used to log the event to stdout in "headless" mode,
// 	// not to print to the terminal in "live" mode
// 	printItemListed(itemListedEvent)
// }

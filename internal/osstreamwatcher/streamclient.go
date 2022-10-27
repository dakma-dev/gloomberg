package osstreamwatcher

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/mitchellh/mapstructure"
	"github.com/nshafer/phx"
)

type StreamEventType string

const (
	Listed          StreamEventType = "item_listed"
	ReceivedOffer   StreamEventType = "item_received_offer"
	ReceivedBid     StreamEventType = "item_received_bid"
	MetadataUpdated StreamEventType = "item_metadata_updated"

	// Sold            StreamEventType = "item_sold"
	// Transferred     StreamEventType = "item_transferred"
	// Cancelled       StreamEventType = "item_cancelle".
)

const StreamAPIEndpoint = "wss://stream.openseabeta.com/socket"

type OSStreamWatcher struct {
	QueueListings chan *models.ItemListedEvent
	socket        *phx.Socket
	channels      map[string]*phx.Channel
	// eventHandler  func(event any)
}

var osStreamWatcher *OSStreamWatcher

func NewStreamWatcher(token string, onError func(error)) *OSStreamWatcher {
	if osStreamWatcher != nil {
		return osStreamWatcher
	}

	endpointURL := fmt.Sprint(StreamAPIEndpoint, "?token=", token)

	endpoint, err := url.Parse(endpointURL)
	if err != nil {
		gbl.Log.Error(err)

		return nil
	}

	socket := phx.NewSocket(endpoint)
	socket.Logger = phx.NewSimpleLogger(phx.LoggerLevel(phx.LogError))

	socket.ReconnectAfterFunc = func(attempt int) time.Duration {
		return time.Duration(attempt) * 2 * time.Second
	}

	if onError == nil {
		onError = func(err error) { gbl.Log.Error(err) }
	}

	socket.OnError(onError)

	socket.OnClose(func() {
		gbl.Log.Warnf("opensea stream socket closed, reconnecting...")

		err := socket.Reconnect()
		if err != nil {
			onError(errors.New("opensea stream socket reconnecting failed: " + err.Error()))
		}
	})

	client := &OSStreamWatcher{
		QueueListings: make(chan *models.ItemListedEvent, 1024),
		socket:        socket,
		channels:      make(map[string]*phx.Channel),
	}

	if err := client.Connect(); err != nil {
		gbl.Log.Error(errors.New("opensea stream socket connection failed: " + err.Error()))
		return nil
	}

	osStreamWatcher = client

	return client
}

func (s *OSStreamWatcher) Connect() error {
	return s.socket.Connect()
}

func (s *OSStreamWatcher) Disconnect() error {
	// s.socket.OnError()
	gbl.Log.Info("Successfully disconnected from socket")

	s.channels = make(map[string]*phx.Channel)

	return s.socket.Disconnect()
}

func (s *OSStreamWatcher) createChannel(topic string) *phx.Channel {
	channel := s.socket.Channel(topic, nil)

	join, err := channel.Join()
	if err != nil {
		gbl.Log.Error(err)

		return nil
	}

	join.Receive("ok", func(response any) {
		gbl.Log.Debugf("joined channel: %s", channel.Topic()) // ), response)
	})

	join.Receive("error", func(response any) {
		gbl.Log.Error("failed to joined channel:", channel.Topic(), response)
	})

	s.channels[topic] = channel

	return channel
}

func (s *OSStreamWatcher) getChannel(topic string) *phx.Channel {
	channel, ok := s.channels[topic]
	if !ok {
		channel = s.createChannel(topic)
	}

	return channel
}

func (s *OSStreamWatcher) on(eventType StreamEventType, collectionSlug string, eventHandler func(payload any)) func() {
	topic := fmt.Sprintf("collection:%s", collectionSlug)

	gbl.Log.Debugf("Fetching channel %s", topic)
	channel := s.getChannel(topic)

	gbl.Log.Debugf("Subscribing to %s events on %s", eventType, topic)
	channel.On(string(eventType), eventHandler)

	return func() {
		gbl.Log.Infof("Unsubscribing from %s events on %s", eventType, topic)

		leave, err := channel.Leave()
		if err != nil {
			gbl.Log.Error("channel.Leave err:", err)
		}

		leave.Receive("ok", func(response any) {
			delete(s.channels, collectionSlug)
			gbl.Log.Infof("Successfully left channel %s listening for %s", topic, eventType)
		})
	}
}

func (s *OSStreamWatcher) OnItemListed(collectionSlug string, eventHandler func(itemListedEvent any)) {
	if eventHandler == nil {
		eventHandler = s.handlerListing
	}

	s.on(Listed, collectionSlug, eventHandler)
}

// func (s StreamClient) OnItemSold(collectionSlug string, eventHandler func(itemSoldEvent any)) {
// 	s.on(Sold, collectionSlug, eventHandler)
// }

// func (s StreamClient) OnItemTransferred(collectionSlug string, eventHandler func(itemTransferredEvent any)) {
// 	s.on(Transferred, collectionSlug, eventHandler)
// }

// func (s StreamClient) OnItemCancelled(collectionSlug string, eventHandler func(itemCancelledEvent any)) {
// 	s.on(Cancelled, collectionSlug, eventHandler)
// }

func (s *OSStreamWatcher) OnItemReceivedBid(collectionSlug string, eventHandler func(itemReceivedBidEvent any)) {
	s.on(ReceivedBid, collectionSlug, eventHandler)
}

func (s *OSStreamWatcher) OnItemReceivedOffer(collectionSlug string, eventHandler func(itemReceivedOfferEvent any)) {
	s.on(ReceivedOffer, collectionSlug, eventHandler)
}

func (s *OSStreamWatcher) OnItemMetadataUpdated(collectionSlug string, eventHandler func(itemMetadataUpdatedEvent any)) {
	s.on(MetadataUpdated, collectionSlug, eventHandler)
}

func (s *OSStreamWatcher) SubscribeToListingsFor(slug string) {
	s.OnItemListed(slug, s.handlerListing)
	gbl.Log.Debugf("subscribed to listings for: %s", slug)
}

func (s *OSStreamWatcher) handlerListing(response any) {
	var itemListedEvent models.ItemListedEvent

	err := mapstructure.Decode(response, &itemListedEvent)
	if err != nil {
		gbl.Log.Error("mapstructure.Decode failed for incoming stream api event", err)
		return
	}

	gbl.Log.Debugf("received event from opensea: %+v", itemListedEvent.BaseStreamMessage.StreamEvent)

	s.QueueListings <- &itemListedEvent
}

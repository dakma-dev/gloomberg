package opensea

import (
	"fmt"
	"net/url"
	"time"

	"github.com/benleb/gloomberg/internal/utils/gbl"
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

type StreamClient struct {
	socket       *phx.Socket
	channels     map[string]*phx.Channel
	eventHandler func(event any)
}

var streamClient *StreamClient

func NewStreamClient(token string, onError func(error)) *StreamClient {
	if streamClient != nil {
		return streamClient
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

	socket.OnError(onError)
	socket.OnClose(func() {
		err := socket.Reconnect()
		if err != nil {
			onError(err)
		}
	})

	client := &StreamClient{
		socket:   socket,
		channels: make(map[string]*phx.Channel),
	}
	if err := client.Connect(); err != nil {
		gbl.Log.Error(err)

		return nil
	}

	streamClient = client

	return client
}

func (s *StreamClient) Connect() error {
	return s.socket.Connect()
}

func (s *StreamClient) Disconnect() error {
	// s.socket.OnError()
	gbl.Log.Info("Successfully disconnected from socket")

	s.channels = make(map[string]*phx.Channel)

	return s.socket.Disconnect()
}

func (s *StreamClient) createChannel(topic string) *phx.Channel {
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

func (s *StreamClient) getChannel(topic string) *phx.Channel {
	channel, ok := s.channels[topic]
	if !ok {
		channel = s.createChannel(topic)
	}

	return channel
}

func (s *StreamClient) on(eventType StreamEventType, collectionSlug string, eventHandler func(payload any)) func() {
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

func (s *StreamClient) OnItemListed(collectionSlug string, eventHandler func(itemListedEvent any)) {
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

func (s *StreamClient) OnItemReceivedBid(collectionSlug string, eventHandler func(itemReceivedBidEvent any)) {
	s.on(ReceivedBid, collectionSlug, eventHandler)
}

func (s *StreamClient) OnItemReceivedOffer(collectionSlug string, eventHandler func(itemReceivedOfferEvent any)) {
	s.on(ReceivedOffer, collectionSlug, eventHandler)
}

func (s *StreamClient) OnItemMetadataUpdated(collectionSlug string, eventHandler func(itemMetadataUpdatedEvent any)) {
	s.on(MetadataUpdated, collectionSlug, eventHandler)
}

// func SubscribeToCollectionSlugs(apiToken string, slugs []string, eventHandler func(itemListedEvent any)) {
//	if client := NewStreamClient(apiToken, func(err error) {
//		gbl.Log.Error(err)
//	}); client != nil {
//		for _, slug := range slugs {
//			SubscribeToCollectionSlug(client, slug, eventHandler)
//		}
//	}
//}

func SubscribeToListingsForCollectionSlug(client *StreamClient, slug string, eventHandler func(itemListedEvent any)) {
	gbl.Log.Debugf("client %+v | streamClient: %+v\n", client, streamClient)

	if client != nil || streamClient != nil {
		if client == nil {
			client = streamClient
		}

		gbl.Log.Debugf("eventHandler %p | client.eventHandler: %p\n", eventHandler, client.eventHandler)

		if eventHandler != nil || client.eventHandler != nil {
			if eventHandler == nil {
				eventHandler = client.eventHandler
			} else {
				client.eventHandler = eventHandler
			}

			client.OnItemListed(slug, eventHandler)
			// client.OnItemReceivedOffer(slug, eventHandler)

			gbl.Log.Debugf("subscribed to listings for: %s", slug)
		}
	}
}

//func SubscribeToEverythingForCollectionSlug(client *StreamClient, slug string, eventHandler func(itemListedEvent any)) {
//	gbl.Log.Debugf("client %+v | streamClient: %+v\n", client, streamClient)
//
//	if client != nil || streamClient != nil {
//		if client == nil {
//			client = streamClient
//		}
//
//		gbl.Log.Debugf("eventHandler %p | client.eventHandler: %p\n", eventHandler, client.eventHandler)
//
//		if eventHandler != nil || client.eventHandler != nil {
//			if eventHandler == nil {
//				eventHandler = client.eventHandler
//			} else {
//				client.eventHandler = eventHandler
//			}
//
//			client.OnItemReceivedBid(slug, eventHandler)
//			gbl.Log.Infof("subscribed to bids for: %s", slug)
//
//			client.OnItemReceivedOffer(slug, eventHandler)
//			gbl.Log.Infof("subscribed to offers for: %s", slug)
//
//			client.OnItemMetadataUpdated(slug, eventHandler)
//			gbl.Log.Infof("subscribed to metadata updates for: %s", slug)
//		}
//	}
//}

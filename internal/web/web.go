package web

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"math/big"
	"net/http"

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/jfyne/live"
)

const (
	wsend       = "send"
	wnewmessage = "newmessage"
)

type EventMessage struct {
	ID             string // Unique ID per message so that we can use `live-update`.
	User           string
	Msg            string
	Time           string
	Typemoji       string
	Price          string
	PricePerItem   float64
	CollectionName string
	TokenID        *big.Int
	ColorPrimary   string
	ColorSecondary string
	Event          *collections.Event
	SalesCount     uint64
	ListingsCount  uint64
	SaLiRa         float64
}

func NewEvent(data interface{}) EventMessage {
	// This can handle both the chat example, and the cluster example.
	switch m := data.(type) {
	case EventMessage:
		return m
	}

	return EventMessage{}
}

type EventStream struct {
	Events      []EventMessage
	queueOutWeb *chan *collections.Event
}

func New(queueWeb *chan *collections.Event) *EventStream {
	return &EventStream{
		queueOutWeb: queueWeb,
	}
}

func (es *EventStream) Start() {
	http.Handle("/", live.NewHttpHandler(live.NewCookieStore("session-name", []byte("ZWh0NGkzdHZxNjY2NjZxNDg1NWJwdjk0NmM1YnA5MkM2NQ")), es.NewEventHandler()))
	http.Handle("/live.js", live.Javascript{})
	http.Handle("/auto.js.map", live.JavascriptMap{})

	gbl.Log.Infof("starting http server...")

	if err := http.ListenAndServe(":42069", nil); err != nil {
		fmt.Printf("error: %s", err)
		gbl.Log.Error(err)
	}
}

func (es *EventStream) NewEventstreamInstance(s live.Socket) *EventStream {
	go func() {
		for event := range *es.queueOutWeb {
			gbl.Log.Debugf("outWeb event: %+v", event)

			if !event.PrintEvent {
				gbl.Log.Infof("outWeb discarded event: %+v", event)
				continue
			}

			data := EventMessage{
				ID: live.NewID(),
				// User: live.SessionID(s.Session()),
				Time:           event.Time.Format("15:04:05"),
				Typemoji:       event.EventType.Icon(),
				Price:          fmt.Sprintf("%6.3f", event.PriceEther),
				PricePerItem:   event.PriceEtherPerItem,
				CollectionName: event.Collection.Name,
				TokenID:        event.TokenID,
				ColorPrimary:   string(event.Collection.Colors.Primary),
				ColorSecondary: string(event.Collection.Colors.Secondary),
				Event:          event,
				SalesCount:     event.Collection.Counters.Sales,
				ListingsCount:  event.Collection.Counters.Listings,
				SaLiRa:         event.Collection.SaLiRa.Value(),
			}

			if err := s.Broadcast(wnewmessage, data); err != nil {
				gbl.Log.Errorf("failed braodcasting new message: %w", err)
			}
		}
	}()

	m, ok := s.Assigns().(*EventStream)

	if !ok {
		return &EventStream{
			Events: []EventMessage{},
			// 	{ID: live.NewID(), User: "Muh", Msg: "Welcome to chat " + live.SessionID(s.Session())},
			// },
		}
	}

	m.Events = []EventMessage{}

	return m
}

// func NewEventstreamInstance(s live.Socket) *EventStream {
// 	m, ok := s.Assigns().(*EventStream)
// 	if !ok {
// 		return &EventStream{
// 			Events: []EventMessage{
// 				{ID: live.NewID(), User: "Muh", Msg: "Welcome to chat " + live.SessionID(s.Session())},
// 			},
// 		}
// 	}

// 	return m
// }

func (es *EventStream) NewEventHandler() live.Handler {
	t, err := template.ParseFiles("www/layout.html", "www/style.html", "www/view.html")
	if err != nil {
		log.Fatal(err)
	}

	handler := live.NewHandler(live.WithTemplateRenderer(t))

	// Set the mount function for this handler.
	handler.HandleMount(func(ctx context.Context, s live.Socket) (interface{}, error) {
		gbl.Log.Debugf("handle mount socket: %+v", s)
		// This will initialise the chat for this socket.
		return es.NewEventstreamInstance(s), nil
	})

	// Handle user sending a message.
	handler.HandleEvent(wsend, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		gbl.Log.Debugf("handle event params: %+v", p)

		m := es.NewEventstreamInstance(s)
		msg := p.String("message")
		if msg == "" {
			return m, nil
		}
		data := EventMessage{
			ID:   live.NewID(),
			User: live.SessionID(s.Session()),
			Msg:  msg,
		}

		if err := s.Broadcast(wnewmessage, data); err != nil {
			return m, fmt.Errorf("failed braodcasting new message: %w", err)
		}
		return m, nil
	})

	// Handle the broadcasted events.
	handler.HandleSelf(wnewmessage, func(ctx context.Context, s live.Socket, data interface{}) (interface{}, error) {
		gbl.Log.Debugf("handle self data: %+v", data)

		m := es.NewEventstreamInstance(s)

		// Here we don't append to messages as we don't want to use
		// loads of memory. `live-update="append"` handles the appending
		// of messages in the DOM.
		m.Events = []EventMessage{NewEvent(data)}
		return m, nil
	})

	gbl.Log.Debugf("handler created: %+v", handler)

	return handler
}

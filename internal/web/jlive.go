package web

import (
	"context"
	"fmt"
	"html/template"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal/cache"
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/nodes"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/charmbracelet/lipgloss"
	"github.com/jfyne/live"
)

const (
	wsend        = "send"
	wnewmessage  = "newmessage"
	msgGasUpdate = "gasupdate"
)

var templateFiles = []string{
	"www/layout.html",
	"www/style.html",
	"www/view.html",
}

type GasInfoMessage struct {
	PriceGwei int
	TipGwei   float64
}

type EventMessage struct {
	ID                  string // Unique ID per message so that we can use `live-update`.
	User                string
	Msg                 string
	Time                string
	Typemoji            string
	NumItems            uint64
	Price               string
	PriceArrowColor     string
	PricePerItem        string
	FloorPrice          string
	PricePerItemColor   lipgloss.Color
	NumItemsColor       lipgloss.Color
	TrendIndicator      string
	TrendIndicatorColor lipgloss.Color
	CollectionName      string
	TokenID             *big.Int
	ColorPrimary        string
	ColorPrimaryFaint   string
	ColorSecondary      string
	To                  string
	ToColor             string
	Event               *collections.Event
	SalesCount          uint64
	ListingsCount       uint64
	SaLiRa              float64
	LinkOpenSea         string
	LinkEtherscan       string
	GasInfo             GasInfoMessage
	EventType           string
	Divider             bool
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
	ID            string
	ListenAddress string
	ctx           context.Context
	Instances     []ESInstance
	GasPrice      int
	GetGasPrice   func() int
	queueOutWeb   *chan *collections.Event
	Nodes         *nodes.Nodes
}

type ESInstance struct {
	Events   []EventMessage
	GasPrice int
}

func New(queueWeb *chan *collections.Event, listenAddress string, nodes *nodes.Nodes, getGasPrice func() int) *EventStream {
	ctx := context.Background()

	eventStream := &EventStream{
		ID:            live.NewID(),
		ListenAddress: listenAddress,
		Instances:     make([]ESInstance, 0),
		ctx:           ctx,
		queueOutWeb:   queueWeb,
		Nodes:         nodes,
		GetGasPrice:   getGasPrice,
	}

	go func() {
		ticker := time.NewTicker(37 * time.Second)
		for range ticker.C {
			eventStream.GasPrice = getGasPrice()
			gbl.Log.Debugf("updated gas price in %p to %d", eventStream, eventStream.GasPrice)
		}
	}()

	return eventStream
}

func (es *EventStream) Start() {
	http.Handle("/", live.NewHttpHandler(live.NewCookieStore("session-name", []byte("ZWh0NGkzdHZxNjY2NjZxNDg1NWJwdjk0NmM1YnA5MkM2NQ")), es.NewEventHandler()))
	http.Handle("/live.js", live.Javascript{})
	http.Handle("/auto.js.map", live.JavascriptMap{})

	gbl.Log.Infof("starting http server...")

	if err := http.ListenAndServe(es.ListenAddress, nil); err != nil {
		fmt.Printf("error: %s", err)
		gbl.Log.Error(err)
	}
}

func (es *EventStream) NewEventstreamInstance(s live.Socket) *ESInstance {
	esInstance, ok := s.Assigns().(*ESInstance)

	if !ok {
		gbl.Log.Warn("could not get eventstream instance")

		return &ESInstance{
			// ID: live.NewID(),
			GasPrice: es.GetGasPrice(),
			Events:   make([]EventMessage, 0),
		}
	}

	return esInstance
}

// func NewEventstreamInstance(s live.Socket) *EventStream {
// 	eventStream, ok := s.Assigns().(*EventStream)

// 	if !ok {
// 		gbl.Log.Error("could not get eventstream instance")
// 		return &EventStream{
// 			ID: live.NewID(),
// 			// GasPrice: es.GasPrice,
// 			Events: make([]EventMessage, 0),
// 		}
// 	}

// 	eventStream.Events = make([]EventMessage, 0)
// 	// es.Events = make([]EventMessage, 0)

// 	// go startWorker(s, es.queueOutWeb)

// 	return eventStream
// }

// func (es *EventStream) NewEventstreamInstance(s live.Socket) *EventStream {
// 	gbl.Log.Infof("es.GasPrice 0: %v", es.GasPrice)

// 	eventStream, ok := s.Assigns().(*EventStream)
// 	// es, ok := s.Assigns().(*EventStream)

// 	if !ok {
// 		// return &EventStream{
// 		// 	ID: live.NewID(),
// 		// 	// Events: []EventMessage{},
// 		// 	Events: make([]EventMessage, 0),
// 		// 	// 	{ID: live.NewID(), User: "Muh", Msg: "Welcome to chat " + live.SessionID(s.Session())},
// 		// 	// },
// 		// }

// 		// gbl.Log.Infof("es.GasPrice 1: %v", es.GasPrice)

// 		return &EventStream{
// 			ID: live.NewID(),
// 			// GasPrice: es.GasPrice,
// 			// Events: make([]EventMessage, 0),
// 		}
// 	}

// 	eventStream.Events = make([]EventMessage, 0)
// 	// es.Events = make([]EventMessage, 0)

// 	// go startWorker(s, es.queueOutWeb)

// 	return es
// }

func startWorker(s *live.Socket, queueOutWeb *chan *collections.Event) {
	gbl.Log.Infof("webWorker started for session %v", live.SessionID((*s).Session()))

	// 	default:
	for event := range *queueOutWeb {
		// for event := range *queueOutWeb {
		gbl.Log.Debugf("webWorker session %+v - event: %+v", live.SessionID((*s).Session()), event)

		// if !event.PrintEvent {
		if event.Discarded != nil && !event.Discarded.PrintInHistory {
			gbl.Log.Debugf("outWeb discarded event: %+v", event)
			continue
		}

		priceEther, _ := nodes.WeiToEther(event.PriceWei).Float64()
		priceEtherPerItem, _ := nodes.WeiToEther(big.NewInt(int64(event.PriceWei.Uint64() / event.TxLogCount))).Float64()

		var to string
		if event.ToENS != "" {
			to = event.ToENS
		} else {
			to = style.ShortenAddress(&event.To.Address)
		}

		var openseaURL string
		if event.Permalink != "" {
			openseaURL = event.Permalink
		} else {
			openseaURL = fmt.Sprintf("https://opensea.io/assets/%s/%d", event.Collection.ContractAddress, event.TokenID)
		}

		var NumItemsColor, pricePerItemColor lipgloss.Color

		switch {
		case event.TxLogCount > 7:
			NumItemsColor = style.AlmostWhiteStyle.GetForeground().(lipgloss.Color)
			pricePerItemColor = style.DarkWhiteStyle.GetForeground().(lipgloss.Color)
		case event.TxLogCount > 4:
			NumItemsColor = style.DarkWhiteStyle.GetForeground().(lipgloss.Color)
			pricePerItemColor = style.LightGrayStyle.GetForeground().(lipgloss.Color)
		case event.TxLogCount > 1:
			NumItemsColor = style.LightGrayStyle.GetForeground().(lipgloss.Color)
			pricePerItemColor = style.GrayStyle.GetForeground().(lipgloss.Color)
		default:
			NumItemsColor = style.GrayStyle.GetForeground().(lipgloss.Color)
			pricePerItemColor = style.DarkGrayStyle.GetForeground().(lipgloss.Color)
		}

		cUp := lipgloss.Color("#99CC99")
		cDown := lipgloss.Color("#CC9999")
		// cError := lipgloss.Color("#DD1010")
		cSteady := lipgloss.Color("#CCCCCC")

		var currentFloorPrice float64

		if fp, err := cache.GetFloor(event.Collection.ContractAddress); err == nil {
			currentFloorPrice = fp
		} else {
			gbl.Log.Debug("failed getting floor price from cache: %s", err)

			currentFloorPrice = (*event.Collection.FloorPrice).Value()
		}

		gbl.Log.Debugf("event.Collection.PreviousFloorPrice: %f | currentFloorPrice: %f | event.Collection.FloorPrice.Value(): %f", event.Collection.PreviousFloorPrice, currentFloorPrice, (*event.Collection.FloorPrice).Value())
		// gbl.Log.Infof("event.Collection.Counters.Sales: %d | event.Collection.Counters.Listings: %d", event.Collection.Counters.Sales, event.Collection.Counters.Listings)

		var trendIndicator string

		var trendIndicatorColor lipgloss.Color

		if currentFloorPrice > 0.0 {
			switch {
			case event.Collection.PreviousFloorPrice < currentFloorPrice:
				trendIndicator = "↑"
				trendIndicatorColor = cUp
			case event.Collection.PreviousFloorPrice > currentFloorPrice:
				trendIndicator = "↓"
				trendIndicatorColor = cDown
			default:
				trendIndicator = "~"
				trendIndicatorColor = cSteady
			}
		} else {
			trendIndicator = "⊗"
			trendIndicatorColor = style.DarkGrayStyle.GetForeground().(lipgloss.Color)
		}

		var salira float64
		if currentSalira, err := cache.GetSalira(event.Collection.ContractAddress); currentSalira != 0.0 && err == nil {
			salira = currentSalira
		}

		data := EventMessage{
			ID:                  live.NewID(),
			Time:                event.Time.Format("15:04:05"),
			Typemoji:            event.EventType.Icon(),
			EventType:           strings.ToLower(event.EventType.String()),
			NumItems:            event.TxLogCount,
			Price:               fmt.Sprintf("%6.3f", priceEther),
			PricePerItem:        fmt.Sprintf("%6.3f", priceEtherPerItem),
			PricePerItemColor:   pricePerItemColor,
			NumItemsColor:       NumItemsColor,
			PriceArrowColor:     string(event.PriceArrowColor),
			FloorPrice:          fmt.Sprintf("%6.3f", currentFloorPrice),
			TrendIndicator:      trendIndicator,
			TrendIndicatorColor: trendIndicatorColor,
			CollectionName:      event.Collection.Name,
			TokenID:             event.TokenID,
			To:                  to,
			ToColor:             string(style.GenerateColorWithSeed(event.To.Address.Hash().Big().Int64())),
			ColorPrimary:        string(event.Collection.Colors.Primary),
			ColorPrimaryFaint:   string(lipgloss.Color(event.Collection.Colors.Primary)),
			ColorSecondary:      string(event.Collection.Colors.Secondary),
			Event:               event,
			SalesCount:          event.Collection.Counters.Sales,
			ListingsCount:       event.Collection.Counters.Listings,
			SaLiRa:              salira,
			LinkOpenSea:         openseaURL,
			LinkEtherscan:       fmt.Sprintf("https://etherscan.io/tx/%s", event.TxHash),
		}

		// gbl.Log.Infof("")
		// gbl.Log.Infof("data: %+v", data)
		// gbl.Log.Infof("")

		gbl.Log.Debugf("%s| before broadcast: %+v", live.SessionID((*s).Session()), data)

		if err := (*s).Broadcast(wnewmessage, data); err != nil {
			gbl.Log.Errorf("failed braodcasting new message: %s", err)
		}
	}
}

func parseTemplates(filenames ...string) *template.Template {
	t, err := template.ParseFiles(filenames...)
	if err != nil {
		gbl.Log.Error(err)
	}

	return t
}

func (es *EventStream) NewEventHandler() live.Handler {
	t := parseTemplates(templateFiles...)

	handler := live.NewHandler(live.WithTemplateRenderer(t))

	handler.HandleError(func(ctx context.Context, err error) {
		gbl.Log.Error("HandleError: %+v", err)
	})

	// Set the mount function for this handler.
	handler.HandleMount(func(ctx context.Context, s live.Socket) (interface{}, error) {
		gbl.Log.Debugf("handle mount socket: %+v", s)

		gbl.Log.Infof("user connected: %+v", live.SessionID(s.Session()))

		// This will initialise the chat for this socket.
		go startWorker(&s, es.queueOutWeb)

		go func() {
			ticker := time.NewTicker(47 * time.Second)
			for range ticker.C {
				// gasPrice := es.GetGasPrice()
				gasPrice := es.GasPrice
				gbl.Log.Debugf("sending %s msg with gasPrice: %d", msgGasUpdate, gasPrice)
				if err := (s).Broadcast(msgGasUpdate, gasPrice); err != nil {
					gbl.Log.Errorf("failed broadcasting new %s msg: %s", msgGasUpdate, err)
				}
			}
		}()

		return es.NewEventstreamInstance(s), nil
	})

	// Handle user sending a message.
	handler.HandleEvent(wsend, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		gbl.Log.Infof("handle event params: %+v", p)

		m := es.NewEventstreamInstance(s)
		msg := p.String("message")
		if msg == "" {
			gbl.Log.Warnf("empty message")
			return m, nil
		}

		gbl.Log.Infof("msg from user %s: %+v", live.SessionID(s.Session()), msg)

		// data := EventMessage{
		// 	ID:       live.NewID(),
		// 	User:     live.SessionID(s.Session()),
		// 	Time:     time.Now().Format("15:04:05"),
		// 	Typemoji: "🗣️",
		// 	Msg:      msg,
		// }

		// if err := s.Broadcast(wnewmessage, data); err != nil {
		// 	gbl.Log.Errorf("failed braodcasting new message: %s", err)
		// 	return m, fmt.Errorf("failed braodcasting new message: %w", err)
		// }
		return m, nil
	})

	// Handle the broadcasted events.
	handler.HandleSelf(wnewmessage, func(ctx context.Context, s live.Socket, data interface{}) (interface{}, error) {
		gbl.Log.Debugf("handle self data: %+v", data)

		m := es.NewEventstreamInstance(s)

		// Here we don't append to Events as we don't want to use
		// loads of memory. `live-update="append"` handles the appending
		// of messages in the DOM.
		m.Events = []EventMessage{NewEvent(data)}

		return m, nil
	})

	handler.HandleSelf(msgGasUpdate, func(ctx context.Context, s live.Socket, data interface{}) (interface{}, error) {
		gbl.Log.Debugf("handle %s msg: %+v", msgGasUpdate, data)

		m := es.NewEventstreamInstance(s)
		m.GasPrice, _ = data.(int)
		m.Events = []EventMessage{{Divider: true}}

		return m, nil
	})

	gbl.Log.Infof("handler created: %+v", handler)

	return handler
}

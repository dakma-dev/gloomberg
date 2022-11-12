package web

import (
	"context"
	"fmt"
	"html/template"
	"math/big"
	"net/http"
	"strings"

	"github.com/benleb/gloomberg/internal/cache"
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/nodes"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/charmbracelet/lipgloss"
	"github.com/jfyne/live"
)

const (
	wsend       = "send"
	wnewmessage = "newmessage"
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
	Events        []EventMessage
	queueOutWeb   *chan *collections.Event
	Nodes         *nodes.Nodes
}

func New(queueWeb *chan *collections.Event, listenAddress string, nodes *nodes.Nodes) *EventStream {
	ctx := context.Background()

	return &EventStream{
		ID:            live.NewID(),
		ListenAddress: listenAddress,
		ctx:           ctx,
		queueOutWeb:   queueWeb,
		Nodes:         nodes,
	}
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

func (es *EventStream) NewEventstreamInstance(s live.Socket) *EventStream {
	eventStream, ok := s.Assigns().(*EventStream)

	if !ok {
		// return &EventStream{
		// 	ID: live.NewID(),
		// 	// Events: []EventMessage{},
		// 	Events: make([]EventMessage, 0),
		// 	// 	{ID: live.NewID(), User: "Muh", Msg: "Welcome to chat " + live.SessionID(s.Session())},
		// 	// },
		// }
		return &EventStream{
			ID: live.NewID(),
			// Events: make([]EventMessage, 0),
		}
	}

	eventStream.Events = make([]EventMessage, 0)

	// go startWorker(s, es.queueOutWeb)

	return eventStream
}

func startWorker(s *live.Socket, queueOutWeb *chan *collections.Event) {
	gbl.Log.Infof("webWorker started for session %v", live.SessionID((*s).Session()))

	// dividerTicker := time.NewTicker(viper.GetDuration("ticker.divider"))

	// for {
	// 	select {
	// 	// case event := <-*queueOutWeb:

	// 	// 	// for event := range *queueOutWeb {
	// 	// 	gbl.Log.Debugf("webWorker session %+v - event: %+v", live.SessionID((*s).Session()), event)

	// 	// 	if !event.PrintEvent {
	// 	// 		gbl.Log.Debugf("outWeb discarded event: %+v", event)
	// 	// 		continue
	// 	// 	}

	// 	// 	priceEther, _ := nodes.WeiToEther(event.PriceWei).Float64()
	// 	// 	priceEtherPerItem, _ := nodes.WeiToEther(big.NewInt(int64(event.PriceWei.Uint64() / event.TxLogCount))).Float64()

	// 	// 	var to string
	// 	// 	if event.ToENS != "" {
	// 	// 		to = event.ToENS
	// 	// 	} else {
	// 	// 		to = style.ShortenAddress(&event.To.Address)
	// 	// 	}

	// 	// 	var openseaURL string
	// 	// 	if event.Permalink != "" {
	// 	// 		openseaURL = event.Permalink
	// 	// 	} else {
	// 	// 		openseaURL = fmt.Sprintf("https://opensea.io/assets/%s/%d", event.Collection.ContractAddress, event.TokenID)
	// 	// 	}

	// 	// 	var TxLogCountColor lipgloss.Color

	// 	// 	switch {
	// 	// 	case event.TxLogCount > 7:
	// 	// 		TxLogCountColor = style.AlmostWhiteStyle.GetForeground().(lipgloss.Color)
	// 	// 	case event.TxLogCount > 4:
	// 	// 		TxLogCountColor = style.DarkWhiteStyle.GetForeground().(lipgloss.Color)
	// 	// 	case event.TxLogCount > 1:
	// 	// 		TxLogCountColor = style.LightGrayStyle.GetForeground().(lipgloss.Color)
	// 	// 	default:
	// 	// 		TxLogCountColor = style.GrayStyle.GetForeground().(lipgloss.Color)
	// 	// 	}

	// 	// 	data := EventMessage{
	// 	// 		ID:              live.NewID(),
	// 	// 		Time:            event.Time.Format("15:04:05"),
	// 	// 		Typemoji:        event.EventType.Icon(),
	// 	// 		EventType:       strings.ToLower(event.EventType.String()),
	// 	// 		TxLogCount:      event.TxLogCount,
	// 	// 		TxLogCountColor: TxLogCountColor,
	// 	// 		Price:           fmt.Sprintf("%6.3f", priceEther),
	// 	// 		PricePerItem:    priceEtherPerItem,
	// 	// 		PriceArrowColor: string(event.PriceArrowColor),
	// 	// 		CollectionName:  event.Collection.Name,
	// 	// 		TokenID:         event.TokenID,
	// 	// 		To:              to,
	// 	// 		ToColor:         string(style.GenerateColorWithSeed(event.To.Address.Hash().Big().Int64())),
	// 	// 		ColorPrimary:    string(event.Collection.Colors.Primary),
	// 	// 		ColorSecondary:  string(event.Collection.Colors.Secondary),
	// 	// 		Event:           event,
	// 	// 		SalesCount:      event.Collection.Counters.Sales,
	// 	// 		ListingsCount:   event.Collection.Counters.Listings,
	// 	// 		SaLiRa:          event.Collection.SaLiRa.Value(),
	// 	// 		LinkOpenSea:     openseaURL,
	// 	// 		LinkEtherscan:   fmt.Sprintf("https://etherscan.io/tx/%s", event.TxHash),
	// 	// 	}

	// 	// 	gbl.Log.Debugf("%s| before broadcast: %+v", live.SessionID((*s).Session()), data)

	// 	// 	if err := (*s).Broadcast(wnewmessage, data); err != nil {
	// 	// 		gbl.Log.Errorf("failed braodcasting new message: %s", err)
	// 	// 	}
	// 	case <-dividerTicker.C:
	// 		data := EventMessage{Divider: true}
	// 		if err := (*s).Broadcast(wnewmessage, data); err != nil {
	// 			gbl.Log.Errorf("failed braodcasting new message: %s", err)
	// 		}

	// 	default:
	for event := range *queueOutWeb {

		// for event := range *queueOutWeb {
		gbl.Log.Debugf("webWorker session %+v - event: %+v", live.SessionID((*s).Session()), event)

		if !event.PrintEvent {
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

		var NumItemsColor lipgloss.Color
		var pricePerItemColor lipgloss.Color

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

		gbl.Log.Infof("event.Collection.PreviousFloorPrice: %f | currentFloorPrice: %f | event.Collection.FloorPrice.Value(): %f", event.Collection.PreviousFloorPrice, currentFloorPrice, (*event.Collection.FloorPrice).Value())
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
			SaLiRa:              event.Collection.SaLiRa.Value(),
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
	// gbl.Log.Infof("webWorker closed for session %+v", live.SessionID((*s).Session()))
}

// func GasLineMessage(es *EventStream, s *live.Socket, handler *live.BaseHandler, webGasTicker *time.Ticker, ethNodes *nodes.Nodes) {
// 	oldGasPrice := 0

// 	for range webGasTicker.C {
// 		gasNode := ethNodes.GetRandomLocalNode()

// 		if gasInfo, err := gasNode.GetCurrentGasInfo(); err == nil && gasInfo != nil {
// 			// gas price
// 			if gasInfo.GasPriceWei.Cmp(big.NewInt(0)) > 0 {
// 				gasPriceGwei, _ := nodes.WeiToGwei(gasInfo.GasPriceWei).Float64()
// 				gasPrice := int(math.Round(gasPriceGwei))

// 				if math.Abs(float64(gasPrice-oldGasPrice)) < 2.0 {
// 					continue
// 				}

// 				oldGasPrice = gasPrice

// 				gasPriceGwei, _ = nodes.WeiToGwei(gasInfo.GasPriceWei).Float64()
// 				gasTipGwei, _ := nodes.WeiToGwei(gasInfo.GasTipWei).Float64()

// 				data := EventMessage{
// 					Time:     time.Now().Format("15:04:05"),
// 					Typemoji: "🛢️", // "🧟",
// 					GasInfo: GasInfoMessage{
// 						PriceGwei: int(math.Round(gasPriceGwei)),
// 						TipGwei:   math.Round(gasTipGwei),
// 					},
// 				}

// 				if err := (*s).Broadcast(wnewmessage, data); err != nil {
// 					gbl.Log.Errorf("failed braodcasting new message: %s", err)
// 				}
// 			}
// 		}
// 	}
// }

func parseTemplates(filenames ...string) *template.Template {
	t, err := template.ParseFiles(filenames...)
	if err != nil {
		gbl.Log.Error(err)
	}

	return t
}

func (es *EventStream) NewEventHandler() live.Handler {
	// t, err := template.ParseFiles("www/layout.html", "www/style.html", "www/view.html")
	// t, err := template.ParseFiles("www/gLayout.html", "www/gStyle.html", "www/gView.html")
	// if err != nil {
	// 	gbl.Log.Error(err)
	// }

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

		// if tickerInterval := viper.GetDuration("ticker.statsbox"); es.Nodes != nil && len(es.Nodes.GetLocalNodes()) > 0 && tickerInterval > 0 {
		// 	// start gasline ticker
		// 	webGasTicker := time.NewTicker(tickerInterval)
		// 	go GasLineMessage(es, &s, handler, webGasTicker, es.Nodes)
		// }

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
		// gbl.Log.Infof("broadcasting to %s: %+v", live.SessionID(s.Session()), data)

		m := es.NewEventstreamInstance(s)

		// Here we don't append to messages as we don't want to use
		// loads of memory. `live-update="append"` handles the appending
		// of messages in the DOM.
		m.Events = []EventMessage{NewEvent(data)}
		return m, nil
	})

	gbl.Log.Infof("handler created: %+v", handler)

	return handler
}

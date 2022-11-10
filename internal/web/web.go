package web

import (
	"context"
	"fmt"
	"html/template"
	"math"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/nodes"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/charmbracelet/lipgloss"
	"github.com/jfyne/live"
	"github.com/spf13/viper"
)

const (
	wsend       = "send"
	wnewmessage = "newmessage"
	glbGasLine  = "gasline"
)

type EventMessage struct {
	ID              string // Unique ID per message so that we can use `live-update`.
	User            string
	Msg             string
	Time            string
	Typemoji        string
	TxLogCount      uint64
	TxLogCountColor lipgloss.Color
	Price           string
	PricePerItem    float64
	PriceArrowColor string
	CollectionName  string
	TokenID         *big.Int
	ColorPrimary    string
	ColorSecondary  string
	To              string
	ToColor         string
	Event           *collections.Event
	SalesCount      uint64
	ListingsCount   uint64
	SaLiRa          float64
	LinkOpenSea     string
	LinkEtherscan   string
	GasLine         string
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
	gbl.Log.Info("webWorker started for session %+v", live.SessionID((*s).Session()))

	for event := range *queueOutWeb {
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

		var TxLogCountColor lipgloss.Color

		switch {
		case event.TxLogCount > 7:
			TxLogCountColor = style.AlmostWhiteStyle.GetForeground().(lipgloss.Color)
		case event.TxLogCount > 4:
			TxLogCountColor = style.DarkWhiteStyle.GetForeground().(lipgloss.Color)
		case event.TxLogCount > 1:
			TxLogCountColor = style.LightGrayStyle.GetForeground().(lipgloss.Color)
		default:
			TxLogCountColor = style.GrayStyle.GetForeground().(lipgloss.Color)
		}

		data := EventMessage{
			ID:              live.NewID(),
			Time:            event.Time.Format("15:04:05"),
			Typemoji:        event.EventType.Icon(),
			TxLogCount:      event.TxLogCount,
			TxLogCountColor: TxLogCountColor,
			Price:           fmt.Sprintf("%6.3f", priceEther),
			PricePerItem:    priceEtherPerItem,
			PriceArrowColor: string(event.PriceArrowColor),
			CollectionName:  event.Collection.Name,
			TokenID:         event.TokenID,
			To:              to,
			ToColor:         string(style.GenerateColorWithSeed(event.To.Address.Hash().Big().Int64())),
			ColorPrimary:    string(event.Collection.Colors.Primary),
			ColorSecondary:  string(event.Collection.Colors.Secondary),
			Event:           event,
			SalesCount:      event.Collection.Counters.Sales,
			ListingsCount:   event.Collection.Counters.Listings,
			SaLiRa:          event.Collection.SaLiRa.Value(),
			LinkOpenSea:     openseaURL,
			LinkEtherscan:   fmt.Sprintf("https://etherscan.io/tx/%s", event.TxHash),
		}

		gbl.Log.Debugf("%s| before broadcast: %+v", live.SessionID((*s).Session()), data)

		if err := (*s).Broadcast(wnewmessage, data); err != nil {
			gbl.Log.Errorf("failed braodcasting new message: %s", err)
		}
	}

	gbl.Log.Info("webWorker closed for session %+v", live.SessionID((*s).Session()))
}

func GasLineMessage(s *live.Socket, webGasTicker *time.Ticker, ethNodes *nodes.Nodes) {
	oldGasPrice := 0

	for range webGasTicker.C {
		gasNode := ethNodes.GetRandomLocalNode()
		gasLine := strings.Builder{}

		if gasInfo, err := gasNode.GetCurrentGasInfo(); err == nil && gasInfo != nil {
			// gas price
			if gasInfo.GasPriceWei.Cmp(big.NewInt(0)) > 0 {
				gasPriceGwei, _ := nodes.WeiToGwei(gasInfo.GasPriceWei).Float64()
				gasPrice := int(math.Round(gasPriceGwei))

				if math.Abs(float64(gasPrice-oldGasPrice)) < 2.0 {
					continue
				}

				oldGasPrice = gasPrice

				// // tip / priority fee
				// var gasTip int
				// if gasInfo.GasTipWei.Cmp(big.NewInt(0)) > 0 {
				// 	gasTipGwei, _ := nodes.WeiToGwei(gasInfo.GasTipWei).Float64()
				// 	gasTip = int(math.Round(gasTipGwei))
				// 	fmt.Printf("gasInfo.GasTipWei: %+v | gasTipGwei: %+v | gasTip: %+v\n", gasInfo.GasTipWei, gasTipGwei, gasTip)
				// }

				intro := style.DarkerGrayStyle.Render("~  ") + style.DarkGrayStyle.Render("gas") + style.DarkerGrayStyle.Render("  ~   ")
				outro := style.DarkerGrayStyle.Render("   ~   ~")
				divider := style.DarkerGrayStyle.Render("   ~   ~   ~   ~   ~   ~   ")

				formattedGas := style.GrayStyle.Render(fmt.Sprintf("%d", gasPrice)) + style.DarkGrayStyle.Render("gw")
				formattedGasAndTip := formattedGas

				// if gasTip > 0 {
				// 	formattedGasAndTip = formattedGas + "|" + style.GrayStyle.Render(fmt.Sprintf("%d", gasTip)) + style.DarkGrayStyle.Render("gw")
				// }

				gasLine.WriteString(intro + formattedGas + divider + formattedGasAndTip + divider + formattedGas + outro)

				data := EventMessage{
					GasLine: gasLine.String(),
				}

				if err := (*s).Broadcast(wnewmessage, data); err != nil {
					gbl.Log.Errorf("failed braodcasting new message: %s", err)
				}
			}
		}

		// *queueOutput <- gasLine.String()
	}
}

func (es *EventStream) NewEventHandler() live.Handler {
	t, err := template.ParseFiles("www/layout.html", "www/style.html", "www/view.html")
	// t, err := template.ParseFiles("www/gLayout.html", "www/gStyle.html", "www/gView.html")
	if err != nil {
		gbl.Log.Error(err)
	}

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

		if tickerInterval := viper.GetDuration("ticker.statsbox"); es.Nodes != nil && len(es.Nodes.GetLocalNodes()) > 0 && tickerInterval > 0 {
			// start gasline ticker
			webGasTicker := time.NewTicker(tickerInterval)
			go GasLineMessage(&s, webGasTicker, es.Nodes)
		}

		return es.NewEventstreamInstance(s), nil
	})

	// Handle user sending a message.
	handler.HandleEvent(wsend, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		gbl.Log.Debugf("handle event params: %+v", p)

		m := es.NewEventstreamInstance(s)
		msg := p.String("message")
		if msg == "" {
			gbl.Log.Warnf("empty message")
			return m, nil
		}

		gbl.Log.Infof("msg from user %s: %+v", live.SessionID(s.Session()), msg)

		data := EventMessage{
			ID:       live.NewID(),
			User:     live.SessionID(s.Session()),
			Time:     time.Now().Format("15:04:05"),
			Typemoji: "🗣️",
			Msg:      msg,
		}

		if err := s.Broadcast(wnewmessage, data); err != nil {
			gbl.Log.Errorf("failed braodcasting new message: %s", err)
			return m, fmt.Errorf("failed braodcasting new message: %w", err)
		}
		return m, nil
	})

	// Handle the broadcasted events.
	handler.HandleSelf(wnewmessage, func(ctx context.Context, s live.Socket, data interface{}) (interface{}, error) {
		gbl.Log.Debugf("handle self data: %+v", data)
		gbl.Log.Debugf("broadcasting to %s: %+v", live.SessionID(s.Session()), data)

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

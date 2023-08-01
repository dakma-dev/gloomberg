package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"html/template"
	"io"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/benleb/gloomberg/internal/external"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gobwas/ws"
)

// type Client struct {
// 	id         string
// 	conn       net.Conn
// 	eventSream *WebsocketsServer
// }

/*
*
wsUpgrader is used to upgrade incomming HTTP requests into a persitent websocket connection.
*/
var (
	// allowedOrigins = []string{"https://localhost:8080", "https://10.0.0.99:8080", "https://mia.home.benleb.de:8080"}

	// wsUpgrader = websocket.Upgrader{
	// 	EnableCompression: true,
	// 	CheckOrigin:       checkOrigin,
	// 	ReadBufferSize:    1024,
	// 	WriteBufferSize:   1024,
	// }.

	ErrEventNotSupported = errors.New("this event type is not supported")
)

// // checkOrigin will check origin and return true if its allowed
// func checkOrigin(r *http.Request) bool {
// 	origin := r.Header.Get("Origin")
// 	for _, o := range allowedOrigins {
// 		if o == origin {
// 			return true
// 		}
// 	}

// 	return false
// }

type WsHub struct {
	gb *gloomberg.Gloomberg
	// websockets
	clients map[*WsClient]bool

	// templates
	templates *template.Template

	// handlers are functions that are used to handle Events
	handlers map[string]MessageHandler

	server *http.Server

	sync.RWMutex

	// listenHost string `mapstructure:"host"`
	// listenPort uint   `mapstructure:"port"`

	// // ws handling
	// clients    map[string]*Client
	// clientList []*Client

	// out chan []byte
}

func NewHub(gb *gloomberg.Gloomberg) *WsHub {
	// s := &WebsocketsServer{
	// 	mu:                          sync.RWMutex{},
	// 	queueWsOutTokenTransactions: eventQueue,

	// 	listenHost: listenHost,
	// 	listenPort: listenPort,

	// 	clients: make(map[string]*Client),

	// 	out: make(chan []byte),
	// }

	// // go s.writer()

	// return s

	hub := WsHub{
		gb:       gb,
		clients:  make(map[*WsClient]bool),
		handlers: make(map[string]MessageHandler),
	}

	tmplFiles := []string{"www/event.tpl.html", "www/recent_own_events.tpl.html"}
	tmpls, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		gbl.Log.Error(err)
	}

	hub.templates = tmpls

	hub.setupEventHandlers()

	// loopy mcLoopface
	hub.broadcaster()

	// for i := 0; i < viper.GetInt("gloomberg.eventhub.numHandler"); i++ {
	// 	go eh.worker(i)
	// }

	return &hub
}

// func (hub *WsHub) StartServer(server *http.Server) {
// 	// start http server
// 	log.Fatal(server.ListenAndServeTLS("", ""))
// }

// ClientsOnline returns the number of clients that are currently connected.
func (wh *WsHub) ClientsOnline() uint64 {
	wh.RLock()
	defer wh.RUnlock()

	return uint64(len(wh.clients))
}

func (wh *WsHub) broadcaster() {
	parsedEventsChannel := wh.gb.SubscribeParsedEvents()
	recentOwnEventsChannel := wh.gb.SubscribeRecentOwnEvents()

	// log.Printf(" 🧚‍♀️ 🧚‍♀️ 🧚‍♀️ parsedEventsChannel | %p | %d  🧚‍♀️ 🧚‍♀️ 🧚‍♀️ ", parsedEventsChannel, len(parsedEventsChannel))

	go func() {
		for {
			var msgType, preparedMessage string
			var rendered bytes.Buffer

			select {
			case parsedEvent := <-parsedEventsChannel:
				// log.Printf("  🧚‍♀️ parsedEvent | %p | %d  🧚‍♀️ 🧚‍♀️ 🧚‍♀️ ", parsedEventsChannel, len(parsedEventsChannel))

				if parsedEvent == nil {
					continue
				}

				if parsedEvent.From == nil || parsedEvent.From.Addresses[0].Address == (common.Address{}) || len(parsedEvent.From.Addresses[0].Address) == 0 || parsedEvent.To == nil || parsedEvent.To.Addresses[0].Address == (common.Address{}) || len(parsedEvent.To.Addresses[0].Address) == 0 {
					continue
				}

				err := wh.templates.ExecuteTemplate(io.MultiWriter(&rendered), "event", parsedEvent)
				if err != nil {
					log.Errorf("❌ rendering template failed: %+v", err)
				}

				log.Debugf("rendered: %s", rendered.String())

				msgType = MsgNewSale

			case recentOwnEvents := <-recentOwnEventsChannel:
				if len(recentOwnEvents) == 0 {
					continue
				}

				err := wh.templates.ExecuteTemplate(io.MultiWriter(&rendered), "recent_own_events", recentOwnEvents)
				if err != nil {
					log.Errorf("❌ rendering template failed: %+v", err)
				}

				log.Debugf("rendered: %s", rendered.String())

				msgType = MsgRecentOwnEvents
			}

			if wh.ClientsOnline() == 0 {
				time.Sleep(time.Second * 1)

				continue
			}

			//
			// html minify for the poor...
			preparedMessage = strings.ReplaceAll(rendered.String(), "\n", "")

			betweenTagsWhitespace := regexp.MustCompile(`> *<`)
			preparedMessage = betweenTagsWhitespace.ReplaceAllString(preparedMessage, "><")

			multiWhitespace := regexp.MustCompile(` {2,}`)
			preparedMessage = multiWhitespace.ReplaceAllString(preparedMessage, " ")

			//
			// create ws message
			msg := MessageGeneric[EventPayload]{
				Type:     msgType,
				Payload:  EventPayload{Message: preparedMessage},
				GasPrice: wh.gb.CurrentGasPriceGwei,
			}

			log.Debugf("msg: %+v", msg)

			marshalledMsg, err := json.Marshal(msg)
			if err != nil {
				gbl.Log.Errorf("error marshalling event: %s", err.Error())

				continue
			}

			// broadcast
			go wh.broadcast(marshalledMsg)

			gbl.Log.Debugf("event sent to client: %s", string(marshalledMsg))
		}
	}()
}

func (wh *WsHub) broadcast(msg json.RawMessage) {
	// broadcast
	wh.RLock()
	clients := wh.clients
	wh.RUnlock()

	for client := range clients {
		client.egress <- msg
	}
}

// setupEventHandlers configures and adds all handlers.
func (wh *WsHub) setupEventHandlers() {
	wh.handlers[MsgCommand] = func(msg Message, _ *WsClient) error {
		gbl.Log.Info("received message: ", msg)

		return nil
	}
}

// routeEvent is used to make sure the correct event goes into the correct handler.
func (wh *WsHub) routeEvent(event Message, wc *WsClient) error {
	// Check if Handler is present in Map
	if handler, ok := wh.handlers[event.Type]; ok {
		// Execute the handler and return any err
		return handler(event, wc)
	}

	return ErrEventNotSupported
}

// serveWS is a HTTP Handler that the has the Manager that allows connections.
func (wh *WsHub) serveWS(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		// handle error
		gbl.Log.Error(err)

		return
	}

	// create Client
	client := NewClient(conn, wh)

	// add the cliented to the hub
	wh.addClient(client)

	go client.readMessages()
	go client.writeMessages()

	ipInfo, err := external.GetIPInfo(r.Context(), client.conn.RemoteAddr())

	switch {
	case err != nil && ipInfo == nil:
		gbl.Log.Infof("⁉️ error getting ip info: %s", err.Error())
	case err != nil && ipInfo != nil:
		gbl.Log.Debugf("⁉️ warning getting ip info: %s", err.Error())

		client.ipInfo = ipInfo
	default:
		gbl.Log.Infof("ip info: %+v", ipInfo)

		client.ipInfo = ipInfo
	}

	gbl.Log.Infof("new client connected: %+v", client)
	wh.gb.PrModf("web", "new client connected: %+v", style.AlmostWhiteStyle.Render(client.String()))
}

// addClient will add clients to our clientList.
func (wh *WsHub) addClient(client *WsClient) {
	gbl.Log.Debugf("    adding client: %+v", client)

	// add Client to locked map
	wh.Lock()
	wh.clients[client] = true
	wh.Unlock()

	wh.gb.CurrentOnlineWebUsers = wh.ClientsOnline()
}

// removeClient will remove the client and clean up.
func (wh *WsHub) removeClient(client *WsClient) {
	wh.Lock()
	defer wh.Unlock()

	var clientID string

	// Check if Client exists, then delete it
	if _, ok := wh.clients[client]; ok {
		clientID = client.String()

		// close connection
		client.conn.Close()
		// remove
		delete(wh.clients, client)
	}

	gbl.Log.Infof("client %s removed", clientID)
}

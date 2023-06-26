package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
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

	// handlers are functions that are used to handle Events
	handlers map[string]MessageHandler

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

	hub := &WsHub{
		gb:       gb,
		clients:  make(map[*WsClient]bool),
		handlers: make(map[string]MessageHandler),
	}

	hub.setupEventHandlers()

	go hub.broadcastTTx()

	return hub
}

// writer writes broadcast messages from chat.out channel.
func (wh *WsHub) broadcastTTx() {
	for ttx := range wh.gb.SubscribeTokenTransactions() {
		if len(wh.clients) == 0 {
			continue
		}

		// event := event

		// var wsEvent *totra.TokenTransaction
		// copier.Copy(&wsEvent, &event)
		// wsEvent.Collection.Source = wsEvent.Collection.Source.String()

		// pushEvent := eventToPushEvent(ttx)

		newSaleMsg := NewEventMessage{
			Message: fmt.Sprintf("New sale of %s for %s", ttx.From.String(), ttx.AmountPaid.String()),
			From:    ttx.From.String(),
		}

		msg := Message{
			Type:    MsgNewSale,
			Payload: newSaleMsg,
		}

		marshalledMsg, err := json.Marshal(msg)
		if err != nil {
			gbl.Log.Errorf("error marshalling event: %s", err.Error())

			continue
		}

		// broadcast
		wh.RLock()
		clients := wh.clients
		wh.RUnlock()

		gbl.Log.Debugf("pushing new event (%db) to %d clients", len(marshalledMsg), len(clients))

		for client := range clients {
			client.egress <- msg
			// if err := wsutil.WriteServerText(client.conn, marshalledMsg); err != nil {
			// 	if errors.Is(err, syscall.EPIPE) {
			// 		gbl.Log.Errorf("client %s disconnected: %s", client.conn.LocalAddr().String(), err.Error())

			// 		// remove client
			// 		client.conn.Close()
			// 		wh.removeClient(client)
			// 	} else {
			// 		gbl.Log.Errorf("sending event to client %v failed: %s | event: %s", client, string(marshalledMsg), err.Error())
			// 	}

			// 	continue
			// }
		}

		gbl.Log.Debugf("event sent to client: %s", string(marshalledMsg))
	}
}

// setupEventHandlers configures and adds all handlers.
func (wh *WsHub) setupEventHandlers() {
	wh.handlers[MsgCommand] = func(msg Message, _ *WsClient) error {
		gbl.Log.Info("received message: ", msg)
		fmt.Println("received message: ", msg)

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
	gbl.Log.Info("New connection")

	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		// handle error
		gbl.Log.Error(err)

		return
	}

	// Create New Client
	client := NewClient(conn, wh)
	// Add the newly created client to the manager
	wh.addClient(client)

	go client.readMessages()

	go client.writeMessages()

	// // We wont do anything yet so close connection again
	// conn.Close()
}

// addClient will add clients to our clientList.
func (wh *WsHub) addClient(client *WsClient) {
	// Lock so we can manipulate
	wh.Lock()
	defer wh.Unlock()

	// Add Client
	wh.clients[client] = true

	gbl.Log.Info("client added")
}

// removeClient will remove the client and clean up.
func (wh *WsHub) removeClient(client *WsClient) {
	wh.Lock()
	defer wh.Unlock()

	// Check if Client exists, then delete it
	if _, ok := wh.clients[client]; ok {
		// close connection
		client.conn.Close()
		// remove
		delete(wh.clients, client)
	}

	gbl.Log.Info("client removed")
}

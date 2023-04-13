package web

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"syscall"

	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// type Client struct {
// 	id         string
// 	conn       net.Conn
// 	eventSream *WebsocketsServer
// }

/*
*
wsUpgrader is used to upgrade incomming HTTP requests into a persitent websocket connection
*/
var (
	// allowedOrigins = []string{"https://localhost:8080", "https://10.0.0.99:8080", "https://mia.home.benleb.de:8080"}

	// wsUpgrader = websocket.Upgrader{
	// 	EnableCompression: true,
	// 	CheckOrigin:       checkOrigin,
	// 	ReadBufferSize:    1024,
	// 	WriteBufferSize:   1024,
	// }

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
	// websockets
	clients map[*WsClient]bool

	// handlers are functions that are used to handle Events
	handlers map[string]EventHandler

	queueWsOutTokenTransactions *chan *totra.TokenTransaction

	sync.RWMutex

	// listenHost string `mapstructure:"host"`
	// listenPort uint   `mapstructure:"port"`

	// // ws handling
	// clients    map[string]*Client
	// clientList []*Client

	// out chan []byte
}

func NewHub(queueWsOutTokenTransactions chan *totra.TokenTransaction) *WsHub {
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
		clients:  make(map[*WsClient]bool),
		handlers: make(map[string]EventHandler),

		queueWsOutTokenTransactions: &queueWsOutTokenTransactions,
	}

	hub.setupEventHandlers()

	go hub.broadcastTTx()

	return hub
}

// writer writes broadcast messages from chat.out channel.
func (wh *WsHub) broadcastTTx() {
	for ttx := range *wh.queueWsOutTokenTransactions {
		if len(wh.clients) == 0 {
			continue
		}

		// event := event

		// var wsEvent *totra.TokenTransaction
		// copier.Copy(&wsEvent, &event)
		// wsEvent.Collection.Source = wsEvent.Collection.Source.String()

		// pushEvent := eventToPushEvent(ttx)

		marshalledEvent, err := json.Marshal(ttx)
		if err != nil {
			gbl.Log.Errorf("error marshalling event: %s", err.Error())

			continue
		}

		// broadcast
		wh.RLock()
		clients := wh.clients
		wh.RUnlock()

		gbl.Log.Debugf("pushing new event (%db) to %d clients", len(marshalledEvent), len(clients))

		for client := range clients {
			if err := wsutil.WriteServerText(client.conn, marshalledEvent); err != nil {
				if errors.Is(err, syscall.EPIPE) {
					gbl.Log.Errorf("client %s disconnected: %s", client.conn.LocalAddr().String(), err.Error())

					// remove client
					client.conn.Close()
					wh.removeClient(client)
				} else {
					gbl.Log.Errorf("sending event to client %v failed: %s | event: %s", client, string(marshalledEvent), err.Error())
				}

				continue
			}
		}

		gbl.Log.Debugf("event sent to client: %s", string(marshalledEvent))
	}
}

// setupEventHandlers configures and adds all handlers
func (wh *WsHub) setupEventHandlers() {
	wh.handlers[EventSendMessage] = func(e Event, wc *WsClient) error {
		gbl.Log.Error(e)

		return nil
	}
}

// routeEvent is used to make sure the correct event goes into the correct handler
func (wh *WsHub) routeEvent(event Event, wc *WsClient) error {
	// Check if Handler is present in Map
	if handler, ok := wh.handlers[event.Type]; ok {
		// Execute the handler and return any err
		if err := handler(event, wc); err != nil {
			return err
		}

		return nil
	}

	return ErrEventNotSupported
}

// serveWS is a HTTP Handler that the has the Manager that allows connections
func (wh *WsHub) serveWS(w http.ResponseWriter, r *http.Request) {
	gbl.Log.Info("New connection")

	// // Begin by upgrading the HTTP request
	// conn, err := wsUpgrader.Upgrade(w, r, nil)
	// if err != nil {
	// 	gbl.Log.Errorf(err)
	// 	return
	// }

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

// addClient will add clients to our clientList
func (wh *WsHub) addClient(client *WsClient) {
	// Lock so we can manipulate
	wh.Lock()
	defer wh.Unlock()

	// Add Client
	wh.clients[client] = true

	gbl.Log.Info("client added")
}

// removeClient will remove the client and clean up
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

// func (wh *WebsocketsHub) Start() {
// 	listenOn := fmt.Sprint(s.listenHost) + ":" + fmt.Sprint(s.listenPort)

// 	gbl.Log.Infof("✅ starting websocket server on %s", listenOn)

// 	go s.writer()

// 	http.Handle("/", http.HandlerFunc(s.wsHandler))

// 	if err := http.ListenAndServe(listenOn, nil); err != nil { //nolint:gosec
// 		gbl.Log.Fatal(err)
// 	}
// }

// func (wh *WebsocketsHub) ClientsConnected() int {
// 	s.mu.RLock()
// 	defer s.mu.RUnlock()

// 	return len(s.clients)
// }

// // func (wh *WebsocketsHub) Broadcast() error {
// // 	var buf bytes.Buffer

// // 	w := wsutil.NewWriter(&buf, ws.StateServerSide, ws.OpText)
// // 	encoder := json.NewEncoder(w)

// // 	r := totra.TokenTransaction{Topic: "muh"}
// // 	if err := encoder.Encode(r); err != nil {
// // 		return err
// // 	}

// // 	if err := w.Flush(); err != nil {
// // 		return err
// // 	}

// // 	s.out <- buf.Bytes()

// // 	return nil
// // }

// // writer writes broadcast messages from chat.out channel.
// func (wh *WebsocketsHub) writer() {
// 	for ttx := range *s.queueWsOutTokenTransactions {
// 		if len(s.clientList) == 0 {
// 			continue
// 		}

// 		// event := event

// 		// var wsEvent *totra.TokenTransaction
// 		// copier.Copy(&wsEvent, &event)
// 		// wsEvent.Collection.Source = wsEvent.Collection.Source.String()

// 		// pushEvent := eventToPushEvent(ttx)

// 		marshalledEvent, err := json.Marshal(ttx)
// 		if err != nil {
// 			gbl.Log.Errorf("error marshalling event: %s", err.Error())

// 			continue
// 		}

// 		// broadcast
// 		s.mu.RLock()
// 		clients := s.clientList
// 		s.mu.RUnlock()

// 		gbl.Log.Infof("pushing new event (%db) to %d clients", len(marshalledEvent), len(clients))

// 		for _, client := range clients {
// 			if err := wsutil.WriteServerText(client.conn, marshalledEvent); err != nil {
// 				if errors.Is(err, syscall.EPIPE) {
// 					gbl.Log.Errorf("client %s disconnected: %s", client.id, err.Error())

// 					// remove client
// 					client.conn.Close()
// 					s.Remove(client)
// 				} else {
// 					gbl.Log.Errorf("sending event to client %v failed: %s | event: %s", client, string(marshalledEvent), err.Error())
// 				}

// 				continue
// 			}
// 		}

// 		gbl.Log.Debugf("event sent to client: %s", string(marshalledEvent))
// 	}
// }

// func (wh *WebsocketsHub) Register(conn net.Conn) *Client {
// 	client := &Client{
// 		eventSream: s,
// 		conn:       conn,
// 		id:         conn.RemoteAddr().String(),
// 	}

// 	gbl.Log.Infof("register client: %v | %p", client, conn)

// 	s.mu.Lock()
// 	s.clientList = append(s.clientList, client)
// 	s.clients[client.id] = client
// 	s.mu.Unlock()

// 	gbl.Log.Infof("register s.client: %d | %p", len(s.clientList), s.clientList[len(s.clientList)-1].conn)

// 	return client
// }

// // Remove removes user from chat.
// func (wh *WebsocketsHub) Remove(client *Client) {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	if removed := s.remove(client); !removed {
// 		gbl.Log.Errorf("removing client %s failed oO", client.id)

// 		return
// 	}
// }

// // mutex must be held.
// func (wh *WebsocketsHub) remove(client *Client) bool {
// 	if _, has := s.clients[client.id]; !has {
// 		return false
// 	}

// 	delete(s.clients, client.id)

// 	i := sort.Search(len(s.clientList), func(i int) bool {
// 		return s.clientList[i].id >= client.id
// 	})
// 	if i >= len(s.clientList) {
// 		panic("chat: inconsistent state")
// 	}

// 	without := make([]*Client, len(s.clientList)-1)
// 	copy(without[:i], s.clientList[:i])
// 	copy(without[i:], s.clientList[i+1:])
// 	s.clientList = without

// 	return true
// }

// func (wh *WebsocketsHub) upgradeToWS(w http.ResponseWriter, r *http.Request) (net.Conn, error) {
// 	conn, _, _, err := ws.UpgradeHTTP(r, w)
// 	if err != nil {
// 		gbl.Log.Errorf("connection uograde failed: %s", err)

// 		w.WriteHeader(http.StatusUpgradeRequired)

// 		if _, err := w.Write([]byte("connection uograde failed")); err != nil {
// 			gbl.Log.Errorf("failed to write response with status %d: %s", http.StatusUpgradeRequired, err)
// 		}

// 		return nil, err
// 	}

// 	return conn, nil
// }

// func (wh *WebsocketsHub) wsHandler(w http.ResponseWriter, r *http.Request) {
// 	conn, err := s.upgradeToWS(w, r)
// 	if err != nil {
// 		gbl.Log.Errorf("connection upgrade failed: %s", err)
// 	}

// 	gbl.Log.Infof("new client connected: %s | %p\n", conn.RemoteAddr(), conn)

// 	s.Register(conn)
// }

// // func eventToPushEvent(event *totra.TokenTransaction) *collections.PushEvent {
// // 	return &collections.PushEvent{
// // 		EventType:       event.EventType,
// // 		CollectionName:  event.Collection.Name,
// // 		ContractAddress: event.Collection.ContractAddress,

// // 		NodeID:          event.NodeID,
// // 		Topic:           event.Topic,
// // 		TxHash:          event.TxHash,
// // 		TokenID:         event.TokenID,
// // 		ENSMetadata:     event.ENSMetadata,
// // 		PriceWei:        event.PriceWei,
// // 		CollectionColor: event.Collection.Colors.Primary,
// // 		TxItemCount:     event.TxLogCount,
// // 		Time:            event.Time,
// // 		From:            event.From,
// // 		FromENS:         event.FromENS,
// // 		To:              event.To,
// // 		ToENS:           event.ToENS,
// // 	}
// // }

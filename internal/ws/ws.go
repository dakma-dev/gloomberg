package ws

// import (
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"net"
// 	"net/http"
// 	"sort"
// 	"sync"
// 	"syscall"

// 	"github.com/benleb/gloomberg/internal/gbl"
// 	"github.com/benleb/gloomberg/internal/nemo/totra"
// 	"github.com/gobwas/ws"
// 	"github.com/gobwas/ws/wsutil"
// )

// type Client struct {
// 	id         string
// 	conn       net.Conn
// 	eventSream *WebsocketsServer
// }

// type WebsocketsServer struct {
// 	mu sync.RWMutex

// 	queueWsOutTokenTransactions *chan *totra.TokenTransaction

// 	listenHost string `mapstructure:"host"`
// 	listenPort uint   `mapstructure:"port"`

// 	// ws handling
// 	clientList []*Client
// 	clients    map[string]*Client

// 	out chan []byte
// }

// func New(listenHost string, listenPort uint, eventQueue *chan *totra.TokenTransaction) *WebsocketsServer {
// 	s := &WebsocketsServer{
// 		mu:                          sync.RWMutex{},
// 		queueWsOutTokenTransactions: eventQueue,

// 		listenHost: listenHost,
// 		listenPort: listenPort,

// 		clients: make(map[string]*Client),

// 		out: make(chan []byte),
// 	}

// 	// go s.writer()

// 	return s
// }

// func (s *WebsocketsServer) Start() {
// 	listenOn := fmt.Sprint(s.listenHost) + ":" + fmt.Sprint(s.listenPort)

// 	gbl.Log.Infof("✅ starting websocket server on %s", listenOn)

// 	go s.writer()

// 	http.Handle("/oldws", http.HandlerFunc(s.wsHandler))

// 	if err := http.ListenAndServe(listenOn, nil); err != nil { //nolint:gosec
// 		gbl.Log.Fatal(err)
// 	}
// }

// func (s *WebsocketsServer) ClientsConnected() int {
// 	s.mu.RLock()
// 	defer s.mu.RUnlock()

// 	return len(s.clients)
// }

// // func (s *WebsocketsServer) Broadcast() error {
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
// func (s *WebsocketsServer) writer() {
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

// func (s *WebsocketsServer) Register(conn net.Conn) *Client {
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
// func (s *WebsocketsServer) Remove(client *Client) {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	if removed := s.remove(client); !removed {
// 		gbl.Log.Errorf("removing client %s failed oO", client.id)

// 		return
// 	}
// }

// // mutex must be held.
// func (s *WebsocketsServer) remove(client *Client) bool {
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

// func (s *WebsocketsServer) upgradeToWS(w http.ResponseWriter, r *http.Request) (net.Conn, error) {
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

// func (s *WebsocketsServer) wsHandler(w http.ResponseWriter, r *http.Request) {
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

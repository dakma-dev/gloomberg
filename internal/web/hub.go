package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/charmbracelet/log"
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

	tmplFiles := []string{"www/event.html"}
	tmpls, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		gbl.Log.Error(err)
	}

	hub.templates = tmpls

	log.Print("")
	log.Printf("Loaded %d templates", len(tmplFiles))
	log.Printf("tmpl: %+v", hub.templates)

	hub.setupEventHandlers()

	go hub.broadcastTTx()

	return hub
}

// writer writes broadcast messages from chat.out channel.
func (wh *WsHub) broadcastTTx() {
	for parsedEvent := range wh.gb.SubscribeParsedEvents() {
		if len(wh.clients) == 0 {
			continue
		}

		// // event := event

		// // var wsEvent *totra.TokenTransaction
		// // copier.Copy(&wsEvent, &event)
		// // wsEvent.Collection.Source = wsEvent.Collection.Source.String()

		// // pushEvent := eventToPushEvent(ttx)

		// // data := struct {
		// // 	Time string
		// // }{
		// // 	Time: time.Now().Format("15:04:05"),
		// // }

		// // items := make(map[string][]int64)
		// items := []struct {
		// 	CollectionName  string
		// 	CollectionColor lipgloss.Color
		// 	Tokens          []struct {
		// 		ID         int64
		// 		Rank       int64
		// 		RankSymbol string
		// 	}
		// }{}

		// for contractAddress, tokenTransfers := range ttx.GetTransfersByContract() {
		// 	collection := wh.gb.CollectionDB.Collections[contractAddress]
		// 	if collection == nil {
		// 		continue
		// 	}

		// 	tokens := make([]struct {
		// 		ID         int64
		// 		Rank       int64
		// 		RankSymbol string
		// 	}, len(tokenTransfers))

		// 	for _, tokenTransfer := range tokenTransfers {
		// 		tokenID := tokenTransfer.Token.ID.Int64()
		// 		tokens = append(tokens, struct {
		// 			ID         int64
		// 			Rank       int64
		// 			RankSymbol string
		// 		}{
		// 			ID:         tokenID,
		// 			Rank:       wh.gb.Ranks[contractAddress][tokenID].Rank,
		// 			RankSymbol: wh.gb.Ranks[contractAddress][tokenID].GetRankSymbol(collection.Metadata.TotalSupply),
		// 		})
		// 	}

		// 	collectionName := collection.Name

		// 	items = append(items, struct {
		// 		CollectionName  string
		// 		CollectionColor lipgloss.Color
		// 		Tokens          []struct {
		// 			ID         int64
		// 			Rank       int64
		// 			RankSymbol string
		// 		}
		// 	}{
		// 		CollectionName:  collectionName,
		// 		CollectionColor: collection.Colors.Primary,
		// 		Tokens:          tokens,
		// 	})
		// }

		// if len(items) == 0 {
		// 	continue
		// }

		// // prepare data from the ttx for rendering
		// ttxData := struct {
		// 	TxHash     string
		// 	Action     string
		// 	ReceivedAt string
		// 	Typemoji   string
		// 	Price      string
		// 	Items      []struct {
		// 		CollectionName  string
		// 		CollectionColor lipgloss.Color
		// 		Tokens          []struct {
		// 			ID         int64
		// 			Rank       int64
		// 			RankSymbol string
		// 		}
		// 	}
		// 	EtherscanURL string
		// }{
		// 	TxHash:       ttx.TxHash.Hex(),
		// 	Action:       cases.Fold().String(ttx.Action.String()),
		// 	ReceivedAt:   ttx.ReceivedAt.Format("15:04:05"),
		// 	Typemoji:     ttx.Action.Icon(),
		// 	Price:        fmt.Sprintf("%6.3f", ttx.GetPrice().Ether()),
		// 	Items:        items,
		// 	EtherscanURL: ttx.GetEtherscanTxURL(),
		// }

		if parsedEvent == nil {
			continue
		}

		var rendered bytes.Buffer
		err := wh.templates.ExecuteTemplate(io.MultiWriter(&rendered), "event", parsedEvent)
		if err != nil {
			log.Errorf("❌ rendering template failed: %+v", err)
		}

		log.Debugf("rendered: %s", rendered.String())

		//
		// html minify for the poor...
		preparedMessage := strings.ReplaceAll(rendered.String(), "\n", "")

		betweenTagsWhitespace := regexp.MustCompile(`> *<`)
		preparedMessage = betweenTagsWhitespace.ReplaceAllString(preparedMessage, "><")

		multiWhitespace := regexp.MustCompile(` {2,}`)
		preparedMessage = multiWhitespace.ReplaceAllString(preparedMessage, " ")

		//
		// create ws message
		msg := MessageGeneric[EventPayload]{
			Type:     MsgNewSale,
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
		wh.RLock()
		clients := wh.clients
		wh.RUnlock()

		for client := range clients {
			client.egress <- marshalledMsg
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

	gbl.Log.Infof("client added: %s", client.conn.RemoteAddr().String())
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

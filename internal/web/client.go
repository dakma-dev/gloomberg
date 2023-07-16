package web

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"strings"
	"syscall"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/external"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gobwas/ws/wsutil"
)

var (
	// pongWait is how long we will await a pong response from client.
	pongWait = 10 * time.Second

	// pingInterval has to be less than pongWait, We cant multiply by 0.9 to get 90% of time.
	pingInterval = (pongWait * 9) / 10
)

type WsClient struct {
	clientID string

	// conn net.Conn
	conn   net.Conn
	ipInfo *external.IPInfo

	hub *WsHub

	// egress is used to avoid concurrent writes on the WebSocket
	egress chan json.RawMessage
}

// NewClient is used to initialize a new Client with all required values initialized.
func NewClient(conn net.Conn, hub *WsHub) *WsClient {
	client := &WsClient{
		conn:   conn,
		hub:    hub,
		egress: make(chan json.RawMessage),
	}

	// Create a unique ID for the Client
	client.clientID = client.createID()

	return client
}

func (wc *WsClient) createID() string {
	var address string
	if idx := strings.LastIndex(wc.conn.RemoteAddr().String(), ":"); idx == -1 {
		gbl.Log.Warn("could not find ':' in remote address")

		address = wc.conn.RemoteAddr().String()
	} else {
		address = wc.conn.RemoteAddr().String()[:idx]
	}

	data := []byte(address + "/" + wc.conn.RemoteAddr().Network())
	hash := crypto.Keccak256Hash(data)

	// create clientID from last 8 bytes of hash
	clientID := hash.Hex()[58:]

	return clientID
}

func (wc *WsClient) String() string {
	if wc.countryFlag() != "" {
		return wc.clientID + " " + wc.countryFlag()
	}

	return wc.clientID
}

func (wc *WsClient) countryFlag() string {
	if wc.ipInfo != nil {
		return internal.CountryFlags[wc.ipInfo.CountryCode]
	}

	return ""
}

// readMessages will start the client to read messages and handle them
// appropriately.
// This is suppose to be ran as a goroutine.
func (wc *WsClient) readMessages() {
	defer func() {
		// Graceful Close the Connection once this
		// function is done
		wc.hub.removeClient(wc)
	}()

	// Loop Forever
	for {
		// ReadMessage is used to read the next message in queue
		// in the connection
		msg, op, err := wsutil.ReadClientData(wc.conn)
		if err != nil {
			// If Connection is closed, we will Receive an error here
			gbl.Log.Debugf("error reading message: %v", err)

			break // Break the loop to close conn & Cleanup
		}

		log.Printf("op: %x | control: %v | data: %v | reserved: %v\n", op, op.IsControl(), op.IsData(), op.IsReserved())

		// Marshal incoming data into a Event struct
		var request Message
		if err := json.Unmarshal(msg, &request); err != nil {
			gbl.Log.Warnf("error marshalling message: %v", err)

			continue
		}

		// Route the Event
		if err := wc.hub.routeEvent(request, wc); err != nil {
			gbl.Log.Warn("Error handeling Message: ", err)
		}
	}
}

// // pongHandler is used to handle PongMessages for the Client.
// func (wc *WsClient) pongHandler(pongMsg string) error {
// 	// Current time + Pong Wait time
// 	log.Println("pong")

// 	return wc.conn.SetReadDeadline(time.Now().Add(pongWait))
// }

// writeMessages is a process that listens for new messages to output to the Client.
func (wc *WsClient) writeMessages() {
	// Create a ticker that triggers a ping at given interval
	ticker := time.NewTicker(pingInterval)
	defer func() {
		ticker.Stop()
		// Graceful close if this triggers a closing
		wc.hub.removeClient(wc)
	}()

	// for {
	// 	select {
	// 	case message, ok := <-wc.egress:
	for message := range wc.egress {
		// // Ok will be false Incase the egress channel is closed
		// if !ok {
		// 	// Manager has closed this connection channel, so communicate that to frontend
		// 	// if err := wc.conn.WriteMessage(websocket.CloseMessage, nil); err != nil {
		// 	if err := wsutil.WriteServerMessage(wc.conn, ws.OpClose, []byte{}); err != nil {
		// 		// Log that the connection is closed and the reason
		// 		log.Println("connection closed: ", err)
		// 	}
		// 	// Return to close the goroutine
		// 	return
		// }

		// data, err := json.Marshal(message)
		// if err != nil {
		// 	gbl.Log.Error("error marshalling message: %+v", message)

		// 	continue
		// }

		if err := wsutil.WriteServerText(wc.conn, message); err != nil {
			if errors.Is(err, syscall.EPIPE) {
				gbl.Log.Errorf("client %s disconnected: %s", wc.conn.LocalAddr().String(), err.Error())

				// remove client
				wc.conn.Close()
				wc.hub.removeClient(wc)
			} else {
				gbl.Log.Errorf("sending event to client %v failed: %s | event: %+v", wc, err.Error(), message)
			}
		}

		gbl.Log.Debugf("sent message: %+v", string(message))

		// case <-ticker.C:
		// 	log.Println("ping")
		// 	// Send the Ping
		// 	if err := wsutil.WriteServerMessage(wc.conn, ws.OpPing, []byte{}); err != nil {
		// 		log.Println("writemsg: ", err)
		// 		return // return to break this goroutine triggeing cleanup
		// 	}
	}
}

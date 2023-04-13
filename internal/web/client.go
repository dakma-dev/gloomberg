package web

import (
	"encoding/json"
	"log"
	"net"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

var (
	// pongWait is how long we will await a pong response from client
	pongWait = 10 * time.Second

	// pingInterval has to be less than pongWait, We cant multiply by 0.9 to get 90% of time
	// The reason why it has to be less than PingRequency is becuase otherwise it will send a new Ping before getting response
	pingInterval = (pongWait * 9) / 10
)

type WsClient struct {
	// conn net.Conn
	conn net.Conn

	hub *WsHub

	// egress is used to avoid concurrent writes on the WebSocket
	egress chan Event
}

// NewClient is used to initialize a new Client with all required values initialized
func NewClient(conn net.Conn, hub *WsHub) *WsClient {
	return &WsClient{
		conn:   conn,
		hub:    hub,
		egress: make(chan Event),
	}
}

// readMessages will start the client to read messages and handle them
// appropriatly.
// This is suppose to be ran as a goroutine
func (wc *WsClient) readMessages() {
	defer func() {
		// Graceful Close the Connection once this
		// function is done
		wc.hub.removeClient(wc)
	}()

	// TODO: Set Max Size of Messages in Bytes
	// wc.connection.SetReadLimit(512)

	// // Configure Wait time for Pong response, use Current time + pongWait
	// // This has to be done here to set the first initial timer.
	// if err := wc.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
	// 	log.Println(err)

	// 	return
	// }

	// // Configure how to handle Pong responses
	// wc.connection.SetPongHandler(wc.pongHandler)

	// Loop Forever
	for {
		// ReadMessage is used to read the next message in queue
		// in the connection
		// _, payload, err := wc.connection.ReadMessage()
		msg, op, err := wsutil.ReadClientData(wc.conn)
		// msgs, err := wsutil.ReadClientMessage(reader, nil)
		if err != nil {
			// If Connection is closed, we will Recieve an error here
			// We only want to log Strange errors, but simple Disconnection
			// if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			// 	log.Printf("error reading message: %v", err)
			// }
			log.Printf("error reading message: %v", err)

			break // Break the loop to close conn & Cleanup
		}

		log.Printf("op: %x | control: %v | data: %v | reserved: %v\n", op, op.IsControl(), op.IsData(), op.IsReserved())

		// Marshal incoming data into a Event struct
		var request Event
		if err := json.Unmarshal(msg, &request); err != nil {
			log.Printf("error marshalling message: %v", err)
			break // Breaking the connection here might be harsh xD
		}
		// Route the Event
		if err := wc.hub.routeEvent(request, wc); err != nil {
			log.Println("Error handeling Message: ", err)
		}
	}
}

// pongHandler is used to handle PongMessages for the Client
func (wc *WsClient) pongHandler(pongMsg string) error {
	// Current time + Pong Wait time
	log.Println("pong")

	return wc.conn.SetReadDeadline(time.Now().Add(pongWait))
}

// writeMessages is a process that listens for new messages to output to the Client
func (wc *WsClient) writeMessages() {
	// Create a ticker that triggers a ping at given interval
	ticker := time.NewTicker(pingInterval)
	defer func() {
		ticker.Stop()
		// Graceful close if this triggers a closing
		wc.hub.removeClient(wc)
	}()

	for {
		select {
		case message, ok := <-wc.egress:
			// Ok will be false Incase the egress channel is closed
			if !ok {
				// Manager has closed this connection channel, so communicate that to frontend
				// if err := wc.conn.WriteMessage(websocket.CloseMessage, nil); err != nil {
				if err := wsutil.WriteServerMessage(wc.conn, ws.OpClose, []byte{}); err != nil {
					// Log that the connection is closed and the reason
					log.Println("connection closed: ", err)
				}
				// Return to close the goroutine
				return
			}

			data, err := json.Marshal(message)
			if err != nil {
				log.Println(err)
				return // closes the connection, should we really
			}
			// Write a Regular text message to the connection
			if err := wsutil.WriteServerMessage(wc.conn, ws.OpText, data); err != nil {
				log.Println(err)
			}
			log.Println("sent message")

			// case <-ticker.C:
			// 	log.Println("ping")
			// 	// Send the Ping
			// 	if err := wsutil.WriteServerMessage(wc.conn, ws.OpPing, []byte{}); err != nil {
			// 		log.Println("writemsg: ", err)
			// 		return // return to break this goroutine triggeing cleanup
			// 	}
		}
	}
}

package web

import "encoding/json"

// Message is the Messages sent over the websocket
// Used to differ between different actions
type Message struct {
	// Type is the message type sent
	Type string `json:"type"`
	// Payload is the data Based on the Type
	Payload any `json:"payload"`
}

// MessageHandler is a function signature that is used to affect messages on the socket and triggered
// depending on the type
type MessageHandler func(message Message, c *WsClient) error

const (
	// MessageSendMessage is the message name for new chat messages sent
	// MsgSendMessage = "send_message"

	MsgNewSale  = "new_event"
	MsgGasPrice = "gas_price"

	MsgCommand = "cmd"
)

type MessagePayload json.RawMessage

// payloads for message types
type NewEventMessage struct {
	Message string `json:"message"`
	From    string `json:"from"`
}

// payloads for message types
type GasPriceMessage struct {
	Normal float64 `json:"normal"`
	Fast   float64 `json:"fast"`
}

// SendMessageMessage is the payload sent in the
// send_message message
type SendMessageMessage struct {
	Message string `json:"message"`
	From    string `json:"from"`
}

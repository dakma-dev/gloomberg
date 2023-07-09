package web

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
)

// Message is the Messages sent over the websocket
// Used to differ between different actions.
type Message struct {
	// Type is the message type sent
	Type string `json:"type"`

	// Payload is the data Based on the Type
	Payload *json.RawMessage `json:"payload"`
}

type MessageGeneric[T interface{}] struct {
	// Type is the message type sent
	Type string `json:"type"`

	// Payload is the data Based on the Type
	Payload T `json:"payload"`

	// Current Gas Price
	GasPrice uint64 `json:"gas_price"`
}

// func (m MessageGeneric[any]) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(m)
// }
// func (m *MessageGeneric[any]) UnmarshalJSON(data []byte) error {
// 	return json.Unmarshal(data, m)
// }

// MessageHandler is a function signature that is used to affect messages on the socket and triggered
// depending on the type.
type MessageHandler func(message Message, c *WsClient) error

const (
	// MessageSendMessage is the message name for new chat messages sent
	// MsgSendMessage = "send_message".

	MsgNewSale  = "new_event"
	MsgGasPrice = "gas_price"

	MsgCommand = "cmd"
)

type UserInfo struct {
	Name    string         `json:"username"`
	Address common.Address `json:"address"`
}

// payloads for message types.
type EventPayload struct {
	Message string `json:"message"`
}

// 	// action performed by the tx
// 	Action totra.TxType `json:"action"`

// 	// the sender of the tx
// 	From *UserInfo `json:"from"`

// 	// the amount of eth/weth transferred in the tx
// 	AmountPaid float64 `json:"amount_paid"`

// 	// marketplace the tx was performed on
// 	Marketplace string `json:"marketplace"`

// 	// token transfers parsed from the tx logs
// 	Transfers []*totra.TokenTransfer `json:"transfers"`

// 	Highlight bool `json:"highlight"`
// }

// // payloads for message types.
// type GasPriceMessage struct {
// 	Normal float64 `json:"normal"`
// 	Fast   float64 `json:"fast"`
// }

// // SendMessageMessage is the payload sent in the
// // send_message message.
// type SendMessageMessage struct {
// 	Message string `json:"message"`
// 	From    string `json:"from"`
// }

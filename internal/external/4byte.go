package external

import (
	"time"
)

type MethodSignatureResponse struct {
	Count    int               `json:"count"`
	Next     interface{}       `json:"next"`
	Previous interface{}       `json:"previous"`
	Results  []MethodSignature `json:"results"`
}
type EventSignatureResponse struct {
	Count    int              `json:"count"`
	Next     interface{}      `json:"next"`
	Previous interface{}      `json:"previous"`
	Results  []EventSignature `json:"results"`
}

type Signature interface {
	SignatureID() int
}

// 	TextSignature() string
// 	HexSignature() string
// 	BytesSignature() string
// }

type MethodSignature struct {
	ID             int    `json:"id"`
	TextSignature  string `json:"text_signature"`
	HexSignature   string `json:"hex_signature"`
	BytesSignature string `json:"bytes_signature"`
}

func (m MethodSignature) SignatureID() int {
	return m.ID
}

type EventSignature struct {
	MethodSignature
	CreatedAt time.Time `json:"created_at"`
}

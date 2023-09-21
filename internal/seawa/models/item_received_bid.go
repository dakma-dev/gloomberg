package models

import "time"

type ItemReceivedBid struct {
	EventType string                 `json:"event_type" mapstructure:"event_type"`
	SentAt    time.Time              `json:"sent_at"    mapstructure:"sent_at"`
	Payload   ItemReceivedBidPayload `json:"payload"    mapstructure:"payload"`
	// Payload ItemGenericPayload `json:"payload"    mapstructure:"payload"`

	Other map[string]interface{} `mapstructure:",remain"`
}

type ItemReceivedBidPayload struct {
	Item         `json:"item"    mapstructure:"item"`
	EventPayload `json:"payload" mapstructure:",squash"`

	CreatedDate time.Time `json:"created_date" mapstructure:"created_date"`

	Other map[string]interface{} `mapstructure:",remain"`
}

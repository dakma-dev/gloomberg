package models

import (
	"time"
)

type ItemReceivedBid struct {
	// EventType string                 `json:"event_type" mapstructure:"event_type"`
	Event   string                 `json:"event_type" mapstructure:"event_type"`
	SentAt  time.Time              `json:"sent_at"    mapstructure:"sent_at"`
	Payload ItemReceivedBidPayload `json:"payload"    mapstructure:"payload"`

	Other map[string]interface{} `mapstructure:",remain"`
}

type ItemReceivedBidPayload struct {
	EventPayload `mapstructure:",squash"`

	Item Item `json:"item" mapstructure:"item"`

	CreatedDate time.Time `json:"created_date" mapstructure:"created_date"`

	Other map[string]interface{} `mapstructure:",remain"`
}

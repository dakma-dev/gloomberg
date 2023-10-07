package models

import "time"

type ItemListed struct {
	EventType string            `json:"event_type" mapstructure:"event_type"`
	SentAt    time.Time         `json:"sent_at"    mapstructure:"sent_at"`
	Payload   ItemListedPayload `json:"payload"    mapstructure:"payload"`

	Other map[string]interface{} `mapstructure:",remain"`
}

type ItemListedPayload struct {
	EventPayload `mapstructure:",squash"`
	Item         Item        `json:"item"            mapstructure:"item"`
	IsPrivate    bool        `json:"is_private"      mapstructure:"is_private"`
	ListingDate  time.Time   `json:"listing_date"    mapstructure:"listing_date"`
	ListingType  interface{} `json:"listing_type"    mapstructure:"listing_type"`

	Other map[string]interface{} `mapstructure:",remain"`
}

package models

import "time"

type TraitOffer struct {
	EventType string                 `json:"event_type" mapstructure:"event_type"`
	SentAt    time.Time              `json:"sent_at"    mapstructure:"sent_at"`
	Payload   collectionOfferPayload `json:"payload"    mapstructure:"payload"`

	Other map[string]interface{} `mapstructure:",remain"`
}

type TraitOfferPayload struct {
	collectionOfferPayload

	TraitCriteria TraitCriteria `json:"trait_criteria"`
}

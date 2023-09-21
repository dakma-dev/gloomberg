package models

import "time"

type CollectionOffer struct {
	EventType string                 `json:"event_type" mapstructure:"event_type"`
	SentAt    time.Time              `json:"sent_at"    mapstructure:"sent_at"`
	Payload   collectionOfferPayload `json:"payload"    mapstructure:"payload"`

	Other map[string]interface{} `mapstructure:",remain"`
}

type collectionOfferPayload struct {
	CollectionCriteria CollectionCriteria `json:"collection_criteria"     mapstructure:"collection_criteria"`
	ContractCriteria   ContractCriteria   `json:"asset_contract_criteria" mapstructure:"asset_contract_criteria"`
	EventPayload       `json:"payload"                 mapstructure:",squash"`

	CreatedDate time.Time `json:"created_date" mapstructure:"created_date"`

	Other map[string]interface{} `mapstructure:",remain"`
}

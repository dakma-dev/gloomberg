package models

import "time"

type ItemGeneric struct {
	EventType string             `json:"event_type" mapstructure:"event_type"`
	SentAt    time.Time          `json:"sent_at"    mapstructure:"sent_at"`
	Payload   ItemGenericPayload `json:"payload"    mapstructure:"payload"`

	Other map[string]interface{} `mapstructure:",remain"`
}

type ItemGenericPayload struct {
	Item        `json:"item"         mapstructure:"item"`
	CreatedDate time.Time `json:"created_date" mapstructure:"created_date"`

	CollectionCriteria `json:"collection_criteria"     mapstructure:"collection_criteria,omitempty"`
	ContractCriteria   `json:"asset_contract_criteria" mapstructure:"asset_contract_criteria,omitempty"`
	TraitCriteria      `json:"trait_criteria"          mapstructure:"trait_criteria,omitempty"`

	EventPayload `json:"payload" mapstructure:",squash"`

	IsPrivate   bool        `json:"is_private"   mapstructure:"is_private"`
	ListingDate time.Time   `json:"listing_date" mapstructure:"listing_date"`
	ListingType interface{} `json:"listing_type" mapstructure:"listing_type"`

	Other map[string]interface{} `mapstructure:",remain"`
}

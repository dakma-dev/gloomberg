package models

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type CollectionOffer struct {
	Event   string    `json:"event"   mapstructure:"event"`
	SentAt  time.Time `json:"sent_at" mapstructure:"sent_at"`
	Payload struct {
		EventType string                 `json:"event_type" mapstructure:"event_type"`
		Payload   collectionOfferPayload `json:"payload"    mapstructure:"payload"`
	} `json:"payload" mapstructure:"payload"`
}

type collectionOfferPayload struct {
	EventPayload `mapstructure:",squash"`

	AssetContractCriteria struct {
		Address common.Address `json:"address" mapstructure:"address"`
	} `json:"asset_contract_criteria" mapstructure:"asset_contract_criteria"`

	CollectionCriteria CollectionSlug `json:"collection_criteria" mapstructure:"collection_criteria"`

	CreatedDate time.Time `json:"created_date" mapstructure:"created_date"`

	Other map[string]interface{} `mapstructure:",remain"`
}

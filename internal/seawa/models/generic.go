package models

import (
	"fmt"
	"math/big"
	"time"

	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/ethereum/go-ethereum/common"
)

type GeneralEvent struct {
	EventType string             `json:"event_type" mapstructure:"event_type"`
	SentAt    time.Time          `json:"sent_at"    mapstructure:"sent_at"`
	Payload   ItemGenericPayload `json:"payload"    mapstructure:"payload"`

	Other map[string]interface{} `mapstructure:",remain"`

	address common.Address `mapstructure:"-"`
}

func (e *GeneralEvent) ItemName() string {
	var name string

	switch {
	case e.Payload.Item.Name != "":
		name = e.Payload.Item.Name
	case degendb.GetEventType(e.EventType) == degendb.CollectionOffer:
		name = e.Payload.Slug
	case e.Payload.NftID != nil:
		name = e.Payload.NftID.String()
	default:
		name = fmt.Sprintf("%+v", e.Payload)
	}

	return name
}

func (e *GeneralEvent) ItemNameLink() string {
	if e.Payload.Item.Permalink != "" {
		return style.TerminalLink(e.Payload.Permalink, e.ItemName())
	}

	return e.ItemName()
}

func (e *GeneralEvent) ContractAddress() *common.Address {
	if e.address != (common.Address{}) {
		return &e.address
	}

	if degendb.GetEventType(e.EventType) == degendb.CollectionOffer {
		e.address = e.Payload.Address
	} else {
		e.address = e.Payload.Item.NftID.ContractAddress()
	}

	return &e.address
}

func (e *GeneralEvent) BasePrice() *price.Price {
	basePrice := price.NewPrice(big.NewInt(0))
	if degendb.GetEventType(e.EventType) != degendb.MetadataUpdated {
		basePrice = price.NewPrice(e.Payload.BasePrice)
	}

	return basePrice
}

type ItemGenericPayload struct {
	Item        `json:"item"         mapstructure:"item,omitempty"`
	CreatedDate time.Time `json:"created_date" mapstructure:"created_date,omitempty"`

	CollectionSlug   `json:"collection_criteria"     mapstructure:"collection_criteria,omitempty"`
	ContractCriteria `json:"asset_contract_criteria" mapstructure:"asset_contract_criteria,omitempty"`
	TraitCriteria    `json:"trait_criteria"          mapstructure:"trait_criteria,omitempty"`

	EventPayload `json:"payload" mapstructure:",squash,omitempty"`

	IsPrivate   bool        `json:"is_private"   mapstructure:"is_private"`
	ListingDate time.Time   `json:"listing_date" mapstructure:"listing_date"`
	ListingType interface{} `json:"listing_type" mapstructure:"listing_type"`

	Other map[string]interface{} `mapstructure:",remain"`
}

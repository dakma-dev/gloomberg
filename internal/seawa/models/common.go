package models

import (
	"math/big"
	"time"

	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/ethereum/go-ethereum/common"
)

type Item struct {
	Chain struct {
		Name string `json:"name" mapstructure:"name"`
	} `json:"chain" mapstructure:"chain"`
	Metadata struct {
		AnimationURL string `json:"animation_url" mapstructure:"animation_url"`
		ImageURL     string `json:"image_url"     mapstructure:"image_url"`
		MetadataURL  string `json:"metadata_url"  mapstructure:"metadata_url"`
		Name         string `json:"name"          mapstructure:"name"`
	} `json:"metadata"`
	NftID     NftID  `json:"nft_id"    mapstructure:"nft_id"`
	Permalink string `json:"permalink" mapstructure:"permalink"`

	Other map[string]interface{} `mapstructure:",remain"`
}

type EventPayload struct {
	OrderHash common.Hash `json:"order_hash" mapstructure:"order_hash"`

	EventTimestamp time.Time `json:"event_timestamp" mapstructure:"event_timestamp"`
	ExpirationDate time.Time `json:"expiration_date" mapstructure:"expiration_date"`

	Collection CollectionSlug `json:"collection" mapstructure:"collection"`

	Maker Account `json:"maker,omitempty" mapstructure:"maker,omitempty"`
	Taker Account `json:"taker,omitempty" mapstructure:"taker,omitempty"`

	BasePrice    *big.Int     `json:"base_price"    mapstructure:"base_price"`
	Quantity     int          `json:"quantity"      mapstructure:"quantity"`
	PaymentToken PaymentToken `json:"payment_token" mapstructure:"payment_token"`

	ProtocolAddress common.Address `json:"protocol_address,omitempty" mapstructure:"protocol_address,omitempty"`
	ProtocolData    ProtocolData   `json:"protocol_data,omitempty"    mapstructure:"protocol_data,omitempty"`

	Other map[string]interface{} `mapstructure:",remain"`
}

func (ep EventPayload) GetPrice() *price.Price {
	return price.NewPrice(ep.BasePrice)
}

type CollectionSlug struct {
	Slug string `json:"slug" mapstructure:"slug"`
}

func (c *CollectionSlug) GetCollectionSlug() string {
	return c.Slug
}

type Account struct {
	Address common.Address `json:"address"`
}

type PaymentToken struct {
	Address  common.Address `json:"address"`
	Decimals int            `json:"decimals"`
	EthPrice string         `json:"eth_price"`
	Name     string         `json:"name"`
	Symbol   string         `json:"symbol"`
	UsdPrice string         `json:"usd_price"`
}

type ProtocolData struct {
	Parameters struct {
		ConduitKey                      string          `json:"conduitKey"`
		Consideration                   []Consideration `json:"consideration,omitempty"         mapstructure:"consideration,omitempty"`
		Counter                         interface{}     `json:"counter"`
		EndTime                         time.Time       `json:"endTime"`
		Offer                           []Consideration `json:"offer,omitempty"                 mapstructure:"offer,omitempty"`
		Offerer                         string          `json:"offerer"`
		OrderType                       int             `json:"orderType"`
		Salt                            string          `json:"salt,omitempty"                  mapstructure:"salt,omitempty"`
		StartTime                       time.Time       `json:"startTime"`
		TotalOriginalConsiderationItems int             `json:"totalOriginalConsiderationItems"`
		Zone                            common.Address  `json:"zone"`
		ZoneHash                        common.Hash     `json:"zoneHash"`
	} `json:"parameters,omitempty"                        mapstructure:"parameters,omitempty"`
	Signature                             string `json:"signature"`
	UseLazyMintAdapterForSharedStorefront bool   `json:"use_lazy_mint_adapter_for_shared_storefront"`
}

type Consideration struct {
	EndAmount            *big.Int       `json:"endAmount"            mapstructure:"endAmount"`
	IdentifierOrCriteria string         `json:"identifierOrCriteria" mapstructure:"identifierOrCriteria"`
	ItemType             int            `json:"itemType"             mapstructure:"itemType"`
	Recipient            common.Address `json:"recipient,omitempty"  mapstructure:"recipient,omitempty"`
	StartAmount          *big.Int       `json:"startAmount"          mapstructure:"startAmount"`
	Token                common.Address `json:"token"                mapstructure:"token"`
}

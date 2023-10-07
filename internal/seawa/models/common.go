package models

import (
	"math/big"
	"time"

	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/ethereum/go-ethereum/common"
)

type Item struct {
	Metadata  `json:"metadata"  mapstructure:"metadata"`
	NftID     `json:"nft_id"    mapstructure:"nft_id"`
	Chain     Chain  `json:"chain"     mapstructure:"chain"`
	Permalink string `json:"permalink" mapstructure:"permalink"`

	Other map[string]interface{} `mapstructure:",remain"`
}

type Chain struct {
	Name string `json:"name" mapstructure:"name"`
}

//
// criteria structs
//

type CollectionCriteria struct {
	Slug string `json:"slug" mapstructure:"slug"`
}

func (c *CollectionCriteria) String() string {
	return c.Slug
}

type ContractCriteria struct {
	Address common.Address `json:"address" mapstructure:"address"`
}

type TraitCriteria struct {
	TraitName string `json:"trait_name"`
	TraitType string `json:"trait_type"`
}

//
// metadata structs
//

type Metadata struct {
	Name            string          `json:"name"             mapstructure:"name"`
	Description     string          `json:"description"      mapstructure:"description"`
	BackgroundColor string          `json:"background_color" mapstructure:"background_color"`
	AnimationURL    string          `json:"animation_url"    mapstructure:"animation_url"`
	ImageURL        string          `json:"image_url"        mapstructure:"image_url"`
	MetadataURL     string          `json:"metadata_url"     mapstructure:"metadata_url"`
	Traits          []degendb.Trait `json:"traits"           mapstructure:"traits"`
}

type EventPayload struct {
	OrderHash common.Hash `json:"order_hash" mapstructure:"order_hash"`

	EventTimestamp time.Time `json:"event_timestamp" mapstructure:"event_timestamp"`
	ExpirationDate time.Time `json:"expiration_date" mapstructure:"expiration_date"`

	CollectionCriteria `json:"collection" mapstructure:"collection"`

	Maker Account `json:"maker,omitempty" mapstructure:"maker,omitempty"`
	Taker Account `json:"taker,omitempty" mapstructure:"taker,omitempty"`

	BasePrice    *big.Int `json:"base_price"    mapstructure:"base_price"`
	Quantity     int      `json:"quantity"      mapstructure:"quantity"`
	PaymentToken `json:"payment_token" mapstructure:"payment_token"`

	ProtocolAddress common.Address `json:"protocol_address,omitempty" mapstructure:"protocol_address,omitempty"`
	ProtocolData    ProtocolData   `json:"protocol_data,omitempty"    mapstructure:"protocol_data,omitempty"`

	// "item" is weird Oo it afaik just exists in the *collection*Offer event and is always empty!?
	Item interface{} `json:"item,omitempty" mapstructure:"item,omitempty"`

	Other map[string]interface{} `mapstructure:",remain"`
}

func (ep EventPayload) GetPrice() *price.Price {
	if ep.BasePrice == nil {
		return price.NewPrice(big.NewInt(0))
	}

	return price.NewPrice(ep.BasePrice)
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
	Parameters                            `json:"parameters,omitempty"                        mapstructure:"parameters,omitempty"`
	Signature                             string `json:"signature"                                   mapstructure:"signature"`
	UseLazyMintAdapterForSharedStorefront bool   `json:"use_lazy_mint_adapter_for_shared_storefront" mapstructure:"use_lazy_mint_adapter_for_shared_storefront"`
}

type Parameters struct {
	ConduitKey                      string          `json:"conduitKey"                      mapstructure:"conduitKey"`
	Consideration                   []Consideration `json:"consideration,omitempty"         mapstructure:"consideration,omitempty"`
	Counter                         interface{}     `json:"counter"                         mapstructure:"counter"`
	EndTime                         time.Time       `json:"endTime"                         mapstructure:"endTime"`
	Offer                           []Consideration `json:"offer,omitempty"                 mapstructure:"offer,omitempty"`
	Offerer                         string          `json:"offerer"                         mapstructure:"offerer"`
	OrderType                       int             `json:"orderType"                       mapstructure:"orderType"`
	Salt                            string          `json:"salt,omitempty"                  mapstructure:"salt,omitempty"`
	StartTime                       time.Time       `json:"startTime"                       mapstructure:"startTime"`
	TotalOriginalConsiderationItems int             `json:"totalOriginalConsiderationItems" mapstructure:"totalOriginalConsiderationItems"`
	Zone                            common.Address  `json:"zone"                            mapstructure:"zone"`
	ZoneHash                        common.Hash     `json:"zoneHash"                        mapstructure:"zoneHash"`
}

type Consideration struct {
	EndAmount            *big.Int       `json:"endAmount"            mapstructure:"endAmount"`
	IdentifierOrCriteria string         `json:"identifierOrCriteria" mapstructure:"identifierOrCriteria"`
	ItemType             int            `json:"itemType"             mapstructure:"itemType"`
	Recipient            common.Address `json:"recipient,omitempty"  mapstructure:"recipient,omitempty"`
	StartAmount          *big.Int       `json:"startAmount"          mapstructure:"startAmount"`
	Token                common.Address `json:"token"                mapstructure:"token"`
}

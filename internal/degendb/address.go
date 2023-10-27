package degendb

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Address struct {
	// Address is the ethereum address for this wallet/collection
	HexAddress string `bson:"_id" json:"hex_address"`

	// Address is the ethereum Address for this wallet/collection
	Address common.Address `json:"address"`

	// Type is the type of this wallet/collection, e.g. EOA or contract
	Type       string `bson:"type,omitempty"        json:"type,omitempty"`
	IsContract bool   `bson:"is_contract,omitempty" json:"is_contract,omitempty"`

	// Domain is the (main/reverseable) domain (e.g. ENS) for this wallet/collection
	Domain string `bson:"domain,omitempty" json:"domain"`

	// Name is the name of this wallet/collection
	Name string `bson:"name,omitempty" json:"name"`

	// Tags is a list of tags associated with this wallet/collection
	Tags []Tag `bson:"tags,omitempty" json:"tags"`

	//
	// Collection data

	// Slugs are the collection slugs of the collection
	Slugs Slugs `bson:"slugs,omitempty" json:"slugs,omitempty"`

	// Description is the description of the collection
	Description string `bson:"description,omitempty" json:"description,omitempty"`

	// TotalSupply is the total supply of the collection
	TotalSupply int `bson:"total_supply,omitempty" json:"total_supply,omitempty"`

	// ImageURL is the URL of the image of the collection
	ImageURL string `bson:"image_url,omitempty" json:"image_url,omitempty"`

	// ExternalURL is the URL of the collection
	ExternalURL string `bson:"external_url,omitempty" json:"external_url,omitempty"`

	// CreatedAt is the time this address was created in the db
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	// UpdatedAt is the time this address was last updated in the db
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`

	// creates to many writes probably
	// // LastSeen is the time this wallet was last seen
	// LastSeen int64 `bson:"last_seen,omitempty" json:"last_seen"`
}

func (a *Address) String() string {
	return a.Address.Hex()
}

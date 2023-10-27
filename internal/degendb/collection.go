package degendb

import (
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
)

type Collection struct {
	// // ID is the unique identifier for this token
	// ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	// Address is the ethereum address for this wallet/collection
	HexAddress string `bson:"_id" json:"hex_address"`

	// Address is the ethereum Address for this wallet/collection
	Address common.Address `json:"address"`

	// Tags is a list of tags associated with this wallet/collection
	Tags []Tag `bson:"tags,omitempty" json:"tags"`

	//
	// Collection metadata
	Metadata CollectionMetadata `bson:"metadata,omitempty" json:"metadata,omitempty"`

	//
	// data related/created by Gloomberg

	// Source keeps track of where this collection came from (stream, wallet, config)
	Source CollectionSource `bson:"source,omitempty" json:"source,omitempty"`

	// IgnorePrinting is a flag to ignore this collection when printing (to the stream)
	IgnorePrinting bool `bson:"ignore_printing,omitempty" json:"ignore_printing,omitempty"`

	// Colors are the generated colors of the collection --> maybe generate on the fly?
	Colors struct {
		Primary   lipgloss.Color `bson:"primary,omitempty"   json:"primary,omitempty"`
		Secondary lipgloss.Color `bson:"secondary,omitempty" json:"secondary,omitempty"`
	} `bson:"colors,omitempty" json:"colors,omitempty"`

	// CreatedAt is the time this collection was created in our db
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	// UpdatedAt is the time this collection was last updated in our db
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`

	// // creates to many writes probably
	// // LastSeen is the time this wallet was last seen
	// LastSeen int64 `bson:"last_seen,omitempty" json:"last_seen"`
}

func (c *Collection) String() string {
	name := c.HexAddress
	if c.Metadata.Name != "" {
		name = c.Metadata.Name
	}

	return name
}

type CollectionMetadata struct {
	// Name is the name of this wallet/collection
	Name string `bson:"name,omitempty" json:"name"`

	// Slugs are the collection slugs of the collection
	Slugs Slugs `bson:"slugs,omitempty" json:"slugs,omitempty"`

	// Description is the description of the collection
	Description string `bson:"description,omitempty" json:"description,omitempty"`

	// TotalSupply is the total supply of the collection
	TotalSupply uint64 `bson:"total_supply,omitempty" json:"total_supply,omitempty"`

	// ImageURL is the URL of the image of the collection
	ImageURL string `bson:"image_url,omitempty" json:"image_url,omitempty"`

	// Symbol is the symbol/shortcode of the collection
	Symbol string `bson:"symbol,omitempty" json:"symbol,omitempty"`

	// TokenURI is the base URI for the tokens metadata
	TokenURI string `bson:"token_uri,omitempty" json:"token_uri,omitempty"`

	// ExternalURL is the URL of the collection
	ExternalURL string `bson:"external_url,omitempty" json:"external_url,omitempty"`
}

package degendb

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Token struct {
	// ID is the unique identifier for this token
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	// Collection is the collection this token belongs to
	Address Address `bson:"collection,omitempty" json:"collection"`

	// CollectionSlugs are the collection slugs of the collection this token belongs to
	CollectionSlugs Slugs `bson:"collection_slugs,omitempty" json:"collection_slugs"`

	// ContractAddress is the address of the contract this token belongs to
	ContractAddress string `bson:"contract_address,omitempty" json:"contract_address"`

	// Token ID used as in the collection
	TokenID int64 `bson:"token_id,omitempty" json:"token_id"`

	// Name is the name of the token
	Name string `bson:"name,omitempty" json:"name"`

	// Ranks are the ranks of the token
	Rank Rank `bson:"ranks,omitempty" json:"ranks"`

	// Score is the score of the token calculated via the open-rarity algorithm
	Score float64 `bson:"score,omitempty" json:"score"`

	// Metadata is the metadata of the token
	Metadata []MetadataAttribute `bson:"metadata,omitempty" json:"metadata"`

	// CreatedAt is the time this token was created
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at"`

	// UpdatedAt is the time this token was last updated
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at"`
}

type MetadataAttribute struct {
	// Name is the name of the attribute
	Name string `json:"name"`

	// Value is the value of the attribute
	// Value interface{} `json:"value"`
	Value string `json:"value"`

	// Data type of the attribute
	TraitType string `json:"-"`
}

type Rank struct {
	// OpenSea is the OpenSea rank of the token
	OpenSea int64 `bson:"opensea,omitempty" json:"opensea"`

	// Blur is the Blur rank of the token
	Blur int64 `bson:"blur,omitempty" json:"blur"`
}

type Slugs struct {
	// OpenSea is the OpenSea slug of the token
	OpenSea string `bson:"opensea,omitempty" json:"opensea"`

	// Blur is the Blur slug of the token
	Blur string `bson:"blur,omitempty" json:"blur"`
}

type TokenMetadata struct {
	Attributes      []MetadataAttribute `json:"-"`
	Name            string              `json:"name"`
	Description     string              `json:"description"`
	Image           string              `json:"image"`
	TokenID         int64               `json:"token_id"`
	ContractAddress common.Address      `json:"-"`
	Score           Score               `json:"score,omitempty"`
}

type Score struct {
	TokenID       int64   `json:"token_id"`
	Rank          int64   `json:"rank"`
	Score         float64 `json:"score"`
	TokenFeatures struct {
		UniqueAttributeCount int `json:"unique_attribute_count"`
	} `json:"token_features"`
	TokenMetadata map[string]interface{} `json:"token_metadata"`
}

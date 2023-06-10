package degendb

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Collection struct {
	// ID is the unique identifier for this collection
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`

	// Address is the address of the contract
	Address common.Address `bson:"address,omitempty" json:"address,omitempty"`

	// Name is the name of the collection
	Name string `bson:"name,omitempty" json:"name,omitempty"`

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

	// CreatedAt is the time this collection was created
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	// UpdatedAt is the time this collection was last updated
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

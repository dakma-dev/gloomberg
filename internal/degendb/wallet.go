package degendb

import (
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Wallet struct {
	// ID is the unique identifier for this wallet
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	// Address is the wallet address
	Address common.Address `bson:"address" json:"address"`

	// Tags is a list of tags associated with this wallet
	Tags []Tag `bson:"tags,omitempty" json:"tags"`

	// LastSeen is the time this wallet was last seen
	LastSeen int64 `bson:"last_seen,omitempty" json:"last_seen"`

	// CreatedAt is the time this wallet was created
	CreatedAt int64 `bson:"created_at,omitempty" json:"created_at"`

	// UpdatedAt is the time this wallet was last updated
	UpdatedAt int64 `bson:"updated_at,omitempty" json:"updated_at"`
}

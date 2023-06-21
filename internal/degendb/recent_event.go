package degendb

import (
	"math/big"
	"time"
)

type RecentEvent struct {
	Timestamp time.Time `bson:"timestamp,omitempty" json:"timestamp,omitempty"`

	// Event Type
	Type EventType `bson:"type,omitempty" json:"type,omitempty"`

	// Event Volume
	AmountWei    *big.Int `bson:"amount_wei,omitempty"    json:"amount_wei,omitempty"`
	AmountTokens uint64   `bson:"amount_tokens,omitempty" json:"amount_tokens,omitempty"`
}

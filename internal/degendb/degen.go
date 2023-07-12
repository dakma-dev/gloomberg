package degendb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tag string

type Degen struct {
	// ID is the unique identifier for this degen
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	// Name is the name of the degen
	Name string `bson:"name,omitempty" json:"name"`

	// Accounts contains other accounts of this degen
	Accounts Accounts `bson:"accounts,omitempty" json:"accounts"`

	// Wallets is a list of wallet addresses associated with this degen
	Wallets []primitive.ObjectID `bson:"wallets,omitempty" json:"wallets"`
	// RawWallets is a list of wallet addresses associated with this degen
	RawWallets []Wallet `bson:"raw_wallets,omitempty" json:"raw_wallets"`

	// Tags is a list of tags associated with this degen
	Tags []Tag `bson:"tags,omitempty" json:"tags"`

	// CreatedAt is the time this degen was created
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at"`

	// UpdatedAt is the time this degen was last updated
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at"`
}

type Accounts struct {
	// Twitter is the Twitter username for this degen
	Twitter string `bson:"twitter,omitempty" json:"twitter"`

	// Telegram is the Telegram username for this degen
	Telegram string `bson:"telegram,omitempty" json:"telegram"`

	// ChatID is the Telegram chat ID for this degen
	TelegramChatID int64 `bson:"telegram_chat_id,omitempty" json:"telegram_chat_id"`
}

package collections

import (
	"math/big"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
)

type User struct {
	Address       common.Address
	OpenseaUserID string
}

type EventType int64

const (
	Sale EventType = iota
	Mint
	Transfer
	Listing
	Purchase
)

func (et EventType) String() string {
	return map[EventType]string{
		Sale: "Sale", Mint: "Mint", Transfer: "Transfer", Listing: "Listing", Purchase: "Purchase",
	}[et]
}

func (et EventType) Icon() string {
	switch et {
	case Sale:
		return "💰"
	case Mint:
		return "Ⓜ️"
	case Transfer:
		return "📦"
	case Listing:
		return "📢"
	case Purchase:
		return "🛒"
	}

	return "⁉️"
}

func (et EventType) ActionName() string {
	switch et {
	case Sale:
		return "sold"
	case Mint:
		return "minted"
	case Transfer:
		return "transferred"
	case Listing:
		return "listed"
	case Purchase:
		return "purchased"
	}

	return "⁉️"
}

type Event struct {
	NodeID    int
	EventType EventType
	Topic     string
	TxHash    common.Hash
	// Collection      *Collection
	Collection      *GbCollection
	TokenID         uint64
	PriceWei        *big.Int
	PricePerItem    *big.Int
	CollectionColor lipgloss.Color
	// MultiItemTx bool
	Permalink   string
	TxItemCount uint
	Time        time.Time
	From        User
	FromENS     string
	To          User
	ToENS       string
	WorkerID    int
}

package collections

import (
	"math/big"
	"time"

	"github.com/benleb/gloomberg/internal/external"
	"github.com/benleb/gloomberg/internal/models/txlogcollector"
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

type EventDiscarded struct {
	PrintInStream  bool
	PrintInHistory bool
	Reasons        []string
	DiscardedBy    string
}

type Event struct {
	NodeID               int
	EventType            EventType
	Topic                string
	TxHash               common.Hash
	TransactionCollector *txlogcollector.TxLogs
	ContractAddress      common.Address
	Collection           *GbCollection
	TokenID              *big.Int
	UniqueTokenIDs       []*big.Int
	ENSMetadata          *external.ENSMetadata
	PriceWei             *big.Int
	PriceArrowColor      lipgloss.Color
	CollectionColor      lipgloss.Color
	Permalink            string
	TxLogCount           uint64
	Time                 time.Time
	From                 User
	FromENS              string
	To                   User
	ToColor              lipgloss.Color
	ToENS                string
	FromAddresses        map[common.Address]bool
	ToAddresses          map[common.Address]bool
	WorkerID             int
	Discarded *EventDiscarded
	IsAcceptedOffer      bool
}

type PushEvent struct {
	NodeID          int
	EventType       EventType
	Topic           string
	TxHash          common.Hash
	CollectionName  string
	ContractAddress common.Address
	TokenID         *big.Int
	ENSMetadata     *external.ENSMetadata
	PriceWei        *big.Int
	PricePerItem    *big.Int
	CollectionColor lipgloss.Color
	Permalink       string
	TxItemCount     uint64
	Time            time.Time
	From            User
	FromENS         string
	To              User
	ToENS           string
}

type EventWithStyle struct {
	Verbose bool

	Source      string
	SourceColor lipgloss.Color

	Time      time.Time
	TimeColor lipgloss.Color

	EventType  EventType
	EventEmoji string

	Marker      string
	MarkerColor lipgloss.Color

	TxHash common.Hash

	CollectionName        string
	CollectionColor       lipgloss.Color
	CollectionTotalSupply uint64
	TokenID               *big.Int

	PriceEther      string
	PriceEtherColor lipgloss.Color
	PriceArrowColor lipgloss.Color
	PriceWei        *big.Int
	PricePerItem    *big.Int

	TxItemCount uint64

	EtherscanURL string
	OpenseaURL   string

	From      User
	FromColor lipgloss.Color
	FromENS   string

	To      User
	ToColor lipgloss.Color
	ToENS   string

	SalesCount    uint64
	ListingsCount uint64
}

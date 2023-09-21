package degendb

import (
	"time"

	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
)

type PreformattedEvent struct {
	TxHash                 common.Hash
	Action                 string
	ReceivedAt             time.Time
	Typemoji               string
	Price                  *price.Price
	TransferredCollections []TransferredCollection
	BlurURL                string
	EtherscanURL           string
	OpenSeaURL             string

	// "attributes"
	IsOwnWallet     bool
	IsOwnCollection bool

	PurchaseOrBidIndicator string

	// fix this with new chawago watcher
	IsWatchUsersWallet bool

	// temporary until we have a better solution
	From        *Degen
	FromAddress common.Address
	To          *Degen
	ToAddress   common.Address

	Colors EventColors
	Other  map[string]interface{}
}

type EventColors struct {
	Time          lipgloss.Color
	Price         lipgloss.Color
	PriceCurrency lipgloss.Color
	PriceArrow    lipgloss.Color
	From          lipgloss.Color
	To            lipgloss.Color

	Collections map[common.Address]CollectionColors
}

type CollectionColors struct {
	Primary   lipgloss.Color
	Secondary lipgloss.Color
}

type TransferredCollection struct {
	CollectionName    string
	TransferredTokens []TransferredToken

	Colors CollectionColors

	// from & to per collection as we print one line per collection...^^
	From string
	To   string
}

type TransferredToken struct {
	ID         int64
	Rank       int64
	RankSymbol string
	Amount     int64
}

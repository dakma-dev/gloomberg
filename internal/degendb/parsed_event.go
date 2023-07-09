package degendb

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
)

type ParsedEvent struct {
	TxHash                 common.Hash
	Action                 string
	ReceivedAt             string
	TimeColor              lipgloss.Color
	Typemoji               string
	Price                  string
	PriceColor             lipgloss.Color
	PriceCurrencyColor     lipgloss.Color
	PriceArrowColor        lipgloss.Color
	TransferredCollections []TransferredCollection
	BlurURL                string
	EtherscanURL           string
	OpenSeaURL             string

	// temporary until we have a better solution
	From      string
	FromColor lipgloss.Color
	To        string
	ToColor   lipgloss.Color
}

type TransferredCollection struct {
	CollectionName    string
	TransferredTokens []TransferredToken

	PrimaryColor   lipgloss.Color
	SecondaryColor lipgloss.Color

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

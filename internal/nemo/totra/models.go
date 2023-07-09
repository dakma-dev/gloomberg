package totra

import (
	"math/big"
	"time"

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/nemo/standard"
	"github.com/benleb/gloomberg/internal/nemo/token"
	"github.com/ethereum/go-ethereum/common"
)

type TokenTransfer struct {
	// sender of a token || erc721: topic[1] | erc1155: topic[2]
	From common.Address `json:"from"`

	// recipient of a token || erc721: topic[2] | erc1155: topic[3]
	To common.Address `json:"to"`

	// type of the token || erc721: topic[0] == Transfer | erc1155: topic[0] == TransferSingle
	Standard standard.Standard `json:"type"`

	// number of tokens transferred || erc721: 1 | erc1155: data[value]
	AmountTokens *big.Int `json:"amount"`

	// the amount of eth/weth transferred in the same tx to the sender of the nft
	AmountEtherReturned *big.Int `json:"amount_ether_returned"`

	Token *token.Token `json:"token"`
}

// HistoryTokenTransaction is the representation of a token transaction in the history.
type HistoryTokenTransaction struct {
	ReceivedAt           time.Time               `json:"time"`
	AmountPaid           *big.Int                `json:"amount_paid"`
	FmtTokensTransferred []string                `json:"fmt_tokens_transferred"`
	Collection           *collections.Collection `json:"collection"`

	TokenTransaction *TokenTransaction `json:"-"`
}

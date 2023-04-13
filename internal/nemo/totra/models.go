package totra

import (
	"math/big"
	"time"

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/nemo/standard"
	"github.com/benleb/gloomberg/internal/nemo/token"
	"github.com/ethereum/go-ethereum/common"
)

// type Token struct {
// 	// id of the token || erc721: topic[3] | erc1155: tx.To()
// 	ID *big.Int `json:"id"`

// 	// address of the token contract || erc721: tx.To() | erc1155: topic[1]
// 	Address common.Address `json:"address"`
// }

// func (t *Token) String() string {
// 	return t.NftID()
// }

// func (t *Token) NftID() string {
// 	network := "ethereum"

// 	return fmt.Sprintf("%s/%s/%s", network, t.Address.String(), t.ID.String())
// }

// func (t *Token) ShortID() string {
// 	return fmt.Sprintf("%s/%s", t.Address.String(), t.ID.String())
// }

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

// // implement Transfer interface
// func (ttf *TokenTransfer) From() *common.Address {
// 	return &ttf.OldFrom
// }

// func (ttf *TokenTransfer) To() *common.Address {
// 	return &ttf.OldTo
// }

// func (ttf *TokenTransfer) TokenAddress() *common.Address {
// 	return &ttf.Token.Address
// }

// func (ttf *TokenTransfer) TokenID() *big.Int {
// 	return ttf.Token.ID
// }

// func (ttf *TokenTransfer) AmountTokens() *big.Int {
// 	return ttf.OldAmountTokens
// }

// HistoryTokenTransaction is the representation of a token transaction in the history.
type HistoryTokenTransaction struct {
	ReceivedAt           time.Time               `json:"time"`
	AmountPaid           *big.Int                `json:"amount_paid"`
	FmtTokensTransferred []string                `json:"fmt_tokens_transferred"`
	Collection           *collections.Collection `json:"collection"`

	TokenTransaction *TokenTransaction `json:"-"`
}

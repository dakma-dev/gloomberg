package token

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Token struct {
	// id of the token || erc721: topic[3] | erc1155: tx.To()
	ID *big.Int `json:"id"`

	// address of the token contract || erc721: tx.To() | erc1155: topic[1]
	Address common.Address `json:"address"`
}

func (t *Token) String() string {
	return t.NftID()
}

func (t *Token) NftID() string {
	network := "ethereum"

	return fmt.Sprintf("%s/%s/%s", network, t.Address.String(), t.ID.String())
}

func (t *Token) ShortID() string {
	return fmt.Sprintf("%s/%s", t.Address.String(), t.ID.String())
}

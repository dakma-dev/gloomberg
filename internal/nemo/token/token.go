package token

import (
	"fmt"
	"math/big"

	"github.com/benleb/gloomberg/internal/style"
	"github.com/ethereum/go-ethereum/common"
)

type Token struct {
	// id of the token || erc721: topic[3] | erc1155: tx.To()
	ID *big.Int `json:"id"`

	// address of the token contract || erc721: tx.To() | erc1155: topic[1]
	Address common.Address `json:"address"`

	// optional name of the token (opensea api source)
	Name string `json:"name,omitempty"`
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

func (t *Token) LipglossedShortID() string {
	return style.AlmostWhiteStyle.Render(t.Address.String()) + style.GrayStyle.Render("/") + style.BoldAlmostWhite(t.ID.String())
}

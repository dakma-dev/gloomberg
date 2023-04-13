package tokentransfer

import (
	"math/big"

	"github.com/benleb/gloomberg/internal/nemo/standard"
	"github.com/ethereum/go-ethereum/common"
)

type TokTransfer interface {
	From() *common.Address
	To() *common.Address

	TokenAddress() *common.Address
	TokenID() *big.Int
	AmountTokens() *big.Int

	Standard() standard.Standard
}

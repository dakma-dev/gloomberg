package wwatcher

import (
	"github.com/benleb/gloomberg/internal/nemo"
	"github.com/ethereum/go-ethereum/common"
)

type MIWCollection struct {
	MIWs         nemo.AddressCollection
	WeightedMIWs map[common.Address]int
}

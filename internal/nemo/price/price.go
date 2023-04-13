package price

import (
	"math/big"

	"github.com/benleb/gloomberg/internal/utils"
)

// Price represents the value/amount of (w)eth transferred in a transaction.
type Price struct {
	value *big.Int
}

func NewPrice(price *big.Int) *Price {
	if price == nil {
		return nil
	}

	return &Price{
		value: price,
	}
}

func (p *Price) Add(itemPrice *Price) *Price {
	return &Price{
		value: big.NewInt(0).Add(p.value, itemPrice.value),
	}
}

func (p *Price) Wei() *big.Int {
	if p == nil {
		return nil
	}

	return p.value
}

func (p *Price) Gwei() float64 {
	gwei, _ := utils.WeiToGwei(p.value).Float64()

	return gwei
}

func (p *Price) Ether() float64 {
	ether, _ := utils.WeiToEther(p.value).Float64()

	return ether
}

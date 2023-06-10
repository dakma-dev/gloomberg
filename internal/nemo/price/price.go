package price

import (
	"math/big"

	"github.com/benleb/gloomberg/internal/utils"
)

// Price represents the value/amount of (w)eth transferred in a transaction.
type Price struct {
	valueWei *big.Int
}

func NewPrice(price *big.Int) *Price {
	if price == nil {
		return nil
	}

	return &Price{
		valueWei: price,
	}
}

func (p *Price) Add(itemPrice *Price) *Price {
	return &Price{
		valueWei: big.NewInt(0).Add(p.valueWei, itemPrice.valueWei),
	}
}

func (p *Price) Wei() *big.Int {
	if p.valueWei == nil {
		return big.NewInt(0)
	}

	return p.valueWei
}

func (p *Price) Gwei() float64 {
	gwei, _ := utils.WeiToGwei(p.Wei()).Float64()

	return gwei
}

func (p *Price) Ether() float64 {
	ether, _ := utils.WeiToEther(p.Wei()).Float64()

	return ether
}

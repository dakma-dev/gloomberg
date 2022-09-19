package wallet

import (
	"math/big"

	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wealdtech/go-ens/v3"
)

// Wallet represents the wallets configured by the user.
type Wallet struct {
	Name          string         `mapstructure:"name"`
	Address       common.Address `mapstructure:"address"`
	ENS           *ens.Name      `mapstructure:"ens"`
	Balance       *big.Int       `mapstructure:"balance"`
	BalanceBefore *big.Int       `mapstructure:"balance_before"`
	BalanceTrend  string         `mapstructure:"balance_trend"`
	Color         lipgloss.Color `mapstructure:"color"`
}

func (w *Wallet) ColoredName(maxWalletNameLength int) string {
	return lipgloss.NewStyle().Foreground(w.Color).Faint(true).Width(maxWalletNameLength).Render(w.Name)
}

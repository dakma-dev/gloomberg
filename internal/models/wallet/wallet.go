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
	ENSName       string         `mapstructure:"ens_name"`
	Color         lipgloss.Color `mapstructure:"color"`
	Balance       *big.Int
	BalanceBefore *big.Int
	BalanceTrend  string
}

func (w *Wallet) ColoredName(maxWalletNameLength int) string {
	return lipgloss.NewStyle().Foreground(w.Color).Faint(true).Width(maxWalletNameLength).Render(w.Name)
}

func (w *Wallet) Render(text string) string {
	// generate the collection color based on the contract address if none given
	return lipgloss.NewStyle().Foreground(w.Color).Render(text)
}

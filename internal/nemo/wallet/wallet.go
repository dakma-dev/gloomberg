package wallet

import (
	"github.com/benleb/gloomberg/internal/nemo/token"
	"math/big"

	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wealdtech/go-ens/v3"
)

// Wallet represents a EOA account/wallet.
type Wallet struct {
	Name          string         `mapstructure:"name"`
	Address       common.Address `mapstructure:"address"`
	ENS           *ens.Name      `mapstructure:"ens"`
	ENSName       string         `mapstructure:"ens_name"`
	Color         lipgloss.Color `mapstructure:"color"`
	Balance       *big.Int
	BalanceBefore *big.Int
	BalanceTrend  string
	Tokens        map[common.Address]map[string]*token.Token
}

func (w *Wallet) ColoredName(maxWalletNameLength int) string {
	return lipgloss.NewStyle().Foreground(w.Color).Faint(true).Width(maxWalletNameLength).MaxWidth(maxWalletNameLength).Render(w.Name)
}

func (w *Wallet) Render(text string) string {
	return lipgloss.NewStyle().Foreground(w.Color).Render(text)
}

package models

import (
	"math/big"

	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wealdtech/go-ens/v3"
)

type (
	Wallets      map[common.Address]*Wallet
	WalletsSlice []*Wallet
)

func (w WalletsSlice) Len() int           { return len(w) }
func (w WalletsSlice) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }
func (w WalletsSlice) Less(i, j int) bool { return w[i].Balance.Uint64() < w[j].Balance.Uint64() }

func (ws *Wallets) Addresses() []common.Address {
	addresses := make([]common.Address, 0)
	for _, w := range *ws {
		addresses = append(addresses, w.Address)
	}

	return addresses
}

func (ws *Wallets) StringAddresses() []string {
	addresses := make([]string, 0)
	for _, w := range *ws {
		addresses = append(addresses, w.Address.String())
	}

	return addresses
}

func (ws *Wallets) FormattedNames() []string {
	names := make([]string, 0)
	for _, w := range *ws {
		names = append(names, lipgloss.NewStyle().Foreground(&w.Color).Render(w.Name))
	}

	return names
}

// Contains returns true if the given string is in the slice.
func (ws *Wallets) Contains(address common.Address) bool {
	for _, walletAddress := range ws.Addresses() {
		if address == walletAddress {
			return true
		}
	}

	return false
}

func (ws *Wallets) GetAll() WalletsSlice {
	slice := make([]*Wallet, 0)
	for _, w := range *ws {
		slice = append(slice, w)
	}

	return slice
}

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

package wallet

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
)

type (
	Wallets     map[common.Address]*Wallet
	WalletGroup []*Wallet
)

// type WalletGroup struct {
// 	Name             string    `mapstructure:"name"`
// 	TelegramUsername string    `mapstructure:"telegram_username"`
// 	Wallets          []*Wallet `mapstructure:"wallets"`
// 	Own              bool      `mapstructure:"own"`
// }

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

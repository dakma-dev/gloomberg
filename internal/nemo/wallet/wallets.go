package wallet

import (
	"sort"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/nemo/token"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
)

type Wallets map[common.Address]*Wallet

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

func (ws *Wallets) ContainsAddressFromSlice(addresses []common.Address) common.Address {
	for _, address := range addresses {
		if (*ws)[address] != nil {
			return address
		}
	}

	return internal.ZeroAddress
}

func (ws *Wallets) SortByBalance() []*Wallet {
	slice := make([]*Wallet, 0)
	for _, w := range *ws {
		slice = append(slice, w)
	}

	sort.Slice(slice, func(i, j int) bool {
		return slice[i].Balance.Uint64() > slice[j].Balance.Uint64()
	})

	return slice
}

func (ws *Wallets) ContainsToken(tokenAddress common.Address, tokenID string) bool {
	for _, w := range *ws {
		if w.Tokens[tokenAddress] != nil {
			if w.Tokens[tokenAddress][tokenID] != nil {
				return true
			}
		}
	}

	return false
}

func (ws *Wallets) GetCollectionTokens(collectionAddress common.Address) []*token.Token {
	collectionTokens := make([]*token.Token, 0)

	for _, w := range *ws {
		if tokens := w.Tokens[collectionAddress]; len(tokens) > 0 {
			for _, token := range tokens {
				collectionTokens = append(collectionTokens, token)
			}
		}
	}

	return collectionTokens
}

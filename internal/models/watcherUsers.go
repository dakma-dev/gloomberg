package models

import "github.com/ethereum/go-ethereum/common"

type WatcherUsers map[common.Address]*WatcherUser

// Contains returns true if the given string is in the slice.
func (wu *WatcherUsers) Contains(address common.Address) bool {
	return (*wu)[address] != nil
}

func (wu *WatcherUsers) ContainsOneOf(addresses map[common.Address]bool) common.Address {
	for address := range addresses {
		if (*wu).Contains(address) {
			return address
		}
	}

	return common.Address{}
}

// WatcherUser representsa user who can own multiple wallets.
type WatcherUser struct {
	Name            string           `mapstructure:"name"`
	WalletAddresses []common.Address `mapstructure:"wallets"`
	TgUsername      string           `mapstructure:"telegram"`
}

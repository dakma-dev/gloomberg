package models

import "github.com/ethereum/go-ethereum/common"

type WatcherUsers map[common.Address]*WatcherUser

// Contains returns true if the given string is in the slice.
func (wu *WatcherUsers) Contains(address common.Address) bool {
	return (*wu)[address] != nil
}

// WatcherUser representsa user who can own multiple wallets.
type WatcherUser struct {
	Name            string           `mapstructure:"name"`
	WalletAddresses []common.Address `mapstructure:"wallets"`
	TgUsername      string           `mapstructure:"telegram"`
}

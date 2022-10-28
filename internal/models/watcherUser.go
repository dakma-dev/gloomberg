package models

import (
	"github.com/ethereum/go-ethereum/common"
)

type NotificationRecipients map[common.Address]*User

// var Recipients NotificationRecipients = make(map[common.Address]*User)

// WatchUser representsa user who can own multiple wallets.
type WatchUser struct {
	Name            string           `mapstructure:"name"`
	WalletAddresses []common.Address `mapstructure:"wallets"`
	TgUsername      string           `mapstructure:"telegram"`
}

// Contains returns true if the given string is in the slice.
func (nr *NotificationRecipients) Contains(address common.Address) bool {
	for wwatcherAddress := range *nr {
		if address == wwatcherAddress {
			return true
		}
	}

	return false
}

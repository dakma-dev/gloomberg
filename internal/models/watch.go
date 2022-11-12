package models

import "github.com/ethereum/go-ethereum/common"

type Watcher struct {
	Groups          map[string]*WatchGroup
	UserAddresses   map[common.Address]*WatchGroup
	WalletAddresses map[common.Address]*WatchGroup
	WatchUsers      WatcherUsers
}

type WatchWallet struct {
	Name    string         `mapstructure:"name"`
	Address common.Address `mapstructure:"address"`
}

type WatchUser struct {
	Name             string         `mapstructure:"name"`
	TelegramUsername string         `mapstructure:"telegram_username"`
	Wallets          []*WatchWallet `mapstructure:"wallets"`

	WalletAddresses []common.Address

	Group *WatchGroup
}

type WatchGroup struct {
	Name           string         `mapstructure:"group"`
	TelegramChatID int64          `mapstructure:"telegram_chat_id"`
	Users          []*WatchUser   `mapstructure:"users"`
	Wallets        []*WatchWallet `mapstructure:"wallets"`

	// addresses []common.Address
	// Contracts      []WatchContract `mapstructure:"contracts"`
}

// Contains returns true if the given string is in the slice.
func (wu *Watcher) Contains(address common.Address) bool {
	return ((*wu).UserAddresses)[address] != nil
}

func (wu *Watcher) ContainsOneOf(addresses map[common.Address]bool) common.Address {
	for address := range addresses {
		if (*wu).Contains(address) {
			return address
		}
	}

	return common.Address{}
}

package watch

import "github.com/ethereum/go-ethereum/common"

type WatcherUsers map[common.Address]*WUser

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

type Watcher struct {
	Groups          map[string]*WGroup
	UserAddresses   map[common.Address]*WGroup
	WalletAddresses map[common.Address]*WGroup
	WatchUsers      WatcherUsers
}

type WWallet struct {
	Name    string         `mapstructure:"name"`
	Address common.Address `mapstructure:"address"`
}

type WUser struct {
	Name             string     `mapstructure:"name"`
	TelegramUsername string     `mapstructure:"telegram_username"`
	Wallets          []*WWallet `mapstructure:"wallets"`

	WalletAddresses []common.Address

	Group *WGroup
}

type WGroup struct {
	Name           string     `mapstructure:"group"`
	TelegramChatID int64      `mapstructure:"telegram_chat_id"`
	Users          []*WUser   `mapstructure:"users"`
	Wallets        []*WWallet `mapstructure:"wallets"`

	// addresses []common.Address
	// Contracts      []WatchContract `mapstructure:"contracts"`
}

// Contains returns true if the given string is in the slice.
func (wu *Watcher) Contains(address common.Address) bool {
	return wu.UserAddresses[address] != nil
}

func (wu *Watcher) ContainsOneOf(addresses map[common.Address]bool) common.Address {
	for address := range addresses {
		if (*wu).Contains(address) {
			return address
		}
	}

	return common.Address{}
}

func (wu *Watcher) ContainsAddressFromSlice(addresses []common.Address) common.Address {
	for _, address := range addresses {
		if (*wu).Contains(address) {
			return address
		}
	}

	return common.Address{}
}

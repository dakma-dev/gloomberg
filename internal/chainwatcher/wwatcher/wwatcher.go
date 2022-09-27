package wwatcher

import (
	"fmt"

	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/benleb/gloomberg/internal/utils/hooks"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type NotificationRecipients map[common.Address]*User

var Recipients NotificationRecipients = make(map[common.Address]*User)

// User representsa user who can own multiple wallets.
type User struct {
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

func InitWatcher() {
	if userConfigs, ok := viper.Get("wwatcher").([]interface{}); ok {
		for _, userConfig := range userConfigs {
			var newUser *User

			gbl.Log.Debugf("%+v\n", userConfig)

			decodeHooks := mapstructure.ComposeDecodeHookFunc(
				hooks.StringToAddressHookFunc(),
			)

			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				DecodeHook: decodeHooks,
				Result:     &newUser,
			})

			err := decoder.Decode(userConfig)
			if err != nil {
				fmt.Printf("\nerr decode: %+v\n", err)

				return
			}

			gbl.Log.Debugf("\n%+v\n", newUser)

			for _, walletAddress := range newUser.WalletAddresses {
				Recipients[walletAddress] = newUser
			}
		}
	}
}

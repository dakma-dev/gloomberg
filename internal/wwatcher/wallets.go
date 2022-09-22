package wwatcher

import (
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/models/wallet"
	"github.com/benleb/gloomberg/internal/server/node"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wealdtech/go-ens/v3"
)

// var bLog = btv.GetBLog()

// func GetWallets(nodes *newnodes.NodeCollection) *models.Wallets {
// 	var wallets models.Wallets

// 	// parse and validate wallets from config
// 	if viper.IsSet("wallets") {
// 		wallets = GetWalletsFromConfig(nodes, viper.GetStringSlice("wallets"))
// 	}

// 	return &wallets
// }

// GetWalletsFromConfig returns a slice of wallets from the configuration file.
func GetWalletsFromConfig(walletsConfig []string, nodes *node.Nodes) *wallet.Wallets {
	var walletsWg sync.WaitGroup

	var wallets wallet.Wallets = make(map[common.Address]*wallet.Wallet)

	for _, walletConfig := range walletsConfig {
		walletsWg.Add(1)

		gbl.Log.Debugf("walletConfig: %s", walletConfig)

		// bind walletConfig to this loop iteration
		walletConfig := walletConfig

		go func() {
			defer walletsWg.Done()

			var (
				userWallet wallet.Wallet

				walletName    string
				walletAddress common.Address
				walletENS     *ens.Name
			)

			client := nodes.GetRandomNode().Client

			// switch walletConfig := walletConfig.(type) {
			// case string:
			configEntry := walletConfig

			if common.IsHexAddress(configEntry) {
				gbl.Log.Debugf("common.IsHexAddress(configEntry): %+v", common.IsHexAddress(configEntry))

				// address string given, trying to resolve name (reverse & forward lookup)
				walletAddress = common.HexToAddress(configEntry)

				name, err := ens.ReverseResolve(client, walletAddress)
				if err != nil {
					fmt.Println(err)
				}

				resolvedAddress, err := ens.Resolve(client, name)
				if err != nil {
					fmt.Println(err)
				}

				if resolvedAddress == walletAddress {
					// valid ENS records, using ens as wallet name
					walletName = name
					walletENS, _ = ens.NewName(client, name)
				} else {
					// invalid ENS records, using address as wallet name
					gbl.Log.Warnf("Fake ENS? Address %s is %s is %s", walletAddress.Hex(), name, resolvedAddress.Hex())
					walletName = WalletShortAddress(walletAddress)
				}
			} else if name := configEntry; strings.HasSuffix(name, ".eth") {
				gbl.Log.Debugf("configEntry is ens: %+v", name)

				// ens domain given, trying to resolve address
				address, err := ens.Resolve(client, name)
				if err != nil {
					// retry with other node
					address, err = ens.Resolve(nodes.GetRandomNode().Client, name)
					if err != nil {
						gbl.Log.Warnf("could not resolve ENS Address %s: %s", name, err.Error())

						return
					}

					// gbl.Log.Warnf("could not resolve ENS Address %s: %s", walletENS, err.Error())

					// return
				}

				walletAddress = address

				walletENS, err = ens.NewName(client, name)

				if err != nil {
					if ens.Domain(name) != "eth" {
						// ens subdomain case

						// hack to create a subdomain
						walletENS, err = ens.NewName(client, ens.Domain(name))

						if err != nil {
							gbl.Log.Errorf("ens.NewName error %s: %s", name, err.Error())
							walletName = WalletShortAddress(walletAddress)
						} else {
							// hack to create a subdomain object
							walletENS.Domain = ens.Domain(name)
							walletENS.Name = name

							resolver, err := walletENS.ResolverAddress()
							if err != nil {
								gbl.Log.Errorf("ens.ResolverAddress error %s: %s", name, err.Error())

								return
							}

							gbl.Log.Debugf("walletENS %s | walletENS.Domain %s | walletENS.Name %s | walletENS.ResolverAddress() %s", walletENS.Name, walletENS.Domain, walletENS.Name, resolver.Hex())

							if _, err = ens.Resolve(client, walletENS.Name); err != nil {
								gbl.Log.Errorf("could not resolve walletENS.Name ENS Address %s: %s", walletENS.Name, err.Error())
								// walletName = WalletShortAddress(walletAddress)
							}

							walletName = walletENS.Name
						}
					} else {
						walletName = WalletShortAddress(walletAddress)
					}
				} else {
					walletName = walletENS.Name
				}
			}

			userWallet = wallet.Wallet{
				Name:    walletName,
				Address: walletAddress,
				ENS:     walletENS,
			}

			gbl.Log.Debugf("userWallet: %+v", userWallet)

			// default:
			// 	decodeHooks := mapstructure.ComposeDecodeHookFunc(hooks.StringToAddressHookFunc())

			// 	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			// 		DecodeHook: decodeHooks,
			// 		Result:     &userWallet,
			// 	})

			// 	err := decoder.Decode(walletConfig)
			// 	if err != nil {
			// 		log.Fatal(err)
			// 	}

			// 	// if viper.GetString("etherscan.apikey") != "" {
			// 	// 	if name := ReverseLookupAndValidate(userWallet.Address); name != "" {
			// 	// 		walletENS, err = ens.NewName(client, name)
			// 	// 		if err != nil {
			// 	// 			walletName = WalletShortAddress(walletAddress)
			// 	// 		} else {
			// 	// 			walletName = walletENS.Name
			// 	// 		}
			// 	// 	}
			// 	// }
			// 	gbl.Log.Debugf("userWallet: %+v", userWallet)
			// }

			userWallet.Color = style.GenerateColorWithSeed(userWallet.Address.Hash().Big().Int64())
			userWallet.Balance = big.NewInt(0)
			userWallet.BalanceBefore = big.NewInt(0)

			wallets[userWallet.Address] = &userWallet
		}()
	}

	walletsWg.Wait()

	return &wallets
}

func WalletShortAddress(address common.Address) string {
	addressBytes := address.Bytes()

	return fmt.Sprint(
		"0x",
		fmt.Sprintf("%0.2x%0.2x", addressBytes[0], addressBytes[1]),
		"…",
		fmt.Sprintf("%0.2x%0.2x", addressBytes[len(addressBytes)-2], addressBytes[len(addressBytes)-1]),
	)
}

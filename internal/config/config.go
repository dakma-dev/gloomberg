package config

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo"
	"github.com/benleb/gloomberg/internal/nemo/collectionsource"
	"github.com/benleb/gloomberg/internal/nemo/provider"
	"github.com/benleb/gloomberg/internal/nemo/wallet"
	"github.com/benleb/gloomberg/internal/nemo/watch"
	"github.com/benleb/gloomberg/internal/rueidica"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/benleb/gloomberg/internal/utils/hooks"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/wealdtech/go-ens/v3"
)

func GetOwnWalletsFromConfig(providerPool *provider.Pool) *wallet.Wallets {
	ownWallets := make(map[common.Address]*wallet.Wallet, 0)
	mu := sync.Mutex{}

	nodesSpinner := style.GetSpinner("setting up own wallets...")
	_ = nodesSpinner.Start()

	var wgWallets sync.WaitGroup

	rawWalletConfig, ok := viper.Get("wallets").([]interface{})
	if !ok {
		gbl.Log.Warnf("reading wallets configuration failed: %+v", ok)

		return nil
	}

	for _, walletConfig := range rawWalletConfig {
		wgWallets.Add(1)

		go func(walletConfig interface{}) {
			defer wgWallets.Done()

			var newWallet *wallet.Wallet

			decodeHooks := mapstructure.ComposeDecodeHookFunc(
				hooks.StringToAddressHookFunc(),
				hooks.StringToLipglossColorHookFunc(),
			)

			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				DecodeHook: decodeHooks,
				Result:     &newWallet,
			})

			err := decoder.Decode(walletConfig)
			if err != nil {
				gbl.Log.Warnf("reading wallet group configuration failed: %+v", err)

				return
			}

			if newWallet.Color == "" {
				newWallet.Color = style.GenerateColorWithSeed(newWallet.Address.Hash().Big().Int64())
			}

			if newWallet.Name == "" {
				newWallet.Name = utils.WalletShortAddress(newWallet.Address)
			}

			newWallet.ENS = &ens.Name{}

			if providerPool != nil {
				if name, err := providerPool.ReverseResolveAddressToENS(context.TODO(), newWallet.Address); err != nil {
					gbl.Log.Debugf("  ❌ name not resolved: %s -> %+v", newWallet.Address, err)
				} else {
					newWallet.ENS.Name = name
					newWallet.Name = name

					gbl.Log.Debugf("  ✅ name resolved: %s -> %+v", newWallet.Address, name)
				}
			}

			newWallet.Balance, newWallet.BalanceBefore = big.NewInt(0), big.NewInt(0)

			gbl.Log.Infof("✅ successfully added own wallet: %s (%s)", newWallet.Render(newWallet.Name), style.ShortenAddressStyled(&newWallet.Address, lipgloss.NewStyle().Foreground(newWallet.Color)))

			mu.Lock()
			ownWallets[newWallet.Address] = newWallet
			mu.Unlock()
		}(walletConfig)
	}

	// wait for all goroutines to finish
	wgWallets.Wait()

	// resolve addresses to ens names if nodes are available
	gbl.Log.Debugf("ethNodes != nil: %v | %+v | %+v", providerPool != nil, providerPool, ownWallets)

	// if providerPool != nil {
	// 	// providerPool.GetENSForAllAddresses((*wallet.Wallets)(&ownWallets))

	// 	for address, w := range ownWallets {
	// 		gbl.Log.Debugf("wallet: %+v", w)

	// 		if name, err := providerPool.ResolveENSForAddress(w.Address); err != nil {
	// 			gbl.Log.Debugf("  ❌ name not resolved: %s -> %+v", address, err)
	// 		} else {
	// 			ownWallets[address].ENS.Name = name
	// 			ownWallets[address].Name = name

	// 			gbl.Log.Debugf("  ✅ name resolved: %s -> %+v", address, name)
	// 		}
	// 	}
	// }

	// build spinner stop msg with all wallet names
	nodesSpinner.StopMessage(fmt.Sprint(
		style.BoldStyle.Render(fmt.Sprint(len(ownWallets))),
		" wallets: ",
		strings.Join((*wallet.Wallets)(&ownWallets).FormattedNames(), ", "),
	) + "\n")

	_ = nodesSpinner.Stop()

	// gbl.Log.Info(pretty.Sprint(ownWallets))

	return (*wallet.Wallets)(&ownWallets)
}

func GetCollectionsFromConfiguration(providerPool *provider.Pool, rueidica *rueidica.Rueidica) []*collections.Collection {
	ownCollections := make([]*collections.Collection, 0)

	for address, collection := range viper.GetStringMap("collections") {
		contractAddress := common.HexToAddress(address)
		currentCollection := collections.NewCollection(contractAddress, "", providerPool, collectionsource.FromConfiguration, rueidica)

		if collection == nil && common.IsHexAddress(address) {
			gbl.Log.Infof("reading collection without details: %+v", address)

			currentCollection = collections.NewCollection(contractAddress, "", providerPool, collectionsource.FromConfiguration, rueidica)

			// general settings
			if viper.Sub("collections."+currentCollection.ContractAddress.String()+".buy") != nil && !viper.IsSet("show.listings") {
				currentCollection.Show.Listings = false
				currentCollection.FetchListings = true
			} else {
				currentCollection.Show.Listings = viper.GetBool("show.listings")
				currentCollection.FetchListings = viper.GetBool("show.listings") // viper.GetBool("fetch.listings")
			}

			currentCollection.Show.Sales = viper.GetBool("show.sales")
			currentCollection.Show.Mints = true // viper.GetBool("show.mints")
			currentCollection.Show.Transfers = viper.GetBool("show.transfers")
		} else {
			gbl.Log.Debugf("reading collection: %+v - %+v", address, collection)

			decodeHooks := mapstructure.ComposeDecodeHookFunc(
				hooks.StringToAddressHookFunc(),
				hooks.StringToDurationHookFunc(),
				hooks.StringToLipglossColorHookFunc(),
			)

			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				DecodeHook: decodeHooks,
				Result:     &currentCollection,
			})

			err := decoder.Decode(collection)
			if err != nil {
				gbl.Log.Errorf("error decoding collection: %+v", err)

				continue
			}
		}

		// // validating buy rule
		// if rule := viper.Sub("collections." + currentCollection.ContractAddress.String() + ".buy"); rule != nil {
		// 	if buyRule := ValidateRawBuyRule(rule, currentCollection.OpenseaSlug); buyRule != nil {
		// 		currentCollection.BuyRule = buyRule
		// 	} else {
		// 		gbl.Log.Infof("〰️ buyRule for %s is invalid", currentCollection.ContractAddress.Hex())
		// 	}
		// }

		gbl.Log.Debugf("currentCollection: %+v", currentCollection)

		ownCollections = append(ownCollections, currentCollection)
	}

	return ownCollections
}

func ValidateRawBuyRule(rule *viper.Viper, slug string) *nemo.BuyRule {
	// initialize buy conditions with "safe" values
	var (
		// checking general key & threshold
		generalPrivateKey = viper.GetString("buy.private_key")
		maxPrice          = 0.0
		minSales          = uint64(5)
		minListings       = uint64(5)
	)

	// set name to ruleKey or a generic one for the catch-all
	name := slug
	// if contractName, err := cache.GetContractName(context.TODO(), common.HexToAddress(rule.GetString("contract_address"))); err == nil && contractName != "" {
	// 	name = contractName
	// }

	rule.Set("name", name)

	// check price threshold value
	if rule.IsSet("threshold") && (rule.GetFloat64("threshold") < 0.0 || rule.GetFloat64("threshold") > 1.0) {
		gbl.Log.Warnf("🤷‍♀️ %s| invalid rule.Threshold (%.3f) value, skipping auto-buy", name, rule.GetFloat64("threshold"))

		return nil
	}

	if viper.IsSet("buy.rules.min_sales") {
		minSales = viper.GetUint64("buy.rules.min_sales")
	}

	if viper.IsSet("buy.rules.min_listings") {
		minListings = viper.GetUint64("buy.rules.min_listings")
	}

	if !rule.IsSet("private_key") {
		rule.Set("private_key", generalPrivateKey)
	}

	if !rule.IsSet("min_sales") {
		rule.Set("min_sales", minSales)
	}

	if !rule.IsSet("min_listings") {
		rule.Set("min_listings", minListings)
	}

	if !rule.IsSet("max_price") {
		rule.Set("max_price", maxPrice)
	}

	var buyRule *nemo.BuyRule
	if err := rule.Unmarshal(&buyRule, viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(
		hooks.StringToAddressHookFunc(),
		hooks.StringToDurationHookFunc(),
		hooks.StringToLipglossColorHookFunc(),
	))); err != nil {
		gbl.Log.Errorf("❌ error unmarshalling buy rule %s: %s", name, err)

		return nil
	}

	return buyRule
}

func GetGeneralBuyRuleFromConfiguration() *nemo.BuyRule {
	// initialize buy conditions with "safe" values
	var (
		// checking general key & threshold
		privateKey  = viper.GetString("buy.private_key")
		threshold   = viper.GetFloat64("buy.threshold")
		maxPrice    = 0.0
		minSales    = uint64(100)
		minListings = uint64(100)
	)

	if viper.IsSet("buy.min_sales") {
		minSales = viper.GetUint64("buy.min_sales")
	}

	if viper.IsSet("buy.min_listings") {
		minListings = viper.GetUint64("buy.min_listings")
	}

	if privateKey == "" || threshold == 0.0 || (threshold < 0.0 || threshold > 1.0) {
		gbl.Log.Warnf("❌ invalid private key (%s) or threshold (%f), skipping general rule", privateKey, threshold)

		return nil
	}

	// create the general/general buy rule
	buyRule := &nemo.BuyRule{
		ID:          0,
		Name:        "*everything*",
		PrivateKey:  privateKey,
		Threshold:   threshold,
		MinSales:    minSales,
		MinListings: minListings,
		MaxPrice:    maxPrice,
	}

	return buyRule
}

// GetWatchRulesFromConfig reads configured users to be notified from config.
func GetWatchRulesFromConfig() *watch.Watcher {
	mu := sync.Mutex{}

	watcher := watch.Watcher{
		UserAddresses:   make(map[common.Address]*watch.WGroup, 0),
		WalletAddresses: make(map[common.Address]*watch.WGroup, 0),
		Groups:          make(map[string]*watch.WGroup, 0),
		WatchUsers:      make(map[common.Address]*watch.WUser, 0),
	}

	watchSpinner := style.GetSpinner("setting up watch rules...")
	_ = watchSpinner.Start()

	rawWatchConfig, ok := viper.Get("watch").([]interface{})
	if !ok {
		gbl.Log.Warnf("watch configuration is not an array, skipping")

		return nil
	}

	for _, group := range rawWatchConfig {
		var newWatchGroup *watch.WGroup

		decodeHooks := mapstructure.ComposeDecodeHookFunc(
			hooks.StringToAddressHookFunc(),
			hooks.StringToLipglossColorHookFunc(),
		)

		decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			DecodeHook: decodeHooks,
			Result:     &newWatchGroup,
		})

		err := decoder.Decode(group)
		if err != nil {
			gbl.Log.Warnf("reading watchGroup configuration failed: %+v", err)

			continue
		}

		mu.Lock()
		watcher.Groups[newWatchGroup.Name] = newWatchGroup

		for _, user := range newWatchGroup.Users {
			for _, userWallet := range user.Wallets {
				watcher.UserAddresses[userWallet.Address] = watcher.Groups[newWatchGroup.Name]
				watcher.WatchUsers[userWallet.Address] = user
			}
		}
		mu.Unlock()

		gbl.Log.Infof("✅ successfully added group: %s", style.BoldStyle.Render(newWatchGroup.Name))
	}

	userNames := make([]string, 0)
	userWallets := make([]common.Address, 0)

	for _, group := range watcher.Groups {
		for _, user := range group.Users {
			for _, userWallet := range user.Wallets {
				userWallets = append(userWallets, userWallet.Address)
			}

			user.Group = group

			userNames = append(userNames, style.BoldStyle.Render(user.Name))
		}
	}

	// build spinner stop msg with all wallet names
	watchSpinner.StopMessage(fmt.Sprint(
		style.BoldStyle.Render(fmt.Sprint(len(userNames))),
		fmt.Sprintf(" watched users with %s wallets in total: ", style.BoldStyle.Render(fmt.Sprint(len(userWallets)))),
		strings.Join(userNames, ", "),
	) + "\n")

	_ = watchSpinner.Stop()

	return &watcher
}

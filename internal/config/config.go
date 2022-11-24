package config

import (
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/benleb/gloomberg/internal/cache"
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/models/wallet"
	"github.com/benleb/gloomberg/internal/nodes"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/benleb/gloomberg/internal/utils/hooks"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func GetNodesFromConfig() *nodes.Nodes {
	// ethNodes := make([]*nodes.Node, 0)
	var ethNodes nodes.Nodes

	nodesSpinner := style.GetSpinner("setting up node connections...")
	_ = nodesSpinner.Start()

	for idx, nodeConfig := range viper.Get("nodes").([]interface{}) {
		var newNode *nodes.Node

		decodeHooks := mapstructure.ComposeDecodeHookFunc(
			hooks.StringToAddressHookFunc(),
			hooks.StringToLipglossColorHookFunc(),
		)

		decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			DecodeHook: decodeHooks,
			Result:     &newNode,
		})

		err := decoder.Decode(nodeConfig)
		if err != nil {
			gbl.Log.Warnf("reading nodes configuration failed: %+v", err)
			continue
		}

		// set a unique node id
		newNode.NodeID = idx

		// set the default node color to be used to color the marker for example
		if newNode.Color == "" {
			newNode.Color = lipgloss.Color("#1A1A1A")
		}

		// use the node id as the default marker
		if newNode.Marker == "" {
			newNode.Marker = fmt.Sprintf(" %d", idx)
		}

		// connect to the endpoint
		if err := newNode.Connect(); err != nil {
			gbl.Log.Warnf("❌ failed to connect to %s | %s:", newNode.Name, newNode.WebsocketsEndpoint)
			gbl.Log.Warnf("%s %s", style.PinkBoldStyle.PaddingLeft(3).Render("→"), err)
		}

		gbl.Log.Infof("✅ successfully connected to %s", style.BoldStyle.Render(newNode.Name))
		ethNodes = append(ethNodes, newNode)
	}

	// get all node names to be shown as a list of connected nodes
	nodeNames := make([]string, 0)
	for _, n := range ethNodes {
		nodeNames = append(nodeNames, style.BoldStyle.Render(n.Name))
	}

	nodesSpinner.StopMessage(
		fmt.Sprint(style.BoldStyle.Render(fmt.Sprint(len(ethNodes))), " nodes connected: ", strings.Join(nodeNames, ", ")) + "\n",
	)

	_ = nodesSpinner.Stop()

	return &ethNodes
}

func GetOwnWalletsFromConfig(ethNodes *nodes.Nodes) *wallet.Wallets {
	ownWallets := make(map[common.Address]*wallet.Wallet, 0)
	mu := sync.Mutex{}

	nodesSpinner := style.GetSpinner("setting up own wallets...")
	_ = nodesSpinner.Start()

	var wgWallets sync.WaitGroup
	for _, walletConfig := range viper.Get("wallets").([]interface{}) {
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

			if newWallet.Name == "" {
				newWallet.Name = utils.WalletShortAddress(newWallet.Address)
			}

			if newWallet.Color == "" {
				newWallet.Color = style.GenerateColorWithSeed(newWallet.Address.Hash().Big().Int64())
			}

			newWallet.Balance, newWallet.BalanceBefore = big.NewInt(0), big.NewInt(0)

			gbl.Log.Infof("✅ successfully added own wallet: %s", newWallet.Render(newWallet.Name))

			mu.Lock()
			ownWallets[newWallet.Address] = newWallet
			mu.Unlock()
		}(walletConfig)
	}

	// wait for all goroutines to finish
	wgWallets.Wait()

	// resolve addresses to ens names if nodes are available
	gbl.Log.Debugf("ethNodes != nil: %v | %+v | %+v", ethNodes != nil, ethNodes, ownWallets)

	if ethNodes != nil {
		ethNodes.ReverseResolveAllENS((*wallet.Wallets)(&ownWallets))
	}

	// build spinner stop msg with all wallet names
	nodesSpinner.StopMessage(fmt.Sprint(
		style.BoldStyle.Render(fmt.Sprint(len(ownWallets))),
		" wallets: ",
		strings.Join((*wallet.Wallets)(&ownWallets).FormattedNames(), ", "),
	) + "\n")

	_ = nodesSpinner.Stop()

	return (*wallet.Wallets)(&ownWallets)
}

func GetCollectionsFromConfiguration(nodes *nodes.Nodes) []*collections.GbCollection {
	ownCollections := make([]*collections.GbCollection, 0)

	for address, collection := range viper.GetStringMap("collections") {
		contractAddress := common.HexToAddress(address)
		currentCollection := collections.NewCollection(contractAddress, "", nodes, models.FromConfiguration)

		if collection == nil && common.IsHexAddress(address) {
			gbl.Log.Infof("reading collection without details: %+v", address)

			currentCollection = collections.NewCollection(contractAddress, "", nodes, models.FromConfiguration)

			// global settings
			currentCollection.Show.Listings = viper.GetBool("show.listings")
			currentCollection.Show.Sales = viper.GetBool("show.sales")
			currentCollection.Show.Mints = viper.GetBool("show.mints")
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

		gbl.Log.Debugf("currentCollection: %+v", currentCollection)

		ownCollections = append(ownCollections, currentCollection)
	}

	return ownCollections
}

// // GetWatcherUsersFromConfig reads configured users to be notified from config
//func GetWatcherUsersFromConfig() *models.WatcherUsers {
//	mu := sync.Mutex{}
//
//	watcherUsers := make(map[string]bool, 0)
//	watcherUsersWallets := make(map[common.Address]*models.WatchUser, 0)
//
//	watcherSpinner := style.GetSpinner("setting up watched users...")
//	_ = watcherSpinner.Start()
//
//	var wgWatcherUsers sync.WaitGroup
//	for _, watcherUser := range viper.Get("watcher.users").([]interface{}) {
//		wgWatcherUsers.Add(1)
//
//		go func(walletConfig interface{}) {
//			defer wgWatcherUsers.Done()
//
//			var newUser *models.WatchUser
//
//			decodeHooks := mapstructure.ComposeDecodeHookFunc(
//				hooks.StringToAddressHookFunc(),
//				hooks.StringToLipglossColorHookFunc(),
//			)
//
//			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
//				DecodeHook: decodeHooks,
//				Result:     &newUser,
//			})
//
//			err := decoder.Decode(walletConfig)
//			if err != nil {
//				gbl.Log.Warnf("reading wallet group configuration failed: %+v", err)
//				return
//			}
//
//			gbl.Log.Debugf("\n%+v\n", newUser)
//
//			mu.Lock()
//			for _, walletAddress := range newUser.WalletAddresses {
//				watcherUsersWallets[walletAddress] = newUser
//			}
//
//			watcherUsers[newUser.Name] = true
//			mu.Unlock()
//
//			gbl.Log.Infof("✅ successfully added user: %s", style.BoldStyle.Render(newUser.Name))
//		}(watcherUser)
//	}
//
//	// wait for all goroutines to finish
//	wgWatcherUsers.Wait()
//
//	userNames := make([]string, 0)
//	for userName := range watcherUsers {
//		userNames = append(userNames, style.BoldStyle.Render(userName))
//	}
//
//	// build spinner stop msg with all wallet names
//	watcherSpinner.StopMessage(fmt.Sprint(
//		style.BoldStyle.Render(fmt.Sprint(len(watcherUsers))),
//		fmt.Sprintf(" watched users with %s wallets in total: ", style.BoldStyle.Render(fmt.Sprint(len(watcherUsersWallets)))),
//		strings.Join(userNames, ", "),
//	) + "\n")
//
//	_ = watcherSpinner.Stop()
//
//	return (*models.WatcherUsers)(&watcherUsersWallets)
//}

// GetWatchRulesFromConfig reads configured users to be notified from config
func GetWatchRulesFromConfig() models.Watcher {
	mu := sync.Mutex{}

	watcher := models.Watcher{
		UserAddresses:   make(map[common.Address]*models.WatchGroup, 0),
		WalletAddresses: make(map[common.Address]*models.WatchGroup, 0),
		Groups:          make(map[string]*models.WatchGroup, 0),
		WatchUsers:      make(map[common.Address]*models.WatchUser, 0),
	}

	watchSpinner := style.GetSpinner("setting up watch rules...")
	_ = watchSpinner.Start()

	for _, group := range viper.Get("watch").([]interface{}) {
		var newWatchGroup *models.WatchGroup

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

	return watcher
}

func GetBuyRulesFromConfiguration() *models.BuyRules {
	buyRules := &models.BuyRules{
		Rules: make(map[common.Address]*models.BuyRule, 0),
	}

	gbl.Log.Info(" ------ buy rules ------")

	// checking global key & threshold
	globalPrivateKey := viper.GetString("buy.privateKey")
	globalThreshold := viper.GetFloat64("buy.globalThreshold")

	if globalPrivateKey == "" || globalThreshold == 0.0 {
		gbl.Log.Warnf("❌ invalid globalPrivateKey (%s) or globalThreshold (%f), skipping global rule", globalPrivateKey, globalThreshold)
	} else {

		buyRule := &models.BuyRule{
			ID:              0,
			ContractAddress: utils.ZeroAddress,
			PrivateKey:      globalPrivateKey,
			Threshold:       globalThreshold,
		}

		if buyRule.Threshold < 0.0 || buyRule.Threshold > 1.0 {
			gbl.Log.Infof("🤷‍♀️ %d| invalid buyRule.Threshold (%.3f) value, skipping auto-buy", buyRule.ID, buyRule.Threshold)
		} else {

			// buyRules = append(buyRules, buyRule)
			buyRules.Rules[utils.ZeroAddress] = buyRule

			gbl.Log.Debugf("parsed buy rule: %+v", buyRule)
		}
	}

	if viper.Get("buy.rules") == nil {
		gbl.Log.Info("no buy rules configured")
		return buyRules
	}

	for _, rule := range viper.Get("buy.rules").([]interface{}) {
		ruleID := len(buyRules.Rules)

		gbl.Log.Debugf("reading buy rule %d: %+v", ruleID, rule)

		decodeHooks := mapstructure.ComposeDecodeHookFunc(
			hooks.StringToAddressHookFunc(),
			hooks.StringToDurationHookFunc(),
			hooks.StringToLipglossColorHookFunc(),
		)

		var buyRule *models.BuyRule

		decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			DecodeHook: decodeHooks,
			Result:     &buyRule,
		})

		err := decoder.Decode(rule)
		if err != nil {
			gbl.Log.Errorf("error decoding collection: %+v", err)

			continue
		}

		buyRule.ID = len(buyRules.Rules)

		if buyRule.PrivateKey == "" {
			if privateKey := viper.GetString("buy.privateKey"); privateKey != "" {
				gbl.Log.Debugf("no private key found for buy rule %d, using the global privatKey", ruleID)

				buyRule.PrivateKey = privateKey
			} else {

				gbl.Log.Warnf("❌ no private key available, skipping rule")

				continue
			}
		}

		if !common.IsHexAddress(buyRule.ContractAddress.String()) || buyRule.ContractAddress == utils.ZeroAddress || buyRule.Threshold == 0 {
			gbl.Log.Warnf("❌ invalid contract address (%s) or threshold (%f), skipping rule", buyRule.ContractAddress.String(), buyRule.Threshold)

			continue
		}

		if buyRule.Threshold < 0.0 || buyRule.Threshold > 1.0 {
			gbl.Log.Infof("🤷‍♀️ %d| invalid buyRule.Threshold (%.3f) value, skipping auto-buy", buyRule.ID, buyRule.Threshold)
			continue
		}

		gbl.Log.Debugf("parsed buy rule: %+v", buyRule)

		// buyRules = append(buyRules, buyRule)
		buyRules.Rules[buyRule.ContractAddress] = buyRule
	}

	for contractAddress, buyRule := range buyRules.Rules {
		name := buyRule.ContractAddress.String()

		if contractName, err := cache.GetContractName(contractAddress); err == nil && contractName != "" {
			name = contractName
		}

		if buyRule.ContractAddress == utils.ZeroAddress {
			name = "*everything*"
		}

		formattedRuleID := fmt.Sprintf("#%d", buyRule.ID)
		percentageOfFloor := fmt.Sprintf("<=%.0f%%", buyRule.Threshold*100)
		// relativeFloorDifference := int((buyRule.Threshold * 100) - 100)

		// percentageBelowFloor := fmt.Sprintf(">=%.0f%%", math.Abs(float64(relativeFloorDifference)))
		gbl.Log.Infof("Rule %s: buy %s if listed for %s of current floor using key %s...", style.BoldStyle.Render(formattedRuleID), style.BoldStyle.Render(name), style.BoldStyle.Render(percentageOfFloor), style.BoldStyle.Render(buyRule.PrivateKey[:4])) //, style.BoldStyle.Render(percentageBelowFloor))
	}

	gbl.Log.Info(" ----------------------- ")

	return buyRules
}

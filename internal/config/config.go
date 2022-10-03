package config

import (
	"fmt"
	"math/big"
	"strings"
	"sync"

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

func GetNodesFromConfig() nodes.Nodes {
	ethNodes := make([]*nodes.Node, 0)

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

		// use the node id as the default marker
		if newNode.Marker == "" {
			newNode.Marker = fmt.Sprintf(" %d", idx)
		}

		// set the default node color to be used to color the marker for example
		if newNode.Color != "" {
			newNode.Marker = lipgloss.NewStyle().Foreground(newNode.Color).Render(newNode.Marker)
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

	return ethNodes
}

func GetOwnWalletsFromConfig(ethNodes nodes.Nodes) *wallet.Wallets {
	ownWallets := make(map[common.Address]*wallet.Wallet, 0)
	mu := sync.Mutex{}

	nodesSpinner := style.GetSpinner("setting up own wallets...")
	_ = nodesSpinner.Start()

	var wgWallets sync.WaitGroup
	for _, walletConfig := range viper.Get("ownWallets").([]interface{}) {
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

	if viper.IsSet("collections") {
		gbl.Log.Infof("config | reading collections from config")

		for address, collection := range viper.GetStringMap("collections") {
			contractAddress := common.HexToAddress(address)
			currentCollection := collections.NewCollection(contractAddress, "", nodes, collections.Configuration)

			if collection == nil && common.IsHexAddress(address) {
				gbl.Log.Infof("reading collection without details: %+v", address)

				currentCollection = collections.NewCollection(contractAddress, "", nodes, collections.Configuration)
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
	}

	return ownCollections
}

// GetWatcherUsersFromConfig reads configured users to be notified from config
func GetWatcherUsersFromConfig() map[common.Address]*models.WatcherUser {
	mu := sync.Mutex{}

	watcherUsers := make(map[string]bool, 0)
	watcherUsersWallets := make(map[common.Address]*models.WatcherUser, 0)

	watcherSpinner := style.GetSpinner("setting up watched users...")
	_ = watcherSpinner.Start()

	var wgWatcherUsers sync.WaitGroup
	for _, watcherUser := range viper.Get("watcher.users").([]interface{}) {
		wgWatcherUsers.Add(1)

		go func(walletConfig interface{}) {
			defer wgWatcherUsers.Done()

			var newUser *models.WatcherUser

			decodeHooks := mapstructure.ComposeDecodeHookFunc(
				hooks.StringToAddressHookFunc(),
				hooks.StringToLipglossColorHookFunc(),
			)

			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				DecodeHook: decodeHooks,
				Result:     &newUser,
			})

			err := decoder.Decode(walletConfig)
			if err != nil {
				gbl.Log.Warnf("reading wallet group configuration failed: %+v", err)
				return
			}

			gbl.Log.Debugf("\n%+v\n", newUser)

			mu.Lock()
			for _, walletAddress := range newUser.WalletAddresses {
				watcherUsersWallets[walletAddress] = newUser
			}

			watcherUsers[newUser.Name] = true
			mu.Unlock()

			gbl.Log.Infof("✅ successfully added user: %s", newUser.Name)
		}(watcherUser)
	}

	// wait for all goroutines to finish
	wgWatcherUsers.Wait()

	userNames := make([]string, 0)
	for userName := range watcherUsers {
		userNames = append(userNames, userName)
	}

	// build spinner stop msg with all wallet names
	watcherSpinner.StopMessage(fmt.Sprint(
		style.BoldStyle.Render(fmt.Sprint(len(watcherUsers))),
		fmt.Sprintf(" watched users with %s wallets in total: ", style.BoldStyle.Render(fmt.Sprint(len(watcherUsersWallets)))),
		strings.Join(userNames, ", "),
	) + "\n")

	_ = watcherSpinner.Stop()

	return watcherUsersWallets
}

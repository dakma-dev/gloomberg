package cmd

import (
	"fmt"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/chainwatcher"
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/config"
	"github.com/benleb/gloomberg/internal/models/wallet"
	"github.com/benleb/gloomberg/internal/output"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/ticker"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func gloomberg(cmd *cobra.Command, args []string, role internal.RoleMap) {
	// print header
	header := style.GetHeader(Version)
	fmt.Println(header)
	gbl.Log.Info(header)

	outputQueues := make(map[string]interface{})
	queueEvents := make(chan *collections.Event, 1024)
	queueOutput = make(chan string, 1024)

	// websockets server
	if viper.GetBool("server.websockets.enabled") {
		role.WsServer = true
	}

	// telegram notifications
	if viper.GetBool("notifications.telegram") {
		role.TelegramNotifications = true
	}

	if role.OutputTerminal {
		outputQueues["terminal"] = make(chan *collections.Event, 1024)

		for workerID := 1; workerID <= viper.GetInt("server.workers.subscription_logs"); workerID++ {
			gbl.Log.Infof("terminlOutput printer worker %d started", workerID)

			go func(workerID int) {
				for outputLine := range queueOutput {
					gbl.Log.Debugf("%d ~ %d | workerOutput outputLine: %s", workerID, len(queueOutput), outputLine)

					if viper.GetBool("log.debug") {
						debugPrefix := fmt.Sprintf("%s ~ %d | ", style.BoldStyle.Render(fmt.Sprintf("%d", workerID)), len(queueOutput))
						outputLine = fmt.Sprint(debugPrefix, outputLine)
					}

					fmt.Println(outputLine)
				}
			}(workerID)
		}
	}

	if role.ChainWatcher {
		// read nodes from config
		ethNodes := config.GetNodesFromConfig()
		// establish connections to the nodes
		ethNodes.ConnectAllNodes()

		// create a new chainserver instance
		cWatcher = chainwatcher.New(ethNodes)

		//
		// get own wallets from config file
		var ownWallets *wallet.Wallets

		if role.OwnWalletWatcher {
			// get wallets from config file, if nodes are provided,
			// we will try to (reverse) resolve the ENS name
			ownWallets = config.GetOwnWalletsFromConfig(ethNodes)
		}

		//
		// format events and print them to stdout
		for workerID := 1; workerID <= viper.GetInt("server.workers.subscription_logs"); workerID++ {
			// format events from queueEvents in a pretty way for terminal (and later other "outputs")
			// go workerEventFormatter(outputWorkerID, &cWatcher.Nodes, &ownWallets, &queueEvents, &queueOutput, &queueOutWS)
			go func(workerID int) {
				for event := range queueEvents {
					gbl.Log.Debugf("%d ~ %d | workerEventFormatter event: %v", workerID, len(queueEvents), event)

					go output.FormatEvent(event, ownWallets, &cWatcher.Nodes, queueOutput)
				}
			}(workerID)
		}

		// subscribe to the chain logs/events and start the workers
		cWatcher.SubscribeToSales(&queueEvents)

		//
		// ticker & stats
		if role.StatsTicker {
			// gasline ticker
			var gasTicker *time.Ticker

			if tickerInterval := viper.GetDuration("ticker.gasline"); len(cWatcher.Nodes.GetLocalNodes()) > 0 && tickerInterval > 0 {
				// initial startup delay
				time.Sleep(tickerInterval / 5)

				// start gasline ticker
				gasTicker = time.NewTicker(tickerInterval)
				go ticker.GasTicker(gasTicker, &cWatcher.Nodes, &queueOutput)
			}

			// statsbox ticker
			stats = ticker.New(gasTicker, ownWallets, &cWatcher.Nodes, len(cWatcher.OwnCollections.UserCollections))

			// start statsbox ticker
			if statsInterval := viper.GetDuration("stats.interval"); viper.GetBool("stats.enabled") {
				stats.StartTicker(statsInterval)
			}
		}
	}

	// loop forever
	select {}
}

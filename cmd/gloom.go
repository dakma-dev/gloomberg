package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/chainwatcher"
	"github.com/benleb/gloomberg/internal/chainwatcher/subscribe"
	"github.com/benleb/gloomberg/internal/chainwatcher/wwatcher"
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/config"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/models/wallet"
	"github.com/benleb/gloomberg/internal/nodes"
	"github.com/benleb/gloomberg/internal/opensea"
	"github.com/benleb/gloomberg/internal/output"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/ticker"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/benleb/gloomberg/internal/ws"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Version string

func gloomberg(_ *cobra.Command, _ []string, role internal.RoleMap) {
	// print header
	header := style.GetHeader(Version)
	fmt.Println(header)
	gbl.Log.Info(header)

	// show listings for own collections if an opensea api key is set
	if viper.IsSet("api_keys.opensea") && !viper.IsSet("show.listings") {
		viper.Set("show.listings", true)
	}

	watchUsers := make(map[common.Address]*models.WatcherUser)

	outputQueues := make(map[string]chan *collections.Event)

	queueEvents := make(chan *collections.Event, 1024)

	queueOutput := make(chan string, 1024)

	// websockets server
	if viper.GetBool("server.websockets.enabled") {
		role.WsServer = true
	}

	// telegram notifications
	if viper.GetBool("notifications.telegram") {
		role.TelegramNotifications = true
	}

	// websockets server
	if viper.GetBool("server.websockets.enabled") {
		role.WsServer = true
	}

	var cWatcher *chainwatcher.ChainWatcher

	var ethNodes nodes.Nodes

	var ownWallets *wallet.Wallets

	if role.ChainWatcher {
		// read nodes from config
		ethNodes = config.GetNodesFromConfig()
		// establish connections to the nodes
		ethNodes.ConnectAllNodes()

		// create a new chainserver instance
		cWatcher = chainwatcher.New(ethNodes)
	}

	//
	// get own wallets from config file
	if role.OwnWalletWatcher {
		// get wallets from config file, if nodes are provided,
		// we will try to (reverse) resolve the ENS name
		ownWallets = config.GetOwnWalletsFromConfig(ethNodes)
	}

	if role.WalletWatcher {
		watchUsers = config.GetWatcherUsersFromConfig()

		//
		// MIWs
		miwSpinner := style.GetSpinner("setting up MIWs...")
		_ = miwSpinner.Start()

		wwatcher.LoadMIWs()

		if len(wwatcher.MIWC.WeightedMIWs) > 0 {
			miwSpinner.StopMessage(fmt.Sprint(fmt.Sprint(style.BoldStyle.Render(fmt.Sprint(len(wwatcher.MIWC.WeightedMIWs))), " MIWs loaded", "\n")))
			_ = miwSpinner.Stop()
		} else {
			_ = miwSpinner.StopFail()
		}
	}

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
		stats := ticker.New(gasTicker, ownWallets, &cWatcher.Nodes, len(cWatcher.OwnCollections.UserCollections))

		// start statsbox ticker
		if statsInterval := viper.GetDuration("stats.interval"); viper.GetBool("stats.enabled") {
			stats.StartTicker(statsInterval)
		}
	}

	if role.OwnCollections {
		//
		// initialize collections
		collectionsSpinner := style.GetSpinner("setting up collections...")
		_ = collectionsSpinner.Start()

		// collection from config file
		collectionsSpinner.Message("setting up config collections...")

		// read collections from config and store in ownCollections
		configCollections := config.GetCollectionsFromConfiguration(&cWatcher.Nodes)
		configCollectionsNumber := len(configCollections)
		gbl.Log.Infof("collections from config: %d", configCollectionsNumber)

		for _, collection := range configCollections {
			cWatcher.OwnCollections.UserCollections[collection.ContractAddress] = collection
		}

		// collections from wallet holdings
		collectionsSpinner.Message("setting up wallet collections...")

		// read collections hold in wallets from opensea and store in currentCollections
		walletCollections := opensea.GetWalletCollections(*ownWallets, cWatcher.OwnCollections, &cWatcher.Nodes)

		for _, collection := range walletCollections {
			if cWatcher.OwnCollections.UserCollections[collection.ContractAddress] == nil {
				cWatcher.OwnCollections.UserCollections[collection.ContractAddress] = collection
			}
		}

		gbl.Log.Infof("collections from wallets: %d", len(cWatcher.OwnCollections.UserCollections)-configCollectionsNumber)

		// print collections from config & wallet holdings
		if len(cWatcher.OwnCollections.UserCollections) > 0 {
			collectionNames := cWatcher.OwnCollections.SortedAndColoredNames()
			collectionsSpinner.StopMessage(fmt.Sprint(style.BoldStyle.Render(fmt.Sprint(len(collectionNames))), " collections: ", strings.Join(collectionNames, ", "), "\n\n"))
		}

		_ = collectionsSpinner.Stop()
	}

	if role.OutputTerminal {
		outputQueues["terminal"] = make(chan *collections.Event, 1024)

		terminalPrinterQueue := make(chan string, 1024)

		go func() {
			gbl.Log.Info("starting terminal printer...")

			for eventLine := range terminalPrinterQueue {
				gbl.Log.Debugf("terminal printer eventLine: %s", eventLine)

				if viper.GetBool("log.debug") {
					debugPrefix := fmt.Sprintf("%d | ", len(terminalPrinterQueue))
					eventLine = fmt.Sprint(debugPrefix, eventLine)
				}

				fmt.Println(eventLine)
			}
		}()

		for workerID := 1; workerID <= viper.GetInt("server.workers.output"); workerID++ {
			gbl.Log.Infof("starting terminal formatter %d...", workerID)

			go func(workerID int) {
				gbl.Log.Infof("terminal formatter %d started", workerID)

				for event := range outputQueues["terminal"] {
					gbl.Log.Debugf("terminal formatter %d (queue: %d): %v", workerID, len(outputQueues["terminal"]), event)

					go output.FormatEvent(event, ownWallets, watchUsers, &cWatcher.Nodes, terminalPrinterQueue)
				}
			}(workerID)
		}
	}

	if role.OsStreamWatcher {
		//
		// subscribe to sales on the stream api for all collections discovered in wallets and configuration
		if openseaToken := viper.GetString("api_keys.opensea"); openseaToken != "" {
			queueListings := make(chan *models.ItemListedEvent, 1024)

			go func() {
				if client := opensea.NewStreamClient(openseaToken, func(err error) {
					gbl.Log.Error(err)
				}); client != nil {
					for _, collection := range cWatcher.OwnCollections.UserCollections {
						gbl.Log.Debugf("%s: collection.Show.Listings: %v | collection.OpenseaSlug: %s", collection.Name, collection.Show.Listings, collection.OpenseaSlug)

						if collection.Show.Listings {
							if collection.OpenseaSlug == "" {
								if slug := opensea.GetCollectionSlug(collection.ContractAddress); slug != "" {
									collection.OpenseaSlug = slug
								} else {
									gbl.Log.Warnf("no slug for collection %s", collection.ContractAddress)

									continue
								}
							}

							go opensea.SubscribeToListingsForCollectionSlug(client, collection.OpenseaSlug, func(response any) {
								var itemListedEvent models.ItemListedEvent

								err := mapstructure.Decode(response, &itemListedEvent)
								if err != nil {
									gbl.Log.Error("mapstructure.Decode failed for incoming stream api event", err)
								}

								gbl.Log.Infof("received event from opensea: %+v", itemListedEvent.BaseStreamMessage.StreamEvent)

								queueListings <- &itemListedEvent
							})

							collection.ResetStats()

							time.Sleep(337 * time.Millisecond)
							// opensea.SubscribeToEverythingForCollectionSlug(client, collection.OpenseaSlug, streamListingsReceiver)
							// time.Sleep(337 * time.Millisecond)
						}
					}
				}
			}()

			// processes new listings from the opensea stream api
			for listingsWorkerID := 1; listingsWorkerID <= viper.GetInt("workers.listings"); listingsWorkerID++ {
				go subscribe.StreamListingsHandler(listingsWorkerID, cWatcher.OwnCollections, &queueListings, &queueEvents)
			}
		}
	}

	// subscribe to the chain logs/events and start the workers
	cWatcher.SubscribeToSales(&queueEvents)

	if role.WsServer {
		queueWS := make(chan *collections.Event, 1024)
		outputQueues["websockets"] = queueWS

		wsServer := ws.New(viper.GetString("server.websockets.host"), viper.GetUint("server.websockets.port"), &queueWS)
		go wsServer.Start()
		fmt.Printf("🔎 wsServer started: %v\n", wsServer)
	}

	//
	// distribution of the events to the outputs
	for workerID := 1; workerID <= viper.GetInt("server.workers.output"); workerID++ {
		go func(workerID int) {
			for event := range queueEvents {
				gbl.Log.Debugf("%d ~ %d | pushing event to outputs: %v", workerID, len(queueEvents), event)

				for _, outputQueue := range outputQueues {
					outputQueue <- event
				}
			}
		}(workerID)
	}

	// for event := range outputQueues {
	// 	for queueName, queue := range outputQueues {
	// 	gbl.Log.Infof("startung output queue workers for: %s", queueName)

	// 	for workerID := 1; workerID <= viper.GetInt("server.workers.output"); workerID++ {

	// 		go func(workerID int) {
	// 			for outputLine := range queue {
	// 				fmt.Println(outputLine)
	// 			}
	// 		}(workerID)

	// 		gbl.Log.Infof("%s output worker %d started", queueName, workerID)
	// 	}
	// }

	// logsReceivedTicker := time.NewTicker(time.Second * 37)
	// for range logsReceivedTicker.C {
	// 	logsPerNodeFormatted := make([]string, 0)
	// 	logsReceivedTotal := uint64(0)

	// 	for _, node := range cWatcher.Nodes {
	// 		logsPerNodeFormatted = append(logsPerNodeFormatted, fmt.Sprintf("%s: %d", node.Name, node.NumLogsReceived))
	// 		logsReceivedTotal += node.NumLogsReceived
	// 	}

	// 	fmt.Printf("logs received: %d || %s\n", logsReceivedTotal, strings.Join(logsPerNodeFormatted, " | "))

	// 	gbl.Log.Infof("logs received: %d", logsReceivedTotal)

	// }

	// loop forever
	select {}
}

// func streamListingsReceiver(response any) {
// 	var itemListedEvent models.ItemListedEvent

// 	err := mapstructure.Decode(response, &itemListedEvent)
// 	if err != nil {
// 		gbl.Log.Error("mapstructure.Decode failed for incoming stream api event", err)
// 	}

// 	gbl.Log.Infof("received event from opensea: %+v", itemListedEvent.BaseStreamMessage.StreamEvent)

// 	// var queueLising chan *models.ItemListedEvent = queues["listings"].(*chan *models.ItemListedEvent)
// 	*queues["listings"].(*chan *models.ItemListedEvent) <- &itemListedEvent
// }

package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal/chainwatcher"
	"github.com/benleb/gloomberg/internal/chainwatcher/subscribe"
	"github.com/benleb/gloomberg/internal/chainwatcher/wwatcher"
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/config"
	"github.com/benleb/gloomberg/internal/gloomclient"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/models/gloomberg"
	"github.com/benleb/gloomberg/internal/models/wallet"
	"github.com/benleb/gloomberg/internal/opensea"
	ossw "github.com/benleb/gloomberg/internal/osstreamwatcher"
	"github.com/benleb/gloomberg/internal/output"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/ticker"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/benleb/gloomberg/internal/web"
	"github.com/benleb/gloomberg/internal/ws"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Version string

func runGloomberg(_ *cobra.Command, _ []string, role gloomberg.RoleMap) {
	// print header
	header := style.GetHeader(Version)
	fmt.Println(header)
	gbl.Log.Info(header)

	// global defaults
	viper.Set("http.timeout", 17*time.Second)

	// show listings for own collections if an opensea api key is set
	if viper.IsSet("api_keys.opensea") && !viper.IsSet("show.listings") {
		viper.Set("show.listings", true)
	}

	// websockets server
	if viper.GetBool("server.websockets.enabled") {
		role.WsServer = true
	}

	// telegram notifications
	if viper.GetBool("notifications.telegram") {
		role.TelegramNotifications = true
	}

	gb := &gloomberg.Gloomberg{
		ChainWatcher: nil,
		CollectionDB: collections.New(),
		OwnWallets:   &wallet.Wallets{},
		WatchUsers:   &models.WatcherUsers{},
		OutputQueues: make(map[string]chan *collections.Event),
		Role:         role,
	}

	queueEvents := make(chan *collections.Event, 1024)

	//
	// connect to ethereum nodes and create the chainwatcher
	if role.ChainWatcher {
		// read nodes from config & establish connections to the nodes
		if ethNodes := config.GetNodesFromConfig(); ethNodes.ConnectAllNodes() != nil {
			gb.Nodes = ethNodes
		}

		// create chainserver
		if cWatcher := chainwatcher.New(gb.Nodes, gb.CollectionDB); cWatcher == nil {
			gbl.Log.Fatal("❌ running chainwatcher failed, exiting")
		} else {
			gb.ChainWatcher = cWatcher
		}
	}

	//
	// websockets client to get events from a server instead directly from the chain (nodes)
	if role.GloomClient {
		gloomclient.ConnectToServer("ws://10.0.0.99:42069/", &queueEvents)
	}

	//
	// get collections from config file
	collectionsSpinner := style.GetSpinner("setting up collections...")
	_ = collectionsSpinner.Start()

	// collection from config file
	collectionsSpinner.Message("setting up config collections...")

	for _, collection := range config.GetCollectionsFromConfiguration(gb.Nodes) {
		gb.CollectionDB.Collections[collection.ContractAddress] = collection
	}

	// print collections from config & wallet holdings
	if len(gb.CollectionDB.Collections) > 0 {
		collectionNames := gb.CollectionDB.SortedAndColoredNames()
		collectionsSpinner.StopMessage(fmt.Sprint(style.BoldStyle.Render(fmt.Sprint(len(collectionNames))), " collections from config: ", strings.Join(collectionNames, ", "), "\n"))
	}

	// stop spinner
	_ = collectionsSpinner.Stop()

	//
	// get own wallets from config file
	if role.OwnWalletWatcher {
		gb.OwnWallets = config.GetOwnWalletsFromConfig(gb.Nodes)
	}

	//
	// initialize collections database
	if role.CollectionDB {
		collectionsSpinner := style.GetSpinner("setting up collections...")
		_ = collectionsSpinner.Start()

		if len(gb.OwnWallets.GetAll()) > 0 {
			// collections from wallet holdings
			collectionsSpinner.Message("setting up wallet collections...")

			// read collections hold in wallets from opensea and store in currentCollections
			gbl.Log.Debugf("gb.OwnWallets: %v | gb.CollectionDB: %+v | gb.Nodes: %+v", gb.OwnWallets, gb.CollectionDB, gb.Nodes)
			walletCollections := opensea.GetWalletCollections(gb.OwnWallets, gb.CollectionDB, gb.Nodes)

			for _, collection := range walletCollections {
				if gb.CollectionDB.Collections[collection.ContractAddress] == nil {
					gb.CollectionDB.Collections[collection.ContractAddress] = collection
				}
			}

			gbl.Log.Infof("collections from wallets: %d", len(walletCollections))
		}

		// print collections from config & wallet holdings
		if len(gb.CollectionDB.Collections) > 0 {
			collectionNames := gb.CollectionDB.SortedAndColoredNames()
			collectionsSpinner.StopMessage(fmt.Sprint(style.BoldStyle.Render(fmt.Sprint(len(collectionNames))), " collections from config & wallets: ", strings.Join(collectionNames, ", "), "\n"))
		}

		_ = collectionsSpinner.Stop()
	}

	//
	// wallet watcher (todo) & MIWs
	if role.WalletWatcher {
		gb.WatchUsers = config.GetWatcherUsersFromConfig()

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

	fmt.Println()

	//
	// print to terminal
	if role.OutputTerminal {
		gb.OutputQueues["terminal"] = make(chan *collections.Event, 1024)

		terminalPrinterQueue := make(chan string, 1024)

		// ticker & stats
		if role.StatsTicker && gb.OutputQueues["terminal"] != nil {
			// gasline ticker
			var gasTicker *time.Ticker

			if tickerInterval := viper.GetDuration("ticker.gasline"); gb.Nodes != nil && len(gb.Nodes.GetLocalNodes()) > 0 && tickerInterval > 0 {
				// initial startup delay
				time.Sleep(tickerInterval / 5)

				// start gasline ticker
				gasTicker = time.NewTicker(tickerInterval)
				go ticker.GasTicker(gasTicker, gb.Nodes, &terminalPrinterQueue)
			}

			// statsbox ticker
			stats := ticker.New(gasTicker, gb.OwnWallets, gb.Nodes, len(gb.CollectionDB.Collections))

			// start statsbox ticker
			if statsInterval := viper.GetDuration("stats.interval"); viper.GetBool("stats.enabled") {
				stats.StartTicker(statsInterval)
			}
		}

		//
		// event formatter for terminal output
		for workerID := 1; workerID <= viper.GetInt("server.workers.output"); workerID++ {
			gbl.Log.Debugf("starting terminal formatter %d...", workerID)

			go func(workerID int) {
				gbl.Log.Debugf("terminal formatter %d started", workerID)

				for event := range gb.OutputQueues["terminal"] {
					gbl.Log.Debugf("terminal formatter %d (queue: %d): %v", workerID, len(gb.OutputQueues["terminal"]), event)

					go output.FormatEvent(gb, event, terminalPrinterQueue)
				}
			}(workerID)
		}

		//
		// terminal printer
		for workerID := 1; workerID <= viper.GetInt("server.workers.output"); workerID++ {
			go func() {
				gbl.Log.Debug("starting terminal printer...")

				for eventLine := range terminalPrinterQueue {
					gbl.Log.Debugf("terminal printer eventLine: %s", eventLine)

					if viper.GetBool("log.debug") {
						debugPrefix := fmt.Sprintf("%d | ", len(terminalPrinterQueue))
						eventLine = fmt.Sprint(debugPrefix, eventLine)
					}

					fmt.Println(eventLine)
				}
			}()
		}
	}

	//
	// opensea stream api watcher
	if role.OsStreamWatcher {
		// subscribe to sales on the stream api for all collections discovered in wallets and configuration
		if openseaToken := viper.GetString("api_keys.opensea"); openseaToken != "" {
			streamWatcher := ossw.NewStreamWatcher(openseaToken, nil)

			if streamWatcher != nil {
				go func() {
					for _, collection := range gb.CollectionDB.Collections {
						if collection.Show.Listings {
							if collection.OpenseaSlug == "" {
								if slug := opensea.GetCollectionSlug(collection.ContractAddress); slug != "" {
									collection.OpenseaSlug = slug
								} else {
									gbl.Log.Warnf("❌ subscribe to listings for collection %s failed: no slug found", collection.ContractAddress)
									continue
								}
							}

							// go client.SubscribeToListingsFor(collection.OpenseaSlug)
							go streamWatcher.OnItemListed(collection.OpenseaSlug, nil)

							collection.ResetStats()

							time.Sleep(337 * time.Millisecond)
						}
					}
				}()
			}
			// processes new listings from the opensea stream api
			for listingsWorkerID := 1; listingsWorkerID <= viper.GetInt("server.workers.listings"); listingsWorkerID++ {
				go subscribe.StreamListingsHandler(listingsWorkerID, gb.CollectionDB, &streamWatcher.QueueListings, &queueEvents)
			}
		}
	}

	//
	// subscribe to the chain logs/events and start the workers
	if role.ChainWatcher {
		gb.ChainWatcher.SubscribeToSales(&queueEvents)
		// gb.ChainWatcher.SubscribeToOrderFulfilled(&queueEvents)
	}

	//
	// websockets server
	if role.WsServer {
		queueWS := make(chan *collections.Event, 1024)
		gb.OutputQueues["websockets"] = queueWS

		wsServer := ws.New(viper.GetString("server.websockets.host"), viper.GetUint("server.websockets.port"), &queueWS)
		go wsServer.Start()
		fmt.Printf("🔎 wsServer started: %v\n", wsServer)
	}

	//
	// web server
	// if role.WsServer {
	queueWeb := make(chan *collections.Event, 1024)
	gb.OutputQueues["web"] = queueWeb

	eventStream := web.New(&queueWeb)
	go eventStream.Start()
	fmt.Printf("🔎 event stream webserver started: %v\n", eventStream)
	// }

	//
	// distribution of the events to the outputs
	for workerID := 1; workerID <= viper.GetInt("server.workers.output"); workerID++ {
		go func(workerID int) {
			for event := range queueEvents {
				gbl.Log.Debugf("%d ~ %d | pushing event to outputs: %v", workerID, len(queueEvents), event)

				for _, outputQueue := range gb.OutputQueues {
					outputQueue <- event
				}
			}
		}(workerID)
	}

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

	// go func() {
	// 	// Run the server
	// 	http.Handle("/", live.NewHttpHandler(live.NewCookieStore("session-name", []byte("ZWh0NGkzdHZxNjY2NjZxNDg1NWJwdjk0NmM1YnA5MkM2NQ")), web.NewEventHandler()))
	// 	http.Handle("/live.js", live.Javascript{})
	// 	http.Handle("/auto.js.map", live.JavascriptMap{})

	// 	if err := http.ListenAndServe(":18080", nil); err != nil {
	// 		fmt.Printf("error: %s", err)
	// 		gbl.Log.Error(err)
	// 	}
	// }()

	// loop forever
	select {}
}

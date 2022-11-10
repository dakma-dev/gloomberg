package cmd

import (
	"fmt"
	"net"
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
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Version string

// liveCmd represents the live command
var liveCmd = &cobra.Command{
	Use:   "live",
	Short: "watch the chain stream",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runGloomberg(cmd, args)
	},
}

func runGloomberg(_ *cobra.Command, _ []string) { //, role gloomberg.RoleMap) {
	// print header
	header := style.GetHeader(Version)
	fmt.Println(header)
	gbl.Log.Info(header)

	// global defaults
	viper.Set("http.timeout", 27*time.Second)

	// show listings for own collections if an opensea api key is set
	if viper.IsSet("api_keys.opensea") && !viper.IsSet("listings.enabled") {
		viper.Set("listings.enabled", true)
		gbl.Log.Infof("listings from opensea: %v", viper.GetBool("listings.enabled"))
	}

	// dump.P(viper.AllSettings())
	fmt.Println()
	fmt.Println()

	gb := &gloomberg.Gloomberg{
		ChainWatcher: nil,
		CollectionDB: collections.New(),
		OwnWallets:   &wallet.Wallets{},
		Watcher:      &models.Watcher{},
		WatchUsers:   &models.WatcherUsers{},
		OutputQueues: make(map[string]chan *collections.Event),
		QueueSlugs:   make(chan common.Address, 1024),
	}

	queueEvents := make(chan *collections.Event, 1024)

	//
	// connect to ethereum nodes and create the chainwatcher
	// if role.ChainWatcher {
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

	//
	// subscribe to the chain logs/events and start the workers
	// gb.ChainWatcher.SubscribeToOrderFulfilled(&queueEvents)
	gb.ChainWatcher.SubscribeToSales(&queueEvents)
	// }

	//
	// websockets client to get events from a server instead directly from the chain (nodes)
	if viper.GetBool("client") {
		gloomclient.ConnectToServer("ws://10.0.0.99:42068/", &queueEvents)
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
	if viper.GetBool("sales.enabled") {
		gb.OwnWallets = config.GetOwnWalletsFromConfig(gb.Nodes)
	}

	//
	// initialize collections database
	if viper.GetBool("sales.enabled") {
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
	if viper.GetBool("sales.enabled") {
		// gb.WatchUsers = config.GetWatcherUsersFromConfig()
		// gb.WatchUsers = config.GetWatchRulesFromConfig()
		watcher := config.GetWatchRulesFromConfig()
		gb.Watcher = &watcher

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

	slugTicker := time.NewTicker(7 * time.Second)
	go chainwatcher.SlugWorker(slugTicker, &gb.QueueSlugs)

	//
	// print to terminal
	if !viper.GetBool("ui.headless") {
		gb.OutputQueues["terminal"] = make(chan *collections.Event, 1024)

		terminalPrinterQueue := make(chan string, 1024)

		// ticker & stats
		// if role.StatsTicker && gb.OutputQueues["terminal"] != nil {
		if gb.OutputQueues["terminal"] != nil {
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
			if statsInterval := viper.GetDuration("ticker.statsbox"); viper.GetBool("stats.enabled") {
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
	if viper.GetBool("listings.enabled") {
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
				go subscribe.StreamListingsHandler(gb, listingsWorkerID, &streamWatcher.QueueListings, &queueEvents)
			}
		}
	}

	//
	// subscribe to the chain logs/events and start the workers
	// if role.ChainWatcher {
	// 	gb.ChainWatcher.SubscribeToSales(&queueEvents)
	// 	// gb.ChainWatcher.SubscribeToOrderFulfilled(&queueEvents)
	// }

	//
	// websockets server
	if viper.GetBool("server.websockets.enabled") {
		queueWS := make(chan *collections.Event, 1024)
		gb.OutputQueues["websockets"] = queueWS

		wsServer := ws.New(viper.GetString("server.websockets.host"), viper.GetUint("server.websockets.port"), &queueWS)
		go wsServer.Start()
		fmt.Printf("📡 websockets server started on %s:%d\n", viper.GetString("server.websockets.host"), viper.GetUint("server.websockets.port"))
	}

	//
	// web ui
	if viper.GetBool("ui.web.enabled") {
		webSpinner := style.GetSpinner("setting up web ui...")
		_ = webSpinner.Start()

		queueWeb := make(chan *collections.Event, 1024)
		gb.OutputQueues["web"] = queueWeb

		listenAddress := viper.GetString("ui.web.host") + ":" + viper.GetString("ui.web.port")
		gb.WebEventStream = web.New(&queueWeb, listenAddress, gb.Nodes)

		go gb.WebEventStream.Start()

		uiURL := fmt.Sprintf("http://%s", listenAddress)
		uiLink := style.TerminalLink(uiURL, style.BoldStyle.Render(uiURL))

		webSpinner.StopMessage(fmt.Sprintf("web ui running: %s", uiLink))

		// stop spinner
		_ = webSpinner.Stop()
	}

	fmt.Println()
	fmt.Println()

	//
	// distribution of the events to the outputs
	for workerID := 1; workerID <= viper.GetInt("server.workers.output"); workerID++ {
		go func(workerID int) {
			for event := range queueEvents {
				gbl.Log.Debugf("%d ~ %d | pushing event to outputs...", workerID, len(queueEvents)) // , event)

				for outputName, outputQueue := range gb.OutputQueues {
					gbl.Log.Debugf("%d ~ %d | pushing event to %s queue", workerID, len(queueEvents), outputName)
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

	// // buy test
	// time.Sleep(10 * time.Second)
	// tx, err := seaport.FulfillBasicOrder(gb, &models.SeaportOrder{}, viper.GetString("buy.privateKey"))
	// if err != nil {
	// 	gbl.Log.Error("❌ purchase failed: ", err)
	// } else {
	// 	gbl.Log.Info("✅ purchase succeeded: ", tx)
	// }

	// loop forever
	select {}
}

func init() {
	rootCmd.AddCommand(liveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// liveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// liveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// main
	liveCmd.Flags().Bool("watch-sales", true, "get sales")
	_ = viper.BindPFlag("sales.enabled", liveCmd.Flags().Lookup("watch-sales"))
	liveCmd.Flags().Bool("watch-listings", false, "get (opensea) listings for own collections")
	_ = viper.BindPFlag("listings.enabled", liveCmd.Flags().Lookup("watch-listings"))

	// websockets server
	liveCmd.Flags().Bool("websockets", false, "enable websockets server")
	_ = viper.BindPFlag("server.websockets.enabled", liveCmd.Flags().Lookup("websockets"))
	liveCmd.Flags().IP("websockets-host", net.IPv4(0, 0, 0, 0), "websockets listen address")
	_ = viper.BindPFlag("server.websockets.host", liveCmd.Flags().Lookup("websockets-host"))
	liveCmd.Flags().Uint16("websockets-port", 42068, "websockets server port")
	_ = viper.BindPFlag("server.websockets.port", liveCmd.Flags().Lookup("websockets-port"))

	// notifications
	liveCmd.Flags().Bool("telegram", false, "send telegram notifications")
	_ = viper.BindPFlag("notifications.telegram.enabled", liveCmd.Flags().Lookup("telegram"))

	// no ui
	liveCmd.Flags().Bool("headless", false, "run without terminal output")
	_ = viper.BindPFlag("ui.headless", liveCmd.Flags().Lookup("headless"))

	// web ui
	liveCmd.Flags().Bool("web-ui", false, "enable web ui")
	_ = viper.BindPFlag("ui.web.enabled", liveCmd.Flags().Lookup("web-ui"))
	liveCmd.Flags().IP("web-ui-host", net.IPv4(0, 0, 0, 0), "web ui listen address")
	_ = viper.BindPFlag("ui.web.host", liveCmd.Flags().Lookup("web-ui-host"))
	liveCmd.Flags().Uint16("web-ui-port", 42069, "web ui port")
	_ = viper.BindPFlag("ui.web.port", liveCmd.Flags().Lookup("web-ui-port"))

	// wallets
	liveCmd.Flags().StringSliceVarP(&ownWallets, "wallets", "w", []string{}, "Own wallet addresses")
	_ = viper.BindPFlag("wallets", liveCmd.Flags().Lookup("wallets"))

	// min value for sales to be shown
	liveCmd.Flags().Float64("min-value", 0.0, "minimum value to show sales")
	_ = viper.BindPFlag("show.min_value", liveCmd.Flags().Lookup("min-value"))

	// what to show
	liveCmd.Flags().Bool("show-mints", false, "Show mints")
	_ = viper.BindPFlag("show.mints", liveCmd.Flags().Lookup("show-mints"))
	liveCmd.Flags().Bool("show-transfers", false, "Show transfers")
	_ = viper.BindPFlag("show.transfers", liveCmd.Flags().Lookup("show-transfers"))
	// liveCmd.Flags().Bool("sales", true, "Show sales?")
	// _ = viper.BindPFlag("show.sales", liveCmd.Flags().Lookup("sales"))
	// liveCmd.Flags().Bool("listings", false, "Show listings?")
	// _ = viper.BindPFlag("show.listings", liveCmd.Flags().Lookup("listings"))

	// worker settings
	viper.SetDefault("server.workers.subscription_logs", 5)
	viper.SetDefault("server.workers.output", 3)
	viper.SetDefault("server.workers.listings", 2)

	viper.SetDefault("opensea.auto_list_min_sales", 50000)

	// ticker
	viper.SetDefault("ticker.statsbox", time.Second*89)
	viper.SetDefault("ticker.gasline", time.Second*37)
	viper.SetDefault("ticker.divider", time.Second*89)

	viper.SetDefault("stats.enabled", true)
	viper.SetDefault("stats.balances", true)
	viper.SetDefault("stats.lines", 5)
}

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/config"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/provider"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/benleb/gloomberg/internal/nemo/wallet"
	"github.com/benleb/gloomberg/internal/nemo/watch"
	"github.com/benleb/gloomberg/internal/nepa"
	"github.com/benleb/gloomberg/internal/opensea"
	"github.com/benleb/gloomberg/internal/pusu"
	"github.com/benleb/gloomberg/internal/rueidica"
	"github.com/benleb/gloomberg/internal/seawa"
	"github.com/benleb/gloomberg/internal/stats"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/ticker"
	"github.com/benleb/gloomberg/internal/trapri"
	"github.com/benleb/gloomberg/internal/utils/slugs"
	"github.com/benleb/gloomberg/internal/utils/wwatcher"
	"github.com/benleb/gloomberg/internal/web"
	"github.com/benleb/gloomberg/internal/ws"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/redis/rueidis"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// liveCmd represents the live command.
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

func runGloomberg(_ *cobra.Command, _ []string) {
	// print header
	header := style.GetHeader(internal.GloombergVersion)
	fmt.Println(header)
	gbl.Log.Info(header)

	// global defaults
	viper.Set("http.timeout", 27*time.Second)

	// show listings for own collections if an opensea api key is set
	// if viper.IsSet("api_keys.opensea") && !viper.IsSet("listings.enabled") {
	if apiKey := viper.GetString("api_keys.opensea"); apiKey != "" && !viper.IsSet("listings.enabled") {
		viper.Set("listings.enabled", true)
		gbl.Log.Infof("listings from opensea: %v", viper.GetBool("listings.enabled"))
	}

	rdb := GetRedisClient()

	gb := &gloomberg.Gloomberg{
		CollectionDB: collections.New(),
		OwnWallets:   &wallet.Wallets{},
		Watcher:      &watch.Watcher{},

		QueueSlugs: make(chan common.Address, 1024),

		Rdb:    rdb,
		Rueidi: rueidica.NewRueidica(rdb),
	}

	// cleanup for redis db/cache
	defer gb.Rdb.Close()

	// compatibility with old config key
	var providerConfig interface{}
	if cfg := viper.Get("provider"); cfg != nil {
		providerConfig = cfg
	} else {
		providerConfig = viper.Get("nodes")
	}

	//
	// init provider pool
	if pool, err := provider.FromConfig(providerConfig); err != nil {
		gbl.Log.Fatal("❌ running provider failed, exiting")
	} else if pool != nil {
		gb.ProviderPool = pool
		gb.ProviderPool.Rueidi = gb.Rueidi
	}

	//
	// queue for everything to print to the console
	terminalPrinterQueue := make(chan string, 256)

	if viper.GetBool("notifications.smart_wallets.enabled") {
		alphaTicker := ticker.NewAlphaScore(gb)
		go alphaTicker.AlphaCallerTicker(gb, time.NewTicker(time.Minute*1))
	}

	// nepa
	queueTokenTransactions := make(chan *totra.TokenTransaction, 256)
	queueWsOutTokenTransactions := make(chan *totra.TokenTransaction, 256)
	queueWsInTokenTransactions := make(chan *totra.TokenTransaction, 256)
	nePa := nepa.NewNePa(gb, queueTokenTransactions)

	// trapri | ttx printer to process and format the token transactions
	for workerID := 1; workerID <= viper.GetInt("server.workers.ttxFormatter"); workerID++ {
		go trapri.TokenTransactionFormatter(gb, queueTokenTransactions, queueWsOutTokenTransactions, queueWsInTokenTransactions, terminalPrinterQueue)
	}

	// start subscribing
	go nePa.Run()

	// if viper.GetBool("websockets.server.enabled") {
	// 	// queueWS := make(chan *collections.Event, 1024)
	// 	// gb.OutputQueues["websockets"] = queueWS

	// 	wsServer := ws.New(viper.GetString("websockets.server.host"), viper.GetUint("websockets.server.port"), &queueWsOutTokenTransactions)
	// 	go wsServer.Start()

	// 	gbl.Log.Infof("📡 websockets server started on %s:%d\n", viper.GetString("websockets.server.host"), viper.GetUint("websockets.server.port"))
	// }

	//
	// websockets client
	if viper.GetBool("websockets.client.enabled") {
		ws.StartWsClient(viper.GetString("websockets.client.url"), &queueWsInTokenTransactions)
	}

	// //
	// // websockets client to get events from a server instead directly from the chain (nodes)
	// if viper.GetBool("client") {
	// 	gloomclient.ConnectToServer("ws://10.0.0.99:42068/", &queueEvents)
	// }

	//
	// get collections from config file
	collectionsSpinner := style.GetSpinner("setting up collections...")
	_ = collectionsSpinner.Start()

	// collection from config file
	collectionsSpinner.Message("setting up config collections...")

	for _, collection := range config.GetCollectionsFromConfiguration(gb.ProviderPool, gb.Rueidi) {
		gb.CollectionDB.RWMu.Lock()
		gb.CollectionDB.Collections[collection.ContractAddress] = collection
		gb.CollectionDB.RWMu.Unlock()

		// // buy rules
		// if buyRule := collection.BuyRule; buyRule != nil {
		// 	gbl.Log.Debugf("🛍️ buy rule for %s: %v", collection.ContractAddress.Hex(), buyRule)
		// 	gb.BuyRules.Rules[collection.ContractAddress] = buyRule
		// } else {
		// 	gbl.Log.Debugf("🛍️ no buy rule for %s", collection.ContractAddress.Hex())
		// }
	}

	// //
	// // general buy rule
	// if buyRule := config.GetGeneralBuyRuleFromConfiguration(); buyRule != nil {
	// 	gb.BuyRules.Rules[utils.ZeroAddress] = buyRule
	// }

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
		gb.OwnWallets = config.GetOwnWalletsFromConfig(gb.ProviderPool)
	}

	//
	// initialize collections database
	if viper.GetBool("sales.enabled") {
		collectionsSpinner := style.GetSpinner("setting up collections...")
		_ = collectionsSpinner.Start()

		if len(*gb.OwnWallets) > 0 {
			// collections from wallet holdings
			collectionsSpinner.Message("setting up wallet collections...")

			// read collections hold in wallets from opensea and store in currentCollections
			gbl.Log.Debugf("gb.OwnWallets: %v | gb.CollectionDB: %+v | gb.ProviderPool: %+v", gb.OwnWallets, gb.CollectionDB, gb.ProviderPool)
			// walletCollections := opensea.GetWalletCollections(gb.OwnWallets, gb.CollectionDB, gb.Nodes)
			walletCollections := opensea.GetWalletCollections(gb)

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

	// for _, buyRule := range gb.BuyRules.Rules {
	// 	percentageOfFloor := fmt.Sprintf("<=%.0f%%", buyRule.Threshold*100)

	// 	out := strings.Builder{}

	// 	// single line
	// 	out.WriteString(fmt.Sprintf("rule %s:", style.BoldStyle.Render(buyRule.Name)))

	// 	if buyRule.MaxPrice > 0.0 {
	// 		out.WriteString(fmt.Sprintf(" max: %sΞ", style.BoldStyle.Render(fmt.Sprintf("%4.2f", buyRule.MaxPrice))))
	// 	}

	// 	if buyRule.MaxPrice == 0.0 && buyRule.Threshold > 0.0 {
	// 		out.WriteString(fmt.Sprintf(" | threshold: %s%% of floor", style.BoldStyle.Render(percentageOfFloor)))
	// 	}

	// 	out.WriteString(fmt.Sprintf(" | min: %ss / %sl", style.BoldStyle.Render(fmt.Sprint(buyRule.MinSales)), style.BoldStyle.Render(fmt.Sprint(buyRule.MinListings))))

	// 	// print buy rules
	// 	gbl.Log.Infof(out.String())
	// }

	//
	// wallet watcher (todo) & MIWs
	if viper.GetBool("sales.enabled") {
		watcher := config.GetWatchRulesFromConfig()
		gb.Watcher = watcher

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
	// start central terminal printer
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

	slugTicker := time.NewTicker(7 * time.Second)
	go slugs.SlugWorker(slugTicker, &gb.QueueSlugs, gb.Rueidi)

	//
	// gasline ticker
	var gasTicker *time.Ticker

	if tickerInterval := viper.GetDuration("ticker.gasline"); gb.ProviderPool != nil && gb.ProviderPool.PreferredProviderAvailable() && tickerInterval > 0 {
		// initial startup delay
		time.Sleep(tickerInterval / 5)

		// start gasline ticker
		gasTicker = time.NewTicker(tickerInterval)
		go stats.GasTicker(gasTicker, gb.ProviderPool, terminalPrinterQueue)
	}

	// manifold ticker
	if viper.GetBool("notifications.manifold.enabled") && (!viper.GetBool("notifications.disabled")) {
		manifoldTicker := time.NewTicker(time.Hour * 1)
		newManifoldTicker := ticker.NewManifoldTicker(gb)

		if viper.GetBool("notifications.manifold.enabled") {
			go newManifoldTicker.ManifoldTicker(manifoldTicker, &terminalPrinterQueue)
			fmt.Println("Manifold notifications started")
		}

		manifoldTickerDakma := time.NewTicker(time.Minute * 1)
		go newManifoldTicker.OneMinuteTicker(manifoldTickerDakma)
	}

	if viper.GetBool("notifications.bluechip.enabled") {
		// blue chip ticker
		newBluechipTicker := ticker.NewBlueChipTicker(gb)
		go newBluechipTicker.BlueChipTicker(time.NewTicker(time.Minute*5), &terminalPrinterQueue)
	}

	//
	// statsbox
	gb.Stats = stats.New(gasTicker, gb.OwnWallets, gb.ProviderPool, gb.Rdb)

	if statsInterval := viper.GetDuration("ticker.statsbox"); viper.GetBool("stats.enabled") {
		gb.Stats.StartTicker(statsInterval, terminalPrinterQueue)
	}

	//
	// subscribe to OpenSea API
	if viper.GetBool("listings.enabled") {
		go runSeawatcher(nil, nil)
	}

	//
	// subscribe to redis pubsub channel to receive events from gloomberg central
	if viper.GetBool("pubsub.listings.subscribe") {
		// subscribe to redis pubsub channel
		go pusu.SubscribeToListings(gb, queueTokenTransactions)

		// initially send all our slugs & events to subscribe to
		gb.SendSlugsToServer()

		go func() {
			err := gb.Rdb.Receive(context.Background(), gb.Rdb.B().Subscribe().Channel(internal.TopicSeaWatcherMgmt).Build(), func(msg rueidis.PubSubMessage) {
				gbl.Log.Debug(fmt.Sprintf("🚇 received msg on %s: %s", msg.Channel, msg.Message))

				var mgmtEvent *seawa.MgmtEvent

				if err := json.Unmarshal([]byte(msg.Message), &mgmtEvent); err != nil {
					gbl.Log.Fatal(fmt.Sprintf("❌ error json.Unmarshal: %+v", err))
				}

				if mgmtEvent.Action == seawa.SendSlugs {
					gbl.Log.Info(fmt.Sprintf("🚇 SendSlugs received on channel %s", msg.Channel))
					gb.SendSlugsToServer()
				}
			})
			if err != nil {
				gbl.Log.Errorf("❌ error subscribing to redis channels %s: %s", internal.TopicSeaWatcherMgmt, err.Error())

				return
			}
		}()

		// go func() {
		// 	// loop over incoming events
		// 	for msg := range ch {
		// 		gbl.Log.Debug(fmt.Sprintf("🚇 received msg on %s: %s", msg.Channel, msg.Payload))

		// 		var mgmtEvent *seawa.MgmtEvent

		// 		if err := json.Unmarshal([]byte(msg.Payload), &mgmtEvent); err != nil {
		// 			gbl.Log.Fatal(fmt.Sprintf("❌ error json.Unmarshal: %+v", err))
		// 		}

		// 		if mgmtEvent.Action == seawa.SendSlugs {
		// 			gbl.Log.Info(fmt.Sprintf("🚇 SendSlugs received on channel %s", msg.Channel))
		// 			gb.SendSlugsToServer()
		// 		}
		// 	}
		// }()
	}

	//
	// web ui
	if viper.GetBool("web.enabled") {
		webSpinner := style.GetSpinner("setting up web ui...")
		_ = webSpinner.Start()

		go web.StartWebUI(queueWsOutTokenTransactions)

		webSpinner.StopMessage(fmt.Sprintf("web ui running %s", "👍"))

		// stop spinner
		_ = webSpinner.Stop()
	}

	// //
	// // web ui
	// if viper.GetBool("web.enabled") {
	// 	webSpinner := style.GetSpinner("setting up web ui...")
	// 	_ = webSpinner.Start()

	// 	queueWeb := make(chan *totra.TokenTransaction, 1024)
	// 	// gb.OutputQueues["web"] = queueWeb

	// 	listenHost := net.ParseIP(viper.GetString("web.host"))
	// 	listenPort := viper.GetUint("web.port")
	// 	listenAddress := net.JoinHostPort(listenHost.String(), strconv.Itoa(int(listenPort)))

	// 	// webJLive := web.New(&queueWeb, listenAddress, gb.Nodes, nil)
	// 	// go webJLive.Start()
	// 	gloomWeb := web.NewGloomWeb(listenAddress, &queueWeb)
	// 	go func() { log.Fatal(gloomWeb.Run()) }()

	// 	uiURL := fmt.Sprintf("https://%s", listenAddress)
	// 	uiLink := style.TerminalLink(uiURL, style.BoldStyle.Render(uiURL))

	// 	webSpinner.StopMessage(fmt.Sprintf("web ui running: %s", uiLink))

	// 	// stop spinner
	// 	_ = webSpinner.Stop()
	// }

	// prometheus metrics
	if viper.GetBool("metrics.enabled") {
		go func() {
			listenHost := net.ParseIP(viper.GetString("metrics.host"))
			listenPort := viper.GetUint("metrics.port")
			listenAddress := net.JoinHostPort(listenHost.String(), strconv.Itoa(int(listenPort)))

			http.Handle("/metrics", promhttp.Handler())

			gbl.Log.Infof("prometheus metrics: http://%s", listenAddress)

			if err := http.ListenAndServe(listenAddress, nil); err != nil { //nolint:gosec
				gbl.Log.Error(err)
			}
		}()
	}

	// loop forever
	select {}
}

func init() { //nolint:gochecknoinits
	rootCmd.AddCommand(liveCmd)

	// main
	liveCmd.Flags().Bool("watch-sales", true, "get sales")
	_ = viper.BindPFlag("sales.enabled", liveCmd.Flags().Lookup("watch-sales"))
	liveCmd.Flags().Bool("watch-listings", false, "get (opensea) listings for own collections")
	_ = viper.BindPFlag("listings.enabled", liveCmd.Flags().Lookup("watch-listings"))

	// websockets server
	liveCmd.Flags().Bool("websockets", false, "enable websockets server")
	_ = viper.BindPFlag("websockets.server.enabled", liveCmd.Flags().Lookup("websockets"))

	liveCmd.Flags().IP("websockets-host", net.IPv4(0, 0, 0, 0), "websockets listen address")
	_ = viper.BindPFlag("websockets.server.host", liveCmd.Flags().Lookup("websockets-host"))
	liveCmd.Flags().Uint16("websockets-port", 42068, "websockets server port")
	_ = viper.BindPFlag("websockets.server.port", liveCmd.Flags().Lookup("websockets-port"))

	// metrics/prometheus
	liveCmd.Flags().Bool("metrics", false, "enable metrics server")
	_ = viper.BindPFlag("metrics.enabled", liveCmd.Flags().Lookup("metrics"))
	liveCmd.Flags().IP("metrics-host", net.IPv4(0, 0, 0, 0), "metrics listen address")
	_ = viper.BindPFlag("metrics.host", liveCmd.Flags().Lookup("metrics-host"))
	liveCmd.Flags().Uint16("metrics-port", 9090, "metrics server port")
	_ = viper.BindPFlag("metrics.port", liveCmd.Flags().Lookup("metrics-port"))

	// notifications
	liveCmd.Flags().Bool("telegram", false, "send telegram notifications")
	_ = viper.BindPFlag("notifications.telegram.enabled", liveCmd.Flags().Lookup("telegram"))

	// no ui
	liveCmd.Flags().Bool("headless", false, "run without terminal output")
	_ = viper.BindPFlag("ui.headless", liveCmd.Flags().Lookup("headless"))

	// lugges
	liveCmd.Flags().Bool("lugges", false, "enable lugges mode")
	_ = viper.BindPFlag("lugges", liveCmd.Flags().Lookup("lugges"))

	// web ui
	liveCmd.Flags().Bool("web-ui", false, "enable web ui")
	_ = viper.BindPFlag("web.enabled", liveCmd.Flags().Lookup("web-ui"))
	liveCmd.Flags().IP("web-ui-host", net.IPv4(0, 0, 0, 0), "web ui listen address")
	_ = viper.BindPFlag("web.host", liveCmd.Flags().Lookup("web-ui-host"))
	liveCmd.Flags().Uint16("web-ui-port", 42069, "web ui port")
	_ = viper.BindPFlag("web.port", liveCmd.Flags().Lookup("web-ui-port"))

	// wallets
	liveCmd.Flags().StringSliceVarP(&ownWallets, "wallets", "w", []string{}, "Own wallet addresses")
	_ = viper.BindPFlag("wallets", liveCmd.Flags().Lookup("wallets"))

	// min value for sales to be shown
	liveCmd.Flags().Float64("min-value", 0.0, "minimum value to show sales")
	_ = viper.BindPFlag("show.min_value", liveCmd.Flags().Lookup("min-value"))

	// what to show
	liveCmd.Flags().Bool("show-mints", false, "Show mints")
	_ = viper.BindPFlag("show.mints", liveCmd.Flags().Lookup("show-mints"))
	liveCmd.Flags().Bool("show-burns", false, "Show burns")
	_ = viper.BindPFlag("show.burns", liveCmd.Flags().Lookup("show-burns"))
	liveCmd.Flags().Bool("show-reburns", false, "Show re-burns")
	_ = viper.BindPFlag("show.reburns", liveCmd.Flags().Lookup("show-reburns"))
	liveCmd.Flags().Bool("show-airdrops", false, "Show airdrops")
	_ = viper.BindPFlag("show.airdrops", liveCmd.Flags().Lookup("show-airdrops"))
	liveCmd.Flags().Bool("show-transfers", false, "Show transfers")
	_ = viper.BindPFlag("show.transfers", liveCmd.Flags().Lookup("show-transfers"))
	liveCmd.Flags().Bool("show-unknown", false, "Show unknown")
	_ = viper.BindPFlag("show.unknown", liveCmd.Flags().Lookup("show-unknown"))

	// worker settings
	viper.SetDefault("server.workers.newHeadHandler", 2)
	viper.SetDefault("server.workers.newLogHandler", 6)
	viper.SetDefault("server.workers.ttxFormatter", 6)
	viper.SetDefault("server.workers.subscription_logs", 2)
	viper.SetDefault("server.workers.listings", 2)
	viper.SetDefault("server.pubsub.listings", 3)
	viper.SetDefault("server.workers.pubsub.listings", 2)

	viper.SetDefault("opensea.auto_list_min_sales", 50000)

	// ticker
	viper.SetDefault("ticker.statsbox", time.Second*93)
	viper.SetDefault("ticker.gasline", time.Second*39)

	viper.SetDefault("stats.enabled", true)
	viper.SetDefault("stats.balances", true)
	viper.SetDefault("stats.lines", 5)
}

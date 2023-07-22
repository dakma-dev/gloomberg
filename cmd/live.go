package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/chawago"
	"github.com/benleb/gloomberg/internal/config"
	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/degendb/degendata"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/provider"
	"github.com/benleb/gloomberg/internal/nemo/token"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/benleb/gloomberg/internal/nemo/wallet"
	"github.com/benleb/gloomberg/internal/nemo/watch"
	"github.com/benleb/gloomberg/internal/nepa"
	"github.com/benleb/gloomberg/internal/opensea"
	"github.com/benleb/gloomberg/internal/pusu"
	seawatcher "github.com/benleb/gloomberg/internal/seawa"
	seawaModels "github.com/benleb/gloomberg/internal/seawa/models"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/ticker"
	"github.com/benleb/gloomberg/internal/trapri"
	"github.com/benleb/gloomberg/internal/utils/wwatcher"
	"github.com/benleb/gloomberg/internal/web"
	"github.com/benleb/gloomberg/internal/ws"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/muesli/termenv"
	"github.com/redis/rueidis"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	termenv.DefaultOutput().ClearScreen()

	// print header
	header := style.GetHeader(internal.GloombergVersion)
	fmt.Println(header)
	gbl.Log.Info(header)

	// global defaults
	viper.Set("http.timeout", 27*time.Second)

	// // show listings for own collections if an opensea api key is set
	// // if viper.IsSet("api_keys.opensea") && !viper.IsSet("listings.enabled") {
	// if apiKey := viper.GetString("api_keys.opensea"); apiKey != "" && !viper.IsSet("listings.enabled") {
	// 	viper.Set("listings.enabled", true)
	// 	gbl.Log.Infof("listings from opensea: %v", viper.GetBool("listings.enabled"))
	// }

	// initialize
	gb.CollectionDB = collections.New()
	gb.OwnWallets = &wallet.Wallets{}
	gb.Watcher = &watch.Watcher{}

	go func() {
		gb.DegenDB = degendb.NewDegenDB()
	}()

	//
	// start central terminal printer
	go func() {
		gbl.Log.Debug("starting terminal printer...")

		printToTerminalChannel := gb.SubscribePrintToTerminal()

		for eventLine := range printToTerminalChannel {
			gbl.Log.Debugf("terminal printer eventLine: %s", eventLine)

			if viper.GetBool("log.debug") {
				debugPrefix := fmt.Sprintf("%d | ", len(printToTerminalChannel))
				eventLine = fmt.Sprint(debugPrefix, eventLine)
			}

			fmt.Println(eventLine)
		}
	}()

	// cleanup for redis db/cache
	// defer gb.Rdb.Close()

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

		// get all node names to be shown as a list of connected nodes
		providers := gb.ProviderPool.GetProviders()
		nodeNames := make([]string, 0)
		for _, n := range providers {
			nodeNames = append(nodeNames, style.BoldStyle.Render(n.Name))
		}

		gb.Pr(fmt.Sprintf("connected to %s providers: %s", style.AlmostWhiteStyle.Render(fmt.Sprint(len(providers))), style.AlmostWhiteStyle.Render(strings.Join(nodeNames, ", "))))
	}

	//
	// queue for everything to print to the console
	terminalPrinterQueue := make(chan string, 256)

	if viper.GetBool("notifications.smart_wallets.enabled") {
		alphaTicker := ticker.NewAlphaScore(gb)
		go alphaTicker.AlphaCallerTicker(gb, time.NewTicker(time.Minute*1))
	}

	// nepa
	queueWsInTokenTransactions := make(chan *totra.TokenTransaction, 256)
	nePa := nepa.NewNePa(gb)

	//
	var seawa *seawatcher.SeaWatcher
	if viper.GetBool("seawatcher.enabled") || viper.GetBool("listings.enabled") {
		var oopenseaAPIKey string
		if key := viper.GetString("api_keys.opensea"); key != "" {
			oopenseaAPIKey = viper.GetString("api_keys.opensea")
		} else if key := viper.GetString("seawatcher.api_key"); key != "" {
			oopenseaAPIKey = viper.GetString("seawatcher.api_key")
		}
		seawa = seawatcher.NewSeaWatcher(oopenseaAPIKey, gb)
	}

	// trapri | ttx printer to process and format the token transactions
	go trapri.TokenTransactionFormatter(gb, seawa)

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
	// collectionsSpinner := style.GetSpinner("setting up collections...")
	// _ = collectionsSpinner.Start()

	// collection from config file
	// collectionsSpinner.Message("setting up config collections...")

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
	// if len(gb.CollectionDB.Collections) > 0 {
	// 	collectionNames := gb.CollectionDB.SortedAndColoredNames()
	// 	collectionsSpinner.StopMessage(fmt.Sprint(style.BoldStyle.Render(fmt.Sprint(len(collectionNames))), " collections from config: ", strings.Join(collectionNames, ", "), "\n"))
	// }

	gb.Pr(fmt.Sprintf("%s collections loaded from config", style.AlmostWhiteStyle.Render(fmt.Sprint(len(gb.CollectionDB.Collections)))))

	// // stop spinner
	// _ = collectionsSpinner.Stop()

	//
	// get own wallets from config file
	if viper.GetBool("sales.enabled") {
		gb.OwnWallets = config.GetOwnWalletsFromConfig(gb.ProviderPool)

		if len(*gb.OwnWallets) > 0 {
			// miwSpinner.StopMessage(fmt.Sprint(fmt.Sprint(style.BoldStyle.Render(fmt.Sprint(len(wwatcher.MIWC.WeightedMIWs))), " MIWs loaded", "\n")))
			// _ = miwSpinner.Stop()
			gb.PrMod("wawa", fmt.Sprintf("%s own wallets: %s", style.AlmostWhiteStyle.Render(fmt.Sprint(len(*gb.OwnWallets))), strings.Join(gb.OwnWallets.FormattedNames(), ", ")))
		}
	}

	//
	// initialize collections database
	if viper.GetBool("sales.enabled") {
		// collectionsSpinner := style.GetSpinner("setting up collections...")
		// _ = collectionsSpinner.Start()

		if len(*gb.OwnWallets) > 0 {
			// collections from wallet holdings
			// collectionsSpinner.Message("setting up wallet collections...")

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

			GetWalletTokens(gb)
		}

		// // print collections from config & wallet holdings
		// if len(gb.CollectionDB.Collections) > 0 {
		// 	collectionNames := gb.CollectionDB.SortedAndColoredNames()
		// 	collectionsSpinner.StopMessage(fmt.Sprint(style.BoldStyle.Render(fmt.Sprint(len(collectionNames))), " collections from config & wallets: ", strings.Join(collectionNames, ", "), "\n"))
		// }

		gb.Pr(fmt.Sprintf("%s collections from config & wallets: ", style.AlmostWhiteStyle.Render(fmt.Sprint(len(gb.CollectionDB.Collections)))))

		// _ = collectionsSpinner.Stop()
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
		watcher := config.GetWatchRulesFromConfig(gb)
		gb.Watcher = watcher

		//
		// MIWs
		// miwSpinner := style.GetSpinner("setting up MIWs...")
		// _ = miwSpinner.Start()

		wwatcher.LoadMIWs()

		if len(wwatcher.MIWC.WeightedMIWs) > 0 {
			// miwSpinner.StopMessage(fmt.Sprint(fmt.Sprint(style.BoldStyle.Render(fmt.Sprint(len(wwatcher.MIWC.WeightedMIWs))), " MIWs loaded", "\n")))
			// _ = miwSpinner.Stop()
			gb.Pr(fmt.Sprintf("%s MIWs loaded", style.AlmostWhiteStyle.Render(fmt.Sprint(len(wwatcher.MIWC.WeightedMIWs)))))
		}
		//  else {
		// 	_ = miwSpinner.StopFail()
		// }
	}

	// old printer active until fully migrated
	go func() {
		gbl.Log.Debug("starting OLD terminal printer...")

		for eventLine := range terminalPrinterQueue {
			gbl.Log.Debugf("OLD terminal printer eventLine: %s", eventLine)

			if viper.GetBool("log.debug") {
				debugPrefix := fmt.Sprintf("%d | ", len(terminalPrinterQueue))
				eventLine = fmt.Sprint(debugPrefix, eventLine)
			}

			fmt.Println(eventLine)
		}
	}()

	// slugTicker := time.NewTicker(7 * time.Second)
	// go slugs.SlugWorker(slugTicker, &gb.QueueSlugs, gb.Rueidi)

	//
	// gasline ticker
	var gasTicker *time.Ticker

	if tickerInterval := viper.GetDuration("ticker.gasline"); gb.ProviderPool != nil && gb.ProviderPool.PreferredProviderAvailable() && tickerInterval > 0 {
		// initial startup delay
		time.Sleep(tickerInterval / 5)

		// start gasline ticker
		gasTicker = time.NewTicker(tickerInterval)
		go gloomberg.GasTicker(gb, gasTicker, gb.ProviderPool, terminalPrinterQueue)
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
	gb.Stats = gloomberg.NewStats(gb, gasTicker, gb.OwnWallets, gb.ProviderPool, gb.Rdb)

	if statsInterval := viper.GetDuration("ticker.statsbox"); viper.GetBool("stats.enabled") {
		go gb.Stats.StartTicker(statsInterval, terminalPrinterQueue)
	}

	//
	// subscribe to OpenSea API
	if viper.GetBool("seawatcher.enabled") || viper.GetBool("listings.enabled") {
		go trapri.OpenseaEventsHandler(gb)

		go gb.SendSlugsToServer()
	}

	//
	// subscribe to redis pubsub channel to receive events from gloomberg central
	if viper.GetBool("seawatcher.pubsub") || viper.GetBool("pubsub.listings.subscribe") {
		// subscribe to redis pubsub channel
		go pusu.SubscribeToListingsViaRedis(gb)

		// initially send all our slugs & events to subscribe to
		go gb.SendSlugsToServer()

		// subscribe to redis pubsub mgmt channel to listen for "SendSlugs" events
		go func() {
			err := gb.Rdb.Receive(context.Background(), gb.Rdb.B().Subscribe().Channel(internal.TopicSeaWatcherMgmt).Build(), func(msg rueidis.PubSubMessage) {
				gbl.Log.Debug(fmt.Sprintf("🚇 received msg on %s: %s", msg.Channel, msg.Message))

				var mgmtEvent *seawaModels.MgmtEvent

				if err := json.Unmarshal([]byte(msg.Message), &mgmtEvent); err != nil {
					gbl.Log.Fatal(fmt.Sprintf("❌ error json.Unmarshal: %+v", err))
				}

				if mgmtEvent.Action == seawaModels.SendSlugs {
					gbl.Log.Info(fmt.Sprintf("🚇 SendSlugs received on channel %s", msg.Channel))
					gb.SendSlugsToServer()
				}
			})
			if err != nil {
				gbl.Log.Errorf("❌ error subscribing to redis channels %s: %s", internal.TopicSeaWatcherMgmt, err.Error())

				return
			}
		}()
	}

	//
	// degendata - ranks
	// ❕ placed on the end to have the least interference with the calls
	//    to opensea that fetch the wallet collections at the beginning
	// ❗️ probably doesn't matter anymore since we're using the redis cache now
	//    waiting for more feedback from the community before moving it up
	go func() {
		if err := degendata.LoadOpenseaRanks(gb); err != nil {
			gbl.Log.Errorf("error loading opensea ranks: %v", err)
		}
	}()

	//
	// web ui
	if viper.GetBool("web.enabled") {
		go web.StartWebUI(gb) //nolint:errcheck

		gb.PrMod("web", "web-ui started")
	}

	// // prometheus metrics
	// if viper.GetBool("metrics.enabled") {
	// 	go func() {
	// 		listenHost := net.ParseIP(viper.GetString("metrics.host"))
	// 		listenPort := viper.GetUint("metrics.port")
	// 		listenAddress := net.JoinHostPort(listenHost.String(), strconv.Itoa(int(listenPort)))

	// 		http.Handle("/metrics", promhttp.Handler())

	// 		gbl.Log.Infof("prometheus metrics: http://%s", listenAddress)

	// 		if err := http.ListenAndServe(listenAddress, nil); err != nil { //nolint:gosec
	// 			gbl.Log.Error(err)
	// 		}
	// 	}()
	// }

	// marmot tasks
	// gb.CreatePeriodicTask("testing", 5*time.Second, func(gb *gloomberg.Gloomberg) {
	// 	log.Printf("testing tasks lol! %+v", len(gb.Ranks))
	// })
	// gb.CreateScheduledTask("testing", time.Now().Add(17*time.Second), func(gb *gloomberg.Gloomberg) {
	// 	log.Printf("testing scheduled tasks lol! %+v", len(gb.Ranks))
	// })

	go func() {
		wawa := chawago.NewWalletWatcher(gb)
		wawa.Watch()

		gb.Prf("wallet watcher started: %+v", wawa)
	}()

	gb.Prf("starting grpc client...")
	go testGRPC()

	// loop forever
	select {}
}

func testGRPC() {
	var opts []grpc.DialOption

	// tls := true

	// if tls {
	// 	// if caFile == "" {
	// 	// 	caFile = "x509/ca_cert.pem"
	// 	// }

	// 	creds := credentials.NewClientTLSFromCert(nil, "")

	// 	opts = append(opts, grpc.WithTransportCredentials(creds))
	// } else {
	// 	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// }

	if creds := gloomberg.GetTLSClientCredentials(); creds != nil {
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	grpcAddress := fmt.Sprintf("%s:%d", viper.GetString("seawatcher.grpc.server"), viper.GetUint("seawatcher.grpc.port"))

	gb.Prf("connecting to gRPC %s...", style.BoldAlmostWhite(grpcAddress))

	conn, err := grpc.Dial(grpcAddress, opts...)
	if err != nil {
		log.Errorf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := seawatcher.NewSeaWatcherClient(conn)

	gb.Prf("subscribing via grpc to: %s", style.BoldAlmostWhite(degendb.Listing.OpenseaEventName()))

	subsriptionRequest := &seawatcher.SubscriptionRequest{EventTypes: []seawatcher.EventType{seawatcher.EventType_ITEM_LISTED}, Collections: gb.CollectionDB.OpenseaSlugs()} //nolint:nosnakecase
	stream, err := client.GetEvents(context.Background(), subsriptionRequest)
	if err != nil {
		log.Errorf("client.GetEvents failed: %v", err)

		return
	}

	for {
		event, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.ErrUnexpectedEOF) || errors.Is(err, io.EOF) {
				log.Errorf("io.EOF error: %v", err)

				break
			}

			log.Errorf("receiving event failed: %v", err)

			time.Sleep(time.Second * 1)
		}

		gb.Prf("🐔 RECEIVED: %+v", event)
	}
}

var degendataPath string

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

	// min value for sales to be shown (single item price)
	liveCmd.Flags().Float64("min-value", 0.0, "minimum value to show sales")
	_ = viper.BindPFlag("show.min_value", liveCmd.Flags().Lookup("min-value"))
	// multiplier for min_value applied to the total price of a tx
	viper.SetDefault("show.min_value_multiplier", 2.0)

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

	// degendb
	liveCmd.Flags().StringVar(&degendataPath, "degendata", "degendata", "path to degendata repo")
	_ = viper.BindPFlag("degendata.path", liveCmd.Flags().Lookup("degendata"))

	// worker settings
	viper.SetDefault("trapri.numOpenSeaEventhandlers", 3)

	viper.SetDefault("gloomberg.eventhub.numHandler", 3)
	viper.SetDefault("gloomberg.eventhub.inQueuesSize", 256)
	viper.SetDefault("gloomberg.eventhub.outQueuesSize", 32)

	viper.SetDefault("etherscan.fetchInterval", time.Second*3)

	// OLD worker settings OLD
	viper.SetDefault("server.workers.newHeadHandler", 2)
	viper.SetDefault("server.workers.newLogHandler", 6)
	viper.SetDefault("server.workers.ttxFormatter", 6)
	viper.SetDefault("server.workers.subscription_logs", 2)
	viper.SetDefault("server.workers.listings", 2)
	viper.SetDefault("server.pubsub.listings", 3)
	viper.SetDefault("server.workers.pubsub.listings", 2)

	// opensea settings
	viper.SetDefault("seawatcher.auto_subscribe_after_sales", 37)

	//
	// timeframes

	// sali.default_timeframe is used for the sales/listings counts shown
	viper.SetDefault("salira.default_timeframe", time.Minute*137)
	viper.SetDefault("salira.timeframes", []time.Duration{
		// time.Minute * 1,
		// time.Minute * 3,
		time.Minute * 13,
		time.Minute * 37,
		time.Minute * 137,
	})

	// ticker
	viper.SetDefault("ticker.statsbox", internal.BlockTime*9)
	viper.SetDefault("ticker.gasline", internal.BlockTime*3)

	// stats settings
	viper.SetDefault("stats.enabled", true)
	viper.SetDefault("stats.balances", true)
	viper.SetDefault("stats.timeframe", time.Minute*13) // 13
	viper.SetDefault("stats.lines", 6)
	// high volume mints detection
	viper.SetDefault("stats.high_volume.check_interval", time.Second*17)
	viper.SetDefault("stats.high_volume.min_checks_below_threshold", 2)
	viper.SetDefault("stats.high_volume.mints.threshold", 37)
}

func GetWalletTokens(gb *gloomberg.Gloomberg) map[common.Address]*token.Token {
	gbTokens := make([]*token.Token, 0)

	for _, w := range *gb.OwnWallets {
		tokensForWallet := opensea.GetTokensFor(w.Address, 2, "")
		gbTokens = append(gbTokens, tokensForWallet...)

		log.Debugf("Wallet %s has %d tokens: %+v", w.Address.String(), len(tokensForWallet), tokensForWallet)

		tokenMapForWallet := make(map[common.Address]map[string]*token.Token)
		for _, t := range tokensForWallet {
			if _, ok := tokenMapForWallet[t.Address]; !ok {
				tokenMapForWallet[t.Address] = make(map[string]*token.Token)
			}

			tokenMapForWallet[t.Address][t.ID.String()] = t
		}

		w.Tokens = tokenMapForWallet

		// honor the rate limit
		time.Sleep(time.Millisecond * 337)
	}

	gb.PrMod("wawa", fmt.Sprintf("found %s tokens in our %s wallets", style.AlmostWhiteStyle.Render(fmt.Sprint(len(gbTokens))), style.AlmostWhiteStyle.Render(fmt.Sprint(len(*gb.OwnWallets)))))

	// create map
	gbTokensMap := make(map[common.Address]*token.Token)
	for _, t := range gbTokens {
		gbTokensMap[t.Address] = t
	}

	return gbTokensMap
}

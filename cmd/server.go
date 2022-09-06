package cmd

import (
	"context"
	"fmt"
	"math"
	"net"
	"time"

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/server"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/subscriptions"
	"github.com/benleb/gloomberg/internal/wwatcher"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serverCmd represents the server command.
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: gbServer, // func(cmd *cobra.Command, args []string) { server() },
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// stats
	viper.SetDefault("stats.enabled", true)
	viper.SetDefault("stats.extended", false)
	viper.SetDefault("stats.gas", true)
	viper.SetDefault("stats.balances", true)
	viper.SetDefault("stats.lines", 5)
	viper.SetDefault("stats.interval", time.Second*90)

	// worker settings
	viper.SetDefault("workers.log_handler", 7)
	viper.SetDefault("workers.listings_handler", 4)
	viper.SetDefault("workers.output", 10)

	// number of retries to resolve an ens name to an address or vice versa
	viper.SetDefault("ens.resolve_max_retries", 5)

	// ipfs gateway to fetch metadata/images
	viper.SetDefault("ipfs.gateway", "https://ipfs.io/ipfs/")

	// default collections and wallet
	viper.SetDefault("collections", []any{map[string]string{"name": "OSF's Red Lite District", "mark": "#6A0F27", "address": "0x513cd71defc801b9c1aa763db47b5df223da77a2"}})
	viper.SetDefault("wallets", []string{"pranksy.eth"})
	viper.SetDefault("wwatcher", []any{map[string]string{}})

	// show desktop notifications
	serverCmd.Flags().Bool("notifications", false, "Show notifications?")
	_ = viper.BindPFlag("show.notifications", serverCmd.Flags().Lookup("notifications"))

	// types of events to show
	serverCmd.Flags().Bool("sales", true, "Show sales?")
	_ = viper.BindPFlag("show.sales", serverCmd.Flags().Lookup("sales"))
	serverCmd.Flags().Bool("mints", false, "Show mints?")
	_ = viper.BindPFlag("show.mints", serverCmd.Flags().Lookup("mints"))
	serverCmd.Flags().Bool("listings", false, "Show listings?")
	_ = viper.BindPFlag("show.listings", serverCmd.Flags().Lookup("listings"))
	serverCmd.Flags().Bool("transfers", false, "Show transfers?")
	_ = viper.BindPFlag("show.transfers", serverCmd.Flags().Lookup("transfers"))
	serverCmd.Flags().Float64("min-price", 0.0, "Minimum price to show sales?")
	_ = viper.BindPFlag("show.min_price", serverCmd.Flags().Lookup("min-price"))

	// process & show just our own collections (from config/wallet)
	serverCmd.Flags().Bool("own", false, "Show only own collections (from config/wallet)")
	_ = viper.BindPFlag("show.own", serverCmd.Flags().Lookup("own"))

	// websockets server
	serverCmd.Flags().Bool("server", false, "Start websockets server")
	_ = viper.BindPFlag("server.enabled", serverCmd.Flags().Lookup("server"))
	serverCmd.Flags().IP("host", net.IPv4(0, 0, 0, 0), "Websockets server port")
	_ = viper.BindPFlag("server.host", serverCmd.Flags().Lookup("host"))
	serverCmd.Flags().Uint16("port", 42069, "Websockets server port")
	_ = viper.BindPFlag("server.port", serverCmd.Flags().Lookup("port"))

	// telegram bot
	serverCmd.Flags().Bool("telegram", false, "Start telegram bot")
	_ = viper.BindPFlag("telegram.enabled", serverCmd.Flags().Lookup("telegram"))
}

func gbServer(_ *cobra.Command, _ []string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*4200*10)
	defer cancel()

	gbl.GetSugaredLogger()

	// wallets *models.Wallets

	ownCollections := collections.New()

	//
	// config
	//
	if viper.GetString("server.host") != net.IPv4(0, 0, 0, 0).String() || viper.GetUint("server.port") != 42069 {
		viper.Set("server.enabled", true)
	}

	if viper.GetFloat64("show.min_price") > 0.0 {
		numGlickerLines := math.Max(float64(viper.GetInt("stats.lines")), 4.0)
		viper.Set("stats.lines", numGlickerLines)
	}

	// show listings for own collections if an opensea api key is set
	if viper.IsSet("api_keys.opensea") && !viper.IsSet("show.listings") {
		viper.Set("show.listings", true)
	}

	// wallet watcher
	wwatcher.InitWatcher()

	// write settings to log
	gbl.Log.Debug(viper.AllSettings())

	// print header
	header := style.GetHeader(Version)
	fmt.Println(header)
	gbl.Log.Info(header)

	//
	// initialize node connections
	//
	nodes := getNodes()

	// if we subscribe to all chain-events, we can do it now
	if !viper.GetBool("show.own") {
		for _, node := range *nodes {
			queue := make(chan types.Log, 1024)
			logQueues[node.NodeID] = &queue
			nodes.SubscribeToAllTransfers(ctx, *logQueues[node.NodeID])
		}
	}

	//
	// MIWs
	//
	miwSpinner := style.GetSpinner("setting up MIWs...")
	_ = miwSpinner.Start()

	wwatcher.LoadMIWs()

	if len(wwatcher.MIWC.WeightedMIWs) > 0 {
		miwSpinner.StopMessage(fmt.Sprint(fmt.Sprint(style.BoldStyle.Render(fmt.Sprint(len(wwatcher.MIWC.WeightedMIWs))), " MIWs loaded", "\n")))
		_ = miwSpinner.Stop()
	} else {
		_ = miwSpinner.StopFail()
	}

	//
	// set up workers print to process events from the chain
	//

	// processes logs from the ethereum chain from our nodes
	for _, node := range *nodes {
		for workerID := 1; workerID <= viper.GetInt("workers.log_handler"); workerID++ {
			go subscriptions.SubscriptionLogsHandler(ctx, node, nodes, ownCollections, logQueues[node.NodeID], queueEvents, queueWS)
		}
	}

	// websockets server
	if viper.GetBool("server.enabled") {
		viper.Set("mode.server", true)

		go server.StartWebsocketServer(&queueWS)
	}

	select {}
}

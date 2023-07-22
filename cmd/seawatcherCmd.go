package cmd

import (
	"github.com/benleb/gloomberg/internal/seawa"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// seaWatcherCmd represents the seawa command.
var seaWatcherCmd = &cobra.Command{
	Use:     "seawatcher",
	Aliases: []string{"seawa"},
	Short:   "receives events from the OpenSea API and pushes them to the redis database",

	Run: runSeawatcher,
}

func init() {
	rootCmd.AddCommand(seaWatcherCmd)

	// grpc
	seaWatcherCmd.Flags().Uint16("grpc-port", 31337, "gRPC server port")
	_ = viper.BindPFlag("seawatcher.grpc.port", seaWatcherCmd.Flags().Lookup("grpc-port"))
	// grpc server
	seaWatcherCmd.Flags().IPVar(&grpcServer, "grpc-listen", nil, "gRPC server listen address")
	_ = viper.BindPFlag("grpc.listen", seaWatcherCmd.Flags().Lookup("grpc-listen"))
	// grpc client
	seaWatcherCmd.Flags().IPVar(&grpcListen, "grpc", nil, "server gRPC client connects to")
	_ = viper.BindPFlag("grpc.server", seaWatcherCmd.Flags().Lookup("grpc"))

	// // metrics/prometheus
	// viper.SetDefault("metrics.enabled", false)
	// viper.SetDefault("metrics.host", net.IPv4(0, 0, 0, 0))
	// viper.SetDefault("metrics.port", 9090)
}

func runSeawatcher(cmd *cobra.Command, _ []string) {
	// find api key
	var apiKey string
	switch {
	case viper.GetString("api_keys.opensea") != "":
		apiKey = viper.GetString("api_keys.opensea")
	case viper.GetString("seawatcher.api_key") != "":
		apiKey = viper.GetString("seawatcher.api_key")
	default:
		log.Fatal("no api key found")
	}

	// start sea watcher & loop forever
	sw := seawa.NewSeaWatcher(apiKey, gb)

	sw.Pr("⚓️ starting seawatcher")

	if viper.GetBool("seawatcher.pubsub") {
		go sw.SubscribeToPubsubMgmt()
		sw.Pr("subscribed to mgmt channel…")

		// publish a "SendSlugs" event to the management channel to request the slugs/events to subscribe to from the clients
		sw.PublishSendSlugs()
		sw.Pr("requested slugs from clients…")
	}

	sw.Prf("seaWatcherCmd.CalledAs(): %+v", cmd.CalledAs())

	select {}
}

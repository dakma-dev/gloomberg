package cmd

import (
	"net"

	"github.com/benleb/gloomberg/internal/seawa"
	"github.com/benleb/gloomberg/internal/web"
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

	// run manifold sns receiver endpoint
	viper.SetDefault("manifoldSNS.enabled", false)
	viper.SetDefault("manifoldSNS.host", net.IPv4(127, 0, 0, 1))
	viper.SetDefault("manifoldSNS.port", viper.GetUint16("web.port")-1)

	// grpc
	seaWatcherCmd.Flags().Uint16("grpc-port", 31337, "gRPC server port")
	_ = viper.BindPFlag("seawatcher.grpc.port", seaWatcherCmd.Flags().Lookup("grpc-port"))
	// grpc server
	seaWatcherCmd.Flags().IPVar(&grpcServerListen, "grpc-listen", nil, "gRPC server listen address")
	_ = viper.BindPFlag("grpc.listen", seaWatcherCmd.Flags().Lookup("grpc-listen"))
	// grpc client
	seaWatcherCmd.Flags().IPVar(&grpcClientHost, "grpc", nil, "server gRPC client connects to")
	_ = viper.BindPFlag("seawatcher.grpc.client.host", seaWatcherCmd.Flags().Lookup("grpc"))
}

func runSeawatcher(_ *cobra.Command, _ []string) {
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

	sw.Pr("⚓️ seawatcher started... 🌊 👀")

	if viper.GetBool("seawatcher.pubsub") {
		go sw.SubscribeToPubsubMgmt()
		sw.Pr("subscribed to mgmt channel…")

		// publish a "SendSlugs" event to the management channel to request the slugs/events to subscribe to from the clients
		go sw.PublishSendSlugs()
		sw.Pr("requested slugs from clients…")
	}

	// sw.Prf("seaWatcherCmd.CalledAs(): %+v", cmd.CalledAs())

	//
	// manifold SNS receiver
	if viper.GetBool("manifoldSNS.enabled") {
		go web.StartmanifoldSNS(gb)
	}

	select {}
}

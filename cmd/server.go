package cmd

import (
	"net"

	"github.com/benleb/gloomberg/internal/server/ws"

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/server"
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
	Run: LFG, // func(cmd *cobra.Command, args []string) { server() },
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(serverCmd)

	// websockets server
	serverCmd.Flags().IP("ws-host", net.IPv4(0, 0, 0, 0), "websockets listen address")
	serverCmd.Flags().Uint16("ws-port", 42069, "websockets server port")
	_ = viper.BindPFlag("server.websockets.host", serverCmd.Flags().Lookup("ws-host"))
	_ = viper.BindPFlag("server.websockets.port", serverCmd.Flags().Lookup("ws-port"))

	// worker settings
	viper.SetDefault("server.workers.subscription_logs", 5)

	// // number of retries to resolve an ens name to an address or vice versa
	// viper.SetDefault("ens.resolve_max_retries", 5)

	// ipfs gateway to fetch metadata/images
	viper.SetDefault("ipfs.gateway", "https://ipfs.io/ipfs/")
}

func LFG(_ *cobra.Command, _ []string) {
	// create a new server instance
	cServer := server.New()

	// create a queue/channel for the received & parsed events
	queueEvents := make(chan *collections.Event, 1024)
	// create a queue/channel for events to be sent out via ws
	queueOutWS := make(chan *collections.Event, 1024)

	// subscribe to the chain logs/events and start the workers
	cServer.Subscribe(&queueEvents)

	// websockets
	wsServer := ws.New(viper.GetString("server.websockets.host"), viper.GetUint("server.websockets.port"), &queueOutWS)
	go wsServer.Start()
}

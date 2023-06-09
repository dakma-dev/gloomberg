package cmd

import (
	"net"

	"github.com/benleb/gloomberg/internal"
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

	// metrics/prometheus
	viper.SetDefault("metrics.enabled", false)
	viper.SetDefault("metrics.host", net.IPv4(0, 0, 0, 0))
	viper.SetDefault("metrics.port", 9090)
}

func runSeawatcher(_ *cobra.Command, _ []string) {
	log.Infof("⚓️ starting seawatcher %s…", internal.GloombergVersion)

	startOpenseaSubscription()

	select {}
}

func startOpenseaSubscription() *seawa.SeaWatcher {
	// start sea watcher & loop forever
	seaWatcher := seawa.NewSeaWatcher(viper.GetString("api_keys.opensea"), gb)

	go seaWatcher.Run()

	// publish a "SendSlugs" event to the management channel to request the slugs/events to subscribe to from the clients
	seaWatcher.PublishSendSlugs()

	return seaWatcher
}

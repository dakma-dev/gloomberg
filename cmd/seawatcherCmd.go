package cmd

import (
	"net"
	"net/http"
	"strconv"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/seawa"
	"github.com/charmbracelet/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// seaWatcherCmd represents the seawa command.
var seaWatcherCmd = &cobra.Command{
	Use:     "seawatcher",
	Aliases: []string{"seawa"},
	Short:   "receives events from the OpenSea API and pushes them to the redis database",

	Run: run,
}

func init() {
	rootCmd.AddCommand(seaWatcherCmd)

	// metrics/prometheus
	viper.SetDefault("metrics.enabled", false)
	viper.SetDefault("metrics.host", net.IPv4(0, 0, 0, 0))
	viper.SetDefault("metrics.port", 9090)
}

func run(_ *cobra.Command, _ []string) {
	log.Infof("⚓️ starting seawatcher %s…", internal.GloombergVersion)

	//
	// init metrics
	if viper.GetBool("metrics.enabled") {
		go func() {
			listenHost := net.ParseIP(viper.GetString("metrics.host"))
			listenPort := viper.GetUint("metrics.port")
			listenAddress := net.JoinHostPort(listenHost.String(), strconv.Itoa(int(listenPort)))

			http.Handle("/metrics", promhttp.Handler())

			log.Infof("⚓️ 📐 metrics: http://%s", listenAddress)

			if err := http.ListenAndServe(listenAddress, nil); err != nil { //nolint:gosec
				log.Errorf("⚓️ 📐 ❌ error starting metrics server: %s", err)
			}
		}()
	}

	// start sea watcher & loop forever
	seaWatcher := seawa.NewSeaWatcher(viper.GetString("api_keys.opensea"), GetRedisClient())

	go seaWatcher.Run()

	// publish a "SendSlugs" event to the management channel to request the slugs/events to subscribe to from the clients
	seaWatcher.PublishSendSlugs()

	select {}
}

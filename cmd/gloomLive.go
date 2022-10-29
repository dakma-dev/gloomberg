package cmd

import (
	"net"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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
	_ = viper.BindPFlag("telegram.enabled", liveCmd.Flags().Lookup("telegram"))

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
	viper.SetDefault("ticker.gasline", time.Second*29)

	viper.SetDefault("stats.enabled", true)
	viper.SetDefault("stats.interval", time.Second*90)
	viper.SetDefault("stats.balances", true)
	viper.SetDefault("stats.lines", 5)
}

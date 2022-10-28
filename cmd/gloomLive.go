package cmd

import (
	"net"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// gloomLiveCmd represents the gloomLive command
var gloomLiveCmd = &cobra.Command{
	Use:   "gloomLive",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runGloomberg(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(gloomLiveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gloomLiveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gloomLiveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// client := gloomberg.RoleMap{
	// 	GloomClient:           true,

	// 	OwnWalletWatcher:      true,
	// 	StatsTicker:           true,
	// 	TelegramBot:           false,
	// 	TelegramNotifications: false,
	// 	WalletWatcher:         true,

	// }

	// main
	gloomLiveCmd.Flags().Bool("sales", true, "get sales")
	_ = viper.BindPFlag("sales.enabled", gloomLiveCmd.Flags().Lookup("sales"))
	gloomLiveCmd.Flags().Bool("listings", false, "get (opensea) listings for own collections")
	_ = viper.BindPFlag("listings.enabled", gloomLiveCmd.Flags().Lookup("listings"))
	gloomLiveCmd.Flags().Bool("websockets", false, "enable websockets server")
	_ = viper.BindPFlag("server.websockets.enabled", gloomLiveCmd.Flags().Lookup("websockets"))

	// websockets options
	gloomLiveCmd.Flags().IP("ws-host", net.IPv4(0, 0, 0, 0), "websockets listen address")
	_ = viper.BindPFlag("server.websockets.host", gloomLiveCmd.Flags().Lookup("ws-host"))
	gloomLiveCmd.Flags().Uint16("ws-port", 42068, "websockets server port")
	_ = viper.BindPFlag("server.websockets.port", gloomLiveCmd.Flags().Lookup("ws-port"))

	// notifications
	gloomLiveCmd.Flags().Bool("telegram", false, "send telegram notifications")
	_ = viper.BindPFlag("telegram.enabled", gloomLiveCmd.Flags().Lookup("telegram"))

	// ui
	gloomLiveCmd.Flags().Bool("headless", false, "run without terminal output")
	_ = viper.BindPFlag("ui.headless", gloomLiveCmd.Flags().Lookup("headless"))
	gloomLiveCmd.Flags().Bool("webui", false, "enable web ui")
	_ = viper.BindPFlag("ui.webui", gloomLiveCmd.Flags().Lookup("webui"))

	// min value for sales to be shown
	gloomLiveCmd.Flags().Float64("min-value", 0.0, "minimum value to show sales?")
	_ = viper.BindPFlag("show.min_value", gloomLiveCmd.Flags().Lookup("min-value"))

	// what to show
	gloomLiveCmd.Flags().Bool("show-mints", false, "Show mints?")
	_ = viper.BindPFlag("show.mints", gloomLiveCmd.Flags().Lookup("show-mints"))
	gloomLiveCmd.Flags().Bool("show-transfers", false, "Show transfers?")
	_ = viper.BindPFlag("show.transfers", gloomLiveCmd.Flags().Lookup("show-transfers"))
	// gloomLiveCmd.Flags().Bool("sales", true, "Show sales?")
	// _ = viper.BindPFlag("show.sales", gloomLiveCmd.Flags().Lookup("sales"))
	// gloomLiveCmd.Flags().Bool("listings", false, "Show listings?")
	// _ = viper.BindPFlag("show.listings", gloomLiveCmd.Flags().Lookup("listings"))

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

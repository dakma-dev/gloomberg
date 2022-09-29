package cmd

import (
	"fmt"
	"time"

	"github.com/benleb/gloomberg/internal"
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
		fmt.Println("gloomLive called")

		// default roles for gloomClient
		role := internal.RoleMap{
			ChainWatcher:          true,
			OsStreamWatcher:       true,
			OutputTerminal:        true,
			OwnCollections:        true,
			OwnWalletWatcher:      true,
			StatsTicker:           true,
			TelegramBot:           false,
			TelegramNotifications: false,
			WalletWatcher:         true,
			WsServer:              false,
		}

		gloomberg(cmd, args, role)
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

	gloomLiveCmd.Flags().Float64("min-value", 0.0, "minimum value to show sales?")
	_ = viper.BindPFlag("show.min_value", gloomLiveCmd.Flags().Lookup("min-value"))

	gloomLiveCmd.Flags().Bool("mints", false, "Show mints?")
	_ = viper.BindPFlag("show.mints", gloomLiveCmd.Flags().Lookup("mints"))

	// notifications
	gloomLiveCmd.Flags().Bool("telegram", false, "Send notifications to telegram?")
	_ = viper.BindPFlag("notifications.telegram.enabled", gloomLiveCmd.Flags().Lookup("telegram"))

	// worker settings
	viper.SetDefault("server.workers.subscription_logs", 5)
	viper.SetDefault("server.workers.output", 3)

	viper.SetDefault("workers.listings", 2)

	viper.SetDefault("opensea.auto_list_min_sales", 50000)

	// ticker
	viper.SetDefault("ticker.statsbox", time.Second*89)
	viper.SetDefault("ticker.gasline", time.Second*17)
	viper.SetDefault("stats.enabled", true)
	viper.SetDefault("stats.interval", time.Second*90)
	viper.SetDefault("stats.balances", true)
	viper.SetDefault("stats.lines", 5)
}

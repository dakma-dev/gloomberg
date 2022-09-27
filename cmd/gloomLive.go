package cmd

import (
	"fmt"

	"github.com/benleb/gloomberg/internal"
	"github.com/spf13/cobra"
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
			OsStreamWatcher:       false,
			OutputTerminal:        true,
			OwnWalletWatcher:      true,
			StatsTicker:           true,
			TelegramBot:           false,
			TelegramNotifications: false,
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
}

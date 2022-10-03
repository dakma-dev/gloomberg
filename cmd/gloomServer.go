package cmd

import (
	"fmt"

	"github.com/benleb/gloomberg/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// gloomServerCmd represents the gloomServer command
var gloomServerCmd = &cobra.Command{
	Use:   "gloomServer",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gloomServer called")

		// default roles for gloomClient
		role := internal.RoleMap{
			ChainWatcher:          true,
			OsStreamWatcher:       false,
			WsServer:              false,
			TelegramBot:           false,
			TelegramNotifications: false,
			OutputTerminal:        false,
		}

		gloomberg(cmd, args, role)
	},
}

func init() {
	rootCmd.AddCommand(gloomServerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gloomServerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gloomServerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// show telegram notifications
	gloomServerCmd.Flags().Bool("telegram", false, "Send notifications to telegram?")
	_ = viper.BindPFlag("notifications.telegram", gloomServerCmd.Flags().Lookup("telegram"))

	// // websockets server
	// gloomServerCmd.Flags().Bool("ws", false, "enable websockets server")
	// gloomServerCmd.Flags().IP("ws-host", net.IPv4(0, 0, 0, 0), "websockets listen address")
	// gloomServerCmd.Flags().Uint16("ws-port", 42069, "websockets server port")
	// _ = viper.BindPFlag("server.websockets.enabled", gloomServerCmd.Flags().Lookup("ws"))
	// _ = viper.BindPFlag("server.websockets.host", gloomServerCmd.Flags().Lookup("ws-host"))
	// _ = viper.BindPFlag("server.websockets.port", gloomServerCmd.Flags().Lookup("ws-port"))
}
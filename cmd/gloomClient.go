package cmd

// import (
// 	"github.com/benleb/gloomberg/internal/models/gloomberg"
// 	"github.com/spf13/cobra"
// )

// // gloomClientCmd represents the gloomClient command
// var gloomClientCmd = &cobra.Command{
// 	Use:   "gloomClient",
// 	Short: "A brief description of your command",
// 	Long: `A longer description that spans multiple lines and likely contains examples
// and usage of using your command. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		// default roles for gloomberg live
// 		role := gloomberg.RoleMap{
// 			ChainWatcher:          false,
// 			GloomClient:           true,
// 			OsStreamWatcher:       true,
// 			OutputTerminal:        true,
// 			CollectionDB:          true,
// 			OwnWalletWatcher:      true,
// 			StatsTicker:           true,
// 			TelegramBot:           false,
// 			TelegramNotifications: false,
// 			WalletWatcher:         true,
// 			WsServer:              false,
// 		}

// 		runGloomberg(cmd, args, role)
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(gloomClientCmd)

// 	// Here you will define your flags and configuration settings.

// 	// Cobra supports Persistent Flags which will work for this command
// 	// and all subcommands, e.g.:
// 	// gloomClientCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	// Cobra supports local flags which will only run when this command
// 	// is called directly, e.g.:
// 	// gloomClientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }

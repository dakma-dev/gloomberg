package cmd

import (
	"fmt"

	"github.com/benleb/gloomberg/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// gloomClientCmd represents the gloomClient command
var gloomClientCmd = &cobra.Command{
	Use:   "gloomClient",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gloomClient called")

		// default roles for gloomClient
		role := internal.RoleMap{}

		// default role settings
		role.OutputTerminal = true

		// opensea stream api to subscribe to new listings for own collections
		if openseaToken := viper.GetString("api_keys.opensea"); openseaToken != "" {
			role.OsStreamWatcher = true
		}

		gloomberg(cmd, args, role)
	},
}

func init() {
	rootCmd.AddCommand(gloomClientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gloomClientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gloomClientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

package cmd

import (
	"fmt"
	"log"

	"github.com/benleb/gloomberg/internal/analytics"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setOwnerCmd represents the setOwner command.
var setOwnerCmd = &cobra.Command{
	Use:   "setOwner",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: setOwner,
}

func init() {
	rootCmd.AddCommand(setOwnerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setOwnerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setOwnerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	setOwnerCmd.Flags().StringVar(&apiKeyMoralis, "moralis", "", "Moralis API Key")
	_ = viper.BindPFlag("api_keys.moralis", setOwnerCmd.Flags().Lookup("moralis"))
}

func setOwner(cmd *cobra.Command, args []string) {
	if !viper.IsSet("api_keys.moralis") {
		log.Fatal("api_keys.moralis not set")
	}

	// print header
	header := style.GetHeader(Version)
	fmt.Println(header)
	gbl.Log.Info(header)

	analytics.GetSetOwner(cmd, args)
}

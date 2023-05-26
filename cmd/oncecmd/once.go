/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package oncecmd

import (
	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/provider"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var lg = internal.BaseLogger

func init() {
	// what to show
	OnceCmd.Flags().Bool("fresh", false, "dont use cached data")
	_ = viper.BindPFlag("fresh", OnceCmd.Flags().Lookup("fresh"))
}

// OnceCmd represents the once command
var OnceCmd = &cobra.Command{
	Use:   "once",
	Short: "This command is intended to quickly implement a one-time task or similar",
	Long:  `A One-time tasks or experiment can easily implemented based the foundation glooomberg provides. Just create a new .go file in the once/ directory for your task or add it to a already available one and call it from this command.`,
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	//
	// initialize some basic resources
	//

	// pool of providers
	pool, err := provider.FromConfig(viper.Get("provider"))
	if err != nil {
		gbl.Log.Fatal("❌ running provider failed, exiting")
	}

	// ethclient.Client
	client := pool.GetProviders()[0].Client

	//
	// if you need other resources besides the ethclient.Client, feel free to initiate them here
	//

	//
	//
	// call your task/experiment/whatever from here on
	//
	//

	// lawless metadata: get the lawless on-chain metadata and save it to a json file
	analyzeLawlessTokenNames(client)

	// TestOwnScorer()
}

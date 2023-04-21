package flotscmd

import (
	"fmt"

	"github.com/benleb/gloomberg/internal"
	"github.com/spf13/cobra"
)

// FlotsCmd represents the flots command
var FlotsCmd = &cobra.Command{
	Use: "flots",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("flots called")
	},
}

var (
	flagRawTransactions []string
	flagBBundleHash     string

	flagPlusBlocks int64

	lo = internal.BasePrinter // .WithPrefix("flots 🤖")
)

func init() {
	// Here you will define your flags and configuration settings.

	// FlotsCmd.PersistentFlags().StringSliceVarP(&flagRawTransactions, "rawtxs", "t", make([]string, 0), "signed transactions")

	// FlotsCmd.PersistentFlags().Int64Var(&flagPlusBlocks, "plusBlocks", 5, "blocks to add to the current block number")
	// _ = viper.BindPFlag("flots.plusBlocks", FlotsCmd.PersistentFlags().Lookup("plusBlocks"))
}

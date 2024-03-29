package flotscmd

import (
	"fmt"

	"github.com/benleb/gloomberg/internal/flots"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// sendBundleCmd represents the callSendBundle command.
var callBundleCmd = &cobra.Command{
	Use:   "callBundle",
	Short: "simulates a bundle of given raw transactions",
	Run: func(cmd *cobra.Command, args []string) {
		// flashbots client
		flots := flots.New()

		//
		// parse raw txs
		if flagRawTransactions == nil {
			log.Fatal("❌ no raw transactions provided!")
		}

		fmt.Println("transactions:")

		rawTxs := make([][]byte, 0)
		for idx, rawtx := range flagRawTransactions {
			rawTxs = append(rawTxs, hexutil.MustDecode(rawtx))
			fmt.Printf("  tx %d: %+v", idx, rawtx)
		}

		//
		// simulate with call bundle
		callBundle := flots.CallBundle(rawTxs)
		fmt.Printf("🟢 call bundle: %+v\n\n\n", callBundle)
	},
}

func init() {
	FlotsCmd.AddCommand(callBundleCmd)

	callBundleCmd.Flags().StringSliceVarP(&flagRawTransactions, "rawtxs", "t", make([]string, 0), "signed transactions (get them from https://flashbots-bundler.surge.sh/rpc for example)")
	_ = callBundleCmd.MarkFlagRequired("rawtxs")

	callBundleCmd.Flags().Int64Var(&flagPlusBlocks, "plusBlocks", 5, "blocks to add to the current block number")
	_ = viper.BindPFlag("flots.plusBlocks", callBundleCmd.Flags().Lookup("plusBlocks"))
}

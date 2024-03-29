package flotscmd

import (
	"fmt"
	"os"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/flots"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// sendBundleCmd represents the callSendBundle command.
var sendBundleCmd = &cobra.Command{
	Use:   "sendBundle",
	Short: "simulates a bundle of given raw transactions and, if successful, sends it to the flashbots network",
	Run: func(cmd *cobra.Command, args []string) {
		// flashbots client
		flots := flots.New()

		//
		// parse raw txs
		if flagRawTransactions == nil {
			log.Fatal("❌ no raw transactions provided!")
		}

		fmt.Println("")
		fmt.Println("transactions:")

		rawTxs := make([][]byte, 0)
		for idx, rawtx := range flagRawTransactions {
			rawTxs = append(rawTxs, hexutil.MustDecode(rawtx))
			fmt.Printf("  tx %d: %+v\n", idx, rawtx)
		}

		//
		// simulate with call bundle
		callBundle := flots.CallBundle(rawTxs)
		fmt.Printf("\n🟢 call bundle: %+v\n", callBundle)

		//
		// send bundle
		bundleHash := flots.SendBundleWithRawTxs(rawTxs)
		fmt.Printf("\n🟢 bundle sent! hash: %s\n\n", bundleHash)

		// store (blocknum + plusBlocks) at time of sending the bundle (not sure if this is needed at all)
		latestBlockPlusWhenSending := flots.LatestBlockPlus()

		// wait for bundle to be mined for (blockTime * plusBlocks) seconds + 1 block as a buffer
		waitBlocks := flots.PlusBlocks.Int64() + 1
		minedUntil := internal.BlockTime * time.Duration(waitBlocks)
		fmt.Printf("bundleStats | mined until blockNum: %d | bundleHash: %s\n", latestBlockPlusWhenSending, bundleHash.String())

		//
		// start a timer that exits the program after (blockTime * plusBlocks) seconds
		go func() {
			// bundle send, now wait for bundle to be mined for (blockTime * plusBlocks) seconds
			killTimer := time.NewTimer(minedUntil)
			<-killTimer.C

			fmt.Print("\n\n")
			fmt.Printf("waited for %d blocks / %.0f seconds - tx is mined or never will be\n", waitBlocks, minedUntil.Seconds())

			bundleStats := flots.GetBundleStats(bundleHash)
			fmt.Printf("bundleStats | mined until blockNum: %d | bundleHash: %s:\n%+v\n", flots.LatestBlock().Int64()+waitBlocks, bundleHash.String(), bundleStats)

			os.Exit(0)
		}()

		//
		// check status periodically
		checkStatusEvery := time.Second * 2

		for {
			bundleStats := flots.GetBundleStats(bundleHash)
			fmt.Printf(" status | %d - %d | %+v | considered: %d | sealed: %d   || sleeping for %.0f sec...\n", flots.LatestBlock().Uint64(), latestBlockPlusWhenSending.Uint64(), bundleHash.String(), len(bundleStats.ConsideredByBuildersAt), len(bundleStats.SealedByBuildersAt), checkStatusEvery.Seconds())

			time.Sleep(checkStatusEvery)
		}
	},
}

func init() {
	FlotsCmd.AddCommand(sendBundleCmd)

	sendBundleCmd.Flags().StringSliceVarP(&flagRawTransactions, "rawtxs", "t", make([]string, 0), "signed transactions (get them from https://flashbots-bundler.surge.sh/rpc for example)")

	_ = callBundleCmd.MarkFlagRequired("rawtxs")

	sendBundleCmd.Flags().Int64Var(&flagPlusBlocks, "plusBlocks", 5, "blocks to add to the current block number")

	_ = viper.BindPFlag("flots.plusBlocks", sendBundleCmd.Flags().Lookup("plusBlocks"))
}

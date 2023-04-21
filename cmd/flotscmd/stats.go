/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package flotscmd

import (
	"fmt"

	"github.com/benleb/gloomberg/internal/flots"
	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	"github.com/spf13/cobra"
)

// statsCmd represents the userStats command
var statsCmd = &cobra.Command{
	Use:     "stats",
	Aliases: []string{"userStats", "bundleStats"},
	Short:   "gets the user stats and, if a bundle hash is provided, the bundle stats for that bundle",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stats called")

		flots := flots.New()

		userStats := flots.GetUserStats()

		// print user statistics
		lo.Info(fmt.Sprintf("High priority: %t", userStats.IsHighPriority))
		lo.Info(fmt.Sprintf("7 day fees: %s ETH", w3.FromWei(userStats.Last7dValidatorPayments, 18)))
		lo.Info(fmt.Sprintf("Total fees: %s ETH", w3.FromWei(userStats.AllTimeValidatorPayments, 18)))

		if flagBBundleHash == "" {
			return
		}

		fmt.Printf("\n\n\n\n")

		bundleHash := common.HexToHash(flagBBundleHash)
		lo.Info(fmt.Sprintf("bundleStats | blockNum: %d | bundleHash: %s\n", flots.LatestBlock().Uint64(), bundleHash.String()))

		bundleStats := flots.GetBundleStats(bundleHash)
		lo.Info(fmt.Sprintf("bundleStats: %+v\n", bundleStats))
	},
}

func init() {
	FlotsCmd.AddCommand(statsCmd)

	statsCmd.Flags().StringVarP(&flagBBundleHash, "bundleHash", "s", "", "bundleHash")
}

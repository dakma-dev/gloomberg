/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package flotscmd

import (
	"fmt"

	"github.com/benleb/gloomberg/internal/flots"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	"github.com/spf13/cobra"
)

// statsCmd represents the userStats command.
var statsCmd = &cobra.Command{
	Use:     "stats",
	Aliases: []string{"userStats", "bundleStats"},
	Short:   "gets the user stats and, if a bundle hash is provided, the bundle stats for that bundle",

	Run: func(cmd *cobra.Command, args []string) {
		// print header
		fmt.Println("\n" + flashBotsTitle + "\n")

		flots := flots.New()

		userStats := flots.GetUserStats()

		// colorize
		var prioStyle lipgloss.Style
		if userStats.IsHighPriority {
			prioStyle = style.TrendGreenStyle.Copy().Bold(true)
		} else {
			prioStyle = style.TrendRedStyle.Copy().Bold(true)
		}

		// print user statistics
		fmt.Printf("  user: %s\n\n", style.Bold(flots.UserAddress().String()))

		fmt.Printf("    high priority: %s\n", prioStyle.Render(fmt.Sprint(userStats.IsHighPriority)))
		fmt.Printf("    7 day fees: %sΞ\n", style.Bold(w3.FromWei(userStats.Last7dValidatorPayments, 18)))
		fmt.Printf("    total fees: %sΞ\n", style.Bold(w3.FromWei(userStats.AllTimeValidatorPayments, 18)))

		// print bundle statistics if a bundleHash was given
		if bundleHash := common.HexToHash(flagBBundleHash); bundleHash != (common.Hash{}) {
			fmt.Printf("\n\n - - - - - - - - - - - - - -\n\n")

			bundleStats := flots.GetBundleStats(bundleHash)

			// colorize
			if bundleStats.IsHighPriority {
				prioStyle = style.TrendGreenStyle.Copy().Bold(true)
			} else {
				prioStyle = style.TrendRedStyle.Copy().Bold(true)
			}

			fmt.Printf("  bundle: %s\n\n", style.Bold(bundleHash.String()))

			fmt.Printf("    receivedAt: %s\n", style.Bold(fmt.Sprint(bundleStats.ReceivedAt)))
			fmt.Printf("    isSimulated: %s\n", style.Bold(fmt.Sprint(bundleStats.IsSimulated)))
			if bundleStats.IsSimulated {
				fmt.Printf("    simulatedAt: %s\n", style.Bold(fmt.Sprint(bundleStats.SimulatedAt)))
			}
			fmt.Printf("    isHighPriority: %s\n", prioStyle.Render(fmt.Sprint(bundleStats.IsHighPriority)))
			fmt.Printf("    consideredByBuildersAt: %s\n", style.Bold(fmt.Sprint(len(bundleStats.ConsideredByBuildersAt))))
			fmt.Printf("    sealedByBuildersAt: %s\n", style.Bold(fmt.Sprint(bundleStats.SealedByBuildersAt)))
		}

		fmt.Printf("\n")
	},
}

func init() {
	FlotsCmd.AddCommand(statsCmd)

	statsCmd.Flags().StringVarP(&flagBBundleHash, "bundleHash", "s", "", "bundleHash")
}

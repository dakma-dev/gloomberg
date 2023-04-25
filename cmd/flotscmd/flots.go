package flotscmd

import (
	"fmt"

	"github.com/benleb/gloomberg/internal/style"
	"github.com/spf13/cobra"
)

var (
	blinkingRobo   = style.BoldStyle.Copy().Blink(true).Render("  🤖  ")
	flashBotsTitle = blinkingRobo + style.Bold("Flashbots") + blinkingRobo
)

// FlotsCmd represents the flots command.
var FlotsCmd = &cobra.Command{
	Use:   "flots",
	Short: "Interact with the Flashbots API/network",
	Long: fmt.Sprintf(`
 %s

Flashbots is a research and development organization formed to mitigate the
negative externalities posed by Maximal Extractable Value (MEV) to stateful
blockchains, starting with Ethereum    more info: https://www.flashbots.net


%s is a command for interacting with the Flashbots API/network.`, flashBotsTitle, style.Bold("gloomberg flots")),

	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("flots called")
	// },
}

var (
	flagRawTransactions []string
	flagBBundleHash     string

	flagPlusBlocks int64
)

func init() {}

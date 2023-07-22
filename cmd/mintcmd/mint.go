package mintcmd

import (
	"github.com/spf13/cobra"
)

// MintCmd represents the mint command.
var MintCmd = &cobra.Command{
	Use:   "mint",
	Short: "Mint something",
	// 	Long: fmt.Sprintf(`%s

	//	Mints the token from somewhere with the configured wallets.`, style.GetSmallHeader(internal.GloombergVersion)),
	//
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println("mint called")
	//	},
}

func init() {}

package mintcmd

import (
	"crypto/ecdsa"

	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

type MintWallet struct {
	privateKey *ecdsa.PrivateKey
	address    *common.Address

	color lipgloss.Color
	tag   string
}

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

var (
	flagURL string

	flagGasFeeCapMultiplier float64
	flagGasTipCapMultiplier float64

	flagPrivateKeys []string
	flagRPCs        []string
)

func init() {}

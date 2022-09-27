package cmd

import (
	"fmt"
	"time"

	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	paramContract string
	paramToken    int
)

// ownerCmd represents the owner command.
var ownerCmd = &cobra.Command{
	Use:   "owner",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: owner,
}

func init() {
	rootCmd.AddCommand(ownerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ownerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ownerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	ownerCmd.Flags().Bool("raw", false, "Only show the owner addresses")
	_ = viper.BindPFlag("owner.raw", ownerCmd.Flags().Lookup("raw"))

	ownerCmd.Flags().StringVar(&apiKeyMoralis, "moralis", "", "Moralis API Key")
	_ = viper.BindPFlag("moralis", ownerCmd.Flags().Lookup("moralis"))

	// apis
	ownerCmd.Flags().StringVar(&paramContract, "contract", "", "Contract address")
	_ = viper.BindPFlag("owner.contract", ownerCmd.Flags().Lookup("contract"))
	// apis
	ownerCmd.Flags().IntVar(&paramToken, "token", -1, "Token ID")
	_ = viper.BindPFlag("owner.token", ownerCmd.Flags().Lookup("token"))
}

func owner(_ *cobra.Command, _ []string) {
	if !viper.GetBool("owner.raw") {
		// print header
		header := style.GetHeader(Version)
		fmt.Println(header)
	}

	if !viper.IsSet("moralis") {
		fmt.Println("moralis key not set")
		gbl.Log.Fatal("moralis key not set")
	}

	if !common.IsHexAddress(paramContract) {
		fmt.Println("contract address not valid")
		gbl.Log.Fatal("contract address not valid")
	}

	contractAddress := common.HexToAddress(paramContract)

	client, _ := createMoralisHTTPClient()

	uniqueOwner := make(map[string]*Owner, 0)
	cursor := ""

	if !viper.GetBool("owner.raw") {
		fmt.Println(lipgloss.NewStyle().Padding(1, 0, 1, 2).Render(fmt.Sprintf("token %d on contract %s...", paramToken, contractAddress.Hex())))
	}

	for {
		ownerResponse, err := getOwnerPage(client, apiKeyMoralis, contractAddress, paramToken, cursor)
		if err != nil {
			gbl.Log.Fatal(err)
		}

		for _, owner := range ownerResponse.Result {
			uniqueOwner[owner.OwnerOf] = owner
		}

		if ownerResponse.Cursor == "" {
			break
		}

		cursor = ownerResponse.Cursor

		time.Sleep(time.Second * 1)
	}

	for address := range uniqueOwner {
		var ownerAddress string
		if viper.GetBool("owner.raw") {
			ownerAddress = address
		} else {
			ownerAddress = lipgloss.NewStyle().PaddingLeft(4).Render(address)
		}

		fmt.Println(ownerAddress)
	}

	if !viper.GetBool("owner.raw") {
		fmt.Println()
		fmt.Println(lipgloss.NewStyle().Padding(1, 0, 1, 2).Render(fmt.Sprintf("token %d on contract %s is owned by: %s", paramToken, contractAddress.Hex(), style.BoldStyle.Render(fmt.Sprintf("%d wallets", len(uniqueOwner))))))
		fmt.Println()
	}
}

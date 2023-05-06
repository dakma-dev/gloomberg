package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/accounts"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	addrPrefix     string
	addrSuffix     string
	pkPrefix       string
	pkSuffix       string
	derivationPath string

	path     accounts.DerivationPath
	parallel int
)

// generateCmd represents the generate command.
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate secure wallets with vanity addresses",

	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("log.logFile", "/tmp/gloomberg-generate.log")
		viper.Set("log.verbose", true)

		path = hdwallet.MustParseDerivationPath(derivationPath)

		// add the 0x prefix if not already there
		if !strings.HasPrefix(addrPrefix, "0x") {
			addrPrefix = "0x" + addrPrefix
		}

		fmt.Println("--")
		fmt.Printf("prefix:\t\t%s\n", addrPrefix)
		fmt.Printf("suffix:\t\t%s\n", addrSuffix)
		fmt.Printf("-> addresses:\t%s...%s\n\n", addrPrefix, addrSuffix)
		fmt.Println("")
		fmt.Printf("privateKey prefix:\t\t%s\n", pkPrefix)
		fmt.Printf("privateKey suffix:\t\t%s\n", pkSuffix)
		fmt.Printf("-> privateKey:\t%s...%s\n\n", pkPrefix, pkSuffix)
		fmt.Printf("path:\t\t%s\n", path.String())
		fmt.Printf("parallel:\t%d\n", parallel)
		fmt.Println("--")

		for i := 1; i <= parallel; i++ {
			go generateMnemonicWallets(i)
		}

		select {}
	},
}

func generateMnemonicWallets(workerID int) {
	fmt.Printf("worker %d started\n", workerID)

	for i := 1; ; i++ {
		if i%100000 == 0 {
			fmt.Printf("worker %d: %d\n", workerID, i)
		}

		mnemonic, _ := hdwallet.NewMnemonic(256)

		wallet, err := hdwallet.NewFromMnemonic(mnemonic)
		if err != nil {
			log.Fatal(err)
		}

		account, err := wallet.Derive(path, false)
		if err != nil {
			log.Fatal(err)
		}

		privateKey, err := wallet.PrivateKeyHex(account)
		if err != nil {
			log.Fatal(err)
		}

		// check privateKey prefix and suffix
		if pkPrefix != "" && !strings.HasPrefix(privateKey, pkPrefix) {
			continue
		} else if pkSuffix != "" && !strings.HasSuffix(privateKey, pkSuffix) {
			continue
		}

		publicKey, err := wallet.PublicKeyHex(account)
		if err != nil {
			log.Fatal(err)
		}

		if addrPrefix != "" && !strings.HasPrefix(account.Address.String(), addrPrefix) {
			continue
		} else if addrSuffix != "" && !strings.HasSuffix(account.Address.String(), addrSuffix) {
			continue
		}

		fmt.Println("")
		fmt.Printf("mnemonic:\t%s\n", mnemonic)
		fmt.Printf("path:\t\t%s\n", path.String())
		fmt.Printf("private:\t%s\n", privateKey)
		fmt.Printf("public:\t\t%s\n", publicKey)
		fmt.Printf("address:\t%s\n", lipgloss.NewStyle().Bold(true).Render(account.Address.String()))
		fmt.Println("")
	}
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	viper.SetDefault("path", "m/44'/60'/0'/0/0")

	generateCmd.Flags().StringVar(&addrPrefix, "prefix", "", "prefix the address must have")
	generateCmd.Flags().StringVar(&addrSuffix, "suffix", "", "suffix the address must have")
	generateCmd.Flags().StringVar(&pkPrefix, "pk-prefix", "", "prefix the privatekey must have")
	generateCmd.Flags().StringVar(&pkSuffix, "pk-suffix", "", "suffix the privatekey must have")
	generateCmd.Flags().StringVar(&derivationPath, "path", "m/44'/60'/0'/0/0", "address derivation path")
	generateCmd.Flags().IntVar(&parallel, "parallel", 4, "number of generators running in parallel")
}

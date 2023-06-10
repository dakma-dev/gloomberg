package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/accounts"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	addrPrefix     string
	addrPrefixes   []string
	addrSuffix     string
	addrSuffixes   []string
	notContains    string
	pkPrefix       string
	pkSuffix       string
	derivationPath string

	path     accounts.DerivationPath
	parallel int
	zeroes   int
	fs       int

	startTime = time.Now()
)

// generateCmd represents the generate command.
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate secure wallets with vanity addresses",

	Run: func(_ *cobra.Command, _ []string) {
		viper.Set("log.logFile", "/tmp/gloomberg-generate.log")
		viper.Set("log.verbose", true)

		path = hdwallet.MustParseDerivationPath(derivationPath)

		// prepare address prefixes/suffixes
		if addrPrefix != "" {
			addrPrefixes = append(addrPrefixes, addrPrefix)
		}

		for i, prefix := range addrPrefixes {
			if !strings.HasPrefix(prefix, "0x") {
				addrPrefixes[i] = "0x" + prefix
			}
		}

		if addrSuffix != "" {
			addrSuffixes = append(addrSuffixes, addrSuffix)
		}

		log.Printf("addrPrefixes: %+v", addrPrefixes)
		log.Printf("addrSuffixes: %+v", addrSuffixes)

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
		if i%10_000 == 0 {
			fmt.Printf("worker %d: %d | running: %+v\n", workerID, i, time.Since(startTime).String())
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

		// check address prefix
		foundPrefix := false
		for _, prefix := range addrPrefixes {
			if strings.HasPrefix(strings.ToLower(account.Address.String()), strings.ToLower(prefix)) {
				foundPrefix = true

				log.Printf("found prefix: %s | %s", prefix, account.Address.String())

				break
			}
		}

		if len(addrPrefixes) > 0 && !foundPrefix {
			continue
		}

		// check address suffix
		foundSuffix := false
		for _, suffix := range addrSuffixes {
			if strings.HasSuffix(strings.ToLower(account.Address.String()), strings.ToLower(suffix)) {
				foundSuffix = true

				log.Printf("found suffix: %s | %s", suffix, account.Address.String())

				break
			}
		}

		if len(addrSuffixes) > 0 && !foundSuffix {
			continue
		}

		// check address notContains
		if notContains != "" && strings.Contains(strings.ToLower(account.Address.String()), strings.ToLower(notContains)) {
			continue
		}

		// check zeroes and fs
		if zeroes > 0 {
			counter := 0
			for _, c := range account.Address.String() {
				if c == '0' {
					counter++
				}
			}

			if counter < zeroes {
				continue
			}
		}

		if fs > 0 {
			counter := 0
			for _, c := range account.Address.String() {
				if c == 'f' {
					counter++
				}
			}

			if counter < fs {
				continue
			}
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
	generateCmd.Flags().StringSliceVar(&addrPrefixes, "prefixes", []string{}, "prefixes of which the address must have one")
	generateCmd.Flags().StringVar(&addrSuffix, "suffix", "", "suffix the address must have")
	generateCmd.Flags().StringSliceVar(&addrSuffixes, "suffixes", []string{}, "suffixes of which the address must have one")
	generateCmd.Flags().StringVar(&notContains, "not-contains", "", "character the address must not contain")
	generateCmd.Flags().StringVar(&pkPrefix, "pk-prefix", "", "prefix the privatekey must have")
	generateCmd.Flags().StringVar(&pkSuffix, "pk-suffix", "", "suffix the privatekey must have")
	generateCmd.Flags().StringVar(&derivationPath, "path", "m/44'/60'/0'/0/0", "address derivation path")
	generateCmd.Flags().IntVar(&parallel, "parallel", 4, "number of generators running in parallel")
	generateCmd.Flags().IntVar(&zeroes, "zeroes", 0, "number of zeroes the address must have")
	generateCmd.Flags().IntVar(&fs, "fs", 0, "number of Fs the address must have")
}

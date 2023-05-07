package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/benleb/gloomberg/internal/chawago"
	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/benleb/gloomberg/internal/nemo/provider"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type TokenERC20 struct {
	Address common.Address
	Name    string
	// Symbol string
	Icon string
	// UniswapPairWETH common.Address
}

var tokensERC20 = []TokenERC20{
	{common.HexToAddress("0x7baece5d47f1bc5e1953fbe0e9931d54dab6d810"), "TURBO", "🌪️"},
	{common.HexToAddress("0xcf1e086a145dd7b5b771b1e6acbbc1d2b58d7e80"), "GENW", "🧬"},
	{common.HexToAddress("0x9bf1D7D63dD7a4ce167CF4866388226EEefa702E"), "BEN", "🌽"},
}

// get them from the config
var addressesToCheck = map[common.Address]map[string]string{
	common.HexToAddress("0x974404e37f......126a8E355690"): {"name": "drastic"},
}

// chawagoCmd represents the chawago command.
var chawagoCmd = &cobra.Command{
	Use:     "chawago",
	Aliases: []string{"chawa"},
	Short:   "chain watcher (go version)",

	Run: func(cmd *cobra.Command, args []string) {
		// queue for raw logs received from providers/ethereum nodes
		qRawLogs := make(chan types.Log, 1024)
		// queue for raw logs received from providers/ethereum nodes
		qPendingTx := make(chan *types.Transaction, 1024)
		// queue for full transactions with logs
		qTxsWithLogs := make(chan chawago.TxWithLogs, 1024)

		// init provider pool
		pool, err := provider.FromConfig(viper.Get("provider"))
		if err != nil {
			log.Fatal("❌ running provider failed, exiting")
		}

		// subscribe
		if _, err = pool.SubscribeToEverything(qRawLogs); err != nil {
			log.Fatal("❌ subscribing to provider failed, exiting")
		}
		if _, err = pool.SubscribeToEverythingPending(qPendingTx); err != nil {
			log.Fatal("❌ subscribing to provider failed, exiting")
		}

		// get transactions for logs
		go chawago.GetTransactionsForLogs(qRawLogs, qTxsWithLogs, pool)

		go chawago.GetPendingTransactions(qPendingTx, qTxsWithLogs, pool)

		currentBlock, _ := pool.BlockNumber(context.TODO())

		// handle received transactions
		for txWithLogs := range qTxsWithLogs {
			if !txWithLogs.Pending && txWithLogs.BlockNumber.Uint64() > currentBlock {
				currentBlock = txWithLogs.BlockNumber.Uint64()
				log.Print("")
			}

			if txWithLogs.Value().Int64() == 0 {
				continue
			}

			// erc20
			for _, token := range tokensERC20 {
				if foundIn := txWithLogs.CheckAddress(token.Address); foundIn >= 0 {
					foundMsg := strings.Builder{}
					foundMsg.WriteString(fmt.Sprintf("%d | %v | ", txWithLogs.BlockNumber.Uint64(), txWithLogs.Pending))

					if token.Icon != "" {
						foundMsg.WriteString(token.Icon + " ")
					}

					if token.Name != "" {
						foundMsg.WriteString("$" + style.Bold(token.Name) + " | ")
					}

					var valueStyle lipgloss.Style

					// found in
					switch foundIn {
					case chawago.Topic1:
						valueStyle = style.TrendGreenStyle.Copy().Bold(true)
					case chawago.Topic2:
						valueStyle = style.TrendRedStyle.Copy().Bold(true)
					}

					// value
					foundMsg.WriteString(valueStyle.Render(fmt.Sprintf("%5.2f", price.NewPrice(txWithLogs.Value()).Ether())) + "Ξ  | ")

					// from -> to
					// foundMsg.WriteString(fmt.Sprintf("%s -> %s", txWithLogs.Sender().Hex(), txWithLogs.To().Hex()))

					// print
					log.Printf(foundMsg.String())
				}
			}

			// wallets
			for address, addressInfo := range addressesToCheck {
				if foundIn := txWithLogs.CheckAddress(address); foundIn >= 0 {
					foundMsg := strings.Builder{}
					foundMsg.WriteString(fmt.Sprintf("%d | %v | ", txWithLogs.BlockNumber.Uint64(), txWithLogs.Pending))

					if addressInfo["icon"] != "" {
						foundMsg.WriteString(addressInfo["icon"] + " ")
					}

					if addressInfo["name"] != "" {
						foundMsg.WriteString(style.Bold(addressInfo["name"]) + " | ")
					}

					var valueStyle lipgloss.Style

					// found in
					switch foundIn {
					case chawago.Topic1:
						valueStyle = style.TrendGreenStyle.Copy().Bold(true)
					case chawago.Topic2:
						valueStyle = style.TrendRedStyle.Copy().Bold(true)
					}

					// value
					foundMsg.WriteString(valueStyle.Render(fmt.Sprintf("%5.2f", price.NewPrice(txWithLogs.Value()).Ether())) + "Ξ  | ")

					// from -> to
					foundMsg.WriteString(fmt.Sprintf("%s -> %s", txWithLogs.Sender().Hex(), txWithLogs.To().Hex()))

					// print
					log.Printf(foundMsg.String())
				}
			}
		}

		// // generate a xid | https://github.com/rs/xid
		// guid := xid.New()
		// log.Printf(guid.String())
	},
}

func init() { rootCmd.AddCommand(chawagoCmd) }

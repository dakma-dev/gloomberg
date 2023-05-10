package cmd

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/abis/uniswapv2"
	"github.com/benleb/gloomberg/internal/abis/uniswapv3"
	"github.com/benleb/gloomberg/internal/chawago"
	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/benleb/gloomberg/internal/nemo/provider"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils"
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
	PairsWETH []common.Address
}

var tokensMapERC20 = map[common.Address]TokenERC20{
	common.HexToAddress("0x3adefbdb5101972823cb6b320d4a269dead1546d"): {common.HexToAddress("0x3adefbdb5101972823cb6b320d4a269dead1546d"), "BLOW", "🌬️", []common.Address{common.HexToAddress("0x9c8d68eE8c97d36DA37Eb58285712A7656E6CF4c")}},
	common.HexToAddress("0x7baece5d47f1bc5e1953fbe0e9931d54dab6d810"): {common.HexToAddress("0x7baece5d47f1bc5e1953fbe0e9931d54dab6d810"), "TURBO", "🌪️", []common.Address{common.HexToAddress("0x7baecE5d47f1BC5E1953FBE0E9931D54DAB6D810")}},
	common.HexToAddress("0x9bf1D7D63dD7a4ce167CF4866388226EEefa702E"): {common.HexToAddress("0x9bf1D7D63dD7a4ce167CF4866388226EEefa702E"), "BEN", "💫", []common.Address{common.HexToAddress("0xdce93ed9ae7c53143e19cf799d156b72d1cc2777")}},
}

type LogTransfer struct {
	From         common.Address
	To           common.Address
	AmountTokens *big.Int
}

// chawagoCmd represents the chawago command.
var chawagoCmd = &cobra.Command{
	Use:     "chawago",
	Aliases: []string{"chawa"},
	Short:   "chain watcher (go version)",

	Run: func(_ *cobra.Command, _ []string) {
		log.SetReportCaller(false)
		log.TimestampStyle = style.GrayStyle

		// queue for raw logs received from providers/ethereum nodes
		qRawLogs := make(chan types.Log, 1024)
		// queue for raw logs received from providers/ethereum nodes
		// qPendingTx := make(chan *types.Transaction, 1024)
		// queue for full transactions with logs
		qTxsWithLogs := make(chan chawago.TxWithLogs, 1024)

		// init provider pool
		pool, err := provider.FromConfig(viper.Get("provider"))
		if err != nil {
			log.Fatal("❌ running provider failed, exiting")
		}

		// erc20ABI, _ := erc20.ERC20MetaData.GetAbi()
		uniswapV2FactoryABI, _ := uniswapv2.FactoryMetaData.GetAbi()
		// uniswapV2PairABI, _ := uniswapv2.PairMetaData.GetAbi()
		uniswapV3PoolABI, _ := uniswapv3.PoolMetaData.GetAbi()

		// subscribe
		if _, err = pool.SubscribeToEverything(qRawLogs); err != nil {
			log.Fatal("❌ subscribing to provider failed, exiting")
		}
		// if _, err = pool.SubscribeToEverythingPending(qPendingTx); err != nil {
		// 	log.Fatal("❌ subscribing to provider failed, exiting")
		// }

		// topicSwapUniswapV2 := uniswapV2PairABI.Events["Swap"].ID
		// topicSwapUniswapV3 := uniswapV3PoolABI.Events["Swap"].ID
		// topicPairCreatedUniswapV2 := uniswapV2FactoryABI.Events["PairCreated"].ID

		// if _, err = pool.SubscribeToTopics(qRawLogs, [][]common.Hash{{topicSwapUniswapV2}, {}, {}, {}}); err != nil {
		// 	log.Fatal("❌ subscribing to provider failed, exiting")
		// }

		// uniswapV3FactoryFilterer, err := uniswapv3.NewFactoryFilterer(internal.UniswapV3FactoryContractAddress, pool.GetProviders()[0].Client)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		pairToTokenMap := make(map[common.Address]TokenERC20)

		for _, token := range tokensMapERC20 {
			for _, pairAddress := range token.PairsWETH {
				pairToTokenMap[pairAddress] = token
			}
		}

		// get transactions for logs
		go chawago.GetTransactionsForLogs(qRawLogs, qTxsWithLogs, pool)

		// go chawago.GetPendingTransactions(qPendingTx, qTxsWithLogs, pool)

		currentBlock, _ := pool.BlockNumber(context.TODO())

		// handle received transactions
		for txWithLogs := range qTxsWithLogs {
			if !txWithLogs.Pending && txWithLogs.BlockNumber.Uint64() > currentBlock {
				currentBlock = txWithLogs.BlockNumber.Uint64()

				log.Print(style.DarkGrayStyle.Render("0x") + style.Gray4Style.Render(txWithLogs.BlockNumber.Text(16)))
			}

			var valueEther, valueToken *big.Int

			// erc20Filterer, err := erc20.NewERC20Filterer(txWithLogs.ContractAddress, pool.GetProviders()[0].Client)
			// if err != nil {
			// 	log.Fatal(err)
			// }

			uniswapV3PoolFilterer, err := uniswapv3.NewPoolFilterer(internal.UniswapV3FactoryContractAddress, pool.GetProviders()[0].Client)
			if err != nil {
				log.Fatal(err)
			}

			uniswapV2FactoryFilterer, err := uniswapv2.NewFactoryFilterer(internal.UniswapV2FactoryContractAddress, pool.GetProviders()[0].Client)
			if err != nil {
				log.Fatal(err)
			}

			// uniswapV2PairFilterer, err := uniswapv2.NewPairFilterer(internal.UniswapV2Router01ContractAddress, pool.GetProviders()[0].Client)
			// if err != nil {
			// 	log.Fatal(err)
			// }

			msg := strings.Builder{}

			for _, rawLog := range txWithLogs.Logs {
				// filter logs without topics
				if len(rawLog.Topics) == 0 {
					continue
				}

				switch {
				// UniswapV2 PairCreated
				case rawLog.Topics[0] == uniswapV2FactoryABI.Events["PairCreated"].ID:
					uniswapV2PairCreated, err := uniswapV2FactoryFilterer.ParsePairCreated(*rawLog)
					if err != nil {
						log.Error(err)

						log.Printf("uniswapV2PairCreated rawLog.Topics: %+v\n", rawLog.Topics)

						continue
					}

					if uniswapV2PairCreated.Token0 != internal.WETHContractAddress && uniswapV2PairCreated.Token1 != internal.WETHContractAddress {
						log.Print("new pair is not a WETH pair")

						continue
					}

					log.Printf("%s 🧑‍🤝‍🧑 %s | %s", style.Bold("PairCreated"), style.TerminalLink(utils.GetEtherscanTokenURL(uniswapV2PairCreated.Pair.Hex()), uniswapV2PairCreated.Pair.Hex()), uniswapV2PairCreated.Token0.Hex())

				// UniswapV3 Swap
				case rawLog.Topics[0] == uniswapV3PoolABI.Events["Swap"].ID:
					uniswapV3Swap, err := uniswapV3PoolFilterer.ParseSwap(*rawLog)
					if err != nil {
						log.Error(err)

						log.Printf("uniswapV3Swap rawLog.Topics: %+v\n", rawLog.Topics)

						continue
					}

					var token *TokenERC20
					if t, ok := tokensMapERC20[uniswapV3Swap.Raw.Address]; ok {
						token = &t
					}
					if t, ok := pairToTokenMap[uniswapV3Swap.Raw.Address]; ok {
						token = &t
					}

					if token == nil {
						log.Debug("token empty")

						continue
					}

					if rawLog.Address == internal.WETHContractAddress {
						valueEther = uniswapV3Swap.Amount0
						valueToken = uniswapV3Swap.Amount1
					} else {
						valueEther = uniswapV3Swap.Amount1
						valueToken = uniswapV3Swap.Amount0
					}

					tokenBuy := valueEther.Sign() > 0

					amountPaid := price.NewPrice(big.NewInt(0).Abs(valueEther))

					tokenPerEther := new(big.Int).Div(big.NewInt(0).Abs(valueToken), big.NewInt(0).Abs(valueEther))
					tokenPerEtherPrice := price.NewPrice(tokenPerEther)

					log.Debug(" %+v | +%+v | %+v", valueEther, amountPaid, tokenBuy)

					// default style
					var redStyle, greenStyle, valueStyle lipgloss.Style
					if amountPaid.Ether() > 1 {
						redStyle = style.TrendRedStyle.Copy()
						greenStyle = style.TrendGreenStyle.Copy()
					} else {
						redStyle = style.TrendLightRedStyle.Copy()
						greenStyle = style.TrendLightGreenStyle.Copy()
					}

					if tokenBuy {
						valueStyle = greenStyle.Bold(false)
					} else {
						valueStyle = redStyle.Bold(false)
					}

					if amountPaid.Ether() > 3 {
						valueStyle = valueStyle.Bold(true)
					}

					tokenStyle := lipgloss.NewStyle().Bold(true).Foreground(style.GenerateColorWithSeed(int64(token.Address.Big().Int64())))
					tokenStyleFaint := tokenStyle.Copy().Faint(true)

					msg.WriteString(token.Icon + " " + valueStyle.Render(fmt.Sprintf("%5.2f", amountPaid.Ether())) + "Ξ")
					msg.WriteString(" | " + style.GrayStyle.Render(fmt.Sprintf("%7.4f", 1/tokenPerEtherPrice.Gwei())) + "Ξ/")
					msg.WriteString(" | " + tokenStyleFaint.Render("$") + tokenStyle.Render(token.Name))
					msg.WriteString(" | " + style.TerminalLink(utils.GetEtherscanTxURL(txWithLogs.Hash().Hex()), "Tx"))

					// // ERC20 transfer
					// case rawLog.Topics[0] == erc20ABI.Events["Transfer"].ID:
					// 	erc20transfer, err := erc20Filterer.ParseTransfer(*rawLog)
					// 	if err != nil {
					// 		log.Error(err)

					// 		log.Printf("uniswapV2PairCreated rawLog.Topics: %+v\n", rawLog.Topics)

					// 		continue
					// 	}

					// 	log.Printf("%s  -->  %s:   %+v", erc20transfer.From, erc20transfer.To, erc20transfer.Value)

				}

				if msg.Len() > 0 {
					log.Print(msg.String())

					break
				}
			}

			// if tokenTransfer != nil && transferredToken != nil && wethTransfer != nil {
			// 	valueStyle := style.LightGrayStyle
			// 	if wethTransfer.To == *txWithLogs.Sender() {
			// 		valueStyle = style.TrendRedStyle.Copy().Bold(true)
			// 	} else if tokenTransfer.To == *txWithLogs.Sender() {
			// 		valueStyle = style.TrendGreenStyle.Copy().Bold(true)
			// 	}

			// 	// log.Print("weth")
			// 	// log.Print(wethTransfer.From)
			// 	// log.Print(wethTransfer.To)

			// 	// log.Print("token")
			// 	// log.Print(tokenTransfer.From)
			// 	// log.Print(tokenTransfer.To)

			// 	// log.Print("tx")
			// 	// log.Print(txWithLogs.Sender())
			// 	// log.Print(txWithLogs.To())

			// 	msg := strings.Builder{}

			// 	tokenStyle := lipgloss.NewStyle().Bold(true).Foreground(style.GenerateColorWithSeed(txWithLogs.ContractAddress.Big().Int64()))
			// 	tokenStyleFaint := tokenStyle.Copy().Faint(true)

			// 	msg.WriteString(transferredToken.Icon + " " + valueStyle.Render(fmt.Sprintf("%5.2f", price.NewPrice(value).Ether())) + "Ξ")
			// 	msg.WriteString(" | " + tokenStyleFaint.Render("$") + tokenStyle.Render(transferredToken.Name))
			// 	msg.WriteString(" | " + style.TerminalLink(utils.GetEtherscanTxURL(txWithLogs.Hash().Hex()), "Tx"))
			// 	msg.WriteString(" | " + style.GrayStyle.Render("0x") + fmt.Sprintf("%s | %v | ", txWithLogs.BlockNumber.Text(16), txWithLogs.Pending))

			// 	log.Print(msg.String())
			// }
		}
	},
}

func init() { rootCmd.AddCommand(chawagoCmd) }

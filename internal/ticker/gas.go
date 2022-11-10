package ticker

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal/nodes"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
)

func GasTicker(gasTicker *time.Ticker, ethNodes *nodes.Nodes, queueOutput *chan string) {
	oldGasPrice := 0

	for range gasTicker.C {
		gasNode := ethNodes.GetRandomLocalNode()
		gasLine := strings.Builder{}

		if viper.GetBool("log.verbose") {
			gasLine.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#1A1A1A")).Render(fmt.Sprint(gasNode.Marker)))
			gasLine.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#111111")).Render("|"))
		}

		gasLine.WriteString(style.GrayStyle.Copy().Faint(true).Render(time.Now().Format("15:04:05")))
		gasLine.WriteString(" " + style.DarkGrayStyle.Render("🧟"))

		gasLine.WriteString("   ")

		if gasInfo, err := gasNode.GetCurrentGasInfo(); err == nil && gasInfo != nil {
			// gas price
			if gasInfo.GasPriceWei.Cmp(big.NewInt(0)) > 0 {
				gasPriceGwei, _ := nodes.WeiToGwei(gasInfo.GasPriceWei).Float64()
				gasPrice := int(math.Round(gasPriceGwei))

				if math.Abs(float64(gasPrice-oldGasPrice)) < 2.0 {
					continue
				}

				oldGasPrice = gasPrice

				// // tip / priority fee
				// var gasTip int
				// if gasInfo.GasTipWei.Cmp(big.NewInt(0)) > 0 {
				// 	gasTipGwei, _ := nodes.WeiToGwei(gasInfo.GasTipWei).Float64()
				// 	gasTip = int(math.Round(gasTipGwei))
				// 	fmt.Printf("gasInfo.GasTipWei: %+v | gasTipGwei: %+v | gasTip: %+v\n", gasInfo.GasTipWei, gasTipGwei, gasTip)
				// }

				intro := style.DarkerGrayStyle.Render("~  ") + style.DarkGrayStyle.Render("gas") + style.DarkerGrayStyle.Render("  ~   ")
				outro := style.DarkerGrayStyle.Render("   ~   ~")
				divider := style.DarkerGrayStyle.Render("   ~   ~   ~   ~   ~   ~   ")

				formattedGas := style.GrayStyle.Render(fmt.Sprintf("%d", gasPrice)) + style.DarkGrayStyle.Render("gw")
				formattedGasAndTip := formattedGas

				// if gasTip > 0 {
				// 	formattedGasAndTip = formattedGas + "|" + style.GrayStyle.Render(fmt.Sprintf("%d", gasTip)) + style.DarkGrayStyle.Render("gw")
				// }

				gasLine.WriteString(intro + formattedGas + divider + formattedGasAndTip + divider + formattedGas + outro)
			}
		}

		*queueOutput <- gasLine.String()
	}
}

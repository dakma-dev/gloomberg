package ticker

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal/server/node"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
)

func GasTicker(gasTicker *time.Ticker, nodes *node.Nodes, queueOutput *chan string) {
	oldGasPrice := 0

	for range gasTicker.C {
		gasNode := nodes.GetRandomLocalNode()
		gasLine := strings.Builder{}

		if viper.GetBool("log.verbose") {
			gasLine.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#1A1A1A")).Render(fmt.Sprint(gasNode.Marker)))
			gasLine.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#111111")).Render("|"))
		}

		gasLine.WriteString(style.GrayStyle.Copy().Faint(true).Render(time.Now().Format("15:04:05")))
		gasLine.WriteString(" " + style.DarkGrayStyle.Render("🧟"))
		// gasLine.WriteString(" ⛽️")

		gasLine.WriteString("   ")

		if gasInfo, err := gasNode.GetCurrentGasInfo(); err == nil && gasInfo != nil {
			// gas price
			if gasInfo.GasPriceWei.Cmp(big.NewInt(0)) > 0 {
				gasPriceGwei, _ := node.WeiToGwei(gasInfo.GasPriceWei).Float64()
				gasPrice := int(math.Round(gasPriceGwei))

				if math.Abs(float64(gasPrice-oldGasPrice)) < 2.0 {
					continue
				}

				oldGasPrice = gasPrice

				gasLine.WriteString(style.DarkerGrayStyle.Render("~  gas  ~   "))
				gasLine.WriteString(style.LightGrayStyle.Render(fmt.Sprintf("%d", gasPrice))) // ⛽️
				gasLine.WriteString(style.DarkGrayStyle.Render("gw"))

				gasLine.WriteString(style.DarkerGrayStyle.Render("   ~   ~   ~   ~   ~   ~   "))
				gasLine.WriteString(style.LightGrayStyle.Render(fmt.Sprintf("%d", gasPrice))) // ⛽️
				gasLine.WriteString(style.DarkGrayStyle.Render("gw"))

				gasLine.WriteString(style.DarkerGrayStyle.Render("   ~   ~   ~   ~   ~   ~   "))
				gasLine.WriteString(style.LightGrayStyle.Render(fmt.Sprintf("%d", gasPrice))) // ⛽️
				gasLine.WriteString(style.DarkGrayStyle.Render("gw"))
				gasLine.WriteString(style.DarkerGrayStyle.Render("   ~   ~"))
			}
			// // tip / priority fee
			// if gasInfo.GasTipWei.Cmp(big.NewInt(0)) > 0 {
			// 	gasTipGwei, _ := weiToGwei(gasInfo.GasTipWei).Float64()
			// 	gasLine.WriteString(style.GrayStyle.Render(fmt.Sprintf("  |  tip: %d", int(math.Ceil(gasTipGwei)))))
			// }
		}

		*queueOutput <- gasLine.String()
	}
}

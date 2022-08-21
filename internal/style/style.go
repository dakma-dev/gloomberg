package style

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"os"
	"strings"

	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	// tui.
	pink                 = lipgloss.AdaptiveColor{Light: "#FF44DD", Dark: "#FF0099"}
	OwnerGreen           = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	Subtle               = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	darkGray             = lipgloss.Color("#333")
	OpenseaToneBlue      = lipgloss.Color("#5f7699")
	PinkStyle            = lipgloss.NewStyle().Foreground(pink)
	TrendGreenStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#66CC66"))
	TrendLightGreenStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#667066"))
	TrendRedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF6666"))
	TrendLightRedStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#806666"))
	AlmostWhiteStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#EEE"))
	DarkWhiteStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#DDD"))
	VeryLightGrayStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#BBB"))
	LightGrayStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#999"))
	GrayStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color("#666"))
	GrayFaintStyle       = GrayStyle.Copy().Faint(true)
	DarkGrayStyle        = lipgloss.NewStyle().Foreground(darkGray)
	BoldStyle            = lipgloss.NewStyle().Bold(true)
	PinkBoldStyle        = BoldStyle.Copy().Foreground(pink)
	OwnerGreenBoldStyle  = BoldStyle.Copy().Foreground(OwnerGreen)
	GrayBoldStyle        = BoldStyle.Copy().Foreground(GrayStyle.GetForeground())
	Sharrow              = BoldStyle.Copy().SetString("→")
	DividerArrowRight    = GrayBoldStyle.SetString("→")
	DividerArrowLeft     = GrayBoldStyle.SetString("←")

	// borders
	// noBorderStyle              = lipgloss.NewStyle().BorderTop(false).BorderBottom(false).BorderLeft(false).BorderRight(false)
	// baseBorderStyle s= lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(darkGray).PaddingRight(2)
	// topBorderStyle             = baseBorderStyle.Copy().BorderTop(true).BorderBottom(false)
	// bottomBorderStyle          = baseBorderStyle.Copy().BorderTop(false).BorderBottom(true)
	// topBottomBorderStyle       = baseBorderStyle.Copy().BorderTop(true).BorderBottom(true)
	// topBottomHiddenBorderStyle = baseBorderStyle.Copy().BorderStyle(lipgloss.HiddenBorder()).BorderTop(true).BorderBottom(true).
)

var ShadesPink = []lipgloss.Color{
	"#ffffff",
	"#ffe5f4",
	"#ffccea",
	"#ffb2e0",
	"#ff99d6",
	"#ff7fcc",
	"#ff66c1",
	"#ff4cb7",
	"#ff32ad",
	"#ff19a3",
	"#ff0099",
}

var PaletteRLD = []lipgloss.Color{
	"#D23469",
	"#400817",
	"#8D1537",
	// "#120407",
	"#6A0F27",
	// "#F6C3DE",
	"#6F2B4E",
	"#A14C7C",
	"#A46C8C",
	"#6C3441",
}

var TitleStyle = lipgloss.NewStyle().
	MarginLeft(1).
	MarginRight(5).
	Padding(0, 1).
	Italic(true).
	Foreground(lipgloss.Color("#FFF7DB")).
	SetString("BTV")

var DescStyle = lipgloss.NewStyle().MarginTop(1)

var InfoStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderTop(true).
	BorderForeground(Subtle)

var baseDivider = lipgloss.NewStyle().
	SetString("•").
	Padding(0, 0).
	Foreground(pink)
var Divider = baseDivider.String()

func GetHeader(version string, commit string) string {
	randColorID, _ := crand.Int(crand.Reader, big.NewInt(int64(len(PaletteRLD))))
	randHeaderID, _ := crand.Int(crand.Reader, big.NewInt(int64(len(headers))))

	headerLogo := headers[randHeaderID.Int64()]
	headerColor := PaletteRLD[randColorID.Int64()]

	header := strings.Builder{}

	headerStyle := lipgloss.NewStyle().Foreground(headerColor).Padding(2, 0, 1, 0)
	subHeaderStyle := DarkGrayStyle.Copy()

	header.WriteString(headerStyle.Render(headerLogo) + "\n")

	subHeader := strings.Builder{}
	subHeader.WriteString(lipgloss.NewStyle().Foreground(headerColor).Bold(true).Render("·"))
	subHeader.WriteString(" " + DarkGrayStyle.Render("gloomberg"))
	subHeader.WriteString(" " + lipgloss.NewStyle().Foreground(lipgloss.Color("#444444")).Render(version))
	// subHeader.WriteString(" " + GrayStyle.Render(version))
	subHeader.WriteString(" " + lipgloss.NewStyle().Foreground(headerColor).Bold(true).Render("|"))
	subHeader.WriteString(" " + DarkGrayStyle.Render("github.com/benleb/gloomberg"))
	subHeader.WriteString(" " + lipgloss.NewStyle().Foreground(headerColor).Bold(true).Render("·"))

	header.WriteString(subHeaderStyle.Render(subHeader.String()))
	// header.WriteString("\n" + "💰 ❌ 💤")

	width, _, err := terminal.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		gbl.Log.Error(err)
	}

	return lipgloss.NewStyle().PaddingBottom(3).Width(width - 4).Align(lipgloss.Center).Render(header.String())
}

func GetHeader2() string {
	// output header
	doc := strings.Builder{}

	var (
		colors = ColorGrid(1, 5)
		title  strings.Builder
	)

	for i, v := range colors {
		const offset = 2

		c := lipgloss.Color(v[0])

		fmt.Fprint(&title, TitleStyle.Copy().MarginLeft(i*offset).Background(c))

		if i < len(colors)-1 {
			title.WriteRune('\n')
		}
	}

	desc := lipgloss.JoinVertical(lipgloss.Left,
		DescStyle.Render("👀 watching the ethereum chain"),
		InfoStyle.Render(fmt.Sprintf("  config %s %s", Divider, BoldStyle.Render(viper.ConfigFileUsed()))),
		InfoStyle.UnsetBorderStyle().Render(fmt.Sprintf("  debug %s %s", Divider, BoldStyle.Render(fmt.Sprint(viper.GetBool("log.debug"))))),
		InfoStyle.UnsetBorderStyle().Render(fmt.Sprintf("  verbose %s %s", Divider, BoldStyle.Render(fmt.Sprint(viper.GetBool("log.verbose"))))),
		InfoStyle.UnsetBorderStyle().Render(fmt.Sprintf("  status %s every %s", Divider, BoldStyle.Render(viper.GetDuration("stats.interval").String()))),
	)

	row := lipgloss.JoinHorizontal(lipgloss.Top, title.String(), desc)
	doc.WriteString(row + "\n")

	return doc.String()
}

func GetPriceShadeColor(txValue *big.Float) lipgloss.Color {
	var priceColor lipgloss.Color

	switch {
	case txValue.Cmp(big.NewFloat(0)) == 0:
		priceColor = "#333333"
	case txValue.Cmp(big.NewFloat(0.25)) < 0:
		priceColor = ShadesPink[0]
	case txValue.Cmp(big.NewFloat(0.5)) < 0:
		priceColor = ShadesPink[1]
	case txValue.Cmp(big.NewFloat(0.75)) < 0:
		priceColor = ShadesPink[2]
	case txValue.Cmp(big.NewFloat(1.0)) < 0:
		priceColor = ShadesPink[3]
	case txValue.Cmp(big.NewFloat(2.0)) < 0:
		priceColor = ShadesPink[5]
	case txValue.Cmp(big.NewFloat(3.0)) < 0:
		priceColor = ShadesPink[6]
	case txValue.Cmp(big.NewFloat(5.0)) < 0:
		priceColor = ShadesPink[8]
	case txValue.Cmp(big.NewFloat(5.0)) >= 0:
		priceColor = ShadesPink[9]
	}

	return priceColor
}

func ColorGrid(xSteps, ySteps int) [][]string {
	x0y0, _ := colorful.Hex("#F25D94")
	x1y0, _ := colorful.Hex("#EDFF82")
	x0y1, _ := colorful.Hex("#643AFF")
	x1y1, _ := colorful.Hex("#14F9D5")

	x0 := make([]colorful.Color, ySteps)
	for i := range x0 {
		x0[i] = x0y0.BlendLuv(x0y1, float64(i)/float64(ySteps))
	}

	x1 := make([]colorful.Color, ySteps)
	for i := range x1 {
		x1[i] = x1y0.BlendLuv(x1y1, float64(i)/float64(ySteps))
	}

	grid := make([][]string, ySteps)

	for x := 0; x < ySteps; x++ {
		y0 := x0[x]

		grid[x] = make([]string, xSteps)
		for y := 0; y < xSteps; y++ {
			grid[x][y] = y0.BlendLuv(x1[x], float64(y)/float64(xSteps)).Hex()
		}
	}

	return grid
}

// ShortenAddress returns a shortened address styled with colors.
func ShortenAddress(address *common.Address) string {
	return fmt.Sprintf(
		"0x%s…%s",
		fmt.Sprintf("%0.2x%0.2x", address.Bytes()[0], address.Bytes()[1]),
		fmt.Sprintf("%0.2x%0.2x", address.Bytes()[len(address.Bytes())-2], address.Bytes()[len(address.Bytes())-1]),
	)
}

// ShortenAddressStyled returns a shortened address styled with colors.
func ShortenAddressStyled(address *common.Address, style lipgloss.Style) string {
	return fmt.Sprintf(
		"%s%s%s%s",
		style.Faint(true).Render("0x"),
		style.Faint(false).Render(fmt.Sprintf("%0.2x%0.2x", address.Bytes()[0], address.Bytes()[1])),
		style.Faint(true).Render("…"),
		style.Faint(false).Render(fmt.Sprintf("%0.2x%0.2x", address.Bytes()[len(address.Bytes())-2], address.Bytes()[len(address.Bytes())-1])),
	)
}

// GenerateColorWithSeed generates a color based on the given seed.
func GenerateColorWithSeed(seed int64) lipgloss.Color {
	rand.Seed(seed)

	//nolint:gosec
	r := rand.Intn(256)
	//nolint:gosec
	g := rand.Intn(256)
	//nolint:gosec
	b := rand.Intn(256)

	color := lipgloss.Color(fmt.Sprintf("#%02x%02x%02x", r, g, b))

	return color
}

func CreateTrendIndicator(before float64, now float64) string {
	var trendIndicator string

	cUp := lipgloss.Color("#99CC99")
	cDown := lipgloss.Color("#CC9999")
	cSteady := lipgloss.Color("#CCCCCC")

	before = toFixed(before, 3)
	now = toFixed(now, 3)

	if now > 0.0 {
		// if before < now {
		// 	trendIndicator = lipgloss.NewStyle().Foreground(cUp).Render("↑")
		// } else if before > now {
		// 	trendIndicator = lipgloss.NewStyle().Foreground(cDown).Render("↓")
		// } else {
		// 	trendIndicator = lipgloss.NewStyle().Foreground(cSteady).Render("~")
		// }
		switch {
		case before < now:
			trendIndicator = lipgloss.NewStyle().Foreground(cUp).Render("↑")
		case before > now:
			trendIndicator = lipgloss.NewStyle().Foreground(cDown).Render("↓")
		default:
			trendIndicator = lipgloss.NewStyle().Foreground(cSteady).Render("~")
		}
	} else {
		trendIndicator = lipgloss.NewStyle().Foreground(cSteady).Faint(true).Render("⊗")
	}

	return trendIndicator
}

// round and toFixed from https://stackoverflow.com/a/29786394/13180763
func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))

	return float64(round(num*output)) / output
}

// TerminalLink formats a link for the terminal using ANSI codes.
func TerminalLink(params ...string) string {
	var text string

	url := params[0]

	if len(params) >= 2 {
		text = params[1]
	} else {
		text = url
	}

	return fmt.Sprintf("\033]8;;%s\033\\%s\033]8;;\033\\", url, text)
}

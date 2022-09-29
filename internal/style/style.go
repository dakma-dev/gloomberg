package style

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"os"
	"strings"

	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/term"
)

var (
	Pink       = lipgloss.AdaptiveColor{Light: "#FF44DD", Dark: "#FF0099"}
	OwnerGreen = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	Subtle     = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	darkGray   = lipgloss.Color("#333")
	darkerGray = lipgloss.Color("#222")

	OpenseaToneBlue      = lipgloss.Color("#5f7699")
	TrendGreenStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#66CC66"))
	TrendLightGreenStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#667066"))
	TrendRedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF6666"))
	TrendLightRedStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#806666"))
	AlmostWhiteStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#EEE"))
	DarkWhiteStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#DDD"))
	VeryLightGrayStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#BBB"))
	LightGrayStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#999"))
	GrayStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color("#666"))
	DarkGrayStyle        = lipgloss.NewStyle().Foreground(darkGray)
	DarkerGrayStyle      = lipgloss.NewStyle().Foreground(darkerGray)
	BoldStyle            = lipgloss.NewStyle().Bold(true)
	PinkBoldStyle        = BoldStyle.Copy().Foreground(Pink)
	OwnerGreenBoldStyle  = BoldStyle.Copy().Foreground(OwnerGreen)
	GrayBoldStyle        = BoldStyle.Copy().Foreground(GrayStyle.GetForeground())
	Sharrow              = BoldStyle.Copy().SetString("→")
	DividerArrowRight    = GrayBoldStyle.SetString("→")
	DividerArrowLeft     = GrayBoldStyle.SetString("←")

	// darkestGray          = lipgloss.Color("#111")
	// DarkestGrayStyle     = lipgloss.NewStyle().Foreground(darkestGray)

	// PinkStyle            = lipgloss.NewStyle().Foreground(pink)
	// GrayFaintStyle       = GrayStyle.Copy().Faint(true)

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

//	SetString("•").
//	Padding(0, 0).
//	Foreground(pink)
//
// var Divider = baseDivider.String()
func GetHeader(version string) string {
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

	width, _, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		gbl.Log.Error(err)
	}

	return lipgloss.NewStyle().PaddingBottom(3).Width(width - 4).Align(lipgloss.Center).Render(header.String())
}

func GetPriceShadeColor(txValue float64) lipgloss.Color {
	var priceColor lipgloss.Color

	switch {
	case txValue >= 0:
		priceColor = "#333333"
	case txValue >= 0.25:
		priceColor = ShadesPink[0]
	case txValue >= 0.5:
		priceColor = ShadesPink[1]
	case txValue >= 0.75:
		priceColor = ShadesPink[2]
	case txValue >= 1.0:
		priceColor = ShadesPink[3]
	case txValue >= 2.0:
		priceColor = ShadesPink[5]
	case txValue >= 3.0:
		priceColor = ShadesPink[6]
	case txValue >= 5.0:
		priceColor = ShadesPink[9]
	}

	return priceColor
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

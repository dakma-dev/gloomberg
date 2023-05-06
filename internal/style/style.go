package style

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"os"
	"strings"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/term"
)

var (
	Pink        = lipgloss.AdaptiveColor{Light: "#FF44DD", Dark: "#FF0099"}
	Subtle      = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	DarkGray    = lipgloss.Color("#333")
	darkerGray  = lipgloss.Color("#222")
	darkestGray = lipgloss.Color("#111")

	OpenseaToneBlue      = lipgloss.Color("#5f7699")
	BlurOrange           = lipgloss.Color("#FF8700")
	TrendGreenStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#66CC66"))
	TrendLightGreenStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#667066"))
	TrendRedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF6666"))
	TrendLightRedStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#806666"))
	AlmostWhiteStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#eeeeee"))
	DarkWhiteStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#dddddd"))
	VeryLightGrayStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#bbbbbb"))
	LightGrayStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#999999"))
	Gray5Style           = lipgloss.NewStyle().Foreground(lipgloss.Color("#555"))
	Gray7Style           = lipgloss.NewStyle().Foreground(lipgloss.Color("#777"))
	Gray8Style           = lipgloss.NewStyle().Foreground(lipgloss.Color("#888"))
	GrayStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color("#666666"))
	DarkGrayStyle        = lipgloss.NewStyle().Foreground(DarkGray)
	DarkerGrayStyle      = lipgloss.NewStyle().Foreground(darkerGray)
	DarkestGrayStyle     = lipgloss.NewStyle().Foreground(darkestGray)
	BoldStyle            = lipgloss.NewStyle().Bold(true)
	PinkBoldStyle        = BoldStyle.Copy().Foreground(Pink)
	GrayBoldStyle        = BoldStyle.Copy().Foreground(GrayStyle.GetForeground())
	Sharrow              = BoldStyle.Copy().SetString("→")
	DividerArrowRight    = LightGrayStyle.Copy().Bold(true).SetString("→")
	DividerArrowLeft     = GrayBoldStyle.SetString("←")

	// darkestGray          = lipgloss.Color("#111")
	// DarkestGrayStyle     = lipgloss.NewStyle().Foreground(darkestGray).

	// PinkStyle            = lipgloss.NewStyle().Foreground(pink)
	// GrayFaintStyle       = GrayStyle.Copy().Faint(true).

	// borders
	// noBorderStyle              = lipgloss.NewStyle().BorderTop(false).BorderBottom(false).BorderLeft(false).BorderRight(false)
	// baseBorderStyle s= lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(darkGray).PaddingRight(2)
	// topBorderStyle             = baseBorderStyle.Copy().BorderTop(true).BorderBottom(false)
	// bottomBorderStyle          = baseBorderStyle.Copy().BorderTop(false).BorderBottom(true)
	// topBottomBorderStyle       = baseBorderStyle.Copy().BorderTop(true).BorderBottom(true)
	// topBottomHiddenBorderStyle = baseBorderStyle.Copy().BorderStyle(lipgloss.HiddenBorder()).BorderTop(true).BorderBottom(true).
)

var ShadesPink = []lipgloss.Color{
	"#fff1f6",
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

// Bold returns a bold string.
func Bold(str string) string {
	return BoldStyle.Render(str)
}

//	SetString("•").
//	Padding(0, 0).
//	Foreground(pink)
//
// var Divider = baseDivider.String().
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

func GetBuyDiffShadeColor(priceDiff float64) lipgloss.Color {
	var priceColor lipgloss.Color

	switch {
	case priceDiff <= 0.03:
		priceColor = ShadesPink[9]
	case priceDiff <= 0.05:
		priceColor = ShadesPink[8]
	case priceDiff <= 0.1:
		priceColor = ShadesPink[4]
	case priceDiff <= 0.25:
		priceColor = ShadesPink[2]
	case priceDiff <= 0.5:
		priceColor = ShadesPink[1]
	default:
		priceColor = "#dddddd"
	}

	return priceColor
}

func GetPriceShadeColor(txValue float64) lipgloss.Color {
	var priceColor lipgloss.Color

	switch {
	case txValue >= 10.0:
		priceColor = ShadesPink[9]
	case txValue >= 5.0:
		priceColor = ShadesPink[8]
	case txValue >= 2.0:
		priceColor = ShadesPink[6]
	case txValue >= 1.0:
		priceColor = ShadesPink[5]
	case txValue >= 0.5:
		priceColor = ShadesPink[4]
	case txValue >= 0.25:
		priceColor = ShadesPink[3]
	case txValue >= 0.1:
		priceColor = ShadesPink[2]
	case txValue >= 0.075:
		priceColor = ShadesPink[1]
	case txValue >= 0.02:
		priceColor = ShadesPink[0]
	default:
		priceColor = "#333333"
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
	// gray out zero address
	if *address == internal.ZeroAddress {
		gray := DarkGrayStyle.Copy().Faint(false).Render
		darkGray := DarkGrayStyle.Copy().Faint(true).Render

		return fmt.Sprint(
			darkGray("0x"),
			gray(fmt.Sprintf("%0.2x", address.Bytes()[0])),
			darkGray("…"),
			gray(fmt.Sprintf("%0.2x", address.Bytes()[len(address.Bytes())-1])),
		)
	}

	return fmt.Sprint(
		style.Faint(true).Render("0x"),
		style.Faint(false).Render(fmt.Sprintf("%0.2x%0.2x", address.Bytes()[0], address.Bytes()[1])),
		style.Faint(true).Render("…"),
		style.Faint(false).Render(fmt.Sprintf("%0.2x%0.2x", address.Bytes()[len(address.Bytes())-2], address.Bytes()[len(address.Bytes())-1])),
	)
}

func EnforceMinLength(str string, minLength int) string {
	if len(str) <= minLength {
		r := []rune(str)
		spacePlaceHolder := make([]rune, minLength-len(str))

		for i := 0; i < len(spacePlaceHolder); i++ {
			spacePlaceHolder[i] = ' '
		}

		return string(append(r, spacePlaceHolder...))
	}

	return str
}

func ShortenCollectionName(collectionName string, numItems int) string {
	maxLength := 25
	if numItems > 1 {
		maxLength -= 3
	}

	collectionName = EnforceMinLength(collectionName, maxLength)

	if len(collectionName) > maxLength {
		return fmt.Sprintf(
			"%s%s",
			collectionName[:maxLength-3],
			"...",
		)
	}

	return collectionName
}

// GenerateColorWithSeed generates a color based on the given seed.
func GenerateColorWithSeed(seed int64) lipgloss.Color {
	rng := rand.New(rand.NewSource(seed)) //nolint:gosec

	return lipgloss.Color(fmt.Sprintf("#%02x%02x%02x", rng.Intn(256), rng.Intn(256), rng.Intn(256)))
}

func CreateTrendIndicator(before float64, now float64) lipgloss.Style {
	var trendIndicatorStyle lipgloss.Style

	cUp := lipgloss.Color("#99CC99")
	cDown := lipgloss.Color("#CC9999")
	cSteady := lipgloss.Color("#555")

	before = toFixed(before, 3)
	now = toFixed(now, 3)

	if now > 0.0 {
		switch {
		case before < now:
			trendIndicatorStyle = lipgloss.NewStyle().Foreground(cUp).SetString("↑")
		case before > now:
			trendIndicatorStyle = lipgloss.NewStyle().Foreground(cDown).SetString("↓")
		default:
			trendIndicatorStyle = lipgloss.NewStyle().Foreground(cSteady).SetString("~")
		}
	} else {
		trendIndicatorStyle = lipgloss.NewStyle().Foreground(cSteady).Faint(true).SetString("⊗")
	}

	return trendIndicatorStyle
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

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

	LightGrayForeground = lipgloss.Color("#BBB")

	OpenSea         = lipgloss.NewStyle().Foreground(OpenseaToneBlue)
	Blur            = lipgloss.NewStyle().Foreground(BlurOrange).Render
	BoldAlmostWhite = AlmostWhiteStyle.Copy().Bold(true).Render

	OpenseaToneBlue      = lipgloss.Color("#5f7699")
	WebUIColor           = lipgloss.Color("#662288")
	BlurOrange           = lipgloss.Color("#FF8700")
	TrendGreenStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#66CC66"))
	TrendLightGreenStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#77A077"))
	TrendRedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF6666"))
	TrendLightRedStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#997777"))
	ReddishPurple        = lipgloss.NewStyle().Foreground(lipgloss.Color("#9F2B68"))
	PurplePower          = lipgloss.NewStyle().Foreground(lipgloss.Color("#5D3FD3"))
	AlmostWhiteStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#eeeeee"))
	DarkWhiteStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#dddddd"))
	VeryLightGrayStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#bbbbbb"))
	LightGray            = lipgloss.Color("#999999")
	LightGrayStyle       = lipgloss.NewStyle().Foreground(LightGray)
	Gray4Style           = lipgloss.NewStyle().Foreground(lipgloss.Color("#444"))
	Gray5Style           = lipgloss.NewStyle().Foreground(lipgloss.Color("#555"))
	Gray6Style           = lipgloss.NewStyle().Foreground(lipgloss.Color("#666"))
	Gray7                = lipgloss.Color("#777")
	Gray7Style           = lipgloss.NewStyle().Foreground(Gray7)
	Gray8Style           = lipgloss.NewStyle().Foreground(lipgloss.Color("#888"))
	GrayStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color("#666666"))
	DarkGrayStyle        = lipgloss.NewStyle().Foreground(DarkGray)
	DarkerGrayStyle      = lipgloss.NewStyle().Foreground(darkerGray)
	DarkestGrayStyle     = lipgloss.NewStyle().Foreground(darkestGray)
	BoldStyle            = lipgloss.NewStyle().Bold(true)
	PinkBoldStyle        = BoldStyle.Copy().Foreground(Pink)
	GrayBoldStyle        = BoldStyle.Copy().Foreground(GrayStyle.GetForeground())
	Sharrow              = lipgloss.NewStyle().SetString("→")
	DividerArrowRight    = LightGrayStyle.Copy().SetString("→")
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

// var ShadesOfGray = []lipgloss.Color{
// 	lipgloss.Color("#111"),
// 	lipgloss.Color("#222"),
// 	lipgloss.Color("#333"),
// 	lipgloss.Color("#444"),
// 	lipgloss.Color("#555"),
// 	lipgloss.Color("#666"),
// 	lipgloss.Color("#777"),
// 	lipgloss.Color("#888"),
// 	lipgloss.Color("#999"),
// 	lipgloss.Color("#aaa"),
// 	lipgloss.Color("#bbb"),
// 	lipgloss.Color("#ccc"),
// 	lipgloss.Color("#ddd"),
// 	lipgloss.Color("#eee"),
// 	lipgloss.Color("#fff"),
// }

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

	headerBaseStyle := lipgloss.NewStyle().Foreground(headerColor)
	headerStyle := headerBaseStyle.Copy().Padding(2, 0, 1, 0)
	headerSeparatorStyle := headerBaseStyle // .Copy().Bold(true)
	subHeaderStyle := DarkGrayStyle.Copy()

	header.WriteString(headerStyle.Render(headerLogo) + "\n")

	subHeader := strings.Builder{}
	subHeader.WriteString(headerBaseStyle.Copy().Bold(true).Render("·"))
	subHeader.WriteString(" " + DarkGrayStyle.Render("gloomberg"))
	subHeader.WriteString(" " + lipgloss.NewStyle().Foreground(lipgloss.Color("#444444")).Render(version))
	subHeader.WriteString(" " + headerSeparatorStyle.Render("|"))
	subHeader.WriteString(" " + DarkGrayStyle.Render("github.com/benleb/gloomberg"))
	subHeader.WriteString(" " + headerSeparatorStyle.Render("·"))

	header.WriteString(subHeaderStyle.Render(subHeader.String()))

	width, _, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		gbl.Log.Error(err)
	}

	return lipgloss.NewStyle().PaddingBottom(3).Width(width - 4).Align(lipgloss.Center).Render(header.String())
}

func GetSmallHeader(version string) string {
	randColorID, _ := crand.Int(crand.Reader, big.NewInt(int64(len(PaletteRLD))))
	randHeaderID, _ := crand.Int(crand.Reader, big.NewInt(int64(len(smallHeaders))))

	headerLogo := smallHeaders[randHeaderID.Int64()]
	headerColor := PaletteRLD[randColorID.Int64()]

	header := strings.Builder{}

	stupidStaticPaddingLeft := 11

	headerBaseStyle := lipgloss.NewStyle().Foreground(headerColor)
	headerStyle := headerBaseStyle.Copy().Padding(2, 0, 1, stupidStaticPaddingLeft)
	headerSeparatorStyle := headerBaseStyle // .Copy().Bold(true)
	subHeaderStyle := DarkGrayStyle.Copy().Padding(0, 0, 0, stupidStaticPaddingLeft-5)

	header.WriteString(headerStyle.Render(headerLogo) + "\n")

	subHeader := strings.Builder{}
	subHeader.WriteString(headerBaseStyle.Copy().Bold(true).Render("·"))
	subHeader.WriteString(" " + DarkGrayStyle.Render("gloomberg"))
	subHeader.WriteString(" " + lipgloss.NewStyle().Foreground(lipgloss.Color("#444444")).Render(version))
	subHeader.WriteString(" " + headerSeparatorStyle.Render("|"))
	subHeader.WriteString(" " + DarkGrayStyle.Render("github.com/benleb/gloomberg"))
	subHeader.WriteString(" " + headerSeparatorStyle.Render("·"))

	header.WriteString(subHeaderStyle.Render(subHeader.String()))

	width, _, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		// gbl.Log.Error(err)

		return lipgloss.NewStyle().Width(80 - 4).Render(header.String())
	}

	// return lipgloss.NewStyle().PaddingBottom(3).Width(width - 4).Align(lipgloss.Center).Render(header.String())
	return lipgloss.NewStyle().Width(width - 4).Render(header.String())
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

// ShortenAdressPTR returns a shortened address styled with colors.
func ShortenAdressPTR(address *common.Address) string {
	return fmt.Sprintf(
		"0x%s…%s",
		fmt.Sprintf("%0.2x%0.2x", address.Bytes()[0], address.Bytes()[1]),
		fmt.Sprintf("%0.2x%0.2x", address.Bytes()[len(address.Bytes())-2], address.Bytes()[len(address.Bytes())-1]),
	)
}

func ShortenAddress(address common.Address) string {
	return fmt.Sprintf(
		"0x%s…%s",
		fmt.Sprintf("%0.2x%0.2x", address.Bytes()[0], address.Bytes()[1]),
		fmt.Sprintf("%0.2x%0.2x", address.Bytes()[len(address.Bytes())-2], address.Bytes()[len(address.Bytes())-1]),
	)
}

func ShortenedTokenIDStyled(tokenID *big.Int, primaryStyle lipgloss.Style, secondaryStyle lipgloss.Style) string {
	shortened := false

	// shorten token id if it's too long
	if tokenID.Cmp(big.NewInt(999_999)) > 0 {
		tokenID = big.NewInt(tokenID.Int64() % 10000)
		shortened = true
	}

	// token id
	prefix := secondaryStyle.Render("#")
	id := primaryStyle.Render(fmt.Sprint(tokenID))

	if shortened {
		id += secondaryStyle.Render("…")
	}

	return prefix + id
}

func FormatAddress(address *common.Address) string {
	style := lipgloss.NewStyle().Foreground(GenerateColorWithSeed(address.Big().Int64()))

	return ShortenAddressStyled(address, style)
}

// GenerateColors generates two colors based on contract address of the collection.
func GenerateAddressColors(address *common.Address) (lipgloss.Color, lipgloss.Color) {
	return GenerateColorWithSeed(address.Big().Int64()), GenerateColorWithSeed(address.Big().Int64() ^ 2)
}

func GenerateAddressStyles(address *common.Address) (lipgloss.Style, lipgloss.Style) {
	primaryColor, secondaryColor := GenerateAddressColors(address)

	return lipgloss.NewStyle().Foreground(primaryColor), lipgloss.NewStyle().Foreground(secondaryColor)
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

// ShortenHashsStyled returns a shortened hash styled with colors.
func ShortenHashStyled(txHash common.Hash) string {
	style := lipgloss.NewStyle().Foreground(GenerateColorWithSeed(txHash.Big().Int64()))

	// gray out zero txHash
	if txHash == internal.ZeroHash {
		gray := DarkGrayStyle.Copy().Faint(false).Render
		darkGray := DarkGrayStyle.Copy().Faint(true).Render

		return fmt.Sprint(
			darkGray("0x"),
			gray(fmt.Sprintf("%0.2x", txHash.Bytes()[0])),
			darkGray("…"),
			gray(fmt.Sprintf("%0.2x", txHash.Bytes()[len(txHash.Bytes())-1])),
		)
	}

	return fmt.Sprint(
		style.Faint(true).Render("0x"),
		style.Faint(false).Render(fmt.Sprintf("%0.2x%0.2x", txHash.Bytes()[0], txHash.Bytes()[1])),
		style.Faint(true).Render("…"),
		style.Faint(false).Render(fmt.Sprintf("%0.2x%0.2x", txHash.Bytes()[len(txHash.Bytes())-2], txHash.Bytes()[len(txHash.Bytes())-1])),
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

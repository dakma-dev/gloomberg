package gloomberg

import (
	"fmt"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

// Pr prints messages from gloomberg to the terminal.
// func (gb *Gloomberg) Pr(message string) {.
func Pr(message string) {
	printToTerminal("🧃", style.Gray5Style.Render("gb"), message) // style.PinkBoldStyle.Render("・"))
}

// Prf formats and prints messages from gloomberg to the terminal.
// func (gb *Gloomberg) Prf(format string, a ...any) {.
func Prf(format string, a ...any) {
	Pr(fmt.Sprintf(format, a...))
}

// func (gb *Gloomberg) PrWarn(message string) {.

// func (gb *Gloomberg) PrWithKeywordAndIcon(icon string, keyword string, message string) {.
func PrWithKeywordAndIcon(icon string, keyword string, message string) {
	printToTerminal(icon, keyword, message)
}

// func (gb *Gloomberg) PrMod(mod string, message string) {.
func PrMod(mod string, message string) {
	printConfigurations := make(map[string]*printConfig)

	for idx, config := range predefinedPrintConfigurations {
		printConfiguration := predefinedPrintConfigurations[idx]

		for _, keyword := range config.Keywords {
			printConfigurations[keyword] = &printConfiguration
		}
	}

	prConfig, ok := printConfigurations[mod]
	if !ok {
		log.Warnf("no print configuration for module %s | message: %s", mod, message)

		return
	}

	icon := prConfig.Icon

	color := style.DarkGray
	if prConfig.Color != "" {
		color = prConfig.Color
	}

	tag := lipgloss.NewStyle().Foreground(color).Render(mod)

	printToTerminal(icon, tag, message)
}

// PrModf formats and prints messages from gloomberg to the terminal.
// func (gb *Gloomberg) PrModf(mod string, format string, a ...any) {.
func PrModf(mod string, format string, a ...any) {
	PrMod(mod, fmt.Sprintf(format, a...))
}

// PrVMod prints messages from gloomberg to the terminal if verbose mode is enabled.
// func (gb *Gloomberg) PrVMod(mod string, message string) {.

// PrVModf formats and prints messages from gloomberg to the terminal if verbose mode is enabled.
// func (gb *Gloomberg) PrVModf(mod string, format string, a ...any) {.

// PrDMod prints messages from gloomberg to the terminal if debug mode is enabled.
// func (gb *Gloomberg) PrDMod(mod string, message string) {.

// PrDModf formats and prints messages from gloomberg to the terminal if debug mode is enabled.
// func (gb *Gloomberg) PrDModf(mod string, format string, a ...any) {.

// func (gb *Gloomberg) printToTerminal(icon string, keyword string, message string) {.
func printToTerminal(icon string, keyword string, message string) {
	if message == "" {
		return
	}

	// WEN...??
	now := time.Now()
	currentTime := now.Format("15:04:05")

	out := strings.Builder{}
	out.WriteString(style.DarkGrayStyle.Render("|"))
	out.WriteString(style.Gray4Style.Render(currentTime))
	out.WriteString(" " + icon)
	out.WriteString(" " + lipgloss.NewStyle().Width(6).Align(lipgloss.Right).Render(keyword))
	out.WriteString("  " + message)

	// gb.In.PrintToTerminal <- out.String()
	TerminalPrinterQueue <- out.String()
}

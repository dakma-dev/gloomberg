// copyright © 2022 benleb <git@benleb.de>
package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/benleb/gloomberg/cmd"
	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/log"
	"github.com/muesli/termenv"
)

var version = "dev"

func main() {
	internal.GloombergVersion = version

	// cmd.BuildDate = buildDate
	// cmd.BuiltBy = builtBy

	// save default foreground color and change it to light gray
	defaultForeground := termenv.ForegroundColor()
	lipgloss.SetColorProfile(termenv.TrueColor)
	termenv.DefaultOutput().SetForegroundColor(termenv.RGBColor(style.LightGrayForeground))

	// signal handler channel
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		sig := <-c

		// ctrl+c handler
		log.Debug(fmt.Sprintf("Got %s signal. Aborting...\n", sig))

		// reset/restore default foreground color
		termenv.DefaultOutput().SetForegroundColor(defaultForeground)
		termenv.DefaultOutput().Reset()

		os.Exit(0)
	}()

	cmd.Execute()
}

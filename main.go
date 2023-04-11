// copyright © 2022 benleb <git@benleb.de>
package main

import (
	"github.com/benleb/gloomberg/cmd"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

var version = "dev" // commit    = "none"
// buildDate = "unknown"
// builtBy   = "unknown"

func main() {
	cmd.Version = version

	// cmd.BuildDate = buildDate
	// cmd.BuiltBy = builtBy

	lipgloss.SetColorProfile(termenv.TrueColor)

	cmd.Execute()
}

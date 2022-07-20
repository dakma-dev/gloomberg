// copyright © 2022 benleb <git@benleb.de>
package main

import (
	"fmt"

	"github.com/benleb/gloomberg/cmd"
)

var (
	version   = "dev"
	commit    = "none"
	buildDate = "unknown"
	builtBy   = "unknown"
)

func main() {
	cmd.Version = version
	cmd.Commit = commit
	cmd.BuildDate = buildDate
	cmd.BuiltBy = builtBy

	fmt.Printf("\n\ngloomberg %s (%s) | %s / %s\n\n", version, commit, buildDate, builtBy)

	cmd.Execute()
}

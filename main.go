// copyright © 2022 benleb <git@benleb.de>
package main

import (
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

	cmd.Execute()
}

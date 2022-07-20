package style

import (
	"time"

	"github.com/theckman/yacspin"
)

// spinnerBaseConfig is the base configuration for the spinners.
var spinnerBaseConfig = yacspin.Config{
	Frequency:       100 * time.Millisecond,
	CharSet:         yacspin.CharSets[11],
	SuffixAutoColon: true,
	Suffix:          " ",
	Message:         "",
	StopCharacter:   "✓",
	Colors:          []string{"fgMagenta"},
	StopColors:      []string{"fgGreen"},
}

// GetSpinner returns a new spinner with default configuration.
func GetSpinner(message string) *yacspin.Spinner {
	spinnerConfig := spinnerBaseConfig
	spinnerConfig.Message = message
	spinnerConfig.StopMessage = message
	spinner, _ := yacspin.New(spinnerConfig)

	return spinner
}

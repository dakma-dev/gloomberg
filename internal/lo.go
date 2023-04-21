package internal

import (
	"github.com/charmbracelet/log"
)

// disabled as charmbracelet/log 0.2.1 is required for this to work
// but that version has a bug that causes gloomberg to randomly print ansi sequences...^^
var (
	// BaseLogger = log.NewWithOptions(os.Stderr, log.Options{
	// 	ReportCaller:    true,
	// 	ReportTimestamp: true,
	// 	TimeFormat:      time.TimeOnly,
	// 	Prefix:          "㏒",
	// })

	BaseLogger = log.New()

	// BasePrinter = log.NewWithOptions(os.Stderr, log.Options{
	// 	ReportCaller:    false,
	// 	ReportTimestamp: false,
	// })

	BasePrinter = log.New()
)

// func FileLogger(file *os.File) *log.Logger {
// 	return log.NewWithOptions(file, log.Options{
// 		ReportCaller:    true,
// 		ReportTimestamp: true,
// 		TimeFormat:      time.DateTime,
// 	})
// }

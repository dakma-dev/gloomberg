package gbl

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

func GetSugaredLogger() *zap.SugaredLogger {
	if Log == nil {
		Log = initSugaredLogger()
	}

	return Log
}

func GetLogsLogger() *zap.SugaredLogger {
	outputPaths := []string{"/tmp/gloomberg-logs.log"}

	config := zap.NewProductionConfig()

	// config = zap.Config{
	// 	Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
	// 	Encoding:         "console",
	// 	OutputPaths:      outputPaths,
	// 	ErrorOutputPaths: outputPaths,
	// 	EncoderConfig: zapcore.EncoderConfig{
	// 		EncodeTime:       zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"), // zapcore.RFC3339TimeEncoder
	// 		MessageKey:       "message",
	// 		LevelKey:         "level",
	// 		EncodeLevel:      zapcore.LowercaseColorLevelEncoder,
	// 		ConsoleSeparator: " ",
	// 	},
	// }

	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	config.Encoding = "console"
	config.OutputPaths = outputPaths
	config.ErrorOutputPaths = outputPaths
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // zapcore.RFC3339TimeEncoder
	config.EncoderConfig.MessageKey = "message"
	config.EncoderConfig.LevelKey = "level"
	config.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	config.EncoderConfig.ConsoleSeparator = " "

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	logsLogger := logger.Sugar()

	// flushes buffer, if any
	err = logger.Sync()

	if err != nil {
		Log.Errorf("flushing log buffer failed: %s", err)

		return nil
	}

	return logsLogger
}

// initSugaredLogger initializes the global logger.
func initSugaredLogger() *zap.SugaredLogger {
	var outputPaths []string
	// outputPaths = append(outputPaths, "/tmp/gloomberg.log")
	// outputPaths = append(outputPaths, "stdout")

	if logFile := viper.GetString("log.log_file"); logFile != "" {
		outputPaths = append(outputPaths, logFile)
	}

	logLevel := zap.WarnLevel

	if viper.GetBool("log.debug") {
		logLevel = zap.DebugLevel
	} else if viper.GetBool("log.verbose") {
		logLevel = zap.InfoLevel
	}

	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(logLevel)
	config.Encoding = "console"
	config.OutputPaths = outputPaths
	config.ErrorOutputPaths = []string{"stderr"}
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // zapcore.RFC3339TimeEncoder
	config.EncoderConfig.MessageKey = "message"
	config.EncoderConfig.LevelKey = "level"
	config.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	config.EncoderConfig.ConsoleSeparator = " "

	// config := GetDefaultConfig("")

	viper.Set("log.level", config.Level)

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	Log = logger.Sugar()

	// flushes buffer, if any
	err = logger.Sync()

	if err != nil {
		Log.Errorf("flushing log buffer failed: %s", err)

		return nil
	}

	return Log
}

package miner

import (
	"github.com/hunterhug/golog/v2"
)

type Level = golog.Level

const (
	DEBUG Level = golog.DebugLevel
	INFO  Level = golog.InfoLevel
	WARN  Level = golog.WarnLevel
	ERROR Level = golog.ErrorLevel
)

// Logger Global logger config for debug
var (
	Logger = golog.New()
)

func init() {
	Logger.InitLogger()
}

// SetLogLevel Set log level
func SetLogLevel(level Level) {
	Logger.SetLevel(level)
	Logger.InitLogger()
}

// Log Return global log
func Log() golog.LoggerInterface {
	return Logger
}

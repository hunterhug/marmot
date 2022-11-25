package miner

import (
	"github.com/hunterhug/golog"
)

type Level = golog.Level

const (
	DEBUG = golog.DebugLevel
	INFO  = golog.InfoLevel
	WARN  = golog.WarnLevel
	ERROR = golog.ErrorLevel
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

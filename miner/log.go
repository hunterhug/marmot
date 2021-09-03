/*
	All right reserved https://github.com/hunterhug/marmot at 2016-2021
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
*/
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

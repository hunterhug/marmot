/*
	All right reserved https://github.com/hunterhug/marmot at 2016-2021
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
*/
package miner

import (
	"github.com/hunterhug/marmot/util/go-logging"
	"os"
)

// Global logger config for debug
var (
	Logger = logging.MustGetLogger("Marmot")

	format = logging.MustStringFormatter(
		"%{color}%{time:2006-01-02 15:04:05.000} %{longpkg}:%{longfunc} [%{level:.5s}]:%{color:reset} %{message}",
	)

	// LevelNames Level name you can refer
	LevelNames = []string{
		"CRITICAL",
		"ERROR",
		"WARNING",
		"NOTICE",
		"INFO",
		"DEBUG",
	}
)

// Init log record
func init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
	logging.SetLevel(logging.INFO, "Marmot")
}

// SetLogLevel Set log level
func SetLogLevel(level string) {
	lvl, _ := logging.LogLevel(level)
	logging.SetLevel(lvl, "Marmot")
}

// Log Return global log
func Log() *logging.Logger {
	return Logger
}

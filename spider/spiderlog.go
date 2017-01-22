package spider

import (
	"github.com/op/go-logging"
	"os"
)

var Logger = logging.MustGetLogger("go_tool_spider")

var format = logging.MustStringFormatter(
	"%{color}%{time:2006-01-02 15:04:05.000} %{longpkg}:%{longfunc} [%{level:.5s}]:%{color:reset} %{message}",
)

// level name you can refer
var LevelNames = []string{
	"CRITICAL",
	"ERROR",
	"WARNING",
	"NOTICE",
	"INFO",
	"DEBUG",
}

// init log record
func init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
	logging.SetLevel(logging.INFO, "go_tool_spider")
}

// set log level
func SetLogLevel(level string) {
	lvl, _ := logging.LogLevel(level)
	logging.SetLevel(lvl, "go_tool_spider")
}

// return global log
func Log() *logging.Logger {
	return Logger
}

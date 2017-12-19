package miner

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/hunterhug/parrot/util"
)

// Wait some secord
func Wait(waittime int) {
	if waittime <= 0 {
		return
	} else {
		// debug
		Logger.Debugf("Wait %d Second.", waittime)
		util.Sleep(waittime)
	}
}

// Header map[string][]string ,can use to copy a http header, so that they are not effect each other
func CopyM(h http.Header) http.Header {
	if h == nil || len(h) == 0 {
		return h
	}
	h2 := make(http.Header, len(h))
	for k, vv := range h {
		vv2 := make([]string, len(vv))
		copy(vv2, vv)
		h2[k] = vv2
	}
	return h2
}

//if a file size small than sizes(KB) ,it will be throw a error
func TooSortSizes(data []byte, sizes float64) error {
	if float64(len(data))/1000 < sizes {
		return errors.New(fmt.Sprintf("FileSize:%d bytes,%d kb < %f kb dead too sort", len(data), len(data)/1000, sizes))
	}
	return nil
}

// Just debug a map
func OutputMaps(info string, args map[string][]string) {
	s := "\n"
	for k, v := range args {
		s = s + fmt.Sprintf("%-25s| %-6s\n", k, strings.Join(v, "||"))
	}
	Logger.Debugf("[GoWorker] %s", s)
}

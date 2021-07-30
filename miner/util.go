/*
	All right reserved https://github.com/hunterhug/marmot at 2016-2021
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
*/
package miner

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/hunterhug/marmot/util"
)

// Wait some second
func Wait(waitTime int) {
	if waitTime <= 0 {
		return
	} else {
		Logger.Debugf("Wait %d Second.", waitTime)
		util.Sleep(waitTime)
	}
}

// CopyM Header map[string][]string ,can use to copy a http header, so that they are not effect each other
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

// TooSortSizes if a file size small than sizes(KB) ,it will be throw a error
func TooSortSizes(data []byte, sizes float64) error {
	if float64(len(data))/1000 < sizes {
		return errors.New(fmt.Sprintf("FileSize:%d bytes,%d kb < %f kb dead too sort", len(data), len(data)/1000, sizes))
	}
	return nil
}

// OutputMaps Just debug a map
func OutputMaps(info string, args map[string][]string) {
	s := "\n"
	for k, v := range args {
		s = s + fmt.Sprintf("%-25s| %-6s\n", k, strings.Join(v, "||"))
	}
	Logger.Debugf("[GoWorker] %s-%s", info, s)
}

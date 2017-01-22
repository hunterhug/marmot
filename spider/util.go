/*
 * Created by 一只尼玛 on 2016/8/13.
 * 功能：
 *	spider tool
 */
package spider

import (
	"errors"
	"fmt"
	"github.com/hunterhug/go_tool/util"
	"net/http"
)

// Wait some secord
func Wait(waittime int) {
	if waittime <= 0 {
		return
	} else {
		// debug
		Logger.Debugf("Stop %d Second～～", waittime)
		util.Sleep(waittime)
	}
}

//Header map[string][]string ,can use to copy a http header, so that they are not effect each other
func CopyM(h http.Header) http.Header {
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

//just debug a map
func OutputMaps(info string, args map[string][]string) {
	Logger.Debugf("%s:%v", info, args)
}

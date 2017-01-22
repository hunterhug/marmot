//常量包
package spider

import "net/http"

const (
	//暂停时间 default wait time
	WaitTime = 5
)

var (
	//浏览器头部 default header ua
	FoxfireLinux = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:46.0) Gecko/20100101 Firefox/46.0"
	SpiderHeader = map[string][]string{
		"User-Agent": {
			FoxfireLinux,
		},
	}
	// http get and post No timeout
	DefaultTimeOut = 0
)

func SetGlobalTimeout(num int) {
	DefaultTimeOut = num
}

// usually a header has ua,host and refer
func NewHeader(ua interface{}, host string, refer interface{}) map[string][]string {
	if ua == nil {
		ua = FoxfireLinux
	}
	if refer == nil {
		h := map[string][]string{
			"User-Agent": {
				ua.(string),
			},
			"Host": {
				host,
			},
		}
		return h
	}
	h := map[string][]string{
		"User-Agent": {
			ua.(string),
		},
		"Host": {
			host,
		},
		"Referer": {
			refer.(string),
		},
	}
	return h
}

//merge Cookie，后来的覆盖前来的
func MergeCookie(before []*http.Cookie, after []*http.Cookie) []*http.Cookie {
	cs := make(map[string]*http.Cookie)

	for _, b := range before {
		cs[b.Name] = b
	}

	for _, a := range after {
		if a.Value != "" {
			cs[a.Name] = a
		}
	}

	res := make([]*http.Cookie, 0, len(cs))

	for _, q := range cs {
		res = append(res, q)

	}

	return res

}

// clone a header
func CloneHeader(h map[string][]string) map[string][]string {
	if h == nil || len(h)==0{
		h = SpiderHeader
	}
	return CopyM(h)
}

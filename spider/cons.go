/*
Copyright 2017 by GoSpider author.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package spider

import "net/http"

const (
	// 暂停时间 default wait time
	WaitTime = 5

	// HTTP方法
	POST     = "POST"
	POSTJSON = "POSTJSON"
	POSTXML  = "POSTXML"
	POSTFILE = "POSTFILE"

	// 实现了!
	PUT     = "PUT"
	PUTJSON = "PUTJSON"
	PUTXML  = "PUTXML"
	PUTFILE = "PUTFILE"

	DELETE = "DELETE"
	GET    = "GET"
	OTHER  = "OTHER"

	CRITICAL = "CRITICAL"
	ERROR    = "ERROR"
	WARNING  = "WARNING"
	NOTICE   = "NOTICE"
	INFO     = "INFO"
	DEBUG    = "DEBUG"

	HTTPFORMContentType = "application/x-www-form-urlencoded"
	HTTPJSONContentType = "application/json"
	HTTPXMLContentType  = "text/xml"
	HTTPFILEContentType = "multipart/form-data"
)

var (
	// 浏览器头部 default header ua
	// 默认的,取消使用！！
	FoxfireLinux = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:46.0) Gecko/20100101 Firefox/46.0"
	SpiderHeader = map[string][]string{
		"User-Agent": {
			FoxfireLinux,
		},
	}
	// http get and post No timeout
	// 不设置时没有超时时间
	DefaultTimeOut = 0
)

// 超时目前只能这样设置全局
func SetGlobalTimeout(num int) {
	DefaultTimeOut = num
}

// usually a header has ua,host and refer
// 浏览器标志，主机名，来源
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

// merge Cookie，后来的覆盖前来的
// 暂时没有用的
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
// 克隆头部，因为是引用
func CloneHeader(h map[string][]string) map[string][]string {
	if h == nil || len(h) == 0 {
		//h = SpiderHeader
		return map[string][]string{}
	}
	return CopyM(h)
}

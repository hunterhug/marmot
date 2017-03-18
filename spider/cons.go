/*
Copyright 2017 hunterhug/一只尼玛.
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
	if h == nil || len(h) == 0 {
		h = SpiderHeader
	}
	return CopyM(h)
}

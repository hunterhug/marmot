/*
Copyright 2017 by GoSpider author. Email: gdccmcm14@live.com
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

import (
	"github.com/hunterhug/GoTool/util"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type SpiderConfig struct {
	Url    string      // Which url we want
	Method string      // Get/Post method
	Header http.Header // Http header
	Data   url.Values  // Sent by form data
	BData  []byte      // Sent by binary data
	Wait   int         // Wait Time
}

type Spider struct {
	*SpiderConfig
	Preurl        string         // Pre url
	Raw           []byte         // Raw data we get
	UrlStatuscode int            // the last url response code, such as 404
	Client        *http.Client   // Our Client
	Fetchtimes    int            // Url fetch number times
	Errortimes    int            // Url fetch error times
	Ipstring      string         // spider ip, just for user to record their proxy ip, default: localhost
	Request       *http.Request  // Debug
	Response      *http.Response // Debug
	mux           sync.RWMutex   // lock, execute concurrently please use spider Pool!
}

// Java Bean Chain pattern
func (config *SpiderConfig) SetHeader(header http.Header) *SpiderConfig {
	config.Header = header
	return config
}

func (config *SpiderConfig) SetHeaderParm(k, v string) *SpiderConfig {
	config.Header.Set(k, v)
	return config
}

// Set Cookie!
// Cookie 这样设置如果有jar != nil 那么同名cookie会和这个一起发送过去
func (config *SpiderConfig) SetCookie(v string) *SpiderConfig {
	config.SetHeaderParm("Cookie", v)
	return config
}

// Set Cookie by file.
func (config *SpiderConfig) SetCookieByFile(file string) (*SpiderConfig, error) {
	haha, err := util.ReadfromFile(file)
	if err != nil {
		return nil, err
	}
	cookie := string(haha)
	cookie = strings.Replace(cookie, " ", "", -1)
	cookie = strings.Replace(cookie, "\n", "", -1)
	cookie = strings.Replace(cookie, "\r", "", -1)
	sconfig := config.SetCookie(cookie)
	return sconfig, nil
}

func (config *SpiderConfig) SetUa(ua string) *SpiderConfig {
	config.Header.Set("User-Agent", ua)
	return config
}

func (config *SpiderConfig) SetRefer(refer string) *SpiderConfig {
	config.Header.Set("Referer", refer)
	return config
}

func (config *SpiderConfig) SetHost(host string) *SpiderConfig {
	config.Header.Set("Host", host)
	return config
}

// SetUrl, at the same time SetHost
func (config *SpiderConfig) SetUrl(url string) *SpiderConfig {
	config.Url = url
	//https://www.zhihu.com/people/
	temp := strings.Split(url, "//")
	if len(temp) >= 2 {
		config.SetHost(strings.Split(temp[1], "/")[0])
	}
	return config
}

func (config *SpiderConfig) SetMethod(method string) *SpiderConfig {
	temp := GET
	switch strings.ToUpper(method) {
	case GET:
		temp = GET
	case POST:
		temp = POST
	case POSTFILE:
		temp = POSTFILE
	case POSTJSON:
		temp = POSTJSON
	case POSTXML:
		temp = POSTXML
	case PUT:
		temp = PUT
	case PUTFILE:
		temp = PUTFILE
	case PUTJSON:
		temp = PUTJSON
	case PUTXML:
		temp = PUTXML
	case DELETE:
		temp = DELETE
	default:
		temp = OTHER
	}
	config.Method = temp
	return config
}

func (config *SpiderConfig) SetWaitTime(num int) *SpiderConfig {
	if num <= 0 {
		num = 1
	}
	config.Wait = num
	return config
}

func (config *SpiderConfig) SetBData(data []byte) *SpiderConfig {
	config.BData = data
	return config
}

func (config *SpiderConfig) SetForm(form url.Values) *SpiderConfig {
	config.Data = form
	return config
}

func (config *SpiderConfig) SetFormParm(k, v string) *SpiderConfig {
	config.Data.Set(k, v)
	return config
}

// Clear data we sent
func (config *SpiderConfig) Clear() *SpiderConfig {
	config.Data = url.Values{}
	config.BData = []byte{}
	return config
}

// All clear include header
func (config *SpiderConfig) ClearAll() *SpiderConfig {
	config.Header = http.Header{}
	config.Data = url.Values{}
	config.BData = []byte{}
	return config
}

// Clear Cookie
func (config *SpiderConfig) ClearCookie() *SpiderConfig {
	config.Header.Del("Cookie")
	return config
}

// Get Cookies
func (sp *Spider) GetCookies() []*http.Cookie {
	if sp.Response != nil {
		return sp.Response.Cookies()
	} else {
		return []*http.Cookie{}
	}
}

// Deprecated
func (sp *Spider) NewHeader(ua interface{}, host string, refer interface{}) {
}

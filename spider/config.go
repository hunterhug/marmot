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
	"context"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/hunterhug/GoTool/util"
)

type Spider struct {
	Url    string      // Which url we want
	Method string      // Get/Post method
	Header http.Header // Http header
	Data   url.Values  // Sent by form data
	BData  []byte      // Sent by binary data
	Wait   int         // Wait Time
	// In order fast chain func call I put the basic config in above.
	////////////////////////////////////////////////////////
	mux      sync.RWMutex   // lock, execute concurrently please use spider Pool!
	Client   *http.Client   // Our Client
	Request  *http.Request  // Debug
	Response *http.Response // Debug
	Raw      []byte         // Raw data we get
	///////////////////////////////////////////////////////
	// The name below is not so good but has already been used in many project, so bear it.
	Preurl        string // Pre url
	UrlStatuscode int    // the last url response code, such as 404
	Fetchtimes    int    // Url fetch number times
	Errortimes    int    // Url fetch error times
	Ipstring      string // spider ip, just for user to record their proxy ip, default: localhost

	// AOP like Java
	Ctx          context.Context
	BeforeAction func(context.Context, *Spider)
	AfterAction  func(context.Context, *Spider)
}

// Java Bean Chain pattern
func (sp *Spider) SetHeader(header http.Header) *Spider {
	sp.Header = header
	return sp
}

// Default Set!
func SetHeader(header http.Header) *Spider {
	return DefaultSpider.SetHeader(header)
}

func (sp *Spider) SetHeaderParm(k, v string) *Spider {
	sp.Header.Set(k, v)
	return sp
}

func SetHeaderParm(k, v string) *Spider {
	return DefaultSpider.SetHeaderParm(k, v)
}

// Set Cookie!
// Cookie 这样设置如果有jar != nil 那么同名cookie会和这个一起发送过去
func (sp *Spider) SetCookie(v string) *Spider {
	sp.SetHeaderParm("Cookie", v)
	return sp
}

func SetCookie(v string) *Spider {
	return DefaultSpider.SetCookie(v)
}

// Set Cookie by file.
func (sp *Spider) SetCookieByFile(file string) (*Spider, error) {
	haha, err := util.ReadfromFile(file)
	if err != nil {
		return nil, err
	}
	cookie := string(haha)
	cookie = strings.Replace(cookie, " ", "", -1)
	cookie = strings.Replace(cookie, "\n", "", -1)
	cookie = strings.Replace(cookie, "\r", "", -1)
	sconfig := sp.SetCookie(cookie)
	return sconfig, nil
}

func SetCookieByFile(file string) (*Spider, error) {
	return DefaultSpider.SetCookieByFile(file)
}

func (sp *Spider) SetUa(ua string) *Spider {
	sp.Header.Set("User-Agent", ua)
	return sp
}

func SetUa(ua string) *Spider {
	return DefaultSpider.SetUa(ua)
}

func (sp *Spider) SetRefer(refer string) *Spider {
	sp.Header.Set("Referer", refer)
	return sp
}

func SetRefer(refer string) *Spider {
	return DefaultSpider.SetRefer(refer)
}

func (sp *Spider) SetHost(host string) *Spider {
	sp.Header.Set("Host", host)
	return sp
}

// SetUrl, at the same time SetHost
func (sp *Spider) SetUrl(url string) *Spider {
	sp.Url = url
	//https://www.zhihu.com/people/
	temp := strings.Split(url, "//")
	if len(temp) >= 2 {
		sp.SetHost(strings.Split(temp[1], "/")[0])
	}
	return sp
}

func SetUrl(url string) *Spider {
	return DefaultSpider.SetUrl(url)
}

func (sp *Spider) SetMethod(method string) *Spider {
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
	sp.Method = temp
	return sp
}

func SetMethod(method string) *Spider {
	return DefaultSpider.SetMethod(method)
}

func (sp *Spider) SetWaitTime(num int) *Spider {
	if num <= 0 {
		num = 1
	}
	sp.Wait = num
	return sp
}

func SetWaitTime(num int) *Spider {
	return DefaultSpider.SetWaitTime(num)
}

func (sp *Spider) SetBData(data []byte) *Spider {
	sp.BData = data
	return sp
}

func SetBData(data []byte) *Spider {
	return DefaultSpider.SetBData(data)
}

func (sp *Spider) SetForm(form url.Values) *Spider {
	sp.Data = form
	return sp
}

func SetForm(form url.Values) *Spider {
	return DefaultSpider.SetForm(form)
}

func (sp *Spider) SetFormParm(k, v string) *Spider {
	sp.Data.Set(k, v)
	return sp
}

func SetFormParm(k, v string) *Spider {
	return DefaultSpider.SetFormParm(k, v)
}

// Set Context so Action can soft
func (sp *Spider) SetContext(ctx context.Context) *Spider {
	sp.Ctx = ctx
	return sp
}

func SetContext(ctx context.Context) *Spider {
	return DefaultSpider.SetContext(ctx)
}

func (sp *Spider) SetBeforeAction(fc func(context.Context, *Spider)) *Spider {
	sp.BeforeAction = fc
	return sp
}

func SetBeforeAction(fc func(context.Context, *Spider)) *Spider {
	return DefaultSpider.SetBeforeAction(fc)
}

func (sp *Spider) SetAfterAction(fc func(context.Context, *Spider)) *Spider {
	sp.AfterAction = fc
	return sp
}

func SetAfterAction(fc func(context.Context, *Spider)) *Spider {
	return DefaultSpider.SetAfterAction(fc)
}

// Clear data we sent
func (sp *Spider) Clear() *Spider {
	sp.Data = url.Values{}
	sp.BData = []byte{}
	return sp
}

func Clear() *Spider {
	return DefaultSpider.Clear()
}

// All clear include header
func (sp *Spider) ClearAll() *Spider {
	sp.Header = http.Header{}
	sp.Data = url.Values{}
	sp.BData = []byte{}
	return sp
}

func ClearAll() *Spider {
	return DefaultSpider.ClearAll()
}

// Clear Cookie
func (sp *Spider) ClearCookie() *Spider {
	sp.Header.Del("Cookie")
	return sp
}

func ClearCookie() *Spider {
	return DefaultSpider.ClearCookie()
}

// Get Cookies
func (sp *Spider) GetCookies() []*http.Cookie {
	if sp.Response != nil {
		return sp.Response.Cookies()
	} else {
		return []*http.Cookie{}
	}
}

func GetCookies() []*http.Cookie {
	return DefaultSpider.GetCookies()
}

// Deprecated
func (sp *Spider) NewHeader(ua interface{}, host string, refer interface{}) {
}

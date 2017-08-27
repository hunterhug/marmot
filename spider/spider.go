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

import (
	"errors"
	"github.com/hunterhug/GoSpider/util"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

// 全局爬虫
var DefaultSpider *Spider

func init() {
	// 初始化全部浏览器, 可能不应该初始化，应该手动,读取文件(故失效)
	UaInit()

	// 默认爬虫
	sp := new(Spider)
	sp.SpiderConfig = new(SpiderConfig)
	sp.Header = http.Header{}
	sp.Data = url.Values{}
	sp.BData = []byte{}
	sp.Client = Client
	// 全局爬虫使用全局客户端
	DefaultSpider = sp

}

// 获取默认Spider
func GetSpider() *Spider {
	return DefaultSpider
}

type SpiderConfig struct {
	Url    string      // now fetch url 这次要抓取的Url
	Method string      // Get Post 请求方法
	Header http.Header // 请求头部
	Data   url.Values  // post form data 表单字段
	BData  []byte      // binary data 文件上传二进制流
	Wait   int         // sleep time 等待时间
}

// 爬虫结构体
type Spider struct {
	*SpiderConfig
	Preurl        string        // pre url 上一次访问的URL
	Raw           []byte        // 抓取到的二进制流
	UrlStatuscode int           // the last url response code,such as 404 响应状态码
	Client        *http.Client  // 真正客户端
	Fetchtimes    int           // url fetch number times 抓取次数
	Errortimes    int           // error times 失败次数
	Ipstring      string        // spider ip,just for user to record their proxyip 代理IP地址，没有代理默认localhost
	Request       *http.Request // 增加方便外部调试
	Response      *http.Response
	mux           sync.RWMutex // 锁，一个爬虫不能并发抓取，并发请建多只爬虫
}

// Java Bean链式结构
func (config *SpiderConfig) SetHeader(header http.Header) *SpiderConfig {
	config.Header = header
	return config
}

func (config *SpiderConfig) SetHeaderParm(k, v string) *SpiderConfig {
	config.Header.Set(k, v)
	return config
}

// Cookie 这样设置如果有jar != nil 那么同名cookie会和这个一起发送过去
func (config *SpiderConfig) SetCookie(v string) *SpiderConfig {
	config.SetHeaderParm("Cookie", v)
	return config
}

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

// SetUrl的同时Set一下Host
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

func (config *SpiderConfig) Clear() *SpiderConfig {
	config.Data = url.Values{}
	config.BData = []byte{}
	return config
}

func (config *SpiderConfig) ClearAll() *SpiderConfig {
	// 全部删除
	config.Header = http.Header{}
	config.Data = url.Values{}
	config.BData = []byte{}
	return config
}

// 可以删除设置的Cookie
func (config *SpiderConfig) ClearCookie() *SpiderConfig {
	config.Header.Del("Cookie")
	return config
}

// 新建一个爬虫，如果ipstring是一个代理IP地址，那使用代理客户端
func NewSpider(ipstring interface{}) (*Spider, error) {
	sp := new(Spider)
	sp.SpiderConfig = new(SpiderConfig)
	sp.Header = http.Header{}
	sp.Data = url.Values{}
	sp.BData = []byte{}
	if ipstring != nil {
		client, err := NewProxyClient(ipstring.(string))
		sp.Client = client
		sp.Ipstring = ipstring.(string)
		return sp, err
	} else {
		client, err := NewClient()
		sp.Client = client
		sp.Ipstring = "localhost"
		return sp, err
	}

}

// 新建爬虫别名函数
func New(ipstring interface{}) (*Spider, error) {
	return NewSpider(ipstring)
}

// 通过官方Client来新建爬虫，方便您更灵活
func NewSpiderByClient(client *http.Client) *Spider {
	sp := new(Spider)
	sp.SpiderConfig = new(SpiderConfig)
	sp.Header = http.Header{}
	sp.Data = url.Values{}
	sp.BData = []byte{}
	sp.Client = client
	return sp
}

// API爬虫，不用保存Cookie，可用于对接各种API，但仍然有默认UA
func NewAPI() *Spider {
	return NewSpiderByClient(NoCookieClient)
}

// auto decide which method
// 自动根据方法调用相应函数，默认GET方法
func (sp *Spider) Go() (body []byte, e error) {
	// 下面这些调用的函数很冗余， 可减少代码量，可是会复制化， 所以放弃
	switch strings.ToUpper(sp.Method) {
	case POST:
		return sp.Post()
	case POSTJSON:
		return sp.PostJSON()
	case POSTXML:
		return sp.PostXML()
	case POSTFILE:
		return sp.PostFILE()
	case PUT:
		return sp.Put()
	case PUTJSON:
		return sp.PutJSON()
	case PUTXML:
		return sp.PutXML()
	case PUTFILE:
		return sp.PutFILE()
	case DELETE:
		return sp.Delete()
	case OTHER:
		return []byte(""), errors.New("Please use method OtherGo(method, content type)")
	default:
		return sp.Get()
	}
}

// Get method,can take a client
// 手动调用方法
func (sp *Spider) Get() (body []byte, e error) {

	sp.mux.Lock()
	defer sp.mux.Unlock()

	// wait but 0 second not
	Wait(sp.Wait)

	//debug,can use SetLogLevel to change
	Logger.Debug("GET url:" + sp.Url)

	//a new request
	request, _ := http.NewRequest("GET", sp.Url, nil)

	//clone a header
	request.Header = CloneHeader(sp.Header)
	sp.Request = request

	//debug the header
	OutputMaps("---------request header--------", request.Header)

	//start request
	if sp.Client == nil {
		// default client
		sp.Client = Client
	}
	response, err := sp.Client.Do(request)
	if err != nil {
		sp.Errortimes++
		return nil, err
	}

	if response != nil {
		defer response.Body.Close()
	}

	//debug
	OutputMaps("----------response header-----------", response.Header)
	Logger.Debugf("Status：%v:%v", response.Status, response.Proto)
	sp.UrlStatuscode = response.StatusCode
	//设置新Cookie
	//Cookieb = MergeCookie(Cookieb, response.Cookies())

	//返回内容 return bytes
	body, e = ioutil.ReadAll(response.Body)
	sp.Raw = body

	sp.Fetchtimes++

	sp.Preurl = sp.Url

	sp.Response = response
	return
}

// 辅助POST
func (sp *Spider) post(method, contenttype string) (body []byte, e error) {
	sp.mux.Lock()
	defer sp.mux.Unlock()
	Wait(sp.Wait)

	Logger.Debug("POST url:" + sp.Url)

	var request = &http.Request{}

	//post data
	if sp.Data != nil {
		pr := ioutil.NopCloser(strings.NewReader(sp.Data.Encode()))
		request, _ = http.NewRequest(method, sp.Url, pr)
	} else {
		request, _ = http.NewRequest(method, sp.Url, nil)
	}
	request.Header = CloneHeader(sp.Header)

	request.Header.Set("Content-Type", contenttype)
	sp.Request = request

	OutputMaps("---------request header--------", request.Header)

	if sp.Client == nil {
		sp.Client = Client
	}
	response, err := sp.Client.Do(request)
	if err != nil {
		sp.Errortimes++
		return nil, err
	}

	if response != nil {
		defer response.Body.Close()
	}

	OutputMaps("----------response header-----------", response.Header)
	Logger.Debugf("Status：%v:%v", response.Status, response.Proto)
	sp.UrlStatuscode = response.StatusCode
	body, e = ioutil.ReadAll(response.Body)
	sp.Raw = body

	//设置新Cookie
	//MergeCookie(Cookieb, response.Cookies())
	sp.Fetchtimes++

	sp.Preurl = sp.Url

	sp.Response = response
	return
}

// 辅助Put
func (sp *Spider) put(method, contenttype string) (body []byte, e error) {
	sp.mux.Lock()
	defer sp.mux.Unlock()
	Wait(sp.Wait)

	Logger.Debug("Put url:" + sp.Url)

	var request = &http.Request{}

	//post data
	if sp.Data != nil {
		pr := ioutil.NopCloser(strings.NewReader(sp.Data.Encode()))
		request, _ = http.NewRequest(method, sp.Url, pr)
	} else {
		request, _ = http.NewRequest(method, sp.Url, nil)
	}
	request.Header = CloneHeader(sp.Header)

	request.Header.Set("Content-Type", contenttype)
	sp.Request = request

	OutputMaps("---------request header--------", request.Header)

	if sp.Client == nil {
		sp.Client = Client
	}
	response, err := sp.Client.Do(request)
	if err != nil {
		sp.Errortimes++
		return nil, err
	}

	if response != nil {
		defer response.Body.Close()
	}

	OutputMaps("----------response header-----------", response.Header)
	Logger.Debugf("Status：%v:%v", response.Status, response.Proto)
	sp.UrlStatuscode = response.StatusCode
	body, e = ioutil.ReadAll(response.Body)
	sp.Raw = body

	//设置新Cookie
	//MergeCookie(Cookieb, response.Cookies())
	sp.Fetchtimes++

	sp.Preurl = sp.Url

	sp.Response = response
	return
}

func (sp *Spider) Delete() (body []byte, e error) {

	sp.mux.Lock()
	defer sp.mux.Unlock()

	// wait but 0 second not
	Wait(sp.Wait)

	//debug,can use SetLogLevel to change
	Logger.Debug("DELETE url:" + sp.Url)

	//a new request
	request, _ := http.NewRequest("DELETE", sp.Url, nil)

	//clone a header
	request.Header = CloneHeader(sp.Header)
	sp.Request = request

	//debug the header
	OutputMaps("---------request header--------", request.Header)

	//start request
	if sp.Client == nil {
		// default client
		sp.Client = Client
	}
	response, err := sp.Client.Do(request)
	if err != nil {
		sp.Errortimes++
		return nil, err
	}

	if response != nil {
		defer response.Body.Close()
	}

	//debug
	OutputMaps("----------response header-----------", response.Header)
	Logger.Debugf("Status：%v:%v", response.Status, response.Proto)
	sp.UrlStatuscode = response.StatusCode
	//设置新Cookie
	//Cookieb = MergeCookie(Cookieb, response.Cookies())

	//返回内容 return bytes
	body, e = ioutil.ReadAll(response.Body)
	sp.Raw = body

	sp.Fetchtimes++

	sp.Preurl = sp.Url

	sp.Response = response
	return
}

// Post附带信息 can take a client
/*
	"application/x-www-form-urlencoded"
	"application/json"
	"text/xml"
	"multipart/form-data"
*/
func (sp *Spider) Post() (body []byte, e error) {
	return sp.post(POST, "application/x-www-form-urlencoded")
}

func (sp *Spider) PostJSON() (body []byte, e error) {
	return sp.post(POST, "application/json")
}

func (sp *Spider) PostXML() (body []byte, e error) {
	return sp.post(POST, "text/xml")
}

func (sp *Spider) PostFILE() (body []byte, e error) {
	return sp.post(POST, "multipart/form-data")

}

// Put
func (sp *Spider) Put() (body []byte, e error) {
	return sp.put(PUT, "application/x-www-form-urlencoded")
}

func (sp *Spider) PutJSON() (body []byte, e error) {
	return sp.put(PUT, "application/json")
}

func (sp *Spider) PutXML() (body []byte, e error) {
	return sp.put(PUT, "text/xml")
}

func (sp *Spider) PutFILE() (body []byte, e error) {
	return sp.put(PUT, "multipart/form-data")

}

// 其他Method
/*
     Method         = "OPTIONS"                ; Section 9.2
                    | "GET"                    ; Section 9.3
                    | "HEAD"                   ; Section 9.4
                    | "POST"                   ; Section 9.5
                    | "PUT"                    ; Section 9.6
                    | "DELETE"                 ; Section 9.7
                    | "TRACE"                  ; Section 9.8
                    | "CONNECT"                ; Section 9.9
                    | extension-method
   extension-method = token
     token          = 1*<any CHAR except CTLs or separators>


// content type
	"application/x-www-form-urlencoded"
	"application/json"
	"text/xml"
	"multipart/form-data"
*/
func (sp *Spider) OtherGo(method, contenttype string) (body []byte, e error) {
	sp.mux.Lock()
	defer sp.mux.Unlock()
	Wait(sp.Wait)

	Logger.Debug("POST url:" + sp.Url)

	var request = &http.Request{}

	//post data
	if sp.Data != nil {
		pr := ioutil.NopCloser(strings.NewReader(sp.Data.Encode()))
		request, _ = http.NewRequest(method, sp.Url, pr)
	} else {
		request, _ = http.NewRequest(method, sp.Url, nil)
	}
	request.Header = CloneHeader(sp.Header)

	request.Header.Set("Content-Type", contenttype)
	sp.Request = request

	OutputMaps("---------request header--------", request.Header)

	if sp.Client == nil {
		sp.Client = Client
	}
	response, err := sp.Client.Do(request)
	if err != nil {
		sp.Errortimes++
		return nil, err
	}

	if response != nil {
		defer response.Body.Close()
	}

	OutputMaps("----------response header-----------", response.Header)
	Logger.Debugf("Status：%v:%v", response.Status, response.Proto)
	sp.UrlStatuscode = response.StatusCode
	body, e = ioutil.ReadAll(response.Body)
	sp.Raw = body

	//设置新Cookie
	//MergeCookie(Cookieb, response.Cookies())
	sp.Fetchtimes++

	sp.Preurl = sp.Url

	sp.Response = response
	return
}

// class method
// 创建新头部快捷方法
func (sp *Spider) NewHeader(ua interface{}, host string, refer interface{}) {
	sp.mux.Lock()
	defer sp.mux.Unlock()
	sp.Header = NewHeader(ua, host, refer)
}

// 将抓到的数据变成字符串
func (sp *Spider) ToString() string {
	if sp.Raw == nil {
		return ""
	}
	return string(sp.Raw)
}

// 将抓到的数据变成字符串，但数据是编码的JSON
func (sp *Spider) JsonToString() (string, error) {
	if sp.Raw == nil {
		return "", nil
	}
	temp, err := util.JsonBack(sp.Raw)
	if err != nil {
		return "", err
	}
	return string(temp), nil
}

// 返回cookie
func (sp *Spider) Cookies() []*http.Cookie {
	if sp.Response != nil {
		return sp.Response.Cookies()
	} else {
		return []*http.Cookie{}
	}
}

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
	"bytes"
	"github.com/hunterhug/GoSpider/util"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

// 全局爬虫
var defaultspider *Spider

func init() {
	spider := new(Spider)
	spider.SpiderConfig = new(SpiderConfig)
	spider.Header = http.Header{}
	spider.Data = url.Values{}
	spider.BData = []byte{}
	spider.Client = Client
	// 全局爬虫使用全局客户端
	defaultspider = spider
}

// 获取默认Spider Todo
// 应该给爬虫对象，一些JavaBean的链式方法
func GetSpider() *Spider {
	return defaultspider
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
	Preurl        string       // pre url 上一次访问的URL
	Raw           []byte       // 抓取到的二进制流
	UrlStatuscode int          // the last url response code,such as 404 响应状态码
	Client        *http.Client // 真正客户端
	Fetchtimes    int          // url fetch number times 抓取次数
	Errortimes    int          // error times 失败次数
	Ipstring      string       // spider ip,just for user to record their proxyip 代理IP地址，没有代理默认localhost
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

func (config *SpiderConfig) SetUa(ua string) *SpiderConfig {
	config.Header.Set("User-Agent", ua)
	return config
}

func (config *SpiderConfig) SetRefer(refer string) *SpiderConfig {
	config.Header.Set("Referer", refer)
	return config
}

func (config *SpiderConfig) setHost(host string) *SpiderConfig {
	config.Header.Set("Host", host)
	return config
}

// SetUrl的同时Set一下Host
func (config *SpiderConfig) SetUrl(url string) *SpiderConfig {
	config.Url = url
	//https://www.zhihu.com/people/
	temp := strings.Split(url, "//")
	if len(temp) > 2 {
		config.setHost(strings.Split(temp[1], "/")[0])
	}
	return config
}

func (config *SpiderConfig) SetMethod(method string) *SpiderConfig {
	temp := GET
	switch method {
	case POST:
		temp = POST
	case POSTFILE:
		temp = POSTFILE
	case POSTJSON:
		temp = POSTJSON
	case PUT:
		temp = PUT
	case POSTXML:
		temp = POSTXML
	default:
	}
	config.Method = temp
	return config
}

func (config *SpiderConfig) SetWaitTime(num int) *SpiderConfig {
	if num > 0 {
		num = 0
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

// 新建一个爬虫，如果ipstring是一个代理IP地址，那使用代理客户端
func NewSpider(ipstring interface{}) (*Spider, error) {
	spider := new(Spider)
	spider.SpiderConfig = new(SpiderConfig)
	spider.Header = http.Header{}
	spider.Data = url.Values{}
	spider.BData = []byte{}
	if ipstring != nil {
		client, err := NewProxyClient(ipstring.(string))
		spider.Client = client
		spider.Ipstring = ipstring.(string)
		return spider, err
	} else {
		client, err := NewClient()
		spider.Client = client
		spider.Ipstring = "localhost"
		return spider, err
	}

}

// 新建爬虫别名函数
func New(ipstring interface{}) (*Spider, error) {
	return NewSpider(ipstring)
}

// 通过官方Client来新建爬虫，方便您更灵活
func NewSpiderByClient(client *http.Client) *Spider {
	spider := new(Spider)
	spider.SpiderConfig = new(SpiderConfig)
	spider.Header = http.Header{}
	spider.Data = url.Values{}
	spider.BData = []byte{}
	spider.Client = client
	return spider
}

// API爬虫，不用保存Cookie，可用于对接各种API，但仍然有默认UA
func NewAPI() *Spider {
	return NewSpiderByClient(NoCookieClient)
}

// auto decide which method
// 自动根据方法调用相应函数，默认GET方法
func (this *Spider) Go() (body []byte, e error) {
	// 下面这些调用的函数很冗余， 可减少代码量，可是会复制化， 所以放弃
	switch strings.ToUpper(this.Method) {
	case POST:
		return this.Post()
	case POSTJSON:
		return this.PostJSON()
	case POSTXML:
		return this.PostXML()
	case POSTFILE:
		return this.PostFILE()
	default:
		return this.Get()
	}
}

// Get method,can take a client
// 手动调用方法
func (this *Spider) Get() (body []byte, e error) {

	this.mux.Lock()
	defer this.mux.Unlock()

	// wait but 0 second not
	Wait(this.Wait)

	//debug,can use SetLogLevel to change
	Logger.Debug("GET url:" + this.Url)

	//a new request
	request, _ := http.NewRequest("GET", this.Url, nil)

	//clone a header
	request.Header = CloneHeader(this.Header)

	//debug the header
	OutputMaps("---------request header--------", request.Header)

	//start request
	if this.Client == nil {
		// default client
		this.Client = Client
	}
	response, err := this.Client.Do(request)
	if err != nil {
		this.Errortimes++
		return nil, err
	}
	defer response.Body.Close()

	//debug
	OutputMaps("----------response header-----------", response.Header)
	Logger.Debugf("Status：%v:%v", response.Status, response.Proto)
	this.UrlStatuscode = response.StatusCode
	//设置新Cookie
	//Cookieb = MergeCookie(Cookieb, response.Cookies())

	//返回内容 return bytes
	body, e = ioutil.ReadAll(response.Body)
	this.Raw = body

	this.Fetchtimes++

	this.Preurl = this.Url
	return
}

// Post附带信息 can take a client
func (this *Spider) Post() (body []byte, e error) {

	this.mux.Lock()
	defer this.mux.Unlock()

	Wait(this.Wait)

	Logger.Debug("POST url:" + this.Url)

	var request = &http.Request{}

	//post data
	if this.Data != nil {
		pr := ioutil.NopCloser(strings.NewReader(this.Data.Encode()))
		request, _ = http.NewRequest("POST", this.Url, pr)
	} else {
		request, _ = http.NewRequest("POST", this.Url, nil)
	}
	request.Header = CloneHeader(this.Header)

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	OutputMaps("---------request header--------", request.Header)

	if this.Client == nil {
		this.Client = Client
	}
	response, err := this.Client.Do(request)
	if err != nil {
		this.Errortimes++
		return nil, err
	}

	defer response.Body.Close()

	OutputMaps("----------response header-----------", response.Header)
	Logger.Debugf("Status：%v:%v", response.Status, response.Proto)
	this.UrlStatuscode = response.StatusCode
	body, e = ioutil.ReadAll(response.Body)
	this.Raw = body

	//设置新Cookie
	//MergeCookie(Cookieb, response.Cookies())
	this.Fetchtimes++

	this.Preurl = this.Url
	return
}

func (this *Spider) PostJSON() (body []byte, e error) {

	this.mux.Lock()
	defer this.mux.Unlock()

	Wait(this.Wait)

	Logger.Debug("POST url:" + this.Url)

	var request = &http.Request{}

	//post data
	if this.Data != nil {
		pr := ioutil.NopCloser(bytes.NewReader(this.BData))
		request, _ = http.NewRequest("POST", this.Url, pr)
	} else {
		request, _ = http.NewRequest("POST", this.Url, nil)
	}
	request.Header = CloneHeader(this.Header)

	request.Header.Set("Content-Type", "application/json")

	OutputMaps("---------request header--------", request.Header)

	if this.Client == nil {
		this.Client = Client
	}
	response, err := this.Client.Do(request)
	if err != nil {
		this.Errortimes++
		return nil, err
	}

	defer response.Body.Close()

	OutputMaps("----------response header-----------", response.Header)
	Logger.Debugf("Status：%v:%v", response.Status, response.Proto)
	this.UrlStatuscode = response.StatusCode
	body, e = ioutil.ReadAll(response.Body)
	this.Raw = body

	//设置新Cookie
	//MergeCookie(Cookieb, response.Cookies())
	this.Fetchtimes++

	this.Preurl = this.Url
	return
}

func (this *Spider) PostXML() (body []byte, e error) {

	this.mux.Lock()
	defer this.mux.Unlock()

	Wait(this.Wait)

	Logger.Debug("POST url:" + this.Url)

	var request = &http.Request{}

	//post data
	if this.Data != nil {
		pr := ioutil.NopCloser(bytes.NewReader(this.BData))
		request, _ = http.NewRequest("POST", this.Url, pr)
	} else {
		request, _ = http.NewRequest("POST", this.Url, nil)
	}
	request.Header = CloneHeader(this.Header)

	request.Header.Set("Content-Type", "text/xml")

	OutputMaps("---------request header--------", request.Header)

	if this.Client == nil {
		this.Client = Client
	}
	response, err := this.Client.Do(request)
	if err != nil {
		this.Errortimes++
		return nil, err
	}

	defer response.Body.Close()

	OutputMaps("----------response header-----------", response.Header)
	Logger.Debugf("Status：%v:%v", response.Status, response.Proto)
	this.UrlStatuscode = response.StatusCode
	body, e = ioutil.ReadAll(response.Body)
	this.Raw = body

	//设置新Cookie
	//MergeCookie(Cookieb, response.Cookies())
	this.Fetchtimes++

	this.Preurl = this.Url
	return
}

func (this *Spider) PostFILE() (body []byte, e error) {

	this.mux.Lock()
	defer this.mux.Unlock()

	Wait(this.Wait)

	Logger.Debug("POST url:" + this.Url)

	var request = &http.Request{}

	//post data
	if this.Data != nil {
		pr := ioutil.NopCloser(bytes.NewReader(this.BData))
		request, _ = http.NewRequest("POST", this.Url, pr)
	} else {
		request, _ = http.NewRequest("POST", this.Url, nil)
	}
	request.Header = CloneHeader(this.Header)

	request.Header.Set("Content-Type", "multipart/form-data")

	OutputMaps("---------request header--------", request.Header)

	if this.Client == nil {
		this.Client = Client
	}
	response, err := this.Client.Do(request)
	if err != nil {
		this.Errortimes++
		return nil, err
	}

	defer response.Body.Close()

	OutputMaps("----------response header-----------", response.Header)
	Logger.Debugf("Status：%v:%v", response.Status, response.Proto)
	this.UrlStatuscode = response.StatusCode
	body, e = ioutil.ReadAll(response.Body)
	this.Raw = body

	//设置新Cookie
	//MergeCookie(Cookieb, response.Cookies())
	this.Fetchtimes++

	this.Preurl = this.Url
	return
}

// class method
// 创建新头部快捷方法
func (this *Spider) NewHeader(ua interface{}, host string, refer interface{}) {
	this.mux.Lock()
	defer this.mux.Unlock()
	this.Header = NewHeader(ua, host, refer)
}

// 将抓到的数据变成字符串
func (this *Spider) ToString() string {
	if this.Raw == nil {
		return ""
	}
	return string(this.Raw)
}

// 将抓到的数据变成字符串，但数据是编码的JSON
func (this *Spider) JsonToString() (string, error) {
	if this.Raw == nil {
		return "", nil
	}
	temp, err := util.JsonBack(this.Raw)
	if err != nil {
		return "", err
	}
	return string(temp), nil
}

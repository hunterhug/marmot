/*
 * Created by 一只尼玛 on 2016/8/12.
 * 功能： 网站爬取功能
 *
 */
package spider

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Spider struct {
	Url           string // the last fetch url
	UrlStatuscode int    // the last url response code,such as 404
	Method        string //Get Post
	Header        http.Header
	Data          url.Values // post data
	Wait          int        // sleep time
	Client        *http.Client
	Fetchtimes    int    // url fetch number times
	Errortimes    int    // error times
	Ipstring      string // spider ip,just for user to record their proxyip
}

func NewSpider(ipstring interface{}) (*Spider, error) {
	spider := new(Spider)
	spider.Header = http.Header{}
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

// auto decide which method
func (this *Spider) Go() (body []byte, e error) {
	if strings.ToLower(this.Method) == "post" {
		return this.Post()
	} else {
		return this.Get()
	}

}

// Get method,can take a client
func (this *Spider) Get() (body []byte, e error) {
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
	this.Fetchtimes++
	return
}

// Post附带信息 can take a client
func (this *Spider) Post() (body []byte, e error) {
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

	//设置新Cookie
	//MergeCookie(Cookieb, response.Cookies())
	this.Fetchtimes++
	return
}

// class method
func (this *Spider) NewHeader(ua interface{}, host string, refer interface{}) {
	this.Header = NewHeader(ua, host, refer)
}

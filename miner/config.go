/*
	版权所有，侵权必究
	署名-非商业性使用-禁止演绎 4.0 国际
	警告： 以下的代码版权归属hunterhug，请不要传播或修改代码
	你可以在教育用途下使用该代码，但是禁止公司或个人用于商业用途(在未授权情况下不得用于盈利)
	商业授权请联系邮箱：gdccmcm14@live.com QQ:459527502

	All right reserved
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
	For more information on commercial licensing please contact hunterhug.
	Ask for commercial licensing please contact Mail:gdccmcm14@live.com Or QQ:459527502
*
*/

package miner

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/hunterhug/marmot/util"
)

// Worker is the main object to sent http request and return result of response
type Worker struct {
	// In order fast chain func call I put the basic config below
	Url          string         // Which url we want
	Method       string         // Get/Post method
	Header       http.Header    // Http header
	Data         url.Values     // Sent by form data
	FileName     string         // FileName which sent to remote
	FileFormName string         // File Form Name which sent to remote
	BData        []byte         // Sent by binary data, can together with File
	Wait         int            // Wait Time
	Client       *http.Client   // Our Client
	Request      *http.Request  // Debug
	Response     *http.Response // Debug
	Raw          []byte         // Raw data we get

	// can ignore ! The name below is not so good but has already been used in many project, so bear it.
	Preurl        string // Pre url
	UrlStatuscode int    // The last url response code, such as 404
	Fetchtimes    int    // Url fetch number times
	Errortimes    int    // Url fetch error times
	Ipstring      string // Worker proxy ip, just for user to record their proxy ip, default: localhost

	// AOP like Java
	Ctx          context.Context
	BeforeAction func(context.Context, *Worker)
	AfterAction  func(context.Context, *Worker)

	mux sync.RWMutex // Lock, execute concurrently please use worker Pool!
}

// Java Bean Chain pattern
func (worker *Worker) SetHeader(header http.Header) *Worker {
	worker.Header = header
	return worker
}

// Default Worker SetHeader!
func SetHeader(header http.Header) *Worker {
	return DefaultWorker.SetHeader(header)
}

func (worker *Worker) SetHeaderParm(k, v string) *Worker {
	worker.Header.Set(k, v)
	return worker
}

func SetHeaderParm(k, v string) *Worker {
	return DefaultWorker.SetHeaderParm(k, v)
}

func (worker *Worker) SetCookie(v string) *Worker {
	worker.SetHeaderParm("Cookie", v)
	return worker
}

func SetCookie(v string) *Worker {
	return DefaultWorker.SetCookie(v)
}

// Set Cookie by file.
func (worker *Worker) SetCookieByFile(file string) (*Worker, error) {
	haha, err := util.ReadfromFile(file)
	if err != nil {
		return nil, err
	}
	cookie := string(haha)
	cookie = strings.Replace(cookie, " ", "", -1)
	cookie = strings.Replace(cookie, "\n", "", -1)
	cookie = strings.Replace(cookie, "\r", "", -1)
	sconfig := worker.SetCookie(cookie)
	return sconfig, nil
}

func SetCookieByFile(file string) (*Worker, error) {
	return DefaultWorker.SetCookieByFile(file)
}

func (worker *Worker) SetUa(ua string) *Worker {
	worker.Header.Set("User-Agent", ua)
	return worker
}

func SetUa(ua string) *Worker {
	return DefaultWorker.SetUa(ua)
}

func (worker *Worker) SetRefer(refer string) *Worker {
	worker.Header.Set("Referer", refer)
	return worker
}

func SetRefer(refer string) *Worker {
	return DefaultWorker.SetRefer(refer)
}

func (worker *Worker) SetHost(host string) *Worker {
	worker.Header.Set("Host", host)
	return worker
}

// SetUrl, at the same time SetHost
func (worker *Worker) SetUrl(url string) *Worker {
	worker.Url = url
	//https://www.zhihu.com/people/
	temp := strings.Split(url, "//")
	if len(temp) >= 2 {
		worker.SetHost(strings.Split(temp[1], "/")[0])
	}
	return worker
}

func SetUrl(url string) *Worker {
	return DefaultWorker.SetUrl(url)
}

func (worker *Worker) SetFileInfo(fileName, fileFormName string) *Worker {
	worker.FileName = fileName
	worker.FileFormName = fileFormName
	return worker
}

func SetFileInfo(fileName, fileFormName string) *Worker {
	return DefaultWorker.SetFileInfo(fileName, fileFormName)
}

func (worker *Worker) SetMethod(method string) *Worker {
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
	worker.Method = temp
	return worker
}

func SetMethod(method string) *Worker {
	return DefaultWorker.SetMethod(method)
}

func (worker *Worker) SetWaitTime(num int) *Worker {
	if num <= 0 {
		num = 1
	}
	worker.Wait = num
	return worker
}

func SetWaitTime(num int) *Worker {
	return DefaultWorker.SetWaitTime(num)
}

func (worker *Worker) SetBData(data []byte) *Worker {
	worker.BData = data
	return worker
}

func SetBData(data []byte) *Worker {
	return DefaultWorker.SetBData(data)
}

func (worker *Worker) SetForm(form url.Values) *Worker {
	worker.Data = form
	return worker
}

func SetForm(form url.Values) *Worker {
	return DefaultWorker.SetForm(form)
}

func (worker *Worker) SetFormParm(k, v string) *Worker {
	worker.Data.Set(k, v)
	return worker
}

func SetFormParm(k, v string) *Worker {
	return DefaultWorker.SetFormParm(k, v)
}

// Set Context so Action can soft
func (worker *Worker) SetContext(ctx context.Context) *Worker {
	worker.Ctx = ctx
	return worker
}

func SetContext(ctx context.Context) *Worker {
	return DefaultWorker.SetContext(ctx)
}

func (worker *Worker) SetBeforeAction(fc func(context.Context, *Worker)) *Worker {
	worker.BeforeAction = fc
	return worker
}

func SetBeforeAction(fc func(context.Context, *Worker)) *Worker {
	return DefaultWorker.SetBeforeAction(fc)
}

func (worker *Worker) SetAfterAction(fc func(context.Context, *Worker)) *Worker {
	worker.AfterAction = fc
	return worker
}

func SetAfterAction(fc func(context.Context, *Worker)) *Worker {
	return DefaultWorker.SetAfterAction(fc)
}

// Clear data we sent
func (worker *Worker) Clear() *Worker {
	worker.Data = url.Values{}
	worker.BData = []byte{}
	return worker
}

func Clear() *Worker {
	return DefaultWorker.Clear()
}

// All clear include header
func (worker *Worker) ClearAll() *Worker {
	worker.Header = http.Header{}
	worker.Data = url.Values{}
	worker.BData = []byte{}
	return worker
}

func ClearAll() *Worker {
	return DefaultWorker.ClearAll()
}

// Clear Cookie
func (worker *Worker) ClearCookie() *Worker {
	worker.Header.Del("Cookie")
	return worker
}

func ClearCookie() *Worker {
	return DefaultWorker.ClearCookie()
}

// Get Cookies
func (worker *Worker) GetCookies() []*http.Cookie {
	if worker.Response != nil {
		return worker.Response.Cookies()
	} else {
		return []*http.Cookie{}
	}
}

func GetCookies() []*http.Cookie {
	return DefaultWorker.GetCookies()
}

// Deprecated
func (worker *Worker) NewHeader(ua interface{}, host string, refer interface{}) {
}

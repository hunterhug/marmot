/*
	All right reserved https://github.com/hunterhug/marmot at 2016-2021
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
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
	*Request
	*Response

	// Which url we want
	Url string

	// Get,Post method
	Method string

	// Our Client
	Client *http.Client

	// Wait Sleep Time
	Wait int

	// Worker proxy ip, just for user to record their proxy ip, default: localhost
	Ip string

	// AOP like Java
	Ctx          context.Context
	BeforeAction func(context.Context, *Worker)
	AfterAction  func(context.Context, *Worker)

	// Http header
	Header http.Header

	// Mux lock
	mux sync.RWMutex
}

type Request struct {
	Data         url.Values    // Sent by form data
	FileName     string        // FileName which sent to remote
	FileFormName string        // File Form Name which sent to remote
	BData        []byte        // Sent by binary data, can together with File
	Request      *http.Request // Debug
}

type Response struct {
	Response           *http.Response // Debug
	Raw                []byte         // Raw data we get
	ResponseStatusCode int            // The last url response code, such as 404
}

// SetHeader Java Bean Chain pattern
func (worker *Worker) SetHeader(header http.Header) *Worker {
	worker.Header = header
	return worker
}

// SetHeader Default Worker SetHeader!
func SetHeader(header http.Header) *Worker {
	return DefaultWorker.SetHeader(header)
}

func (worker *Worker) SetHeaderParam(k, v string) *Worker {
	worker.Header.Set(k, v)
	return worker
}

func SetHeaderParam(k, v string) *Worker {
	return DefaultWorker.SetHeaderParam(k, v)
}

func (worker *Worker) SetCookie(v string) *Worker {
	worker.SetHeaderParam("Cookie", v)
	return worker
}

func SetCookie(v string) *Worker {
	return DefaultWorker.SetCookie(v)
}

// SetCookieByFile Set Cookie by file.
func (worker *Worker) SetCookieByFile(file string) (*Worker, error) {
	haha, err := util.ReadFromFile(file)
	if err != nil {
		return nil, err
	}
	cookie := string(haha)
	cookie = strings.Replace(cookie, " ", "", -1)
	cookie = strings.Replace(cookie, "\n", "", -1)
	cookie = strings.Replace(cookie, "\r", "", -1)
	return worker.SetCookie(cookie), nil
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

// SetUrl at the same time SetHost
func (worker *Worker) SetUrl(url string) *Worker {
	worker.Url = url
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

func (worker *Worker) SetFormParam(k, v string) *Worker {
	worker.Data.Set(k, v)
	return worker
}

func SetFormParam(k, v string) *Worker {
	return DefaultWorker.SetFormParam(k, v)
}

// SetContext Set Context so Action can soft
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
// I suggest use Clone() to avoid clear
func (worker *Worker) Clear() *Worker {
	worker.Request = newRequest()
	worker.Response = new(Response)
	return worker
}

func Clear() *Worker {
	return DefaultWorker.Clear()
}

// ClearAll All clear include header
func (worker *Worker) ClearAll() *Worker {
	worker.Clear()
	worker.Header = http.Header{}
	return worker
}

func ClearAll() *Worker {
	return DefaultWorker.ClearAll()
}

// ClearCookie Clear Cookie
func (worker *Worker) ClearCookie() *Worker {
	worker.Header.Del("Cookie")
	return worker
}

func ClearCookie() *Worker {
	return DefaultWorker.ClearCookie()
}

// GetCookies Get Cookies
func (worker *Worker) GetCookies() []*http.Cookie {
	if worker.Response != nil && worker.Response.Response != nil {
		return worker.Response.Response.Cookies()
	} else {
		return []*http.Cookie{}
	}
}

// GetResponseStatusCode Get ResponseStatusCode
func (worker *Worker) GetResponseStatusCode() int {
	if worker.Response != nil && worker.Response.Response != nil {
		return worker.Response.ResponseStatusCode
	} else {
		return 0
	}
}

func GetCookies() []*http.Cookie {
	return DefaultWorker.GetCookies()
}

func GetResponseStatusCode() int {
	return DefaultWorker.GetResponseStatusCode()
}

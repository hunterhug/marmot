/*
	All right reserved https://github.com/hunterhug/marmot at 2016-2021
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
*/
package miner

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/hunterhug/marmot/util"
	"mime/multipart"
)

// NewWorkerByClient New Worker by Your Client
func NewWorkerByClient(client *http.Client) *Worker {
	worker := new(Worker)
	worker.Request = newRequest()
	worker.Header = http.Header{}
	worker.Response = new(Response)

	worker.Client = client
	return worker
}

// NewAPI New API Worker, No Cookie Keep, share the same http client
func NewAPI() *Worker {
	return NewWorkerByClient(NoCookieClient)
}

// NewWorker New a worker, if ipString is a proxy address, New a proxy client. evey time gen a new http client!
// Proxy address such as:
// 		http://[user]:[password@]ip:port, [] stand it can choose or not. case: socks5://127.0.0.1:1080
func NewWorker(proxyIpString interface{}) (*Worker, error) {
	if proxyIpString != nil {
		client, err := NewProxyClient(strings.ToLower(proxyIpString.(string)))
		if err != nil {
			return nil, err
		}

		worker := NewWorkerByClient(client)
		worker.Ip = proxyIpString.(string)
		return worker, err
	} else {
		client := NewClient()
		worker := NewWorkerByClient(client)
		worker.Ip = "localhost"
		return worker, nil
	}
}

// NewWorkerWithProxy Alias func
func NewWorkerWithProxy(proxyIpString interface{}) (*Worker, error) {
	return NewWorker(proxyIpString)
}

// NewWorkerWithNoProxy Alias func
func NewWorkerWithNoProxy() *Worker {
	w, _ := NewWorker(nil)
	return w
}

// New Alias Name for NewWorker
func New(ipString interface{}) (*Worker, error) {
	return NewWorker(ipString)
}

// fast new request
func newRequest() *Request {
	req := new(Request)
	req.Data = url.Values{}
	req.BData = []byte{}
	return req
}

// Clone Should Clone a new worker if you want to use repeat
func (worker *Worker) Clone() *Worker {
	cloneWorker := NewWorkerByClient(worker.Client)
	cloneWorker.Ip = worker.Ip
	return cloneWorker
}

// Go Auto decide which method, Default Get.
func (worker *Worker) Go() (body []byte, e error) {
	switch strings.ToUpper(worker.Method) {
	case POST:
		return worker.Post()
	case POSTJSON:
		return worker.PostJSON()
	case POSTXML:
		return worker.PostXML()
	case POSTFILE:
		return worker.PostFILE()
	case PUT:
		return worker.Put()
	case PUTJSON:
		return worker.PutJSON()
	case PUTXML:
		return worker.PutXML()
	case PUTFILE:
		return worker.PutFILE()
	case DELETE:
		return worker.Delete()
	case OTHER:
		return []byte(""), errors.New("please use method OtherGo(method, contentType string) or OtherGoBinary(method, contentType string)")
	default:
		return worker.Get()
	}
}

func (worker *Worker) GoByMethod(method string) (body []byte, e error) {
	return worker.SetMethod(method).Go()
}

// ToString This make effect only your worker exec serial! Attention!
// Change Your Raw data To string
func (worker *Worker) ToString() string {
	if worker.Raw == nil {
		return ""
	}
	return string(worker.Raw)
}

// JsonToString This make effect only your worker exec serial! Attention!
// Change Your JSON'like Raw data to string
func (worker *Worker) JsonToString() (string, error) {
	if worker.Raw == nil {
		return "", nil
	}
	temp, err := util.JsonBack(worker.Raw)
	if err != nil {
		return "", err
	}
	return string(temp), nil
}

// Main method I make!
func (worker *Worker) sent(method, contentType string, binary bool) (body []byte, e error) {
	// Lock it for save
	worker.mux.Lock()
	defer worker.mux.Unlock()

	// Before Action we can change or add something before Go()
	if worker.BeforeAction != nil {
		worker.BeforeAction(worker.Ctx, worker)
	}

	// Wait if must
	if worker.Wait > 0 {
		Wait(worker.Wait)
	}

	// For debug
	Logger.Debugf("[GoWorker] %s %s", method, worker.Url)

	// New a Request
	var request *http.Request
	var err error

	// If binary value is true and BData is not empty
	// suit for POSTJSON(), POSTFILE()
	if len(worker.BData) != 0 && binary {
		pr := bytes.NewReader(worker.BData)
		request, err = http.NewRequest(method, worker.Url, pr)
	} else if len(worker.Data) != 0 { // such POST() from table form
		pr := strings.NewReader(worker.Data.Encode())
		request, err = http.NewRequest(method, worker.Url, pr)
	} else {
		request, err = http.NewRequest(method, worker.Url, nil)
	}

	if err != nil {
		return nil, err
	}

	// Close avoid EOF
	// For client requests, setting this field prevents re-use of
	// TCP connections between requests to the same hosts, as if
	// Transport.DisableKeepAlives were set.
	// maybe you want long connection
	//request.Close = true

	// Clone Header, I add some HTTP header!
	request.Header = CloneHeader(worker.Header)

	// In fact content type must not empty
	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}
	worker.Request.Request = request

	// Debug for RequestHeader
	OutputMaps("Request header", request.Header)

	// Tolerate abnormal way to create a Worker
	if worker.Client == nil {
		worker.Client = Client
	}

	// Do it
	response, err := worker.Client.Do(request)
	if err != nil {
		return nil, err
	}

	// Close it attention response may be nil
	if response != nil {
		//response.Close = true
		defer response.Body.Close()
	}

	// Debug
	OutputMaps("Response header", response.Header)
	Logger.Debugf("[GoWorker] %v %s", response.Proto, response.Status)

	// Read output
	body, e = ioutil.ReadAll(response.Body)
	worker.Raw = body

	worker.ResponseStatusCode = response.StatusCode
	worker.Response.Response = response

	// After action
	if worker.AfterAction != nil {
		worker.AfterAction(worker.Ctx, worker)
	}

	return
}

// Get method
func (worker *Worker) Get() (body []byte, e error) {
	worker.Clear()
	return worker.sent(GET, "", false)
}

func (worker *Worker) Delete() (body []byte, e error) {
	worker.Clear()
	return worker.sent(DELETE, "", false)
}

// Post Almost include bellow:
/*
	"application/x-www-form-urlencoded"
	"application/json"
	"text/xml"
	"multipart/form-data"
*/
func (worker *Worker) Post() (body []byte, e error) {
	return worker.sent(POST, HTTPFORMContentType, false)
}

func (worker *Worker) PostJSON() (body []byte, e error) {
	return worker.sent(POST, HTTPJSONContentType, true)
}

func (worker *Worker) PostXML() (body []byte, e error) {
	return worker.sent(POST, HTTPXMLContentType, true)
}

func (worker *Worker) PostFILE() (body []byte, e error) {
	return worker.sentFile(POST)

}

func (worker *Worker) sentFile(method string) ([]byte, error) {
	if worker.FileName == "" || worker.FileFormName == "" {
		return nil, errors.New("fileName or fileFormName must not empty")
	}

	if len(worker.BData) == 0 {
		return nil, errors.New("BData must not empty")
	}

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile(worker.FileFormName, worker.FileName)
	if err != nil {
		return nil, err
	}

	_, err = fileWriter.Write(worker.BData)
	if err != nil {
		return nil, err
	}

	contentType := bodyWriter.FormDataContentType()

	err = bodyWriter.Close()
	if err != nil {
		return nil, err
	}

	worker.SetBData(bodyBuf.Bytes())

	return worker.sent(method, contentType, true)
}

func (worker *Worker) Put() (body []byte, e error) {
	return worker.sent(PUT, HTTPFORMContentType, false)
}

func (worker *Worker) PutJSON() (body []byte, e error) {
	return worker.sent(PUT, HTTPJSONContentType, true)
}

func (worker *Worker) PutXML() (body []byte, e error) {
	return worker.sent(PUT, HTTPXMLContentType, true)
}

func (worker *Worker) PutFILE() (body []byte, e error) {
	return worker.sentFile(PUT)
}

/*
OtherGo Method

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


Content Type

	"application/x-www-form-urlencoded"
	"application/json"
	"text/xml"
	"multipart/form-data"
*/
func (worker *Worker) OtherGo(method, contentType string) (body []byte, e error) {
	return worker.sent(method, contentType, false)
}

func (worker *Worker) OtherGoBinary(method, contentType string) (body []byte, e error) {
	return worker.sent(method, contentType, true)
}
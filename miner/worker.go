package miner

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/hunterhug/parrot/util"
)

// New a worker, if ipstring is a proxy address, New a proxy client.
// Proxy address such as:
// 		http://[user]:[password@]ip:port, [] stand it can choose or not. case: socks5://127.0.0.1:1080
func NewWorker(ipstring interface{}) (*Worker, error) {
	worker := new(Worker)
	worker.Header = http.Header{}
	worker.Data = url.Values{}
	worker.BData = []byte{}
	if ipstring != nil {
		client, err := NewProxyClient(strings.ToLower(ipstring.(string)))
		worker.Client = client
		worker.Ipstring = ipstring.(string)
		return worker, err
	} else {
		client, err := NewClient()
		worker.Client = client
		worker.Ipstring = "localhost"
		return worker, err
	}

}

// Alias Name for NewWorker
func New(ipstring interface{}) (*Worker, error) {
	return NewWorker(ipstring)
}

// New Worker by Your Client
func NewWorkerByClient(client *http.Client) *Worker {
	worker := new(Worker)
	worker.Header = http.Header{}
	worker.Data = url.Values{}
	worker.BData = []byte{}
	worker.Client = client
	return worker
}

// New API Worker, No Cookie Keep.
func NewAPI() *Worker {
	return NewWorkerByClient(NoCookieClient)
}

// Auto decide which method, Default Get.
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
		return []byte(""), errors.New("please use method OtherGo(method, content type)")
	default:
		return worker.Get()
	}
}

func (worker *Worker) GoByMethod(method string) (body []byte, e error) {
	return worker.SetMethod(method).Go()
}

// This make effect only your worker exec serial! Attention!
// Change Your Raw data To string
func (worker *Worker) ToString() string {
	if worker.Raw == nil {
		return ""
	}
	return string(worker.Raw)
}

// This make effect only your worker exec serial! Attention!
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
func (worker *Worker) sent(method, contenttype string, binary bool) (body []byte, e error) {
	// Lock it for save
	worker.mux.Lock()
	defer worker.mux.Unlock()

	// Before FAction we can change or add something before Go()
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
	var request = &http.Request{}

	// If binary parm value is true and BData is not empty
	// suit for POSTJSON(), POSTFILE()
	if len(worker.BData) != 0 && binary {
		pr := ioutil.NopCloser(bytes.NewReader(worker.BData))
		request, _ = http.NewRequest(method, worker.Url, pr)
	} else if len(worker.Data) != 0 { // such POST() from table form
		pr := ioutil.NopCloser(strings.NewReader(worker.Data.Encode()))
		request, _ = http.NewRequest(method, worker.Url, pr)
	} else {
		request, _ = http.NewRequest(method, worker.Url, nil)
	}

	// Clone Header, I add some HTTP header!
	request.Header = CloneHeader(worker.Header)

	// In fact contenttype must not empty
	if contenttype != "" {
		request.Header.Set("Content-Type", contenttype)
	}
	worker.Request = request

	// Debug for RequestHeader
	OutputMaps("Request header", request.Header)

	// Tolerate abnormal way to create a Worker
	if worker.Client == nil {
		worker.Client = Client
	}

	// Do it
	response, err := worker.Client.Do(request)
	if err != nil {
		// I count Error time
		worker.Errortimes++
		return nil, err
	}

	// Close it attention response may be nil
	if response != nil {
		defer response.Body.Close()
	}

	// Debug
	OutputMaps("Response header", response.Header)
	Logger.Debugf("[GoWorker] %v %s", response.Proto, response.Status)

	// Read output
	body, e = ioutil.ReadAll(response.Body)
	worker.Raw = body

	worker.UrlStatuscode = response.StatusCode
	worker.Preurl = worker.Url
	worker.Response = response
	worker.Fetchtimes++

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
	return worker.sent(POST, HTTPFILEContentType, true)

}

// Put
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
	return worker.sent(PUT, HTTPFILEContentType, true)

}

/*
Other Method

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
func (worker *Worker) OtherGo(method, contenttype string) (body []byte, e error) {
	return worker.sent(method, contenttype, true)
}

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
	"bytes"
	"errors"
	"github.com/hunterhug/GoTool/util"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// New a spider, if ipstring is a proxy address, New a proxy client.
// Proxy address such as: http://[user]:[password@]ip:port, [] stand it can choose or not
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

// Alias Name for NewSpider
func New(ipstring interface{}) (*Spider, error) {
	return NewSpider(ipstring)
}

// New Spider by Your Client
func NewSpiderByClient(client *http.Client) *Spider {
	sp := new(Spider)
	sp.SpiderConfig = new(SpiderConfig)
	sp.Header = http.Header{}
	sp.Data = url.Values{}
	sp.BData = []byte{}
	sp.Client = client
	return sp
}

// New API Spider, No Cookie Keep.
func NewAPI() *Spider {
	return NewSpiderByClient(NoCookieClient)
}

// Auto decide which method, Default Get.
func (sp *Spider) Go() (body []byte, e error) {
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
		return []byte(""), errors.New("please use method OtherGo(method, content type)")
	default:
		return sp.Get()
	}
}

// This make effect only your spider exec serial! Attention!
// Change Your Raw data To string
func (sp *Spider) ToString() string {
	if sp.Raw == nil {
		return ""
	}
	return string(sp.Raw)
}

// This make effect only your spider exec serial! Attention!
// Change Your JSON'like Raw data to string
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

// Main method I make!
func (sp *Spider) sent(method, contenttype string, binary bool) (body []byte, e error) {
	sp.mux.Lock()
	defer sp.mux.Unlock()
	Wait(sp.Wait)

	Logger.Debugf("[GoSpider] %s url: %s", method, sp.Url)

	var request = &http.Request{}

	if len(sp.BData) != 0 && binary {
		pr := ioutil.NopCloser(bytes.NewReader(sp.BData))
		request, _ = http.NewRequest(method, sp.Url, pr)
	} else if len(sp.Data) != 0 {
		pr := ioutil.NopCloser(strings.NewReader(sp.Data.Encode()))
		request, _ = http.NewRequest(method, sp.Url, pr)
	} else {
		request, _ = http.NewRequest(method, sp.Url, nil)
	}

	request.Header = CloneHeader(sp.Header)

	if contenttype != "" {
		request.Header.Set("Content-Type", contenttype)
	}
	sp.Request = request

	OutputMaps("Request header", request.Header)

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

	OutputMaps("Response header", response.Header)
	Logger.Debugf("[GoSpider] Status：%v:%v", response.Status, response.Proto)
	sp.UrlStatuscode = response.StatusCode

	body, e = ioutil.ReadAll(response.Body)
	sp.Raw = body
	sp.Fetchtimes++
	sp.Preurl = sp.Url
	sp.Response = response
	return
}

// Get method
func (sp *Spider) Get() (body []byte, e error) {
	sp.Clear()
	return sp.sent(GET, "", false)
}

func (sp *Spider) Delete() (body []byte, e error) {
	sp.Clear()
	return sp.sent(DELETE, "", false)
}

// Post Almost include bellow:
/*
	"application/x-www-form-urlencoded"
	"application/json"
	"text/xml"
	"multipart/form-data"
*/
func (sp *Spider) Post() (body []byte, e error) {
	return sp.sent(POST, HTTPFORMContentType, false)
}

func (sp *Spider) PostJSON() (body []byte, e error) {
	return sp.sent(POST, HTTPJSONContentType, true)
}

func (sp *Spider) PostXML() (body []byte, e error) {
	return sp.sent(POST, HTTPXMLContentType, true)
}

func (sp *Spider) PostFILE() (body []byte, e error) {
	return sp.sent(POST, HTTPFILEContentType, true)

}

// Put
func (sp *Spider) Put() (body []byte, e error) {
	return sp.sent(PUT, HTTPFORMContentType, false)
}

func (sp *Spider) PutJSON() (body []byte, e error) {
	return sp.sent(PUT, HTTPJSONContentType, true)
}

func (sp *Spider) PutXML() (body []byte, e error) {
	return sp.sent(PUT, HTTPXMLContentType, true)
}

func (sp *Spider) PutFILE() (body []byte, e error) {
	return sp.sent(PUT, HTTPFILEContentType, true)

}

// Other Method
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
	return sp.sent(method, contenttype, true)
}

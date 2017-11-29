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
	"net/http"
	"net/url"
)

// Global Spider
var DefaultSpider *Spider

func init() {
	UaInit()

	// New a spider
	sp := new(Spider)
	sp.SpiderConfig = new(SpiderConfig)
	sp.Header = http.Header{}
	sp.Data = url.Values{}
	sp.BData = []byte{}
	sp.Client = Client

	// Global spider!
	DefaultSpider = sp

}

// This make effect only your spider exec serial! Attention!
// Change Your Raw data To string
func ToString() string {
	return DefaultSpider.ToString()
}

// This make effect only your spider exec serial! Attention!
// Change Your JSON'like Raw data to string
func JsonToString() (string, error) {
	return DefaultSpider.JsonToString()
}

func Get() (body []byte, e error) {
	return DefaultSpider.Get()
}

func Delete() (body []byte, e error) {
	return DefaultSpider.Delete()
}

func Go() (body []byte, e error) {
	return DefaultSpider.Go()
}

func OtherGo(method, contenttype string) (body []byte, e error) {
	return DefaultSpider.OtherGo(method, contenttype)
}

func Post() (body []byte, e error) {
	return DefaultSpider.Post()
}

func PostJSON() (body []byte, e error) {
	return DefaultSpider.PostJSON()
}

func PostFILE() (body []byte, e error) {
	return DefaultSpider.PostFILE()
}

func PostXML() (body []byte, e error) {
	return DefaultSpider.PostXML()
}

func Put() (body []byte, e error) {
	return DefaultSpider.Put()
}
func PutJSON() (body []byte, e error) {
	return DefaultSpider.PutJSON()
}

func PutFILE() (body []byte, e error) {
	return DefaultSpider.PutFILE()
}

func PutXML() (body []byte, e error) {
	return DefaultSpider.PutXML()
}

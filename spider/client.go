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
	"github.com/hunterhug/GoTool/util"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

// Cookie record Jar
func NewJar() *cookiejar.Jar {
	cookieJar, _ := cookiejar.New(nil)
	return cookieJar
}

var (
	// Default Client to ask get or post
	// Save Cookie, No timeout!
	Client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			Logger.Debugf("[GoSpider] Redirect:%v", req.URL)
			return nil
		},
		Jar: NewJar(),
	}

	// Default Client
	// Not Save Cookie
	NoCookieClient = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			Logger.Debugf("[GoSpider] Redirect:%v", req.URL)
			return nil
		},
	}
)

// New a Proxy client, Default save cookie, Can timeout
func NewProxyClient(proxystring string) (*http.Client, error) {
	proxy, err := url.Parse(proxystring)
	if err != nil {
		return nil, err
	}

	// This a alone client, diff from global client.
	client := &http.Client{
		// Allow redirect
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			Logger.Debugf("[GoSpider] Redirect:%v", req.URL)
			return nil
		},
		// Allow proxy
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
		// Allow keep cookie
		Jar: NewJar(),
		// Allow Timeout
		Timeout: util.Second(DefaultTimeOut),
	}
	return client, nil
}

// New a client, diff from proxy client
func NewClient() (*http.Client, error) {
	client := &http.Client{
		// Allow redirect
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			Logger.Debugf("[GoSpider] Redirect:%v", req.URL)
			return nil
		},
		Jar:     NewJar(),
		Timeout: util.Second(DefaultTimeOut),
	}
	return client, nil
}

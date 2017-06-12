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

// 功能： 网络COOKIE功能
import (
	"github.com/hunterhug/GoSpider/util"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

//cookie record
// 记录Cookie
func NewJar() *cookiejar.Jar {
	cookieJar, _ := cookiejar.New(nil)
	return cookieJar
}

var (
	//default client to ask get or post
	// 默认的官方客户端，带cookie,方便使用，没有超时时间
	Client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			Logger.Debugf("-----------Redirect:%v------------", req.URL)
			return nil
		},
		Jar: NewJar(),
	}

	// 没有cookie的客户端
	NoCookieClient = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			Logger.Debugf("-----------Redirect:%v------------", req.URL)
			return nil
		},
	}
)

// a proxy client
// 带代理客户端，全部有带cookie
func NewProxyClient(proxystring string) (*http.Client, error) {
	proxy, err := url.Parse(proxystring)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		// allow redirect
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			Logger.Debugf("-----------Redirect:%v------------", req.URL)
			return nil
		},
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
		Jar:     NewJar(),
		Timeout: util.Second(DefaultTimeOut),
	}
	return client, nil
}

// a client
// 不带代理客户端
func NewClient() (*http.Client, error) {
	client := &http.Client{
		// allow redirect
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			Logger.Debugf("-----------Redirect:%v------------", req.URL)
			return nil
		},
		Jar:     NewJar(),
		Timeout: util.Second(DefaultTimeOut),
	}
	return client, nil
}

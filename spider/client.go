/*
 * Created by 一只尼玛 on 2016/8/12.
 * 功能： 网络COOKIE功能
 *
 */
package spider

import (
	"github.com/hunterhug/go_tool/util"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

//cookie record
func NewJar() *cookiejar.Jar {
	cookieJar, _ := cookiejar.New(nil)
	return cookieJar
}

var (
	//default client to ask get or post
	Client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			Logger.Debugf("-----------Redirect:%v------------", req.URL)
			return nil
		},
		Jar: NewJar(),
	}
	//每次访问携带的cookie not use
	Cookieb = []*http.Cookie{} //map[string][]string
)

// a proxy client
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
		Jar: NewJar(),
		Timeout: util.Second(DefaultTimeOut),
	}
	return client, nil
}

// a client
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

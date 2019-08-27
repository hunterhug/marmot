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
	"crypto/tls"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/hunterhug/marmot/util"
	"github.com/hunterhug/marmot/proxy"
	//"golang.org/x/net/proxy" // see https://github.com/golang/net
)

// Cookie record Jar
func NewJar() *cookiejar.Jar {
	cookieJar, _ := cookiejar.New(nil)
	return cookieJar
}

// Default Client
var (
	// Save Cookie, No timeout!
	Client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			Logger.Debugf("[GoWorker] Redirect:%v", req.URL)
			return nil
		},
		Jar: NewJar(),
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// Not Save Cookie
	NoCookieClient = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			Logger.Debugf("[GoWorker] Redirect:%v", req.URL)
			return nil
		},
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
)

// New a Proxy client, Default save cookie, Can timeout
// We should support some proxy way such as http(s) or socks
func NewProxyClient(proxyString string) (*http.Client, error) {
	proxyUrl, err := url.Parse(proxyString)
	if err != nil {
		return nil, err
	}

	prefix := strings.Split(proxyString, ":")[0]

	// setup a http transport
	httpTransport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}

	// http://
	// https://
	// socks5://
	switch prefix {
	case "http", "https":
		httpTransport.Proxy = http.ProxyURL(proxyUrl)
	case "socks5":
		// create a socks5 dialer
		dialer, err := proxy.FromURL(proxyUrl, proxy.Direct)
		if err != nil {
			return nil, err
		}
		httpTransport.Dial = dialer.Dial
	default:
		return nil, errors.New("this proxy way not allow:" + prefix)
	}

	// This a alone client, diff from global client.
	client := &http.Client{
		// Allow redirect
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			Logger.Debugf("[GoWorker] Redirect:%v", req.URL)
			return nil
		},
		// Allow proxy: http, https, socks5
		Transport: httpTransport,
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
			Logger.Debugf("[GoWorker] Redirect:%v", req.URL)
			return nil
		},
		Jar:     NewJar(),
		Timeout: util.Second(DefaultTimeOut),
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	return client, nil
}

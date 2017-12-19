package miner

import (
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/hunterhug/parrot/util"
	"golang.org/x/net/proxy" // see https://github.com/golang/net
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
	}

	// Not Save Cookie
	NoCookieClient = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			Logger.Debugf("[GoWorker] Redirect:%v", req.URL)
			return nil
		},
	}
)

// New a Proxy client, Default save cookie, Can timeout
// We should support some proxy way such as http(s) or socks
func NewProxyClient(proxystring string) (*http.Client, error) {
	proxyUrl, err := url.Parse(proxystring)
	if err != nil {
		return nil, err
	}

	prefix := strings.Split(proxystring, ":")[0]

	// setup a http transport
	httpTransport := &http.Transport{}

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
	}
	return client, nil
}

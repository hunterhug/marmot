/*
	Proxy Spider
	You first should own a remote machine, Then in your local tap:
		`ssh -ND 1080 ubuntu@remoteIp`
	It will gengerate socks5 proxy client in your local, which port is 1080
*/

package main

import (
	"fmt"
	"os"

	"github.com/hunterhug/GoSpider/query"
	"github.com/hunterhug/GoSpider/spider"
)

func init() {
	spider.SetLogLevel(spider.DEBUG)
}

func main() {
	// You can use a lot of proxy ip such "https/http/socks5"
	proxy_ip := "socks5://127.0.0.1:1080"

	url := "https://www.google.com"

	sp, err := spider.New(proxy_ip)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	body, err := sp.SetUa(spider.RandomUa()).SetUrl(url).SetMethod(spider.GET).Go()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(parse(body))
	}
}

// Parse HTML page
func parse(data []byte) string {
	doc, err := query.QueryBytes(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return doc.Find("title").Text()
}

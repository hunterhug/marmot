package main

import (
	"fmt"
	"github.com/hunterhug/GoSpider/spider"
)

func main() {
	// Use Default Spider, You can Also New One:
	// sp:=spider.New(nil)
	spider.SetLogLevel(spider.DEBUG)
	_, err := spider.SetUrl("https://www.whitehouse.gov").Go()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(spider.ToString())
	}
}

package main

import (
	"fmt"
	"github.com/hunterhug/GoSpider/spider"
)

func main() {
	// Use Default Spider, You can Also New One:
	// sp:=spider.New(nil)
	spider.SetLogLevel(spider.DEBUG)
	spider.SetUrl("http://www.google.com")
	_, err := spider.Go()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(spider.ToString())
	}
}

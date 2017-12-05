/*
	New A Spider
*/
package main

// Example
import (
	"fmt"

	"github.com/hunterhug/GoSpider/spider"
)

func main() {
	// 1. New a spider
	sp, _ := spider.New(nil)
	// 2. Set a URL And Fetch
	html, err := sp.SetUrl("https://www.whitehouse.gov").SetUa(spider.RandomUa()).SetMethod(spider.GET).Go()
	if err != nil {
		fmt.Println(err.Error())
	}
	// 4.Print content equal to fmt.Println(sp.ToString())
	fmt.Println(string(html))
}

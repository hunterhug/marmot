package main

import (
	"fmt"
	"strings"

	"github.com/hunterhug/GoSpider/query"
	"github.com/hunterhug/GoSpider/spider"
)

func main() {
	// We can debug, to see whether SetBeforeAction make sense
	spider.SetLogLevel(spider.DEBUG)

	// The url we want
	url := "https://www.whitehouse.gov"

	// IAM we can NewAPI
	sp := spider.NewAPI()

	// Before we make some change, And every GET Or POST it will action
	sp.SetBeforeAction(func(this *spider.Spider) {
		fmt.Println("Before Action, I will add a HTTP header")
		this.SetHeaderParm("GoSpider", "v2")
		this.SetHeaderParm("DUDUDUU", "DUDU")
	})

	sp.SetAfterAction(func(this *spider.Spider) {
		fmt.Println("After Action, I just print this sentence")
	})

	// Let's Go
	body, err := sp.SetUrl(url).GoByMethod(spider.GET)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// Parse We want
		fmt.Printf("Output:\n %s\n", parse(body))
	}
}

// Parse HTML page
func parse(data []byte) string {
	doc, err := query.QueryBytes(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(doc.Find("#hero-caption").Text())
	// return doc.Find("title").Text()
}

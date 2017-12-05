package main

import (
	"context"
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

	// We can aop by context
	// ctx, cancle := context.WithCancel(context.Background())
	// ctx := context.TODO()
	// sp.SetContext(ctx)

	// Before we make some change, And every GET Or POST it will action
	sp.SetBeforeAction(func(ctx context.Context, this *spider.Spider) {
		fmt.Println("Before Action, I will add a HTTP header")
		this.SetHeaderParm("GoSpider", "v2")
		this.SetHeaderParm("DUDUDUU", "DUDU")
		// select {
		// case <-ctx.Done():
		// 	fmt.Println(ctx.Err()) // block in here util cancle()
		// 	os.Exit(1)
		// }
	})

	// we cancle it after 5 secord
	// go func() {
	// 	util.Sleep(5)
	// 	cancle()
	// }()

	sp.SetAfterAction(func(ctx context.Context, this *spider.Spider) {
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

	// for {
	//  in here we loop util cancle() success
	// }
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

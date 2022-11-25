package main

import (
	"fmt"
	"github.com/hunterhug/marmot/miner"
)

/*
New A Worker!
*/
func main() {
	// 1. New a worker
	worker, _ := miner.New(nil)

	// 2. Set a URL And Fetch
	html, err := worker.SetUrl("https://www.bing.com").SetUa(miner.RandomUa()).SetMethod(miner.GET).Go()
	if err != nil {
		fmt.Println(err.Error())
	}

	// 3.Print content equal to fmt.Println(worker.ToString())
	fmt.Println(string(html))
}

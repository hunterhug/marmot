/*
	New A  Worker!
*/
package main

// Example
import (
	"fmt"

	"github.com/hunterhug/marmot/miner"
)

func main() {
	// 1. New a worker
	worker, _ := miner.New(nil)
	// 2. Set a URL And Fetch
	html, err := worker.SetUrl("https://www.whitehouse.gov").SetUa(miner.RandomUa()).SetMethod(miner.GET).Go()
	if err != nil {
		fmt.Println(err.Error())
	}
	// 4.Print content equal to fmt.Println(worker.ToString())
	fmt.Println(string(html))
}

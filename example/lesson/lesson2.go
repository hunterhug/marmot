/*
	All right reserved https://github.com/hunterhug/marmot at 2016-2020
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
*/
package main

import (
	"fmt"
	"github.com/hunterhug/marmot/miner"
)

/*
	New A  Worker!
*/
func main() {
	// 1. New a worker
	worker, _ := miner.New(nil)
	// 2. Set a URL And Fetch
	html, err := worker.SetUrl("http://www.github.com/hunterhug").SetUa(miner.RandomUa()).SetMethod(miner.GET).Go()
	if err != nil {
		fmt.Println(err.Error())
	}
	// 4.Print content equal to fmt.Println(worker.ToString())
	fmt.Println(string(html))
}

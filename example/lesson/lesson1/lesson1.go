/*
	All right reserved https://github.com/hunterhug/marmot at 2016-2021
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
	Most Simple: Use Default Worker!
*/
func main() {
	miner.SetLogLevel(miner.DEBUG)

	// Use Default Worker, You can Also New One:
	//worker, _ := miner.New(nil)
	//worker = miner.NewWorkerWithNoProxy()
	//worker = miner.NewAPI()
	//worker, _ = miner.NewWorkerWithProxy("socks5://127.0.0.1:1080")
	worker := miner.Clone()
	_, err := worker.SetUrl("https://www.bing.com").Go()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(worker.ToString())
	}
}

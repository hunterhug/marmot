/*
	All right reserved https://github.com/hunterhug/marmot at 2016-2020
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
*/

/*
Package miner is the core of this project, use to request for http api.

Example:

	package main

	import (
	"fmt"

	"github.com/hunterhug/marmot/miner"
	)

	func main() {
		// Use Default Worker, You can Also New One:
		// worker:=miner.New(nil)
		miner.SetLogLevel(miner.DEBUG)
		_, err := miner.SetUrl("https://github.com/hunterhug").Go()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(miner.ToString())
		}
	}
*/
package miner

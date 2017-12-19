// Package miner is the core of this project, use to request for http api.
//
// Example:
//
//	package main
//
//	import (
//	"fmt"
//
//	"github.com/hunterhug/marmot/miner"
//	)
//
//	func main() {
//		// Use Default Worker, You can Also New One:
//		// worker:=miner.New(nil)
//		miner.SetLogLevel(miner.DEBUG)
//		_, err := miner.SetUrl("https://www.whitehouse.gov").Go()
//		if err != nil {
//			fmt.Println(err.Error())
//		} else {
//			fmt.Println(miner.ToString())
//		}
//	}
package miner

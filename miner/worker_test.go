/*
	All right reserved https://github.com/hunterhug/marmot at 2016-2020
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
*/
package miner

import (
	"fmt"
	"testing"
)

func TestWorker(t *testing.T) {
	// Global log record
	SetLogLevel("debug")

	// GLOBAL TIMEOUT
	SetGlobalTimeout(3)

	// a new spider without proxy
	// NewWorker(nil)
	worker, err := NewWorker(nil)

	//proxy := "http://smart:smart2016@104.128.121.46:808"
	//worker, err := NewWorker(proxy)

	if err != nil {
		panic(err)
	}
	// method can be get and post
	worker.SetMethod(GET)
	// wait times,can zero
	worker.SetWaitTime(1)
	// which url fetch
	worker.SetUrl("https://github.com/hunterhug")

	worker.SetUa(RandomUa())

	// go!fetch url --||
	body, err := worker.Go()
	if err != nil {
		Log().Error(err.Error())
	} else {
		// bytes get!
		fmt.Printf("%s", string(body))
	}

	// if file size small than 500KB
	err = TooSortSizes(body, 500)
	Log().Error(err.Error())
}

package miner

import (
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
	worker.SetUrl("http://www.cjhug.me")

	//worker.SetUa(spider.RandomUa())

	// go!fetch url --||
	body, err := worker.Go()
	if err != nil {
		Log().Error(err.Error())
	} else {
		// bytes get!
		// fmt.Printf("%s", string(body))
	}

	// if filesize small than 500KB
	err = TooSortSizes(body, 500)
	Log().Error(err.Error())
}

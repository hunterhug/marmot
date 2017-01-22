package spider

import (
	"testing"
	"fmt"
)

func TestSpider(t *testing.T) {

	// global log record
	//SetLogLevel("DEBUg")
	SetLogLevel("debug")
	// GLOBAL TIMEOUT
	SetGlobalTimeout(3)

	// a new spider without proxy
	// NewSpider(nil)
	proxy := "http://smart:smart2016@104.128.121.46:808"
	spiders,err := NewSpider(proxy)
	if err!=nil{
		panic(err)
	}
	// method can be get and post
	spiders.Method = "get"
	// wait times,can zero
	spiders.Wait = 2
	// which url fetch
	spiders.Url = "http://www.goole.com"

	// a new header,default ua, no refer
	spiders.NewHeader(nil, "www.google.com", nil)


	// go!fetch url --||
	body, err := spiders.Go()
	if err != nil {
		Log().Error(err)
	} else {
		// bytes get!
		fmt.Printf("%s", string(body))
	}

	Log().Debugf("%#v",spiders)

	// if filesize small than 500KB
	err = TooSortSizes(body, 500)

	Log().Error(err.Error())
}

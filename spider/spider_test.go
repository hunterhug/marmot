/*
Copyright 2017 hunterhug/一只尼玛.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package spider

import (
	"fmt"
	"testing"
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
	spiders, err := NewSpider(proxy)
	if err != nil {
		panic(err)
	}
	// method can be get and post
	spiders.Method = "get"
	// wait times,can zero
	spiders.Wait = 2
	// which url fetch
	spiders.Url = "http://www.baidu.com"

	// a new header,default ua, no refer
	spiders.NewHeader(nil, "www.baidu.com", nil)

	// go!fetch url --||
	body, err := spiders.Go()
	if err != nil {
		Log().Error(err.Error())
	} else {
		// bytes get!
		fmt.Printf("%s", string(body))
	}

	Log().Debugf("%#v", spiders)

	// if filesize small than 500KB
	err = TooSortSizes(body, 500)

	Log().Error(err.Error())
}

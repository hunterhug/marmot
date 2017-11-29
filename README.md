# Project: Marmot(Tubo)

[![GitHub forks](https://img.shields.io/github/forks/hunterhug/GoSpider.svg?style=social&label=Forks)](https://github.com/hunterhug/GoSpider/network)
[![GitHub stars](https://img.shields.io/github/stars/hunterhug/GoSpider.svg?style=social&label=Stars)](https://github.com/hunterhug/GoSpider/stargazers)
[![GitHub last commit](https://img.shields.io/github/last-commit/hunterhug/GoSpider.svg)](https://github.com/hunterhug/GoSpider)
[![Go Report Card](https://goreportcard.com/badge/github.com/hunterhug/GoSpider)](https://goreportcard.com/report/github.com/hunterhug/GoSpider)
[![GitHub issues](https://img.shields.io/github/issues/hunterhug/GoSpider.svg)](https://github.com/hunterhug/GoSpider/issues)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/hunterhug/GoSpider/master/LICENSE)

[中文介绍](/doc/Chinese.md)

![Marmot](/doc/tubo.png)

## 1. Introduction

World-Wide-Web robot, also known as spiders and crawlers. The principle is to falsify network data by constructing appointed HTTP protocol data packet, then request resource to the specified host, goal is to access the data returned. 
There are a large number of web information, human's hand movement such as `copy-paste data` from web page is `time-consuming` and `laborious`, thus inspired the data acquisition industry.

Batch access to public network data does not break the law, but because there is no difference, no control, very violent means will lead to other services is not stable, therefore, most of the resources provider will filtering some data packets(falsify), 
in this context,  batch small data acquisition has become a problem. Integrated with various requirements, such as various API development, automated software testing(all this have similar technical principle). So this project come into the world(very simple).

The `Marmot` is very easy to understand, just like Python's library `requests`(Not yet Smile~ --| ). By enhancing native Golang HTTP library, help you deal with some trivial logic (such as collecting information, checking parameters), and add some fault-tolerant mechanisms (such as add lock, close time flow, ensure the high concurrent run without accident).
It provides a human friendly API interface, you can reuse it often. Very convenient to support `Cookie Persistence`, `Crawler Proxy Settings`, as well as others general settings, such as  `HTTP request header settings, timeout/pause settings, data upload/post settings`.
It support all of the HTTP methods `POST/PUT/GET/DELETE/...` and has built-in spider pool and browser UA pool, easy to develop UA+Cookie persistence distributed spider.

In addition, It also provides third party tool package. The library is simple and practical, just a few lines of code to replace the previous `Spaghetti code`, has been applied in some large projects such as `Full Golang Automatic Amazon Distributed crawler|spider`, has withstood the test of two thousand long acting proxy IP and high concurrency, single machine every day to get millions of data.

The main uses: WeChat development/ API docking / automated test / rush ticket scripting / site monitoring / vote plug-in / data crawling

Now We support Default Spider, You can easy use:

```
package main

import (
	"fmt"
	"github.com/hunterhug/GoSpider/spider"
)

func main() {
	// Use Default Spider, You can Also New One:
	// sp:=spider.New(nil)
	spider.SetLogLevel(spider.DEBUG)
	spider.SetUrl("http://www.google.com")
	_, err := spider.Go()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(spider.ToString())
	}
}
```

## 2. How To Use

You can get it by:

```
go get -v github.com/hunterhug/GoSpider/spider
```

Or make your GOPATH sub dir: `/src/github.com/hunterhug`, and

```
cd src/github.com/hunterhug
git clone https://github.com/hunterhug/GoSpider
```

## 3. Example

The most simple example such following, more see `example` dir:

```go
package main

// Example
import (
	"fmt"
	"github.com/hunterhug/GoSpider/spider"
)

func main() {
	// 1. New a spider
	sp, _ := spider.New(nil)
	// 2. Set a URL 
	sp.SetUrl("http://www.cjhug.me").SetUa(spider.RandomUa()).SetMethod(spider.GET)
	// 3. Fetch
	html, err := sp.Go()
	if err != nil {
		fmt.Println(err.Error())
	}
	// 4.Print content equal to fmt.Println(sp.ToString())
	fmt.Println(string(html))
}
```

More detail Example is:

```go
package main

import (
	// 1:import package alias boss
	boss "github.com/hunterhug/GoSpider/spider"
	"github.com/hunterhug/GoTool/util"
)

func init() {
	// 2:Optional global setting
	boss.SetLogLevel(boss.DEBUG) // optional, set log to debug
	boss.SetGlobalTimeout(3)     // optional, http request timeout time

}
func main() {

	log := boss.Log() // optional, spider log you can choose to use

	// 3: Must new a spider object, three ways
	//sp, err := boss.NewSpider("http://smart:smart2016@104.128.121.46:808") // proxy spider, format: protocol://user(optional):password(optional)@ip:port
	//sp, err := boss.NewSpider(nil)  // normal spider, default keep Cookie
	//sp, err := boss.NewAPI() // API spider, not keep Cookie
	sp, err := boss.New(nil) // NewSpider alias
	if err != nil {
		panic(err)
	}

	// 4: Set the request Method/URL and some others, can chain set, only SetUrl is required.
	// SetUrl: required, the Url
	// SetMethod: optional, HTTP method: POST/GET/..., default GET
	// SetWaitTime: optional, HTTP request wait/pause time
	sp.SetUrl("http://www.google.com").SetMethod(boss.GET).SetWaitTime(2)
	sp.SetUa(boss.RandomUa())                 // optional, browser user agent: IE/Firefox...
	sp.SetRefer("http://www.google.com")      // optional, url refer
	sp.SetHeaderParm("diyheader", "lenggirl") // optional, some other diy http header
	//sp.SetBData([]byte("file data"))  // optional, if you want post JSON data or upload file
	//sp.SetFormParm("username","jinhan") // optional: if you want post form
	//sp.SetFormParm("password","123")

	// 5: Start Run
	//sp.Get()             // default GET
	//sp.Post()            // POST form request data, data can fill by SetFormParm()
	//sp.PostJSON()        // POST JSON dara, use SetBData()
	//sp.PostXML()         // POST XML, use SetBData()
	//sp.PostFILE()        // POST to Upload File, data in SetBData() too
	//sp.OtherGo("OPTIONS", "application/x-www-form-urlencoded") // Other http method, Such as OPTIONS etcd
	body, err := sp.Go() // if you use SetMethod(), otherwise equal to Get()
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Infof("%s", string(body)) // Print return data
	}

	log.Debugf("%#v", sp) // if you not set log as debug, it will not appear

    // You must Clear it! If you want to POST Data by SetFormParm()/SetBData() again
	// After get the return data by post data, you can clear the data you fill
	sp.Clear()

	//sp.ClearAll() // you can also want to clear all, include http header you set

	// Spider pool for concurrent, every Spider Object is serial as the browser. if you want collateral execution, use this.
	boss.Pool.Set("myfirstspider", sp)
	if poolspider, ok := boss.Pool.Get("myfirstspider"); ok {
		go func() {
			poolspider.SetUrl("http://www.baidu.com")
			data, _ := poolspider.Get()
			log.Info(string(data))
		}()
		util.Sleep(10)
	}
}
```

Easy to use, you just need to `New` one `Spider`, and `SetUrl`, then make some settings and `sp.Go()`.

### 3.1 The First Step

There are three kinds of spider:

1. `sp, err := boss.NewSpider("http://smart:smart2016@104.128.121.46:808") ` // proxy spider, format: `protocol://user(optional):password(optional)@ip:port`, alias to`New()`
2. `sp, err := boss.NewSpider(nil)`  // normal spider, default keep Cookie, alias to `New()`
3. `sp, err := boss.NewAPI()` // API spider, not keep Cookie

### 3.2 The Second Step

Camouflage our spider:

1. `sp.SetUrl("http://www.cjhug.me")`  // required: set url you want to
2. `sp.SetMethod(boss.GET)`  // optional: set http method `POST/GET/PUT/POSTJSON` and so on
3. `sp.SetWaitTime(2)`                         // optional: set timeout of http request
4. `sp.SetUa(boss.RandomUa())`                 // optional: set http browser user agent, you can see spider/config/ua.txt
5. `sp.SetRefer("http://www.baidu.com")`       // optional: set http request Refer
6. `sp.SetHeaderParm("diyheader", "lenggirl")` // optional: set http diy header
7. `sp.SetBData([]byte("file data"))` // optional: set binary data for post or put
8. `sp.SetFormParm("username","jinhan")` // optional: set form data for post or put 
9. `sp.SetCookie("xx=dddd")` // optional: you can set a init cookie, some website you can login and F12 copy the cookie

### 3.3 The Third Step

Run our spider:

1. `body, err := sp.Go()` // if you use SetMethod(), auto use following ways, otherwise use Get()
2. `body, err := sp.Get()` // default
3. `body, err := sp.Post()` // post form request, data fill by SetFormParm()
4. `body, err := sp.PostJSON()` // post JSON request, data fill by SetBData()
5. `body, err := sp.PostXML()` // post XML request, data fill by SetBData()
6. `body, err := sp.PostFILE()` // upload file, data fill by SetBData()
7. `body, err := sp.Delete()` // you know!
8. `body, err := sp.Put()` // ones http method...
9. `body, err := sp.PutJSON()` // put JSON request
10. `body, err := sp.PutXML()`
11. `body, err := sp.PutFILE()`
12. `body, err := sp.OtherGo("OPTIONS", "application/x-www-form-urlencoded")` // Other http method, Such as OPTIONS etcd.

### 3.4 The Fourth Step

Deal the return data, all data will be return as binary:

1. `fmt.Println(string(html))` // type change directly
2. `fmt.Println(sp.ToString())` // use spider method, after http response, data will keep in the field Raw, just use ToString
3. `fmt.Println(sp.JsonToString())` // some json data will include chinese and other multibyte character, such as `我爱你,我的小绵羊`,`사랑해`

Attention: after every request for a url, the next request you can cover your http request header, otherwise header you set still exist,
if just want clear post data, use `Clear()`, and want clear header too please use `ClearAll()` .

More see the code source.

## 4. Project Application

It has already used in many project(although some is very simple) :

1. [Full Golang Automatic Amazon Distributed crawler|spider](https://github.com/hunterhug/AmazonBigSpider) // Just see [Picture](doc/amazon.md)
2. [Jiandan Distributed articles spider](https://github.com/hunterhug/jiandan)
3. [Zhihu API tool](https://github.com/hunterhug/GoZhihu)
4. [Picture helper](/example/taobao/README.md)
5. [Jiandan Picure helper](/example/jiandanmeizi/README.md)
6. [Music Download](/example/music/README.md) // see example dir
7. [GoTaobao](https://github.com/hunterhug/GoTaoBao)
8. A lot closed source... 

Project change you can see [log](/doc/log.md)

Install development environment you can refer:[GoSpider-docker](https://github.com/hunterhug/GoSpider-docker)

# LICENSE

```
Copyright 2017 by GoSpider author. Email: gdccmcm14@live.com
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License
```

Welcome Add PR/issues. 

For questions, please email: gdccmcm14@live.com.

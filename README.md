# Project: Marmot(Tubo)

[中文介绍](Chinese.md)

![Marmot](tubo.png)

## 1. Introduction

World-Wide-Web robot, also known as spiders and crawlers, the principle is falsifying network data by constructing appointed HTTP protocol data packet, request resource to the specified host, access to the data returned. 
There are a large number of Web information, human hand movement `copy-paste` data from web page is `time-consuming` and `laborious`, it inspired the crawler industry.

Batch access to public network data does not break the law, but because there is no difference, no control, very violent means will lead to other services is not stable, therefore, most of the resources provider will filtering some data packets(falsify), 
in this context,  batch small data acquisition has become a problem.Integrated with various requirements, such as various API development, automated software testing, they have similar technical principle so I write this project.

The `Marmot` is a human friendly gesture `Golang Library` which was developed using object-oriented method, easy to understand. By enhancing native Golang HTTP library package, help users deal with some trivial logic (such as collecting information, detection parameters), and add some fault-tolerant mechanisms (such as add lock, close time flow, ensure the high concurrent run without accident).

This library provides a excellent API interface, you can reuse it, very convenient to support `Cookie Persistence`, `Crawler Proxy Settings`, as well as general settings such as the `HTTP request header settings, timeout/pause settings, data upload/web form post settings`, 
support all of the HTTP methods such as `POST/PUT/GET/DELETE`, and also has built-in crawler pool and browser UA pool, easy development UA+Cookie persistence distributed crawler.

In addition, also provides third party tool package, such as support for `mysql/postgresql/redis/cassandra/hbase` and so on. The library is simple and practical, just a few lines of code to replace the previous `Spaghetti code`. 
has been applied in some large projects such as large-distributed crawler: `Projet: Amazon(USA/Japan/Germany/UK) `, has withstood the test of two thousand long acting proxy IP and high concurrency, single machine every day to get millions of data.

The main uses: WeChat development/ API docking / automated test / rush ticket scripting / site monitoring / vote plug-in / data crawling

## 2. How To Use

Just like Python's library `requests`, you can get it by:

```
go get -v github.com/hunterhug/GoSpider
```

Or make your GOPATH sub dir: `/src/github.com/hunterhug`, and

```
cd src/github.com/hunterhug
git clone https://github.com/hunterhug/GoSpider
```

all import package save in `vendor` default, if some panic when run, please move `the file in vendor` to your `GOPATH`. you can choose use `godep` (Optional)

```
godep restore
```

## 3. Project Structure(modularization)

```
    ---example   some example
    ---query     html content parse, only two function
    ---spider    core download module
    ---store     store tool
        ---myredis 
        ---mysql
        ---myetcd
        ---mydb  one database Orm(use xorm)
        ---myhbase
        ---mycassandra
    ---util      other tool
        --- image  picture cut
        --- open   open CAPTCHA picture
        crypto.go encryption..
        file.go  file operation
        time.go  time...
    ---vendor    some dependency package
    ---GOPATH    some must move to your GOPATH
```

## 4. Example

The most simple example such follow, more see `example` dir:

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
	sp.SetUrl("http://www.lenggirl.com").SetUa(spider.RandomUa()).SetMethod(spider.GET)
	// 3. Fetch
	html, err := sp.Go()
	if err != nil {
		fmt.Println(err.Error())
	}
	// 4.Print content equal to fmt.Println(sp.ToString())
	fmt.Println(string(html))
}
```

Detail Example is:

```go
package main

import (
	// 1:import package alias boss
	boss "github.com/hunterhug/GoSpider/spider"
	"github.com/hunterhug/GoSpider/util"
)

func init() {
	// 2:Optional global setting
	boss.SetLogLevel(boss.DEBUG) // optional, set log to debug
	boss.SetGlobalTimeout(3)     // optional, http request timeout time

}
func main() {

	log := boss.Log() // optional, spider log you can choose to use

	// 3: Must new a spider object, three ways
	//spiders, err := boss.NewSpider("http://smart:smart2016@104.128.121.46:808") // proxy spider, format: protocol://user(optional):password(optional)@ip:port
	//spiders, err := boss.NewSpider(nil)  // normal spider, default keep Cookie
	//spiders, err := boss.NewAPI() // API spider, not keep Cookie
	spiders, err := boss.New(nil) // NewSpider alias
	if err != nil {
		panic(err)
	}

	// 4: Set the request Method/URL and some others, can chain set, only SetUrl is required.
	// SetUrl: required, the Url
	// SetMethod: optional, HTTP method: POST/GET/..., default GET
	// SetWaitTime: optional, HTTP request wait/pause time
	spiders.SetUrl("http://www.google.com").SetMethod(boss.GET).SetWaitTime(2)
	spiders.SetUa(boss.RandomUa())                 // optional, browser user agent: IE/Firefox...
	spiders.SetRefer("http://www.google.com")      // optional, url refer
	spiders.SetHeaderParm("diyheader", "lenggirl") // optional, some other diy http header
	//spiders.SetBData([]byte("file data"))  // optional, if you want post JSON data or upload file
	//spiders.SetFormParm("username","jinhan") // optional: if you want post form
	//spiders.SetFormParm("password","123")

	// 5: Start Run
	//spiders.Get()             // default GET
	//spiders.Post()            // POST form request data, data can fill by SetFormParm()
	//spiders.PostJSON()        // POST JSON dara, use SetBData()
	//spiders.PostXML()         // POST XML, use SetBData()
	//spiders.PostFILE()        // POST to Upload File, data in SetBData() too
	body, err := spiders.Go() // if you use SetMethod(), otherwise equal to Get()
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Infof("%s", string(body)) // Print return data
	}

	log.Debugf("%#v", spiders) // if you not set log as debug, it will not appear

	spiders.Clear() // after get the return data by post data, you can clear the data you fill
	//spiders.ClearAll() // you can also want to clear all, include  http header you set

	// spider pool for concurrent, every Spider Object is serial such as the browser. if you want collateral execution, use this.
	boss.Pool.Set("myfirstspider", spiders)
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

Easy to use, you just need to `New` one `Spider`, and `SetUrl`, then add some http header and `spiders.Go()`.

### 4.1 The First Step

There are three kind of spider:

1. `spiders, err := boss.NewSpider("http://smart:smart2016@104.128.121.46:808") ` // proxy spider, format: `protocol://user(optional):password(optional)@ip:port` alias to`New()`
2. `spiders, err := boss.NewSpider(nil)`  // normal spider, default keep Cookie alias to `New()`
3. `spiders, err := boss.NewAPI()` // API spider, not keep Cookie

### 4.2 The Second Step

Camouflage our spider:

1. `spiders.SetUrl("http://www.lenggirl.com")`  // required: set url you want to
2. `spiders.SetMethod(boss.GET)`  // optional: set http method `POST/GET/PUT/POSTJSON` and so on
3. `spiders.SetWaitTime(2)`                         // optional: set timeout of http request
4. `spiders.SetUa(boss.RandomUa())`                 // optional: set http browser user agent, supply 445/see spider/config/ua.txt
5. `spiders.SetRefer("http://www.baidu.com")`       // optional: set http request Refer
6. `spiders.SetHeaderParm("diyheader", "lenggirl")` // optional: set http diy header
7. `spiders.SetBData([]byte("file data"))` // optional: set binary data for post or put
8. `spiders.SetFormParm("username","jinhan")` // optional: set form data for post or put 
9. `spiders.SetCookie("xx=dddd")` // optional: you can set a init cookie, some website you can login and F12 copy the cookie

### 4.3 The Third Step

Run our spider:

1. `body, err := spiders.Go()` // if you use SetMethod(), auto use following ways, otherwise use Get()
2. `body, err := spiders.Get()` // default
3. `body, err := spiders.Post()` // post form request, data fill by SetFormParm()
4. `body, err := spiders.PostJSON()` // post JSON request, data fill by SetBData()
5. `body, err := spiders.PostXML()` // post XML request, data fill by SetBData()
6. `body, err := spiders.PostFILE()` // upload file, data fill by SetBData()
7. `body, err := spiders.Delete()` // you know!
8. `body, err := spiders.Put()` // a http method maybe
9. `body, err := spiders.PutJSON()` // put JSON request
10. `body, err := spiders.PutXML()`
11. `body, err := spiders.PutFILE()`
12. `body, err := spiders.OtherGo("OPTIONS", "application/x-www-form-urlencoded")` // Other http method, Such as OPTIONS etcd.

### 4.4 The Fourth Step

Deal the return data, all data will be return as binary:

1. `fmt.Println(string(html))` // type change directly
2. `fmt.Println(sp.ToString())` // use spider method, after http response, data will keep in the field Raw, just use ToString
3. `fmt.Println(sp.JsonToString())` // some json data will include chinese and other multibyte character, such as `我爱你,我的小绵羊`,`사랑해`

Attention: After every request for a url, the next request you can cover your http request header, otherwise header you set still exist,
if just want clear post data, use `Clear()`, and want clear header too please use `ClearAll()` .

More see the code source.

## 5. Project Application

It has already used in many project(although some is very simple) :

1. [Full Golang Automatic Amazon Distributed crawler|spider](https://github.com/hunterhug/AmazonBigSpider) // Just see [Picture](doc/amazon.md)
2. [Jiandan Distributed articles spider](https://github.com/hunterhug/jiandan)
3. [Zhihu API tool](https://github.com/hunterhug/zhihuxx)
4. [Picture helper](/example/taobao/README.md)
5. [Jiandan Picure helper](/example/jiandanmeizi/README.md)
6. Music Download // see example file dir
7. A lot closed source... 

Project change you can see [log](doc/log.md)

Install development environment you can refer(still chinese):

[gospider-env](http://www.lenggirl.com/tool/gospider-env.html)

[GoSpider-docker](https://github.com/hunterhug/GoSpider-docker)

# LICENSE

Welcome Add PR/issues, Use Apache License (you can use if you want but you should add this in every code file)

```
Copyright 2017 by GoSpider author.
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

For questions, please email: gdccmcm14@live.com.
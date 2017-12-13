# Project: Marmot

[![GitHub forks](https://img.shields.io/github/forks/hunterhug/GoWorker.svg?style=social&label=Forks)](https://github.com/hunterhug/GoWorker/network)
[![GitHub stars](https://img.shields.io/github/stars/hunterhug/GoWorker.svg?style=social&label=Stars)](https://github.com/hunterhug/GoWorker/stargazers)
[![GitHub last commit](https://img.shields.io/github/last-commit/hunterhug/GoWorker.svg)](https://github.com/hunterhug/GoWorker)
[![Go Report Card](https://goreportcard.com/badge/github.com/hunterhug/GoWorker)](https://goreportcard.com/report/github.com/hunterhug/GoWorker)
[![GitHub issues](https://img.shields.io/github/issues/hunterhug/GoWorker.svg)](https://github.com/hunterhug/GoWorker/issues)

[中文介绍](/doc/Chinese.md)

HTTP Download Helper, Supports Many Features such as Cookie Persistence, HTTP(S) and SOCKS5 Proxy....

![Marmot](/doc/tubo.png)

[TOC]

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

Now We support Default Worker, You can easy use:

`lesson1.go`

```go
package main

import (
	"fmt"

	"github.com/hunterhug/marmot/miner"
)

func main() {
	// Use Default Worker, You can Also New One:
	// worker:=miner.New(nil)
	miner.SetLogLevel(miner.DEBUG)
	_, err := miner.SetUrl("https://www.whitehouse.gov").Go()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(miner.ToString())
	}
}
```

## 2. How To Use

You can get it by:

```
go get -v github.com/hunterhug/marmot/miner
```

Or make your GOPATH sub dir: `/src/github.com/hunterhug`, and

```
cd src/github.com/hunterhug
git clone https://github.com/hunterhug/marmot
```

Suggest Golang1.8+

## 3. Example

The most simple example such below, more see `example/lesson` dir:

`lesson2.go`

```go
package main

// Example
import (
	"fmt"

	"github.com/hunterhug/marmot/miner"
)

func main() {
	// 1. New a worker
	worker, _ := miner.New(nil)
	// 2. Set a URL And Fetch
	html, err := worker.SetUrl("https://www.whitehouse.gov").SetUa(miner.RandomUa()).SetMethod(miner.GET).Go()
	if err != nil {
		fmt.Println(err.Error())
	}
	// 4.Print content equal to fmt.Println(worker.ToString())
	fmt.Println(string(html))
}
```

More detail Example is:

`lesson3.go`

```go
package main

import (
	// 1:import package
	"github.com/hunterhug/marmot/miner"
	"github.com/hunterhug/parrot/util"
)

func init() {
	// 2:Optional global setting
	miner.SetLogLevel(miner.DEBUG) // optional, set log to debug
	miner.SetGlobalTimeout(3)      // optional, http request timeout time

}

func main() {

	log := miner.Log() // optional, miner log you can choose to use

	// 3: Must new a Worker object, three ways
	//worker, err := miner.NewWorker("http://smart:smart2016@104.128.121.46:808") // proxy format: protocol://user(optional):password(optional)@ip:port
	//worker, err := miner.NewWorker(nil)  // normal worker, default keep Cookie
	//worker, err := miner.NewAPI() // API worker, not keep Cookie
	worker, err := miner.New(nil) // NewWorker alias
	if err != nil {
		panic(err)
	}

	// 4: Set the request Method/URL and some others, can chain set, only SetUrl is required.
	// SetUrl: required, the Url
	// SetMethod: optional, HTTP method: POST/GET/..., default GET
	// SetWaitTime: optional, HTTP request wait/pause time
	worker.SetUrl("http://cjhug.me/fuck.html").SetMethod(miner.GET).SetWaitTime(2)
	worker.SetUa(miner.RandomUa())                // optional, browser user agent: IE/Firefox...
	worker.SetRefer("https://www.whitehouse.gov") // optional, url refer
	worker.SetHeaderParm("diyheader", "lenggirl") // optional, some other diy http header
	//worker.SetBData([]byte("file data"))    // optional, if you want post JSON data or upload file
	//worker.SetFormParm("username","jinhan") // optional: if you want post form
	//worker.SetFormParm("password","123")

	// 5: Start Run
	//worker.Get()             // default GET
	//worker.Post()            // POST form request data, data can fill by SetFormParm()
	//worker.PostJSON()        // POST JSON dara, use SetBData()
	//worker.PostXML()         // POST XML, use SetBData()
	//worker.PostFILE()        // POST to Upload File, data in SetBData() too
	//worker.OtherGo("OPTIONS", "application/x-www-form-urlencoded") // Other http method, Such as OPTIONS etcd
	body, err := worker.Go() // if you use SetMethod(), otherwise equal to Get()
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Infof("%s", string(body)) // Print return data
	}

	log.Debugf("%#v", worker.GetCookies) // if you not set log as debug, it will not appear

	// You must Clear it! If you want to POST Data by SetFormParm()/SetBData() again
	// After get the return data by post data, you can clear the data you fill
	worker.Clear()
	//worker.ClearAll() // you can also want to clear all, include http header you set

	// Worker pool for concurrent, every Worker Object is serial as the browser. if you want collateral execution, use this.
	miner.Pool.Set("myfirstworker", worker)
	if w, ok := miner.Pool.Get("myfirstworker"); ok {
		go func() {
			data, _ := w.SetUrl("http://cjhug.me/fuck.html").Get()
			log.Info(string(data))
		}()
		util.Sleep(10)
	}
}
```

Last example

`lesson4.go`

```go
package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/hunterhug/marmot/expert"
	"github.com/hunterhug/marmot/miner"
)

func main() {
	// We can debug, to see whether SetBeforeAction make sense
	miner.SetLogLevel(miner.DEBUG)

	// The url we want
	url := "https://www.whitehouse.gov"

	// IAM we can NewAPI
	worker := miner.NewAPI()

	// We can aop by context
	// ctx, cancle := context.WithCancel(context.Background())
	// ctx := context.TODO()
	// worker.SetContext(ctx)

	// Before we make some change, And every GET Or POST it will action
	worker.SetBeforeAction(func(ctx context.Context, this *miner.Worker) {
		fmt.Println("Before Action, I will add a HTTP header")
		this.SetHeaderParm("Marmot", "v2")
		this.SetHeaderParm("DUDUDUU", "DUDU")
		// select {
		// case <-ctx.Done():
		// 	fmt.Println(ctx.Err()) // block in here util cancle()
		// 	os.Exit(1)
		// }
	})

	// we cancle it after 5 secord
	// go func() {
	// 	util.Sleep(5)
	// 	cancle()
	// }()

	worker.SetAfterAction(func(ctx context.Context, this *miner.Worker) {
		fmt.Println("After Action, I just print this sentence")
	})

	// Let's Go
	body, err := worker.SetUrl(url).GoByMethod(miner.GET)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// Parse We want
		fmt.Printf("Output:\n %s\n", parse(body))
	}

	// for {
	//  in here we loop util cancle() success
	// }
}

// Parse HTML page
func parse(data []byte) string {
	doc, err := expert.QueryBytes(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(doc.Find("#hero-caption").Text())
	// return doc.Find("title").Text()
}

```

Easy to use, you just need to `New` one `Worker`, and `SetUrl`, then make some settings and `worker.Go()`.

### 3.1 The First Step

There are four kinds of worker:

1. `worker, err := miner.NewWorker("http://smart:smart2016@104.128.121.46:808") ` // proxy worker, format: `protocol://user(optional):password(optional)@ip:port`, alias to`New()`, support http(s), socks5
2. `worker, err := miner.NewWorker(nil)`  // normal worker, default keep Cookie, alias to `New()`
3. `worker := miner.NewAPI()` // API worker, will not keep Cookie
4. `worker, err := miner.NewWorkerByClient(&http.Client{})` // You can also pass a `http.Client` if you want

### 3.2 The Second Step

Camouflage our worker:

1. `worker.SetUrl("https://www.whitehouse.gov")`  // required: set url you want to
2. `worker.SetMethod(miner.GET)`  // optional: set http method `POST/GET/PUT/POSTJSON` and so on
3. `worker.SetWaitTime(2)`                         // optional: set timeout of http request
4. `worker.SetUa(miner.RandomUa())`                 // optional: set http browser user agent, you can see miner/config/ua.txt
5. `worker.SetRefer("https://www.whitehouse.gov")`       // optional: set http request Refer
6. `worker.SetHeaderParm("diyheader", "lenggirl")` // optional: set http diy header
7. `worker.SetBData([]byte("file data"))` // optional: set binary data for post or put
8. `worker.SetFormParm("username","jinhan")` // optional: set form data for post or put 
9. `worker.SetCookie("xx=dddd")` // optional: you can set a init cookie, some website you can login and F12 copy the cookie
10. `worker.SetCookieByFile("/root/cookie.txt")` // optional: set cookie which store in a file

### 3.3 The Third Step

Run our worker:

1. `body, err := worker.Go()` // if you use SetMethod(), auto use following ways, otherwise use Get()
2. `body, err := worker.Get()` // default
3. `body, err := worker.Post()` // post form request, data fill by SetFormParm()
4. `body, err := worker.PostJSON()` // post JSON request, data fill by SetBData()
5. `body, err := worker.PostXML()` // post XML request, data fill by SetBData()
6. `body, err := worker.PostFILE()` // upload file, data fill by SetBData()
7. `body, err := worker.Delete()` // you know!
8. `body, err := worker.Put()` // ones http method...
9. `body, err := worker.PutJSON()` // put JSON request
10. `body, err := worker.PutXML()`
11. `body, err := worker.PutFILE()`
12. `body, err := worker.OtherGo("OPTIONS", "application/x-www-form-urlencoded")` // Other http method, Such as OPTIONS etcd.
13. `body, err := worker.GoByMethod("POST")` // you can override SetMethod() By this, equal SetMethod() then Go()

### 3.4 The Fourth Step

Deal the return data, all data will be return as binary, You can immediately store it into a new variable:

1. `fmt.Println(string(html))` // type change directly
2. `fmt.Println(worker.ToString())` // use spider method, after http response, data will keep in the field `Raw`, just use ToString
3. `fmt.Println(worker.JsonToString())` // some json data will include chinese and other multibyte character, such as `我爱你,我的小绵羊`,`사랑해`

Attention: after every request for a url, the next request you can cover your http request header, otherwise header you set still exist,
if just want clear post data, use `Clear()`, and want clear HTTP header too please use `ClearAll()` .

Here is a example `lesson6.go` again:

```go
package main

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/hunterhug/marmot/expert"
	"github.com/hunterhug/marmot/miner"
	"github.com/hunterhug/parrot/util"
)

// Num of miner, We can run it at the same time to crawl data fast
var MinerNum = 5

// You can update this decide whether to proxy
var ProxyAddress interface{}

func main() {
	// You can Proxy!
	// ProxyAddress = "socks5://127.0.0.1:1080"

	fmt.Println(`
		Welcome: Input "url" and picture keep "dir"

		
		`)
	for {
		fmt.Println("---------------------------------------------")
		url := util.Input(`URL(Like: "http://publicdomainarchive.com")`, "http://publicdomainarchive.com")
		dir := util.Input(`DIR(Default: "./picture")`, "./picture")
		fmt.Printf("You will keep %s picture in dir %s\n", url, dir)
		fmt.Println("---------------------------------------------")

		// Start Catch
		err := CatchPicture(url, dir)
		if err != nil {
			fmt.Println("Error:" + err.Error())
		}
	}
}

// Come on!
func CatchPicture(picture_url string, dir string) error {
	// Check valid
	_, err := url.Parse(picture_url)
	if err != nil {
		return err
	}

	// Make dir!
	err = util.MakeDir(dir)
	if err != nil {
		return err
	}

	// New a worker to get url
	worker, _ := miner.New(ProxyAddress)

	result, err := worker.SetUrl(picture_url).SetUa(miner.RandomUa()).Get()
	if err != nil {
		return err
	}

	// Find all picture
	pictures := expert.FindPicture(string(result))

	// Empty, What a pity!
	if len(pictures) == 0 {
		return errors.New("empty")
	}

	// Devide pictures into several worker
	xxx, _ := util.DevideStringList(pictures, MinerNum)

	// Chanel to info exchange
	chs := make(chan int, len(pictures))

	// Go at the same time
	for num, imgs := range xxx {

		// Get pool miner
		worker_picture, ok := miner.Pool.Get(util.IS(num))
		if !ok {
			// No? set one!
			worker_temp, _ := miner.New(ProxyAddress)
			worker_picture = worker_temp
			worker_temp.SetUa(miner.RandomUa())
			miner.Pool.Set(util.IS(num), worker_temp)
		}

		// Go save picture!
		go func(imgs []string, worker *miner.Worker, num int) {
			for _, img := range imgs {

				// Check, May be Pass
				_, err := url.Parse(img)
				if err != nil {
					continue
				}

				// Change Name of our picture
				filename := strings.Replace(util.ValidFileName(img), "#", "_", -1)

				// Exist?
				if util.FileExist(dir + "/" + filename) {
					fmt.Println("File Exist：" + dir + "/" + filename)
					chs <- 0
				} else {

					// Not Exsit?
					imgsrc, e := worker.SetUrl(img).Get()
					if e != nil {
						fmt.Println("Download " + img + " error:" + e.Error())
						chs <- 0
						return
					}

					// Save it!
					e = util.SaveToFile(dir+"/"+filename, imgsrc)
					if e == nil {
						fmt.Printf("SP%d: Keep in %s/%s\n", num, dir, filename)
					}
					chs <- 1
				}
			}
		}(imgs, worker_picture, num)
	}

	// Every picture should return
	for i := 0; i < len(pictures); i++ {
		<-chs
	}

	return nil
}
```

More see the code source.

## 4. Project Application

It has already used in many project(although some is very simple) :

1. [Full Golang Automatic Amazon Distributed crawler|spider](https://github.com/hunterhug/AmazonBigSpider) // Just see [Picture](doc/amazon.md)
2. [Jiandan Distributed articles spider](https://github.com/hunterhug/jiandan)
3. [Zhihu API tool](https://github.com/hunterhug/GoZhihu)
4. [GoTaobao](https://github.com/hunterhug/GoTaoBao)
5. A lot closed source...

Project change you can see [log](/doc/log.md), Install development environment you can refer:[GoSpider-docker](https://github.com/hunterhug/GoSpider-docker)
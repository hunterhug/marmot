# Project: Marmot | HTTP Download

[![GitHub forks](https://img.shields.io/github/forks/hunterhug/marmot.svg?style=social&label=Forks)](https://github.com/hunterhug/marmot/network)
[![GitHub stars](https://img.shields.io/github/stars/hunterhug/marmot.svg?style=social&label=Stars)](https://github.com/hunterhug/marmot/stargazers)
[![GitHub last commit](https://img.shields.io/github/last-commit/hunterhug/marmot.svg)](https://github.com/hunterhug/marmot)
[![Go Report Card](https://goreportcard.com/badge/github.com/hunterhug/marmot)](https://goreportcard.com/report/github.com/hunterhug/marmot)
[![GitHub issues](https://img.shields.io/github/issues/hunterhug/marmot.svg)](https://github.com/hunterhug/marmot/issues)
[![License](https://img.shields.io/badge/license-Apache%202-4EB1BA.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

[中文介绍](/README_ZH.md)

HTTP Download Helper, Supports Many Features such as Cookie Persistence, HTTP(S) and SOCKS5 Proxy....

![Marmot](logo.png)

## 1. Introduction

World-Wide-Web robot, also known as spiders and crawlers. The principle is to falsify network data by constructing appointed HTTP protocol data packet, then request resource to the specified host, goal is to access the data returned. 
There are a large number of web information, human's hand movement such as `copy-paste data` from web page is `time-consuming` and `laborious`, thus inspired the data acquisition industry.

Batch access to public network data does not break the law, but because there is no difference, no control, very violent means will lead to other services is not stable, therefore, most of the resources provider will filtering some data packets(falsify), 
in this context,  batch small data acquisition has become a problem. Integrated with various requirements, such as various API development, automated software testing(all this have similar technical principle). So this project come into the world(very simple).

The `Marmot` is very easy to understand, just like Python's library `requests`(Not yet Smile~ --| ). By enhancing native Golang HTTP library, help you deal with some trivial logic (such as collecting information, checking parameters), and add some fault-tolerant mechanisms (such as add lock, close time flow, ensure the high concurrent run without accident).
It provides a human friendly API interface, you can reuse it often. Very convenient to support `Cookie Persistence`, `Crawler Proxy Settings`, as well as others general settings, such as  `HTTP request header settings, timeout/pause settings, data upload/post settings`.
It support all of the HTTP methods `POST/PUT/GET/DELETE/...` and has built-in spider pool and browser UA pool, easy to develop UA+Cookie persistence distributed spider.

The library is simple and practical, just a few lines of code to replace the previous `Spaghetti code`, has been applied in some large projects.

The main uses: `WeChat development`/ `API docking` / `Automated test` / `Rush Ticket Scripting` / `Vote Plug-in` / `Data Crawling`

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
	_, err := miner.Clone().SetUrl("https://github.com/hunterhug").Go()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(miner.ToString())
	}
}
```

See the [example](example) dir. such as lesson or practice.

## 2. How To Use

You can get it by:

```
go get -v github.com/hunterhug/marmot/miner
```

### 2.1 The First Step

There are four kinds of worker:

1. `worker, err := miner.NewWorker("http://smart:smart2016@104.128.121.46:808") ` // proxy worker, format: `protocol://user(optional):password(optional)@ip:port`, alias to`New()`, support http(s), socks5
2. `worker, err := miner.NewWorker(nil)`  // normal worker, default keep Cookie, alias to `New()`
3. `worker := miner.NewAPI()` // API worker, will not keep Cookie
4. `worker, err := miner.NewWorkerByClient(&http.Client{})` // You can also pass a `http.Client` if you want

if you want to use worker twice, you can call `Clone()` method to clone a new worker, it can isolate the request and response of http, otherwise, you should deal concurrent program carefully.

### 2.2 The Second Step

Camouflage our worker:

1. `worker.SetUrl("https://github.com/hunterhug")`  // required: set url you want to
2. `worker.SetMethod(miner.GET)`  // optional: set http method `POST/GET/PUT/POSTJSON` and so on
3. `worker.SetWaitTime(2)`                         // optional: set timeout of http request
4. `worker.SetUa(miner.RandomUa())`                 // optional: set http browser user agent, you can see miner/config/ua.txt
5. `worker.SetRefer("https://github.com/hunterhug")`       // optional: set http request Refer
6. `worker.SetHeaderParm("diyheader", "diy")` // optional: set http diy header
7. `worker.SetBData([]byte("file data"))` // optional: set binary data for post or put
8. `worker.SetFormParm("username","jinhan")` // optional: set form data for post or put 
9. `worker.SetCookie("xx=dddd")` // optional: you can set a init cookie, some website you can login and F12 copy the cookie
10. `worker.SetCookieByFile("/root/cookie.txt")` // optional: set cookie which store in a file

### 2.3 The Third Step

Run our worker:

1. `body, err := worker.Go()` // if you use SetMethod(), auto use following ways, otherwise use Get()
2. `body, err := worker.Get()` // default
3. `body, err := worker.Post()` // post form request, data fill by SetFormParm()
4. `body, err := worker.PostJSON()` // post JSON request, data fill by SetBData()
5. `body, err := worker.PostXML()` // post XML request, data fill by SetBData()
6. `body, err := worker.PostFILE()` // upload file, data fill by SetBData(), and should set SetFileInfo(fileName, fileFormName string)
7. `body, err := worker.Delete()` // you know!
8. `body, err := worker.Put()` // ones http method...
9. `body, err := worker.PutJSON()` // put JSON request
10. `body, err := worker.PutXML()`
11. `body, err := worker.PutFILE()`
12. `body, err := worker.OtherGo("OPTIONS", "application/x-www-form-urlencoded")` // Other http method, Such as OPTIONS etc., can not sent binary.
13. `body, err := worker.OtherGoBinary("OPTIONS", "application/x-www-form-urlencoded")` // Other http method, Such as OPTIONS etc., just sent binary.
14. `body, err := worker.GoByMethod("POST")` // you can override SetMethod() By this, equal SetMethod() then Go()

### 2.4 The Fourth Step

Deal the return data, all data will be return as binary, You can immediately store it into a new variable:

1. `fmt.Println(string(html))` // type change directly
2. `fmt.Println(worker.ToString())` // use spider method, after http response, data will keep in the field `Raw`, just use ToString
3. `fmt.Println(worker.JsonToString())` // some json data will include chinese and other multibyte character, such as `我爱你,我的小绵羊`,`사랑해`

Attention: 

After every request for a url, the next request you should cover your http request header, otherwise http header you set still exist,
if just want clear post data, use `Clear()`, and want clear HTTP header too please use `ClearAll()`, but I suggest use `Clone()` to avoid this.

### 2.5 Other

Hook:

1. `SetBeforeAction(fc func(context.Context, *Worker))`
2. `SetAfterAction(fc func(context.Context, *Worker))`

# License

```
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

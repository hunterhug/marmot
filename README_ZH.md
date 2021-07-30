# 项目代号：土拨鼠

[![GitHub forks](https://img.shields.io/github/forks/hunterhug/marmot.svg?style=social&label=Forks)](https://github.com/hunterhug/marmot/network)
[![GitHub stars](https://img.shields.io/github/stars/hunterhug/marmot.svg?style=social&label=Stars)](https://github.com/hunterhug/marmot/stargazers)
[![GitHub last commit](https://img.shields.io/github/last-commit/hunterhug/marmot.svg)](https://github.com/hunterhug/marmot)
[![Go Report Card](https://goreportcard.com/badge/github.com/hunterhug/marmot)](https://goreportcard.com/report/github.com/hunterhug/marmot)
[![GitHub issues](https://img.shields.io/github/issues/hunterhug/marmot.svg)](https://github.com/hunterhug/marmot/issues)
[![License](https://img.shields.io/badge/license-Apache%202-4EB1BA.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

[English README](/README.md)

![Marmot](logo.png)

>万维网网络机器人，又称蜘蛛，爬虫，原理主要是通过构造符合HTTP协议的网络数据包，向指定主机请求资源，获取返回的数据.万维网有大量的公开信息，人力采集数据费时费力，故激发了爬虫的产业化.
批量获取公开网络数据并不违反，但由于无差别性，无节制，十分暴力的手段会导致对方服务的不稳定，因此，大部分资源提供商对数据包进行了某些过滤，在此背景下，小批量数据获取成为了难题.
综合各种需求，如各种API对接，自动化测试等原理均一样，故开发了此爬虫库.

>土拨鼠项目是一个人类友好姿势的代码库，开发采用面向对象的方式，易于理解.通过对Golang原生HTTP库的封装，帮用户处理了一些琐碎逻辑(如收集信息，检测参数)，并加入了一些容错机制(如加锁，及时关闭流)，保证了爬虫高并发的安全.此库提供了大量优美的API接口，复用率高，十分方便地支持Cookie接力，爬虫代理设置，以及一般的HTTP请求设置如头部设置，超时，暂停设置，数据设置等，支持全部的HTTP方法如POST/PUT/GET/DELETE等，内置爬虫池和浏览器UA池，易于开发多UA多Cookie分布式爬虫.

>该库简单实用，短短几行代码即可取代以往杂乱无章的面包条代码片段，已经应用在某些大项目中。

>该库主要用途：微信开发/API对接/自动化测试/抢票脚本/网站监控/点赞插件/数据爬取

## 一. 下载

自己封装的 `Golang` 爬虫下载库，支持各种代理模式和伪装功能，你只需通过该方式获取库：

```
go get -v github.com/hunterhug/marmot/miner
```

代码结构：

```
├── miner   核心库（HTTP请求封装）
├── expert  解析库（HTML解析封装）
├── example Example示例库
    ├── lession   示例
    ├── practice  练习    
├── tool    小工具
    ├── wx   微信开发相关接口
├── proxy   Golang官方代理库
└── util    基础库，为了避免外部依赖包失效，某些核心依赖包放置于此
```

以下是几个实战例子：

1. [多线程批量抓图片](/example/practice/pictures/README.md)。
2. [模拟上传文件](/example/practice/upload/README.md)。
3. [微信开发相关：如微信登录，小程序开发](/tool/wx/README.md)。

## 二. 使用

此库可模拟上传文件，模拟表单提交，模拟各种各样的操作。

`lesson1.go`

```go
package main

import (
	"fmt"
	"github.com/hunterhug/marmot/miner"
)

func main() {
	miner.SetLogLevel(miner.DEBUG)

	// 使用默认的全局矿工，当然你可以自己建一个
	//worker, _ := miner.New(nil)
	//worker = miner.NewWorkerWithNoProxy()
	//worker = miner.NewAPI()
	//worker, _ = miner.NewWorkerWithProxy("socks5://127.0.0.1:1080")
	worker := miner.Clone()
	_, err := worker.SetUrl("https://www.gov.cn").Go()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(worker.ToString())
	}
}
```

官方部分示例已经合进本库，参见 [example](example) 文件夹。

### 第一步

矿工有四种类型：

1. `miner.NewWorker("http://user:password@104.128.121.46:808") `  // 代理矿工，默认自动化Cookie接力 格式:`协议://代理帐号(可选):代理密码(可选)@ip:port`，支持http(s)，socks5，别名函数 `New()`
2. `miner.NewWorker(nil)`   // 正常矿工，默认自动化Cookie接力，别名函数`New()`
3. `miner.NewAPI()` // API矿工，默认Cookie不接力，主要用来对接服务端 API
4. `miner.NewWorkerByClient(&http.Client{})`    // 可自定义客户端

如果你想使用一个矿工多次，使用前请先调用 `Clone()` 方法，它会创建新的矿工来进行 HTTP 请求的数据隔离，避免你并发调用多次导致的混乱。

### 第二步

模拟矿工设置头部:

1. `worker.SetUrl("https://www.gov.cn")`  // 设置HTTP请求要抓取的网址，**必须**
2. `worker.SetMethod(miner.GET)`  // 设置HTTP请求的方法:`POST/GET/PUT/POSTJSON`等
3. `worker.SetWaitTime(2)` // 设置HTTP请求超时时间
4. `worker.SetUa(miner.RandomUa())`                // 设置HTTP请求浏览器标志，本项目提供445个浏览器标志，可选设置
5. `worker.SetRefer("http://www.baidu.com")`       // 设置HTTP请求Refer头
6. `worker.SetHeaderParam("diyheader", "diy")` // 设置HTTP请求自定义头部
7. `worker.SetBData([]byte("file data"))` // HTTP请求需要上传数据
8. `worker.SetFormParam("username","jinhan")` // HTTP请求需要提交表单
9. `worker.SetCookie("xx=dddd")` // HTTP请求设置cookie，某些网站需要登录后F12复制cookie

### 第三步

矿工启动方式有：

1. `body, err := worker.Go()` // 如果设置SetMethod()，会调用下方对应的方法，否则使用Get()
2. `body, err := worker.Get()` // 默认
3. `body, err := worker.Post()` // POST表单请求，数据在SetFormParam()
4. `body, err := worker.PostJSON()` // 提交JSON请求，数据在SetBData()
5. `body, err := worker.PostXML()` // 提交XML请求，数据在SetBData()
6. `body, err := worker.PostFILE()` // 提交文件上传请求，文件二进制数据通过SetBData()设置，然后设置SetFileInfo(fileName，fileFormName string) 表明文件名和表单field Name
7. `body, err := worker.Delete()` 
8. `body, err := worker.Put()`
9. `body, err := worker.PutJSON()`
10. `body, err := worker.PutXML()`
11. `body, err := worker.PutFILE()`
12. `body, err := worker.OtherGo("OPTIONS", "application/x-www-form-urlencoded")` // 其他自定义的HTTP方法，不能模拟二进制
13. `body, err := worker.OtherGoBinary("OPTIONS", "application/x-www-form-urlencoded")` // 其他自定义的HTTP方法，模拟二进制
14. `body, err := worker.GoByMethod("POST")` // 等同于 SetMethod() 然后 Go()

### 第四步

每次下载会返回 `[]byte`，请自行解析，调试时可以使用以下方法：

1. `fmt.Println(string(html))` // 每次抓取后会返回二进制数据，直接类型转化
2. `fmt.Println(worker.ToString())` // http响应后二进制数据也会保存在矿工对象的Raw字段中，使用ToString可取出来
3. `fmt.Println(worker.JsonToString())` // 如果获取到的是JSON数据，请采用此方法转义回来，不然字符串会乱码

注意：每次下载后，需要使用以下方法将表单等数据重置（我推荐每次使用 `Clone()` 来避免这种做法，也就是克隆一个新的矿工来进行下一次操作）：

1. `Clear()` // 清除表单和二进制数据
2. `ClearAll()` // 清除表单和二进制数据，还清除全部HTTP头部
3. `ClearCookie()` // 可清除Cookie

### 其他

勾子:

1. `SetBeforeAction(fc func(context.Context, *Worker))` 爬虫动作前可AOP注入。
2. `SetAfterAction(fc func(context.Context, *Worker))` 爬虫动作完成后。

# License

```
Copyright [2016-2021] [github.com/hunterhug]

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
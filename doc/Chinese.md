# 项目代号：土拨鼠

![土拨](tubo.png)

>万维网网络机器人,又称蜘蛛,爬虫,原理主要是通过构造符合HTTP协议的网络数据包,向指定主机请求资源,获取返回的数据.万维网有大量的公开信息,人力采集数据费时费力,故激发了爬虫的产业化.
批量获取公开网络数据并不违反,但由于无差别性,无节制,十分暴力的手段会导致对方服务的不稳定,因此,大部分资源提供商对数据包进行了某些过滤,在此背景下,小批量数据获取成为了难题.
综合各种需求，如各种API对接,自动化测试等原理均一样，故开发了此爬虫库.

>土拨鼠项目是一个人类友好姿势的代码库,开发采用面向对象的方式,易于理解.通过对Golang原生HTTP库的封装,帮用户处理了一些琐碎逻辑(如收集信息,检测参数),并加入了一些容错机制(如加锁,及时关闭流),保证了爬虫高并发的安全.此库提供了大量优美的API接口,复用率高,十分方便地支持Cookie接力,爬虫代理设置,以及一般的HTTP请求设置如头部设置,超时,暂停设置,数据设置等,支持全部的HTTP方法如POST/PUT/GET/DELETE等,内置爬虫池和浏览器UA池,易于开发多UA多Cookie分布式爬虫.

>该库简单实用,短短几行代码即可取代以往杂乱无章的面包条代码片段,已经应用在某些大项目中:如`大型亚马逊分布式爬虫(美国/日本/德国/英国)`,经受住两千代理IP超长时间高并发的考验,单台机器每天获取上百万数据.

>该库主要用途： 微信开发/API对接/自动化测试/抢票脚本/网站监控/点赞插件/数据爬取

## 一.下载

自己封装的爬虫库,类似于Python的requests,你只需通过该方式获取库

```
go get -v github.com/hunterhug/marmot/miner
```

## 二.使用

请参考最新的英文版本, 中文版本不再更新...

官方部分示例已经合进本库，参见`example`文件夹

```go
package main

// 示例
import (
	"fmt"
	"github.com/hunterhug/marmot/miner"
)

func main() {
	// 1.新建一个矿工
	worker, _ := miner.New(nil)
	// 2.设置网址
	worker.SetUrl("http://www.cjhug.me").SetUa(worker.RandomUa()).SetMethod(worker.GET)
	// 3.抓取网址
	html, err := worker.Go()
	if err != nil {
		fmt.Println(err.Error())
	}
	// 4.打印内容,等同于fmt.Println(worker.ToString())
	fmt.Println(string(html))
}
```

详细具体步骤如下：

```go
package main

import (
	// 第一步：引入库
	"github.com/hunterhug/marmot/miner"
	"github.com/hunterhug/parrot/util"
)

func init() {
	// 第二步：可选设置全局
	miner.SetLogLevel(miner.DEBUG) // 设置全局矿工日志,可不设置,设置debug可打印出http请求轨迹
	miner.SetGlobalTimeout(3)      // 矿工超时时间,可不设置

}

func main() {

	log := miner.Log() // 矿工为你提供的日志工具,可不用

	// 第三步： 必须新建一个矿工对象
	//worker, err := miner.NewWorker("http://smart:smart2016@104.128.121.46:808") // 代理IP格式: 协议://代理帐号(可选):代理密码(可选)@ip:port
	//worker, err := miner.NewWorker(nil)  // 正常矿工 默认带Cookie
	//worker, err := miner.NewAPI() // API矿工 默认不带Cookie
	worker, err := miner.New(nil) // NewWorker同名函数
	if err != nil {
		panic(err)
	}

	// 第四步：设置抓取方式和网站,可链式结构设置,只有SetUrl是必须的
	// SetUrl:Url必须设置
	// SetMethod:HTTP方法可以是POST或GET,可不设置,默认GET,传错值默认为GET
	// SetWaitTime:暂停时间,可不设置,默认不暂停
	worker.SetUrl("http://cjhug.me/fuck.html").SetMethod(miner.GET).SetWaitTime(2)
	worker.SetUa(miner.RandomUa())                //设置随机浏览器标志
	worker.SetRefer("http://cjhug.me/fuck.html")  // 设置Refer头
	worker.SetHeaderParm("diyheader", "lenggirl") // 自定义头部
	//worker.SetBData([]byte("file data")) // 如果你要提交JSON数据/上传文件
	//worker.SetFormParm("username","jinhan") // 提交表单
	//worker.SetFormParm("password","123")

	// 第五步：开始爬
	//worker.Get()             // 默认GET
	//worker.Post()            // POST表单请求,数据在SetFormParm()
	//worker.PostJSON()        // 提交JSON请求,数据在SetBData()
	//worker.PostXML()         // 提交XML请求,数据在SetBData()
	//worker.PostFILE()        // 提交文件上传请求,数据在SetBData()
	body, err := worker.Go() // 如果设置SetMethod(),采用,否则Get()
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Infof("%s", string(body)) // 打印获取的数据
	}

	log.Debugf("%#v", worker.GetCookies()) // 不设置全局log为debug是不会出现这个东西的

	worker.Clear() // 爬取完毕后可以清除POST的表单数据/文件数据/JSON数据
	//worker.ClearAll() // 爬取完毕后可以清除设置的Http头部和POST的表单数据/文件数据/JSON数据

	// 矿工池子
	miner.Pool.Set("myfirstworker", worker)
	if poolworkerider, ok := miner.Pool.Get("myfirstworker"); ok {
		go func() {
			poolworkerider.SetUrl("http://cjhug.me/fuck.html")
			data, _ := poolworkerider.Get()
			log.Info(string(data))
		}()
		util.Sleep(10)
	}

}
```

使用特别简单,先`New`一个`Worker`, 即土拨鼠矿工, 然后`SetUrl`,适当加头部,最后`worker.Go()`即可。

### 第一步

矿工有四种类型:

1. `worker, err := miner.NewWorker("http://smart:smart2016@104.128.121.46:808") ` // 代理IP矿工 默认自动化Cookie接力 格式:`协议://代理帐号(可选):代理密码(可选)@ip:port`, 支持http(s),socks5, 别名函数`New()`
2. `worker, err := miner.NewWorker(nil)`  // 正常矿工 默认自动化Cookie接力 别名函数`New()`
3. `worker, err := miner.NewAPI()` // API矿工 默认Cookie不接力
4. `worker, err := miner.NewWorkerByClient(&http.Client{})` //  可自定义客户端

### 第二步

模拟矿工设置头部:

1. `worker.SetUrl("http://www.cjhug.me")`  // 设置Http请求要抓取的网址,必须
2. `worker.SetMethod(miner.GET)`  // 设置Http请求的方法:`POST/GET/PUT/POSTJSON`等
3. `worker.SetWaitTime(2)` // 设置Http请求超时时间
4. `worker.SetUa(miner.RandomUa())`                // 设置Http请求浏览器标志,本项目提供445个浏览器标志，可选设置
5. `worker.SetRefer("http://www.baidu.com")`       // 设置Http请求Refer头
6. `worker.SetHeaderParm("diyheader", "lenggirl")` // 设置Http请求自定义头部
7. `worker.SetBData([]byte("file data"))` // Http请求需要上传数据
8. `worker.SetFormParm("username","jinhan")` // Http请求需要提交表单
9. `worker.SetCookie("xx=dddd")` // Http请求设置cookie, 某些网站需要登录后F12复制cookie

### 第三步

矿工启动方式有：

1. `body, err := worker.Go()` // 如果设置SetMethod(),采用下方对应的方法,否则Get()
2. `body, err := worker.Get()` // 默认
3. `body, err := worker.Post()` // POST表单请求,数据在SetFormParm()
4. `body, err := worker.PostJSON()` // 提交JSON请求,数据在SetBData()
5. `body, err := worker.PostXML()` // 提交XML请求,数据在SetBData()
6. `body, err := worker.PostFILE()` // 提交文件上传请求,数据在SetBData()
7. `body, err := worker.Delete()` 
8. `body, err := worker.Put()`
9. `body, err := worker.PutJSON()` 
10. `body, err := worker.PutXML()`
11. `body, err := worker.PutFILE()`
12. `body, err := worker.OtherGo("OPTIONS", "application/x-www-form-urlencoded")` // 其他自定义的HTTP方法
13. `body, err := worker.GoByMethod("POST")` // you can override SetMethod() By this, equal SetMethod() then Go()

### 第四步

爬取到的数据：

1. `fmt.Println(string(html))` // 每次抓取后会返回二进制数据
2. `fmt.Println(worker.ToString())` // http响应后二进制数据也会保存在矿工对象的Raw字段中,使用ToString可取出来
3. `fmt.Println(worker.JsonToString())` // 如果获取到的是JSON数据,请采用此方法转义回来,不然会乱码

注意：每次抓取网站后,下次请求你可以覆盖原先的头部,但是没覆盖的头部还是上次的,所以清除头部或请求数据,请使用`Clear()`(只清除Post数据)或者`ClearAll()`(还清除http头部)

[API参考](doc/api.md),更多自行查看源代码

## 三.项目应用

该爬虫库已经在多个项目中使用

1. [亚马逊大型分布式爬虫](https://github.com/hunterhug/AmazonBigSpider) // 查看[图文](/doc/amazon.md)
2. [煎蛋分布式文章爬虫](https://github.com/hunterhug/jiandan)
3. [知乎全能API小工具](https://github.com/hunterhug/GoZhihu)
4. [网易无损音乐下载](/example/practice/music/README.md)
5. [天猫淘宝有关的采集小程序](https://github.com/hunterhug/GoTaoBao)
6. ...

版本日志信息见[日志](/doc/log.md), 爬虫环境安装请参考:[Docker快速版本](https://github.com/hunterhug/GoSpider-docker)

示例编译二进制方法：

Linux二进制

```bash
cd main

# 64位
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -x -o zhihu_linux_amd64 main.go

# 32位
GOOS=linux GOARCH=386 go build -ldflags "-s -w" -x -o zhihu_linux_386 main.go
```

Windows二进制

```bash
# 64位
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -x -o zhihu_windows_amd64.exe main.go

# 32位
GOOS=windows GOARCH=386 go build -ldflags "-s -w" -x -o zhihu_windows_386.exe main.go
```

如果你觉得项目帮助到你,欢迎请我喝杯咖啡

微信
![微信](https://raw.githubusercontent.com/hunterhug/hunterhug.github.io/master/static/jpg/wei.png)

支付宝
![支付宝](https://raw.githubusercontent.com/hunterhug/hunterhug.github.io/master/static/jpg/ali.png)


问题咨询请发邮件:gdccmcm14@live.com.
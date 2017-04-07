# Golang Spider

Golang爬虫封装包，组件化开发，支持Cookie持久，用户代理，多浏览器模拟等，封装了redis和mysql,可敏捷开发。
爬虫需谨慎，有风险

项目代号：土拨鼠（tubo）

![土拨](tubo.png)

已完成特大亚马逊分布式爬虫：https://github.com/hunterhug/AmazonBigSpider

其他大示例见[http://www.github.com/hunterhug/GoSpiderExample](http://www.github.com/hunterhug/GoSpiderExample)

## 一.下载

自己封装的爬虫库，类似于Python的requests，又不像，你只需通过该方式获取库

```
go get -v github.com/hunterhug/GoSpider
```

或者新建 你的GOPATH路径/src/github.com/hunterhug

```
cd src/github.com/hunterhug
git clone https://github.com/hunterhug/GoSpider
```

默认所有第三方库已经保存在vendor，如果使用包冲突了，请把vendor下的包移到GOPATH下，谨记！！！GOPATH文件夹下的包为不适宜放在vendor下，请手动移动

文件目录（组件化开发）

```
    ---example   爬虫示例，新爬虫已经转移到新仓库
    ---query     内容解析库，只封装了两个方法
    ---spider    爬虫核心库
    ---store     存储库
        ---myredis
        ---mysql
    ---util      杂项工具
    ---vendor    第三方依赖包
    ---GOPATH    不宜放在vendor的包
```

## 二.使用
HelloWorld Simple一般情况，看代码注释
```go
package main

import (
	// 第一步：引入库
	boss "github.com/hunterhug/GoSpider/spider"
)

func main() {

	// 第二步：可选设置全局
	// 设置全局爬虫日志，可不设置，设置debug可打印出http请求轨迹
	boss.SetLogLevel("debug")

	// 爬虫超时时间，可不设置，默认超长时间
	boss.SetGlobalTimeout(3)

	// 爬虫为你提供的日志工具，可不用
	log := boss.Log()

	// 第三步： 必须新建一个爬虫对象，nil表示不使用代理IP，可选代理
	// 也可以使用boss.New(nil),同名函数
	// 代理使用：spiders, err := boss.NewSpider("http://smart:smart2016@104.128.121.46:808")
	spiders, err := boss.NewSpider(nil)

	if err != nil {
		panic(err)
	}

	// 第四步：设置抓取方式和网站，可链式结构设置
	// SetUrl:Url必须设置
	// SetMethod:HTTP方法可以是POST或GET，可不设置，默认GET，传错值默认为GET
	// SetWaitTime:暂停时间，可不设置，默认不暂停
	// SetHeaderParm：自定义头部，可不设置，默认UA是火狐
	spiders.SetUrl("http://www.zhihu.com").SetMethod("get").SetWaitTime(2)
	spiders.SetHeaderParm("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:46.0) Gecko/20100101 Firefox/46.0")
	// 可用spiders.SetHeaderParm("User-Agent", boss.RandomUa())设置随机浏览器标志
	spiders.SetHeaderParm("Host", "www.zhihu.com")

	// 第五步：开始爬
	// 可使用spiders.Get()或spiders.Post()
	body, err := spiders.Go()
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Infof("%s", string(body)) // 打印获取的数据
	}

	// 不设置全局log为debug是不会出现这个东西的
	log.Debugf("%#v", spiders)
}
```

表单提交（如知乎登陆）
```go
package main

import (
	// 第一步：引入库
	"flag"
	"fmt"
	boss "github.com/hunterhug/GoSpider/spider"
	"github.com/hunterhug/GoSpider/util"
	"strings"
)

// go run loginzhihu.go -email=122233 -password=44646
var (
	password = flag.String("password", "", "zhihu password you must set")
	email    = flag.String("email", "", "zhihu email you must set")
)

func init() {
	flag.Parse()
	if *password == "" || *email == "" {
		pw, e := util.ReadfromFile(util.CurDir() + "/data/password.txt")
		if e != nil {
			fmt.Println("命令行为空，且文件也出错" + e.Error())
			panic(0)
		}
		zhihupw := strings.Split(string(pw), ",")
		if len(zhihupw) != 2 {
			fmt.Println("文件中必须有email,password")
			panic(0)
		}
		*password = strings.TrimSpace(zhihupw[1])
		*email = strings.TrimSpace(zhihupw[0])
	}
	fmt.Printf("账号:%s,密码:%s\n", *email, *password)
}
func main() {
	// 第一步：可选设置全局
	boss.SetLogLevel("debug") // 设置全局爬虫日志，可不设置，设置debug可打印出http请求轨迹
	boss.SetGlobalTimeout(3)  // 爬虫超时时间，可不设置，默认超长时间
	log := boss.Log()         // 爬虫为你提供的日志工具，可不用

	// 第二步： 新建一个爬虫对象，nil表示不使用代理IP，可选代理
	spiders, err := boss.NewSpider(nil) // 也可以使用boss.New(nil),同名函数

	if err != nil {
		panic(err)
	}

	// 第三步：设置头部
	spiders.SetMethod(boss.POST).SetUrl("https://www.zhihu.com/login/email").SetUa(boss.RandomUa())
	spiders.SetHost("www.zhihu.com").SetRefer("https://www.zhihu.com")
	spiders.SetFormParm("email", *email).SetFormParm("password", *password)

	// 第四步：开始爬
	body, err := spiders.Post()
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Info(spiders.ToString()) // 打印获取的数据，也可以string(body)
		// 待处理,json数据带有\\u
		context,_ := util.JsonEncode2(body)
		// 登陆成功
		log.Info(string(context))
	}
}
```

二进制提交(如文件上传,JSON上传），解析文件见[helloworld](example/helloworld/README.md)

## 三.例子
### 1.入门

a. 基础简单示例<br/>
b. 中级知乎登录

见[helloworld](example/helloworld/README.md)

### 2.示例项目
a. 任意图片下载,见[图片下载](example/taobao/README.md)

## 四.备注
1. 爬虫对象默认保存网站cookie（爬虫很少不会用的没有cookie的）
2. 不设置Header User-Agent标志默认会使用火狐浏览器标志（这样是为了覆盖Go官方的头部）
3. 项目管理

此库采用[Glide](https://github.com/Masterminds/glide)方式管理第三方库（使用者可以忽略,中国防火长城让我爪机，最终完全弃用，长城太猛）

```
$ glide init                              # 创建工作区
$ open glide.yaml                         # 编辑glide.yaml文件
$ glide get github.com/Masterminds/cookoo # get下库然后会自动写入glide.yaml
$ glide install                           # 安装,没有glide.lock,会先运行glide up

# work, work, work
$ go build                                # 试试可不可以跑
$ glide up                                # 更新库，创建glide.lock
```

# Log
20170405
1. 简单就是美
2. 核心功能完成
3. 示例转移到另外的仓库

20170404
1. 增加存储库redis和mysql
2. 优化

20170330
1. 抽离SpiderConfig出来，重构解耦，链式配置可直接传SpiderConfig，默认逐链覆盖
2. POST之后获取JSON数据可能被编码成\u9a8c，增加JsonToString爬虫对象方法获取数据
3. 例子重构

20170329

1. 增加默认爬虫
2. 单只爬虫加锁

20170319

1. 新增多User-Agent全局变量，默认支持几百个浏览器标志
2. 增加随机User-Agent函数，可以随机提取一个标志
3. 新增多浏览器池Pool，可以模拟若干个浏览器

20170318

1. 新增glide管理第三方库
2. 更新若干函数
3. 修改README.md等
4. 增加基础实例
5. 增加任意图片下载示例（淘宝有特殊处理）
6. 知乎登陆

# LICENSE

欢迎加功能(PR/issues)，请遵循Apache License协议(即可随意使用但每个文件下都需加此申明）

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
# Golang Spider

## 一.下载

自己封装的爬虫库，类似于Python的requests，又不像，你只需通过该方式获取库

```
go get -u -v github.com/hunterhug/GoSpider
```

或者新建 你的GOPATH路径/src/github.com/hunterhug

```
cd src/github.com/hunterhug
git clone https://github.com/hunterhug/GoSpider
```

此库采用[Glide](https://github.com/Masterminds/glide)方式管理第三方库（贡献者可以查看）

```
$ glide init                              # 创建工作区
$ open glide.yaml                         # 编辑glide.yaml文件
$ glide get github.com/Masterminds/cookoo # get下库然后会自动写入glide.yaml
$ glide install                           # 安装,没有glide.lock,会先运行glide up

# work, work, work
$ go build                                # 试试可不可以跑
$ glide up                                # 更新库，创建glide.lock
```

默认所有第三方库已经保存在vendor

## 二.使用
HelloWorld Simple，看代码注释
```go
package main

import (
	"fmt"
	// 第一步：引入库
	boss "github.com/hunterhug/GoSpider/spider"
)

func main() {

	// 第二步：可选设置全局
	//boss.SetLogLevel("debug")   // 设置全局爬虫日志，可不设置，设置debug可打印出http请求轨迹
	boss.SetGlobalTimeout(3) // 爬虫超时时间，可不设置，默认超长时间
	log := boss.Log()        // 爬虫为你提供的日志工具，可不用

	// 第三步： 新建一个爬虫对象，nil表示不使用代理IP，可选代理
	spiders, err := boss.NewSpider(nil) // 也可以使用boss.New(nil),同名函数
	//spiders, err := boss.NewSpider("http://smart:smart2016@104.128.121.46:808")

	if err != nil {
		panic(err)
	}

	// 第四步：设置抓取方式和网站
	//spiders.Method = "get"  // HTTP方法可以是POST或GET，可不设置，默认GET
	//spiders.Wait = 2        // 暂停时间，可不设置，默认不暂停
	spiders.Url = "http://www.lenggirl.com" // 抓取的网址，必须

	// 第五步：自定义头部，可不设置，默认UA是火狐
	spiders.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:46.0) Gecko/20100101 Firefox/46.0")
	//spiders.Header.Set("Host","www.lenggirl.com")

	// 第六步：开始爬
	body, err := spiders.Go() // 可使用spiders.Get()或spiders.Post()
	if err != nil {
		log.Error(err.Error())
	} else {
		fmt.Printf("%s", string(body)) // 打印获取的数据
	}

	// 不设置全局log为debug是不会出现这个东西的
	log.Debugf("%#v", spiders)
}
```

其他如表单提交（如知乎登陆）,二进制提交(如文件上传,JSON上传），解析文件见[helloworld](example/helloworld/README.md)

## 三.例子
### 1.入门

a. 基础简单示例<br/>
b. 中级知乎登录

见[helloworld](example/helloworld/README.md)

### 2.示例项目
a. 任意图片下载,见[图片下载](example/taobao/README.md)

## 四.备注
1. 爬虫对象默认保存网站cookie
2. 不设置Header User-Agent标志默认会使用火狐浏览器标志

# Log
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

# 待做
1. JavaBean链式配置爬虫待写
2. 抽离CatchConfig出来，重构解耦，链式配置可直接传CatchConfig，默认逐链覆盖
2. Redis 分布式抓取示例

![](girl.jpg)

# 联系

1. QQ：569929309 一只尼玛
2. 公众号：lenggirlcom(搬砖的陈大师)
3. 个人网站:www.lenggirl.com

# LICENSE

欢迎加功能(PR/issues)，请遵循Apache License协议(即可随意使用但每个文件下都需加此申明）

```
Copyright 2017 hunterhug/一只尼玛.
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
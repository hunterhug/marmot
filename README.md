# 项目代号：土拨鼠（tubo）

[前言](doc/pre.md)

![土拨](tubo.png)

用途： 微信开发/API对接/自动化测试/抢票脚本/网站监控/点赞插件/数据爬取

## 一.下载

自己封装的爬虫库,类似于Python的requests,你只需通过该方式获取库

```
go get -v github.com/hunterhug/GoSpider
```

或者新建 你的GOPATH路径`/src/github.com/hunterhug`

```
cd src/github.com/hunterhug
git clone https://github.com/hunterhug/GoSpider
```

默认所有第三方库已经保存在vendor,如果使用包冲突了,请把vendor下的包移到GOPATH下,谨记！！


以下godep可选,vendor中已经带全第三方库

```
godep restore
```

文件目录（组件化开发）

```
    ---example   爬虫示例,新爬虫已经转移到新仓库
    ---query     内容解析库,只封装了两个方法
    ---spider    爬虫核心库
    ---store     存储库
        ---myredis 
        ---mysql
        ---myetcd
        ---mydb  关系型数据库Orm(使用xorm)
        ---myhbase
        ---mycassandra
    ---util      杂项工具
        --- image 图片切割库
    ---vendor    第三方依赖包
    ---GOPATH    不宜放在vendor的包,请手动移动到你的GOPATH路径下
```

## 二.使用

最简单示例,更多移动到[官方示例](https://github.com/hunterhug/GoSpiderExample)

```go
package main

// 示例
import (
	"fmt"
	"github.com/hunterhug/GoSpider/spider"
)

func main() {
	// 1.新建爬虫
	sp, _ := spider.New(nil)
	// 2.设置网址
	sp.SetUrl("http://www.lenggirl.com").SetUa(spider.RandomUa()).SetMethod(spider.PUT) // 我的网站不允许PUT请改为GET
	// 3.抓取网址
	html, err := sp.Go()
	if err != nil {
		fmt.Println(err.Error())
	}
	// 4.打印内容,等同于fmt.Println(sp.ToString())
	fmt.Println(string(html))
}
```

使用特别简单,先`New`一只`Spider`,然后`SetUrl`,适当加头部,最后`spiders.Go()`即可。

### 第一步

爬虫有三种类型:

1. `spiders, err := boss.NewSpider("http://smart:smart2016@104.128.121.46:808") ` // 代理IP爬虫 默认自动化Cookie接力 格式:`协议://代理帐号(可选):代理密码(可选)@ip:port` 别名函数`New()`
2. `spiders, err := boss.NewSpider(nil)`  // 正常爬虫 默认自动化Cookie接力 别名函数`New()`
3. `spiders, err := boss.NewAPI()` // API爬虫 默认Cookie不接力

### 第二步

模拟爬虫设置头部:

1. `spiders.SetUrl("http://www.lenggirl.com")`  // 设置Http请求要抓取的网址,必须
2. `spiders.SetMethod(boss.GET)`  // 设置Http请求的方法:`POST/GET/PUT/POSTJSON`等
3. `spiders.SetWaitTime(2)` // 设置Http请求超时时间
4. `spiders.SetUa(boss.RandomUa())`                // 设置Http请求浏览器标志,本项目提供445个浏览器标志，可选设置
5. `spiders.SetRefer("http://www.baidu.com")`       // 设置Http请求Refer头
6. `spiders.SetHeaderParm("diyheader", "lenggirl")` // 设置Http请求自定义头部
7. `spiders.SetBData([]byte("file data"))` // Http请求需要上传数据
8. `spiders.SetFormParm("username","jinhan")` // Http请求需要提交表单
9. `spiders.SetCookie("xx=dddd")` // Http请求设置cookie, 某些网站需要登录后F12复制cookie

### 第三步

爬虫启动方式有：
1. `body, err := spiders.Go()` // 如果设置SetMethod(),采用下方对应的方法,否则Get()
2. `body, err := spiders.Get()` // 默认
3. `body, err := spiders.Post()` // POST表单请求,数据在SetFormParm()
4. `body, err := spiders.PostJSON()` // 提交JSON请求,数据在SetBData()
5. `body, err := spiders.PostXML()` // 提交XML请求,数据在SetBData()
6. `body, err := spiders.PostFILE()` // 提交文件上传请求,数据在SetBData()
7. `body, err := spiders.Delete()` 
8. `body, err := spiders.Put()`
9. `body, err := spiders.PutJSON()` 
10. `body, err := spiders.PutXML()`
11. `body, err := spiders.PutFILE()`
12. `body, err := spiders.OtherGo("OPTIONS", "application/x-www-form-urlencoded")` // 其他自定义的HTTP方法

### 第四步

爬取到的数据：

1. `fmt.Println(string(html))` // 每次抓取后会返回二进制数据
2. `fmt.Println(sp.ToString())` // http响应后二进制数据也会保存在爬虫对象的Raw字段中,使用ToString可取出来
3. `fmt.Println(sp.JsonToString())` // 如果获取到的是JSON数据,请采用此方法转义回来,不然会乱码

注意：每次抓取网站后,下次请求你可以覆盖原先的头部,但是没覆盖的头部还是上次的,所以清除头部或请求数据,请使用`Clear()`(只清除Post数据)或者`ClearAll()`(还清除http头部)

[API参考](doc/api.md),更多自行查看源代码

## 三.项目应用

该爬虫库已经在多个项目中使用

1. [煎蛋分布式文章爬虫](https://github.com/hunterhug/jiandan)
2. [知乎全能API小工具](https://github.com/hunterhug/zhihuxx)
3. [亚马逊分布式爬虫](https://github.com/hunterhug/AmazonBigSpider)
4. ...

如果你觉得项目帮助到你,欢迎请我喝杯咖啡

微信
![微信](https://raw.githubusercontent.com/hunterhug/hunterhug.github.io/master/static/jpg/wei.png)

支付宝
![支付宝](https://raw.githubusercontent.com/hunterhug/hunterhug.github.io/master/static/jpg/ali.png)

版本日志信息见[日志](doc/log.md)

爬虫环境安装请参考:[环境配置](http://www.lenggirl.com/tool/gospider-env.html)

问题咨询请发邮件.

# LICENSE

欢迎加功能(PR/issues),请遵循Apache License协议(即可随意使用但每个文件下都需加此申明）

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

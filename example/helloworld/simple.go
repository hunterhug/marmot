/*
Copyright 2017 by GoSpider author.
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

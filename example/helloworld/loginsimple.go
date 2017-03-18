/*
Copyright 2017 hunterhug/一只尼玛.
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
	"flag"
	"fmt"
	boss "github.com/hunterhug/GoSpider/spider"
	"github.com/hunterhug/go_tool/util"
	"strconv"
	"strings"
)

var (
	password = flag.String("password", "", "zhihu password you must set")
	email    = flag.String("email", "", "zhihu email you must set")
)

func init() {
	flag.Parse()
}
func main() {
	// 第二步：可选设置全局
	boss.SetLogLevel("debug") // 设置全局爬虫日志，可不设置，设置debug可打印出http请求轨迹
	boss.SetGlobalTimeout(3)  // 爬虫超时时间，可不设置，默认超长时间
	log := boss.Log()         // 爬虫为你提供的日志工具，可不用

	// 第三步： 新建一个爬虫对象，nil表示不使用代理IP，可选代理
	spiders, err := boss.NewSpider(nil) // 也可以使用boss.New(nil),同名函数

	if err != nil {
		panic(err)
	}

	// 第四步：设置抓取方式和网站
	spiders.Method = "post"
	//spiders.Wait = 2        // 暂停时间，可不设置，默认不暂停
	spiders.Url = "https://www.zhihu.com/login/email" // 抓取的网址，必须

	// 第五步：自定义头部，可不设置，默认UA是火狐
	spiders.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:46.0) Gecko/20100101 Firefox/46.0")
	spiders.Header.Set("Referer", "https://www.zhihu.com")
	spiders.Header.Set("Host", "www.zhihu.com")
	// 相当于以下方法
	//spiders.NewHeader("Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:46.0) Gecko/20100101 Firefox/46.0","www.zhihu.com","https://www.zhihu.com")
	spiders.Data.Set("email", *email)
	spiders.Data.Set("password", *password)

	// 第六步：开始爬
	body, err := spiders.Post()
	if err != nil {
		log.Error(err.Error())
	} else {
		//log.Infof("%s", body) // 打印获取的数据

		// 待处理,json数据带有\\u
		sUnicodev := strings.Split(string(body), "\\u")
		var context string
		for _, v := range sUnicodev {
			if len(v) < 1 {
				continue
			}
			temp, err := strconv.ParseInt(v, 16, 32)
			if err != nil {
				context += fmt.Sprintf("%s", v)
			} else {
				context += fmt.Sprintf("%c", temp)
			}
		}
		// 登陆成功
		log.Info(context)
		util.SaveToFile(util.CurDir()+"/data/back.json", body)
	}
	spiders.Url = "https://www.zhihu.com"
	index, e := spiders.Get()
	if e != nil {
		log.Error(e.Error())
	} else {
		//log.Info(string(index))
		util.SaveToFile(util.CurDir()+"/data/index.html", index)
	}
}

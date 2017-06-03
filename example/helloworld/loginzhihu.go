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
	"flag"
	"fmt"
	boss "github.com/hunterhug/GoSpider/spider"
	"github.com/hunterhug/GoSpider/util"
	"strings"
)

// 知乎登录有验证码！！
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
	spiders.SetUrl("https://www.zhihu.com/login/email").SetRefer("https://www.zhihu.com/").SetUa(boss.RandomUa())
	spiders.SetFormParm("email", *email).SetFormParm("password", *password)

	// 第四步：开始爬
	body, err := spiders.Post()
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Info(spiders.ToString()) // 打印获取的数据，也可以string(body)
		// 待处理,json数据带有\\u
		context, _ := util.JsonBack(body)
		// 登陆成功
		log.Info(string(context))
	}
}

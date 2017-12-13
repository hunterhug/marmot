/*
	版权所有，侵权必究
	署名-非商业性使用-禁止演绎 4.0 国际
	警告： 以下的代码版权归属hunterhug，请不要传播或修改代码
	你可以在教育用途下使用该代码，但是禁止公司或个人用于商业用途(在未授权情况下不得用于盈利)
	商业授权请联系邮箱：gdccmcm14@live.com QQ:459527502
	All right reserved
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
	For more information on commercial licensing please contact hunterhug.
	Ask for commercial licensing please contact Mail:gdccmcm14@live.com Or QQ:459527502
	2017.7 by hunterhug
*/

// loginsimple.go 知乎表单POST登陆，运行`go run loginzhihu.go -email=122233 -password=44646`,或者data/password.txt填入密码，
// 并`go run loginzhihu.go`

package main

import (
	// 第一步：引入库
	"flag"
	"fmt"
	"strings"

	boss "github.com/hunterhug/marmot/miner"
	"github.com/hunterhug/parrot/util"
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
	spiders, err := boss.NewWorker(nil) // 也可以使用boss.New(nil),同名函数

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

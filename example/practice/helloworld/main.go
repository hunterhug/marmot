/*
	All right reserved https://github.com/hunterhug/marmot at 2016-2020
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
*/
package main

import (
	// 第一步：引入库
	"github.com/hunterhug/marmot/miner"
	"github.com/hunterhug/marmot/util"
)

func init() {
	// 第二步：可选设置全局
	miner.SetLogLevel(miner.DEBUG) // 设置全局矿工日志,可不设置,设置debug可打印出http请求轨迹
	miner.SetGlobalTimeout(3)      // 矿工超时时间,可不设置

}

func main() {

	log := miner.Log() // 矿工为你提供的日志工具,可不用

	// 第三步： 必须新建一个矿工对象
	// 代理IP格式: 协议://代理帐号(可选):代理密码(可选)@ip:port
	//worker, err := miner.NewWorker("http://smart:smart2016@104.128.121.46:808")
	//worker, err := miner.NewWorker(nil)  // 正常矿工 默认带Cookie
	//worker := miner.NewAPI() // API矿工 默认不带Cookie
	worker, err := miner.New(nil) // NewWorker同名函数
	if err != nil {
		panic(err)
	}

	// 第四步：设置抓取方式和网站,可链式结构设置,只有SetUrl是必须的
	// SetUrl:Url必须设置
	// SetMethod:HTTP方法可以是POST或GET,可不设置,默认GET,传错值默认为GET
	// SetWaitTime:暂停时间,可不设置,默认不暂停
	worker.SetUrl("https://github.com/hunterhug").SetMethod(miner.GET).SetWaitTime(2)
	worker.SetUa(miner.RandomUa())                  //设置随机浏览器标志
	worker.SetRefer("https://github.com/hunterhug") // 设置Refer头
	worker.SetHeaderParm("diyheader", "diy")        // 自定义头部
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

	// 爬取完毕后可以清除POST的表单数据/文件数据/JSON数据
	worker.Clear()

	// 爬取完毕后可以清除设置的Http头部和POST的表单数据/文件数据/JSON数据
	//worker.ClearAll()

	// 矿工池子
	miner.Pool.Set("worker1", worker)
	if pools, ok := miner.Pool.Get("worker1"); ok {
		go func() {
			pools.SetUrl("https://github.com/hunterhug")
			data, _ := pools.Get()
			log.Info(string(data))
		}()
		util.Sleep(10)
	}

}

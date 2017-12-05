/*
	More detail Example
*/
package main

import (
	// 1:import package alias boss
	boss "github.com/hunterhug/GoSpider/spider"
	"github.com/hunterhug/GoTool/util"
)

func init() {
	// 2:Optional global setting
	boss.SetLogLevel(boss.DEBUG) // optional, set log to debug
	boss.SetGlobalTimeout(3)     // optional, http request timeout time

}

func main() {

	log := boss.Log() // optional, spider log you can choose to use

	// 3: Must new a spider object, three ways
	//sp, err := boss.NewSpider("http://smart:smart2016@104.128.121.46:808") // proxy spider, format: protocol://user(optional):password(optional)@ip:port
	//sp, err := boss.NewSpider(nil)  // normal spider, default keep Cookie
	//sp, err := boss.NewAPI() // API spider, not keep Cookie
	sp, err := boss.New(nil) // NewSpider alias
	if err != nil {
		panic(err)
	}

	// 4: Set the request Method/URL and some others, can chain set, only SetUrl is required.
	// SetUrl: required, the Url
	// SetMethod: optional, HTTP method: POST/GET/..., default GET
	// SetWaitTime: optional, HTTP request wait/pause time
	sp.SetUrl("https://www.whitehouse.gov").SetMethod(boss.GET).SetWaitTime(2)
	sp.SetUa(boss.RandomUa())                 // optional, browser user agent: IE/Firefox...
	sp.SetRefer("https://www.whitehouse.gov") // optional, url refer
	sp.SetHeaderParm("diyheader", "lenggirl") // optional, some other diy http header
	//sp.SetBData([]byte("file data"))    // optional, if you want post JSON data or upload file
	//sp.SetFormParm("username","jinhan") // optional: if you want post form
	//sp.SetFormParm("password","123")

	// 5: Start Run
	//sp.Get()             // default GET
	//sp.Post()            // POST form request data, data can fill by SetFormParm()
	//sp.PostJSON()        // POST JSON dara, use SetBData()
	//sp.PostXML()         // POST XML, use SetBData()
	//sp.PostFILE()        // POST to Upload File, data in SetBData() too
	//sp.OtherGo("OPTIONS", "application/x-www-form-urlencoded") // Other http method, Such as OPTIONS etcd
	body, err := sp.Go() // if you use SetMethod(), otherwise equal to Get()
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Infof("%s", string(body)) // Print return data
	}

	log.Debugf("%#v", sp.GetCookies) // if you not set log as debug, it will not appear

	// You must Clear it! If you want to POST Data by SetFormParm()/SetBData() again
	// After get the return data by post data, you can clear the data you fill
	sp.Clear()
	//sp.ClearAll() // you can also want to clear all, include http header you set

	// Spider pool for concurrent, every Spider Object is serial as the browser. if you want collateral execution, use this.
	boss.Pool.Set("myfirstspider", sp)
	if poolspider, ok := boss.Pool.Get("myfirstspider"); ok {
		go func() {
			poolspider.SetUrl("https://www.whitehouse.gov")
			data, _ := poolspider.Get()
			log.Info(string(data))
		}()
		util.Sleep(10)
	}
}

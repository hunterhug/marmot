package main

/*
	More detail Example
*/
import (
	// 1:import package
	"github.com/hunterhug/marmot/miner"
	"github.com/hunterhug/marmot/util"
)

func init() {
	// 2:Optional global setting
	miner.SetLogLevel(miner.DEBUG) // optional, set log to debug
	miner.SetGlobalTimeout(3)      // optional, http request timeout time

}

func main() {
	log := miner.Log() // optional, miner log you can choose to use

	// 3: Must new a Worker object, three ways
	//worker, err := miner.NewWorker("http://xx:xxxx@103.118.111.26:808") // proxy format: protocol://user(optional):password(optional)@ip:port
	//worker, err := miner.NewWorker(nil)  // normal worker, default keep Cookie
	//worker := miner.NewAPI() // API worker, not keep Cookie
	worker, err := miner.New(nil) // NewWorker alias
	if err != nil {
		panic(err)
	}

	// 4: Set the request Method/URL and some others, can chain set, only SetUrl is required.
	// SetUrl: required, the Url
	// SetMethod: optional, HTTP method: POST/GET/..., default GET
	// SetWaitTime: optional, HTTP request wait/pause time
	worker.SetUrl("https://www.bing.com").SetMethod(miner.GET).SetWaitTime(2)
	worker.SetUa(miner.RandomUa())             // optional, browser user agent: IE/Firefox...
	worker.SetRefer("https://www.bing.com")    // optional, url refer
	worker.SetHeaderParam("diy_header", "diy") // optional, some other diy http header
	//worker.SetBData([]byte("file data"))    // optional, if you want post JSON data or upload file
	//worker.SetFormParam("username","root") // optional: if you want post form
	//worker.SetFormParam("password","123")

	// 5: Start Run
	//worker.Get()             // default GET
	//worker.Post()            // POST form request data, data can fill by SetFormParam()
	//worker.PostJSON()        // POST JSON dara, use SetBData()
	//worker.PostXML()         // POST XML, use SetBData()
	//worker.PostFILE()        // POST to Upload File, data in SetBData() too
	//worker.OtherGo("OPTIONS", "application/x-www-form-urlencoded") // Other http method, Such as OPTIONS etcd
	body, err := worker.Go() // if you use SetMethod(), otherwise equal to Get()
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Infof("%s", string(body)) // Print return data
	}

	log.Debugf("%#v", worker.GetCookies) // if you not set log as debug, it will not appear

	// You must Clear it! If you want to POST Data by SetFormParam()/SetBData() again
	// After get the return data by post data, you can clear the data you fill
	// I suggest use Clone()
	worker.Clear()
	//worker.ClearAll() // you can also want to clear all, include http header you set

	// Worker pool for concurrent, every Worker Object is serial as the browser. if you want collateral execution, use this.
	miner.Pool.Set("myfirstworker", worker)
	if w, ok := miner.Pool.Get("myfirstworker"); ok {
		go func() {
			data, _ := w.Clone().SetUrl("https://www.bing.com").Get()
			log.Info(string(data))
		}()
		util.Sleep(10)
	}
}

// 
// 	Copyright 2017 by marmot author: gdccmcm14@live.com.
// 	Licensed under the Apache License, Version 2.0 (the "License");
// 	you may not use this file except in compliance with the License.
// 	You may obtain a copy of the License at
// 		http://www.apache.org/licenses/LICENSE-2.0
// 	Unless required by applicable law or agreed to in writing, software
// 	distributed under the License is distributed on an "AS IS" BASIS,
// 	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// 	See the License for the specific language governing permissions and
// 	limitations under the License
//

/*
	More detail Example
*/

package main

import (
	// 1:import package
	"github.com/hunterhug/marmot/miner"
	"github.com/hunterhug/parrot/util"
)

func init() {
	// 2:Optional global setting
	miner.SetLogLevel(miner.DEBUG) // optional, set log to debug
	miner.SetGlobalTimeout(3)      // optional, http request timeout time

}

func main() {

	log := miner.Log() // optional, miner log you can choose to use

	// 3: Must new a Worker object, three ways
	//worker, err := miner.NewWorker("http://smart:smart2016@104.128.121.46:808") // proxy format: protocol://user(optional):password(optional)@ip:port
	//worker, err := miner.NewWorker(nil)  // normal worker, default keep Cookie
	//worker, err := miner.NewAPI() // API worker, not keep Cookie
	worker, err := miner.New(nil) // NewWorker alias
	if err != nil {
		panic(err)
	}

	// 4: Set the request Method/URL and some others, can chain set, only SetUrl is required.
	// SetUrl: required, the Url
	// SetMethod: optional, HTTP method: POST/GET/..., default GET
	// SetWaitTime: optional, HTTP request wait/pause time
	worker.SetUrl("https://hunterhug.github.io").SetMethod(miner.GET).SetWaitTime(2)
	worker.SetUa(miner.RandomUa())                // optional, browser user agent: IE/Firefox...
	worker.SetRefer("https://hunterhug.github.io") // optional, url refer
	worker.SetHeaderParm("diyheader", "diy")      // optional, some other diy http header
	//worker.SetBData([]byte("file data"))    // optional, if you want post JSON data or upload file
	//worker.SetFormParm("username","jinhan") // optional: if you want post form
	//worker.SetFormParm("password","123")

	// 5: Start Run
	//worker.Get()             // default GET
	//worker.Post()            // POST form request data, data can fill by SetFormParm()
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

	// You must Clear it! If you want to POST Data by SetFormParm()/SetBData() again
	// After get the return data by post data, you can clear the data you fill
	worker.Clear()
	//worker.ClearAll() // you can also want to clear all, include http header you set

	// Worker pool for concurrent, every Worker Object is serial as the browser. if you want collateral execution, use this.
	miner.Pool.Set("myfirstworker", worker)
	if w, ok := miner.Pool.Get("myfirstworker"); ok {
		go func() {
			data, _ := w.SetUrl("https://hunterhug.github.io/fuck.html").Get()
			log.Info(string(data))
		}()
		util.Sleep(10)
	}
}

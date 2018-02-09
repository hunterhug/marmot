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

package main

import (
	"fmt"
	"github.com/hunterhug/parrot/util"
	"github.com/hunterhug/marmot/tool"
)

// Num of miner, We can run it at the same time to crawl data fast
var MinerNum = 5

// You can update this decide whether to proxy
var ProxyAddress interface{}

func main() {
	// You can Proxy!
	// ProxyAddress = "socks5://127.0.0.1:1080"

	fmt.Println(`Welcome: Input "url" and picture keep "dir"`)
	fmt.Println("---------------------------------------------")
	url := util.Input(`URL(Like: "http://publicdomainarchive.com")`, "http://publicdomainarchive.com")
	dir := util.Input(`DIR(Default: "./picture")`, "./picture")
	fmt.Printf("You will keep %s picture in dir %s\n", url, dir)
	fmt.Println("---------------------------------------------")

	// Start Catch
	err := tool.DownloadHTMLPictures(url, dir, MinerNum, ProxyAddress)
	if err != nil {
		fmt.Println("Error:" + err.Error())
	}
}

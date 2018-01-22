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
	Proxy  Worker!
	You first should own a remote machine, Then in your local tap:
		`ssh -ND 1080 ubuntu@remoteIp`
	It will gengerate socks5 proxy client in your local, which port is 1080
*/

package main

import (
	"fmt"
	"os"

	"github.com/hunterhug/marmot/expert"
	"github.com/hunterhug/marmot/miner"
)

func init() {
	miner.SetLogLevel(miner.DEBUG)
}

func main() {
	// You can use a lot of proxy ip such "https/http/socks5"
	proxy_ip := "socks5://127.0.0.1:1080"

	url := "https://www.google.com"

	worker, err := miner.New(proxy_ip)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	body, err := worker.SetUa(miner.RandomUa()).SetUrl(url).SetMethod(miner.GET).Go()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(parse(body))
	}
}

// Parse HTML page
func parse(data []byte) string {
	doc, err := expert.QueryBytes(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return doc.Find("title").Text()
}

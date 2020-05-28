/*
	All right reserved https://github.com/hunterhug/marmot at 2016-2020
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
*/
package main

import (
	"fmt"
	"github.com/hunterhug/marmot/tool"
	"github.com/hunterhug/marmot/util"
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
	url := util.Input(`URL(Like: "https://tu.enterdesk.com")`, "https://tu.enterdesk.com")
	dir := util.Input(`DIR(Default: "./picture")`, "./picture")
	fmt.Printf("You will keep %s picture in dir %s\n", url, dir)
	fmt.Println("---------------------------------------------")

	// Start Catch
	err := tool.DownloadHTMLPictures(url, dir, MinerNum, ProxyAddress)
	if err != nil {
		fmt.Println("Error:" + err.Error())
	}
}

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
*
*/

package main

import (
	"fmt"
	"github.com/hunterhug/marmot/util"
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

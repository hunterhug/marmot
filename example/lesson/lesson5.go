/*
	All right reserved https://github.com/hunterhug/marmot at 2016-2020
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
*/
package main

/*
	Proxy  Worker!
	You first should own a remote machine, Then in your local tap:
		`ssh -ND 1080 ubuntu@remoteIp`
	It will gengerate socks5 proxy client in your local, which port is 1080
*/
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
	proxyIp := "socks5://127.0.0.1:1080"

	url := "https://www.google.com"

	worker, err := miner.New(proxyIp)
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

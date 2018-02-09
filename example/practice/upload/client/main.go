// Date: 18-2-9

package main

import (
	"fmt"
	"github.com/hunterhug/marmot/miner"
)

func postFile(filename string, targetUrl string) {
	worker, _ := miner.New(nil)
	result, err := worker.SetUrl(targetUrl).SetBData([]byte("dddd")).SetFileInfo(filename+".xxxx", "uploadfile").PostFILE()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(result))
	}
}

// sample usage
func main() {
	target_url := "http://127.0.0.1:1789/upload"
	filename := "./doc.go"
	postFile(filename, target_url)
}

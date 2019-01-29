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
	"context"
	"fmt"
	"strings"

	"github.com/hunterhug/marmot/expert"
	"github.com/hunterhug/marmot/miner"
)

func main() {
	// We can debug, to see whether SetBeforeAction make sense
	miner.SetLogLevel(miner.DEBUG)

	// The url we want
	url := "https://hunterhug.github.io"

	// IAM we can NewAPI
	worker := miner.NewAPI()

	// We can aop by context
	// ctx, cancle := context.WithCancel(context.Background())
	// ctx := context.TODO()
	// worker.SetContext(ctx)

	// Before we make some change, And every GET Or POST it will action
	worker.SetBeforeAction(func(ctx context.Context, this *miner.Worker) {
		fmt.Println("Before Action, I will add a HTTP header")
		this.SetHeaderParm("Marmot", "v2")
		this.SetHeaderParm("DUDUDUU", "DUDU")
		// select {
		// case <-ctx.Done():
		// 	fmt.Println(ctx.Err()) // block in here util cancle()
		// 	os.Exit(1)
		// }
	})

	// we cancle it after 5 secord
	// go func() {
	// 	util.Sleep(5)
	// 	cancle()
	// }()

	worker.SetAfterAction(func(ctx context.Context, this *miner.Worker) {
		fmt.Println("After Action, I just print this sentence")
	})

	// Let's Go
	body, err := worker.SetUrl(url).GoByMethod(miner.GET)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// Parse We want
		fmt.Printf("Output:\n %s\n", Myparse(body))
	}

	// for {
	//  in here we loop util cancle() success
	// }
}

// Parse HTML page
func Myparse(data []byte) string {
	doc, err := expert.QueryBytes(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(doc.Find("#hero-caption").Text())
	// return doc.Find("title").Text()
}

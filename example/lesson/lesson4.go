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
	"context"
	"fmt"
	"github.com/hunterhug/marmot/util"
	"strings"

	"github.com/hunterhug/marmot/expert"
	"github.com/hunterhug/marmot/miner"
)

func main() {
	// We can debug, to see whether SetBeforeAction make sense
	miner.SetLogLevel(miner.DEBUG)

	// The url we want
	url := "https://github.com/hunterhug"

	// IAM we can NewAPI
	worker := miner.NewAPI()

	// We can aop by context
	ctx, cancel := context.WithCancel(context.Background())
	//ctx := context.TODO()
	worker.SetContext(ctx)

	// we cancel it after 5 secord
	go func() {
		fmt.Println("I stop and sleep 5")
		util.Sleep(5)
		fmt.Println("I wake up after sleep 5")
		cancel()
	}()

	// Before we make some change, And every GET Or POST it will action
	worker.SetBeforeAction(func(ctx context.Context, this *miner.Worker) {
		fmt.Println("Before Action, I will add a HTTP header, then sleep wait cancel")
		this.SetHeaderParm("Marmot", "v2")
		this.SetHeaderParm("DUDUDUU", "DUDU")
		select {
		case <-ctx.Done(): // block in here util cancel()
			//fmt.Println(ctx.Err())
			fmt.Println("after sleep, i do action.")
		}
	})

	worker.SetAfterAction(func(ctx context.Context, this *miner.Worker) {
		fmt.Println("After Action, I just print this sentence")
	})

	// Let's Go
	body, err := worker.SetUrl(url).GoByMethod(miner.GET)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// Parse We want
		fmt.Printf("Output:\n %s\n", MyParse(body))
	}

}

// Parse HTML page
func MyParse(data []byte) string {
	doc, err := expert.QueryBytes(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(doc.Find("title").Text())
}

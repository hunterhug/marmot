/*
	All right reserved https://github.com/hunterhug/marmot at 2016-2020
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
*/
package miner

import (
	"context"
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	worker, _ := New(nil)
	fmt.Printf("%#v\n", worker)

	if worker.BeforeAction == nil {
		fmt.Println("good")
	}

	worker.BeforeAction = func(c context.Context, worker *Worker) {
		worker.SetHeaderParm("Marmot", "v2")
	}
}

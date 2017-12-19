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

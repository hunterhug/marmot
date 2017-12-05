package spider

import (
	"fmt"
	"testing"
)

func TestSpider1(t *testing.T) {
	sp, _ := New(nil)
	fmt.Printf("%#v\n", sp)

	if sp.BeforeAction == nil {
		fmt.Println("good")
	}

	sp.BeforeAction = func(sp *Spider) {
		sp.SetHeaderParm("GoSpider", "v2")
	}
}

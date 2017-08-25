/*
   Created by jinhan on 17-8-25.
   Tip:
   Update:
*/
package util

import (
	"fmt"
	"testing"
)

func TestWalkDir(t *testing.T) {
	f, e := WalkDir("../util", "go")
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Printf("%#v", f)
	}
}

func TestListDirOnlyName(t *testing.T) {
	f, e := ListDirOnlyName("../util", "go")
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Printf("%#v", f)
	}
}

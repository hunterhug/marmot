package slack

import (
	"fmt"
	"testing"
)

func TestSentMessage(t *testing.T) {
	hook := "https://hooks.slack.com/services/Tss"
	m := "😊 Successfully"
	a, err := SentMessage(hook, m)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(a))
	}
}

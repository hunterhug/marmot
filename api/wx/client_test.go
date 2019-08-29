package wx

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestNewGZClient(t *testing.T) {
	var callback = func(w http.ResponseWriter, r *http.Request) {
		callbackMessage, err := GZClientTemplateCallBack(r)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("%#v", callbackMessage)
	}

	http.HandleFunc("/", callback)
	go http.ListenAndServe("127.0.0.0:8000", nil)

	c, err := NewGZClient("", "")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	message := TemplateMessage{
		OpenId:     "56789",
		TemplateId: "1234",
		Url:        "http://www.baidu.com",
		Data: map[string]TemplateMessageData{
			"keyword1": {
				"我爱你", "#173177",
			},
			"keyword2": {
				"我爱你2", "#173177",
			},
		},
	}
	err = c.TemplateMessageSend(message)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	time.Sleep(50 * time.Second)
}

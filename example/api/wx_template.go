package main

import (
	"fmt"
	. "github.com/hunterhug/marmot/api/wx"
	"github.com/hunterhug/marmot/miner"
	"net/http"
	"time"
)

func main() {
	// 初始化公众号客户端
	miner.SetLogLevel(miner.DEBUG)
	c, err := NewGZClient("a", "b")
	if err != nil {
		fmt.Println(err.Error())
	}

	// 准备模板回调
	var callback = func(w http.ResponseWriter, r *http.Request) {
		callbackMessage, err := GZClientTemplateCallBack(r)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("%#v", callbackMessage)
	}

	http.HandleFunc("/", callback)
	go http.ListenAndServe("127.0.0.0:8000", nil)

	// 发送模板消息
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

	t := time.NewTicker(5 * time.Second)
	for {
		<-t.C
	}
}

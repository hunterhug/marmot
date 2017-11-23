package slack

import (
	"github.com/hunterhug/GoSpider/spider"
	"fmt"
	"time"
)

var Slack *spider.Spider

func init() {
	Slack = spider.NewAPI()
	//spider.SetLogLevel("debug")
}

func SentMessage(hook string, message string) ([]byte, error) {
	s := `{"text":"PgToEs: %s | %s"}`
	times := time.Now().UTC().Format("2006-01-02 15:04:05")
	s = fmt.Sprintf(s, times, message)
	Slack.SetUrl(hook)
	Slack.SetBData([]byte(s))
	fmt.Println(hook, s)
	return Slack.PostJSON()
}

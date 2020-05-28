package wx

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hunterhug/marmot/miner"
)

var (
	wxStateDeveloper = "developer"
	wxStateTrial     = "trial"
	wxStateFormal    = "formal"
)

type Message struct {
	ToUser           string                 `json:"to_user"`
	TemplateId       string                 `json:"template_id"`
	Page             string                 `json:"page"`
	MiniProgramState string                 `json:"miniprogram_state"`
	Lang             string                 `json:"lang"`
	Data             map[string]interface{} `json:"data"`
}

type ErrorRsp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (e *ErrorRsp) Error() string {
	return fmt.Sprintf("%d_%v", e.ErrCode, e.ErrMsg)
}

// https://developers.weixin.qq.com/miniprogram/dev/api/open-api/subscribe-message/wx.requestSubscribeMessage.html
func SendMessage(token string, openId string, templateId, page string, data map[string]string, state string) error {
	if token == "" || openId == "" || templateId == "" || state == "" {
		return errors.New("empty")
	}

	if state != wxStateDeveloper && state != wxStateFormal && state != wxStateTrial {
		return errors.New("state wrong")
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s", token)
	m := new(Message)
	m.ToUser = ""
	m.TemplateId = templateId
	m.Page = page
	m.Lang = "zh_CN"
	mm := map[string]interface{}{}
	for k, v := range data {
		mm[k] = map[string]string{"value": v}
	}
	m.Data = mm
	m.MiniProgramState = state
	raw, err := json.Marshal(m)
	if err != nil {
		return err
	}

	worker := miner.NewAPI()
	body, err := worker.SetUrl(url).SetBData(raw).PostJSON()
	if err != nil {
		return err
	}

	miner.Logger.Infof("wx send message result:%v", string(body))

	if worker.UrlStatuscode != 200 {
		return errors.New(fmt.Sprintf("wx send message http status:%d", worker.UrlStatuscode))
	}

	e := ErrorRsp{}
	err = json.Unmarshal(body, &e)
	if err != nil {
		return err
	}

	if e.ErrCode != 0 {
		return errors.New(e.ErrMsg)
	}
	return nil
}

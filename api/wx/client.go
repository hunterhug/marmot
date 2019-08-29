/*

微信公众号SDK

*/
package wx

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/hunterhug/marmot/miner"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var (
	GZToken                  = "https://api.weixin.qq.com/cgi-bin/token"
	GZTemplateApiSetIndustry = "https://api.weixin.qq.com/cgi-bin/template/api_set_industry"
	GZTemplateApiMessageSend = "https://api.weixin.qq.com/cgi-bin/message/template/send"
)

type GZClient struct {
	AppId      string
	Secret     string
	CreateTime int64
	lock       sync.Mutex
	lock2      sync.Mutex
	TokenRsp
}

type TokenRsp struct {
	AccessToken string `json:"access_token"`
	ErrorRsp
}

type ErrorRsp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (e ErrorRsp) Error() string {
	return fmt.Sprintf("%d-%s", e.ErrCode, e.ErrMsg)
}

// 初始化公众号客户端
func NewGZClient(AppId, Secret string) (*GZClient, error) {
	c := new(GZClient)
	c.AppId = AppId
	c.Secret = Secret

	err := c.GetOrRefreshToken()
	if err != nil {
		return c, err
	}
	return c, nil
}

// 获取刷新令牌
func (c *GZClient) GetOrRefreshToken() error {
	// 十分钟才刷新一次Token
	if c.CreateTime+10*60 > time.Now().Unix() {
		return nil
	}
	c.lock.Lock()
	defer c.lock.Unlock()
	raw, err := miner.NewAPI().SetUrl(fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s", GZToken, c.AppId, c.Secret)).Get()
	if err != nil {
		return err
	}

	token := new(TokenRsp)
	err = json.Unmarshal(raw, token)
	if err != nil {
		return err
	}

	if token.ErrCode != 0 {
		return token.ErrorRsp
	}

	c.CreateTime = time.Now().Unix()
	c.TokenRsp = *token
	return nil
}

// 模板消息：https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1433751277
// 模板消息仅用于公众号向用户发送重要的服务通知，只能用于符合其要求的服务场景中，如信用卡刷卡通知，商品购买成功通知等。不支持广告等营销类消息以及其它所有可能对用户造成骚扰的消息。
// 用户必须关注公众号

// https://api.weixin.qq.com/cgi-bin/template/api_set_industry?access_token=ACCESS_TOKEN
//
// POST
//	{
//		"industry_id1":"1",
//		"industry_id2":"4"
//	}
//
// 此等接口在后台获取即可，以下仅为示例
func (c *GZClient) TemplateSetIndustry(codes []string) error {
	c.lock2.Lock()
	defer c.lock2.Unlock()
	c.GetOrRefreshToken()

	inputMap := map[string]string{}
	for k, code := range codes {
		inputMap[fmt.Sprintf("industry_id%d", k)] = code
	}
	input, err := json.Marshal(inputMap)
	if err != nil {
		return err
	}

	cApi := miner.NewAPI()
	raw, err := cApi.SetUrl(fmt.Sprintf("%s?access_token=%s", GZTemplateApiSetIndustry, c.AccessToken)).SetBData(input).PostJSON()
	if err != nil {
		return err
	}

	if cApi.UrlStatuscode != 200 {
		return errors.New(fmt.Sprintf("http status:%d", cApi.UrlStatuscode))
	}

	e := ErrorRsp{}
	err = json.Unmarshal(raw, &e)
	if err != nil {
		return err
	}

	if e.ErrCode != 0 {
		return e
	}
	return nil
}

// 模板消息
// url和miniprogram都是非必填字段，若都不传则模板无跳转；
// 若都传，会优先跳转至小程序。开发者可根据实际需要选择其中一种跳转方式即可。当用户的微信客户端版本不支持跳小程序时，将会跳转至url。
type TemplateMessage struct {
	OpenId      string                         `json:"touser"`                // 接收者openid
	TemplateId  string                         `json:"template_id"`           // 模板ID
	Url         string                         `json:"url,omitempty"`         // 模板跳转链接（海外帐号没有跳转能力）
	MiniProgram *TemplateMessageMiniProgram    `json:"miniprogram,omitempty"` // 跳转的小程序
	Data        map[string]TemplateMessageData `json:"data"`                  // 模板数据
}

type TemplateMessageMiniProgram struct {
	AppId    string `json:"appid"`    // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）
	PagePath string `json:"pagepath"` // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar），要求该小程序已发布，暂不支持小游戏
}

type TemplateMessageData struct {
	Value string `json:"value"`           // 值
	Color string `json:"color,omitempty"` // 模板内容字体颜色，不填默认为黑色
}

// 模板消息发送
func (c *GZClient) TemplateMessageSend(message TemplateMessage) error {
	c.lock2.Lock()
	defer c.lock2.Unlock()
	c.GetOrRefreshToken()

	input, err := json.Marshal(message)
	if err != nil {
		return err
	}

	cApi := miner.NewAPI()
	raw, err := cApi.SetUrl(fmt.Sprintf("%s?access_token=%s", GZTemplateApiMessageSend, c.AccessToken)).SetBData(input).PostJSON()
	if err != nil {
		return err
	}

	if cApi.UrlStatuscode != 200 {
		return errors.New(fmt.Sprintf("http status:%d", cApi.UrlStatuscode))
	}

	e := ErrorRsp{}
	err = json.Unmarshal(raw, &e)
	if err != nil {
		return err
	}

	if e.ErrCode != 0 {
		return e
	}
	return nil
}

// 模板消息回调
// 在模版消息发送任务完成后，微信服务器会将是否送达成功作为通知，发送到开发者中心中填写的服务器配置地址中。
const GZClientTemplateSentSuccess = "success"

type GZClientTemplateCallBackRequestBody struct {
	ToUserName   string `xml:"ToUserName,CDATA"`   // 公众号微信号
	FromUserName string `xml:"FromUserName,CDATA"` // 接收模板消息的用户的openid
	CreateTime   int64  `xml:"CreateTime"`         // 创建时间
	MsgType      string `xml:"MsgType,CDATA"`      // 消息类型是事件
	Event        string `xml:"Event,CDATA"`        // 事件为模板消息发送结束
	MsgID        int    `xml:"MsgID"`              // 消息id
	Status       string `xml:"Status,CDATA"`       // 发送状态为成功 （用户设置拒绝接收公众号消息:failed:user block）
}

// 各种服务器需要传入http.Request
func GZClientTemplateCallBack(request *http.Request) (*GZClientTemplateCallBackRequestBody, error) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	result := new(GZClientTemplateCallBackRequestBody)
	err = xml.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}

	if result.Status != GZClientTemplateSentSuccess {
		return result, errors.New(result.Status)
	}

	return result, nil
}

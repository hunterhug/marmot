package wx

import (
	"encoding/json"
	"fmt"
	"github.com/hunterhug/marmot/miner"
)

type UserInfo struct {
	NickName string `json:"nickname"`
	OpenId   string `json:"openid"`
	Img      string `json:"headimgurl"`
	UnionId  string `json:"unionid"`
	Sex      int64  `json:"sex"`
	City     string `json:"city"`
	Province string `json:"province"`
	Country  string `json:"country"`
}

type webAccessToken struct {
	AccessToken string `json:"access_token"`
	OpenId      string `json:"openid"`
}

func Login(appId, appSecret, code string) (info *UserInfo, err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		appId, appSecret, code)

	api := miner.NewAPI()
	data, err := api.SetUrl(url).Get()
	if err != nil {
		return nil, err
	}

	wErr := new(ErrorRsp)
	err = json.Unmarshal(data, wErr)
	if err != nil {
		return
	}

	if wErr.ErrCode != 0 {
		return nil, wErr
	}

	accessToken := ""
	openId := ""
	wToken := new(webAccessToken)
	err = json.Unmarshal(data, wToken)
	if err != nil {
		return
	}

	miner.Logger.Infof("wx login token:%#v", wToken)

	accessToken = wToken.AccessToken
	openId = wToken.OpenId

	url = fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN", accessToken, openId)
	data, err = api.SetUrl(url).Get()
	if err != nil {
		return
	}

	miner.Logger.Infof("%#v", string(data))

	err = json.Unmarshal(data, wErr)
	if err != nil {
		return
	}

	if wErr.ErrCode != 0 {
		return nil, wErr
	}

	wInfo := new(UserInfo)
	err = json.Unmarshal(data, wInfo)
	if err != nil {
		return
	}

	return wInfo, nil
}

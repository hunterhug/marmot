package wx

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hunterhug/marmot/miner"
)

type Token struct {
	AccessToken string `json:"access_token"`
}

// GlobalToken https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html
func GlobalToken(appId, appSecret string) (token string, err error) {
	if appId == "" || appSecret == "" {
		return "", errors.New("empty")
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appId, appSecret)
	api := miner.NewAPI()
	raw, err := api.Clone().SetUrl(url).Get()
	if err != nil {
		return "", err
	}

	miner.Logger.Infof("wx token:%v", string(raw))
	t := new(Token)
	err = json.Unmarshal(raw, t)
	if err != nil {
		return "", err
	}

	if t.AccessToken == "" {
		return "", errors.New("empty")
	}

	return t.AccessToken, nil
}

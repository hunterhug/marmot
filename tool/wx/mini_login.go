package wx

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hunterhug/marmot/miner"
	"strings"
)

type miniAccessToken struct {
	AccessToken string `json:"session_key"`
	OpenId      string `json:"openid"`
}

type MiniUserInfo struct {
	NickName  string                 `json:"nickName"`
	OpenId    string                 `json:"openId"`
	Img       string                 `json:"avatarUrl"`
	UnionId   string                 `json:"unionId"`
	Sex       int64                  `json:"sex"`
	City      string                 `json:"city"`
	Province  string                 `json:"province"`
	Country   string                 `json:"country"`
	Watermark map[string]interface{} `json:"watermark"`
}

func MiniLogin(appId, appSecret, code, encryptedData, iv string) (info *MiniUserInfo, err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
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

	wToken := new(miniAccessToken)
	err = json.Unmarshal(data, wToken)
	if err != nil {
		return
	}

	miner.Logger.Infof("wx login token:%#v", wToken)

	accessToken := wToken.AccessToken
	//openId := wToken.OpenId

	raw, err := DecryptWXOpenData(accessToken, encryptedData, iv)
	if err != nil {
		return
	}

	miner.Logger.Infof("%#v", string(data))

	temp := strings.Split(string(raw), "}")
	tempL := len(temp)
	if tempL < 2 {
		err = errors.New("not a json")
		return
	}
	temp2 := strings.Join(temp[:tempL-1], "}")
	raw = []byte(temp2 + "}")
	uInfo := new(MiniUserInfo)
	err = json.Unmarshal(raw, uInfo)
	if err != nil {
		return
	}

	if uInfo.Watermark == nil {
		err = errors.New(fmt.Sprintf("wx login err:%s", "Watermark wrong nil"))
		return
	}

	temp3, ok := uInfo.Watermark["appid"]
	if !ok {
		err = errors.New(fmt.Sprintf("wx login err:%s", "Watermark wrong app id not found"))
		return
	}

	temp4 := fmt.Sprintf("%v", temp3)
	if temp4 != appId {
		err = errors.New(fmt.Sprintf("wx login err:%s", "Watermark wrong app id not match"))
		return
	}

	return uInfo, nil

}

func DecryptWXOpenData(sessionKey, encryptData, iv string) ([]byte, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(encryptData)
	if err != nil {
		return nil, err
	}
	sessionKeyBytes, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}
	dataBytes, err := AesDecrypt(decodeBytes, sessionKeyBytes, ivBytes)
	if err != nil {
		return nil, err
	}
	return dataBytes, nil

}

func AesDecrypt(crypted, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	for i, ch := range origData {
		if ch == '\x0e' {
			origData[i] = ' '
		}
	}
	return origData, nil
}

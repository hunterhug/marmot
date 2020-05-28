package wx

import (
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	appId := ""
	appSecret := ""
	code := "xxx"
	info, err := Login(appId, appSecret, code)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(info)
}

func TestMiniLogin(t *testing.T) {
	appId := ""
	appSecret := ""
	code := "xxx"
	encryptedData := "afqaf"
	iv := "ssss"
	info, err := MiniLogin(appId, appSecret, code, encryptedData, iv)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(info)
}

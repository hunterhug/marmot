# 微信开发相关

## 微信第三方登录

适用于网页端，移动端APP的微信登录。参考[官方文档](https://developers.weixin.qq.com/doc/oplatform/Website_App/WeChat_Login/Wechat_Login.html)。

需要客户端和服务端联调。

逻辑如下：

1.客户端先调用以下接口，微信用户允许授权第三方应用后，微信将会携带 `CODE` 并且回调服务端 `http://127.0.0.1:9999`：

https://open.weixin.qq.com/connect/qrconnect?appid=wx01fdsffsds&redirect_uri=http://127.0.0.1:9999&response_type=code&scope=snsapi_login,snsapi_userinfo&state=test#wechat_redirect

2.服务端收到回调，会连续调用以下链接获取到用户信息。

https://api.weixin.qq.com/sns/oauth2/access_token?appid=wx0189ce76eadccf91&secret=00cc512fc031fcdsfsdfba01c8a41f05b4b5&code=CODE&grant_type=authorization_code

https://api.weixin.qq.com/sns/userinfo?access_token=accessToken&openid=openid&lang=zh_CN

你只需使用该 `SDK` 实现登录即可：

```go
	appId := ""
	appSecret := ""
	code := "xxx" // 客户端传给你的，客户端可以是Web前端，IOS，Android
	info, err := Login(appId, appSecret, code)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(info)
```

## 小程序开发

### 小程序微信登录

[小程序登录](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/login.html)区别于网页登录。

需要客户端和服务端联调。

逻辑如下：

1.客户端先调用 `wx.login()` 获取临时登录凭证 `code` 并且 [获取用户信息](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/user-info/wx.getUserInfo.html) 获取 `encryptedData` 和 `iv` 并回传到开发者服务器。

2.服务端使用该 `code` 调用 [`auth.code2Session`](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html) 获取解密密钥，然后解密用户信息。

你只需使用该 `SDK` 实现登录即可：

```go
	appId := ""
	appSecret := ""
	code := "xxx"  // 小程序前端传给你的
	encryptedData := "afqaf"  // 小程序前端传给你的
	iv := "ssss"  // 小程序前端传给你的
	info, err := MiniLogin(appId, appSecret, code, encryptedData, iv)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(info)
```

### 小程序发送 [消息订阅](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/subscribe-message/wx.requestSubscribeMessage.html)。

完全在服务端执行，不需要客户端参与。

1.先获取全局 [`token`](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html)：

```go
	appId := ""
	appSecret := ""
	token, err := GlobalToken(appId, appSecret)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("token is:", token)
```

2.发送[订阅消息](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html)：

```go
	token, _ := GlobalToken(appId, appSecret)
	openId := "sss"  // 接收者（用户）的 openid
	templateId := ""  // 所需下发的订阅模板id
	page := ""  // 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	data := map[string]string{"thing1": "2222", "thing7": "sss", "thing3": "dddd"}
	state := wxStateFormal // 跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版

	err = SendMessage(token, openId, templateId, page, data, state)
	if err != nil {
		fmt.Println("send err:", err.Error())
		return
	}
```

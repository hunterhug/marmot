package wx

var (
	GZToken = "https://api.weixin.qq.com/cgi-bin/token"
)

type GZClient struct {
	AppId  string `json:"appid"`
	Secret string `json:"secret"`
}

func NewGZClient(AppId, Secret string) *GZClient {
	c := new(GZClient)
	c.AppId = AppId
	c.Secret = Secret
	return c
}

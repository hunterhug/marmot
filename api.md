# 核心代码剖析

API使用请看具体示例，这里介绍两个爬虫对象,核心代码spider/spider.go里：

```go
// 新建一个爬虫，如果ipstring是一个代理IP地址，那使用代理客户端
func NewSpider(ipstring interface{}) (*Spider, error) {
	spider := new(Spider)
	spider.SpiderConfig = new(SpiderConfig)
	spider.Header = http.Header{}
	spider.Data = url.Values{}
	spider.BData = []byte{}
	if ipstring != nil {
		client, err := NewProxyClient(ipstring.(string))
		spider.Client = client
		spider.Ipstring = ipstring.(string)
		return spider, err
	} else {
		client, err := NewClient()
		spider.Client = client
		spider.Ipstring = "localhost"
		return spider, err
	}

}
```

可以传入ipstring，表示使用代理，默认开启cookie记录，cookie会一直在内存中更新，默认有头部，如果要自定义http client客户端,使用：

```go
// 通过官方Client来新建爬虫，方便您更灵活
func NewSpiderByClient(client *http.Client) *Spider {
	spider := new(Spider)
	spider.SpiderConfig = new(SpiderConfig)
	spider.Header = http.Header{}
	spider.Data = url.Values{}
	spider.BData = []byte{}
	spider.Client = client
	return spider
}
```

官方的http.Client是这么用的，看spider/client.go

```go
//cookie record
// 记录Cookie
func NewJar() *cookiejar.Jar {
	cookieJar, _ := cookiejar.New(nil)
	return cookieJar
}

var (
	//default client to ask get or post
	// 默认的官方客户端，带cookie,方便使用，没有超时时间，不带cookie的客户端不提供
	Client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			Logger.Debugf("-----------Redirect:%v------------", req.URL)
			return nil
		},
		Jar: NewJar(),
	}
)
```

该客户端重定向打印日志，支持cookie持久，你也可以设置超时时间，代理，SSH等。。。。
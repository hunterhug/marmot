package main

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/hunterhug/GoSpider/query"
	"github.com/hunterhug/GoSpider/spider"
	"github.com/hunterhug/GoTool/util"
)

// Num of spider, We can run it at the same time to crawl data fast
var SpiderNum = 5

// You can update this decide whether to proxy
var ProxyAddress interface{}

func main() {
	// You can Proxy!
	// ProxyAddress = "socks5://127.0.0.1:1080"

	fmt.Println(`
		Welcome: Input "url" and picture keep "dir"

		
		`)
	for {
		fmt.Println("---------------------------------------------")
		url := util.Input(`URL(Like: "http://publicdomainarchive.com")`, "http://publicdomainarchive.com")
		dir := util.Input(`DIR(Default: "./picture")`, "./picture")
		fmt.Printf("You will keep %s picture in dir %s\n", url, dir)
		fmt.Println("---------------------------------------------")

		// Start Catch
		err := CatchPicture(url, dir)
		if err != nil {
			fmt.Println("Error:" + err.Error())
		}
	}
}

// Come on!
func CatchPicture(picture_url string, dir string) error {
	// Check valid
	_, err := url.Parse(picture_url)
	if err != nil {
		return err
	}

	// Make dir!
	err = util.MakeDir(dir)
	if err != nil {
		return err
	}

	// New a sp to get url
	sp, _ := spider.New(ProxyAddress)

	result, err := sp.SetUrl(picture_url).SetUa(spider.RandomUa()).Get()
	if err != nil {
		return err
	}

	// Find all picture
	pictures := query.FindPicture(string(result))

	// Empty, What a pity!
	if len(pictures) == 0 {
		return errors.New("empty")
	}

	// Devide pictures into several sp
	xxx, _ := util.DevideStringList(pictures, SpiderNum)

	// Chanel to info exchange
	chs := make(chan int, len(pictures))

	// Go at the same time
	for num, imgs := range xxx {

		// Get pool spider
		sp_picture, ok := spider.Pool.Get(util.IS(num))
		if !ok {
			// No? set one!
			sp_temp, _ := spider.New(ProxyAddress)
			sp_picture = sp_temp
			sp_temp.SetUa(spider.RandomUa())
			spider.Pool.Set(util.IS(num), sp_temp)
		}

		// Go save picture!
		go func(imgs []string, sp *spider.Spider, num int) {
			for _, img := range imgs {

				// Check, May be Pass
				_, err := url.Parse(img)
				if err != nil {
					continue
				}

				// Change Name of our picture
				filename := strings.Replace(util.ValidFileName(img), "#", "_", -1)

				// Exist?
				if util.FileExist(dir + "/" + filename) {
					fmt.Println("File Exist：" + dir + "/" + filename)
					chs <- 0
				} else {

					// Not Exsit?
					imgsrc, e := sp.SetUrl(img).Get()
					if e != nil {
						fmt.Println("Download " + img + " error:" + e.Error())
						chs <- 0
						return
					}

					// Save it!
					e = util.SaveToFile(dir+"/"+filename, imgsrc)
					if e == nil {
						fmt.Printf("SP%d: Keep in %s/%s\n", num, dir, filename)
					}
					chs <- 1
				}
			}
		}(imgs, sp_picture, num)
	}

	// Every picture should return
	for i := 0; i < len(pictures); i++ {
		<-chs
	}

	return nil
}

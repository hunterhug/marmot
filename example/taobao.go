package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/hunterhug/GoSpider/spider"
	"github.com/hunterhug/GoSpider/query"
	"github.com/hunterhug/go_tool/util"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(`欢迎使用淘宝天猫图片下载小工具，taobao.txt位置在` + util.CurDir() + "/taobao.txt" + `,运行EXE即可`)
	fmt.Println("链接如：tmall.com/item.htm?id=523350171126&skuId=3120562159704,tmall")
	fmt.Println("---------------以上详情页中图片会保存在tmall目录-----------------------")
	c, e := util.ReadfromFile(util.CurDir() + "/taobao.txt")
	if e != nil {
		fmt.Println("打开taobao.txt出错" + e.Error())
	} else {
		urls := strings.Split(string(c), "\n")
		for _, url := range urls {
			url := strings.Replace(strings.TrimSpace(url), "\r", "", -1)
			downlod(url)
		}

	}
	fmt.Println("请手动关闭选框...")
	util.Sleep(100)
}

func md55(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

func downlod(urlmany string) {
	temp := strings.Split(urlmany, ",")
	url := temp[0]
	filename := util.TodayString(3)
	if len(temp) >= 2 {
		filename = temp[1]
	}
	dir := util.CurDir() + "/" + filename
	util.MakeDir(dir)
	s, e := spider.NewSpider(nil)
	if e != nil {

	} else {
		s.Url = url
		dudu := "detail.tmall.com"
		if strings.Contains(url, "item.taobao.com") {
			dudu = "item.taobao.com"
		}
		s.NewHeader("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.99 Safari/537.36", dudu, nil)
		content, err := s.Get()
		if err != nil {

		} else {
			//fmt.Println(string(content))
			docm, err := query.QueryBytes(content)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				//fmt.Println(string(content))
				docm.Find("img").Each(func(num int, node *goquery.Selection) {
					img, e := node.Attr("src")
					if e == false {
						img, e = node.Attr("data-src")
					}
					if e && img != "" {
						if strings.Contains(img, ".gif") {
							return
						}
						fmt.Println("原始文件：" + img)
						r, _ := regexp.Compile(`([\d]{1,4}x[\d]{1,4})`)
						imgdudu := r.FindStringSubmatch(img)
						sizes := "720*720"
						if len(imgdudu) == 2 {
							sizes = imgdudu[1]
						}
						temp := strings.Replace(img, sizes, "720x720", -1)
						filename := md55(temp)
						if util.FileExist(dir + "/" + filename + ".jpg") {
							fmt.Println("文件存在：" + dir + "/" + filename)
						} else {
							fmt.Println("下载:" + temp)
							s.Url = "http:" + temp
							imgsrc, e := s.Get()
							if e != nil {
								fmt.Println("下载出错" + temp + ":" + e.Error())
								return
							}
							e = util.SaveToFile(dir+"/"+filename+".jpg", imgsrc)
							if e == nil {
								fmt.Println("成功保存在" + dir + "/" + filename)
							}
							util.Sleep(2)
							fmt.Println("暂停两秒")
						}
					}
				})

			}

		}
	}

}

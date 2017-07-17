/*
Copyright 2017 by GoSpider author.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/hunterhug/GoSpider/query"
	"github.com/hunterhug/GoSpider/spider"
	"github.com/hunterhug/GoSpider/util"
	"regexp"
	"strings"
)

var txt = flag.String("config", "taobao.csv", "你需要指定taobao.csv的文件地址")

func main() {
	flag.Parse()
	spider.SetGlobalTimeout(2)
	//spider.SetLogLevel("debug")
	fmt.Println(`
	-------------------------------
	欢迎使用淘宝天猫图片下载小工具
	需指定taobao.csv所在位置
	taobao.csv按行写入 淘宝链接,文件夹
	使用方法：
	go run taobao.go -config=taobao.csv
	taobao.exe -config=taobao.csv

	联系QQ：569929309
	一只尼玛
	----------------------------------
	`)
	fmt.Println("链接如：https://item.taobao.com/item.htm?id=40066362090,taobao")
	fmt.Println("---------------以上详情页中图片会保存在taobao目录-----------------------")
	c, e := util.ReadfromFile(*txt)
	if e != nil {
		fmt.Println("打开taobao.csv出错:" + e.Error())
		flag.PrintDefaults()
	} else {
		urls := strings.Split(string(c), "\n")
		for _, url := range urls {
			url := strings.Replace(strings.TrimSpace(url), "\r", "", -1)
			if strings.HasPrefix(url, "#") {
				fmt.Println("跳过" + url)
				continue
			}
			fmt.Println("下载:" + url)
			downlod(url)
		}

	}
	fmt.Println("请手动关闭选框...")
	util.Sleep(100)
}

func downlod(urlmany string) {
	temp := strings.Split(urlmany, ",")
	url := temp[0]
	filename := util.TodayString(3)
	if len(temp) >= 2 {
		filename = temp[1]
	}
	dir := "./image/" + filename
	util.MakeDir(dir)
	s, e := spider.NewSpider(nil)
	if e != nil {

	} else {
		// url:http://a.com   https://a.com/jj
		s.Url = url
		urlhost := strings.Split(url, "//")
		if len(urlhost) != 2 {
			fmt.Println("网站错误：开头必须为http://或https://")
			return
		}
		dudu := strings.Split(urlhost[1], "/")[0]
		s.NewHeader("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.99 Safari/537.36", dudu, url)
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
						temp := img
						if strings.Contains(url, "taobao.com") || strings.Contains(url, "tmall.com") {
							r, _ := regexp.Compile(`([\d]{1,4}x[\d]{1,4})`)
							imgdudu := r.FindStringSubmatch(img)
							sizes := "720*720"
							if len(imgdudu) == 2 {
								sizes = imgdudu[1]
							}
							temp = strings.Replace(img, sizes, "720x720", -1)
						}
						filename := util.Md5(temp)
						if util.FileExist(dir + "/" + filename + ".jpg") {
							fmt.Println("文件存在：" + dir + "/" + filename)
						} else {
							fmt.Println("下载:" + temp)
							s.Url = temp
							if strings.HasPrefix(temp, "//") {
								s.Url = "http:" + temp
							}
							imgsrc, e := s.Get()
							if e != nil {
								fmt.Println("下载出错" + temp + ":" + e.Error())
								return
							}
							e = util.SaveToFile(dir+"/"+filename+".jpg", imgsrc)
							if e == nil {
								fmt.Println("成功保存在" + dir + "/" + filename)
							}
							util.Sleep(1)
							fmt.Println("暂停两秒")
						}
					}
				})

			}

		}
	}

}

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
limitations under the License
*/
//
//  说明：保存图片而已，会自动去重！！！
//
package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/hunterhug/GoSpider/query"
	"github.com/hunterhug/GoSpider/spider"
	"github.com/hunterhug/GoSpider/util"
	"strings"
)

var (
	types = map[int]string{
		1: "ooxx", // 妹子图
		2: "pic",  // 无聊图
	}

	// 要抓哪一种
	Which = 1

	url     = "http://jandan.net/" + types[Which]
	urlpage = "http://jandan.net/" + types[Which] + "/page-%d"

	// 保存在统一文件
	saveroot = false
	// 根据页数保存在很多文件夹下
	savehash = true

	// 保存的地方
	rootdir = "D:\\jiandan\\jiandansum\\" + types[Which]
	// 根据页数分图片保存，不然图片太大了,我简称它hash（之前版本不是用page分而是hash）！
	// 图片太大硬盘会爆！
	hashdir = "D:\\jiandan\\jiandanpage\\" + types[Which]
)

func init() {
	if savehash == false && saveroot == false {
		fmt.Println("这种是不行的：savehash==false && saveroot==false！必须有一个为true")
	}
	// 设置日志和超时时间
	spider.SetLogLevel("info")
	// 有些图片好大！
	spider.SetGlobalTimeout(100)
	// 图片集中地大本营
	util.MakeDir(rootdir)
}

// 单只爬虫，请耐心爬取好吗
func main() {
	// 初始化爬虫
	client, _ := spider.NewSpider(nil)
	// 随机UA
	client.SetUa(spider.RandomUa())

	// 开始抓取
	client.SetUrl(url)
	data, e := client.Go()

	// 首页都抓出错，直接结束
	if e != nil {
		spider.Log().Panic(e.Error())
	}

	// 保存在本地看看
	//util.SaveToFile(util.CurDir()+"/index.html", data)

	// 解析查看页数
	doc, _ := query.QueryBytes(data)
	temp := doc.Find(".current-comment-page").Text()
	pagenum := strings.Replace(strings.Split(temp, "]")[0], "[", "", -1)
	spider.Log().Info(pagenum)

	num, e := util.SI(pagenum)
	if e != nil {
		spider.Log().Panic(e.Error())
	}

	// 页数知道后，建文件夹！！
	for i := num; i >= 1; i-- {
		util.MakeDir(hashdir + "/" + util.IS(i))
	}

	// 循环抓取开始
	for i := num; i >= 1; i-- {
		index := fmt.Sprintf(urlpage, i)
		client.SetUrl(index)
		data, e = client.Go()
		if e != nil {
			spider.Log().Errorf("列表页 %s 抓取出错:%s", index, e.Error())
			continue
		}
		spider.Log().Infof("列表页 %s 完成!", index)
		//util.SaveToFile(rootdir+"/"+util.ValidFileName(index)+".html", data)
		doc, _ = query.QueryBytes(data)
		doc.Find(".view_img_link").Each(func(num int, node *goquery.Selection) {
			imgurl, ok := node.Attr("href")
			if !ok {
				return
			}
			//spider.Log().Infof("img:%s", imgurl)

			// 去重 处理
			temp := strings.Split(imgurl, ".")
			tempnum := len(temp)
			if tempnum <= 1 {
				return
			}

			// 后缀
			houzui := temp[tempnum-1]
			// 文件名
			filename := util.Md5(imgurl) + "." + houzui
			// 大本营文件路径
			filedir := rootdir + "/" + filename
			// 页数分图
			hashfiledir := hashdir + "/" + util.IS(i) + "/" + filename

			// 下面每次都会去扫描
			// 大本营存在？
			exist := util.FileExist(filedir)

			// hash存在？
			exist2 := util.FileExist(hashfiledir)

			// 不保存hash且大本营存在
			if !savehash && exist {
				return
			}
			// 如果要保存hash
			if savehash {
				// hash不存在，大本营存在
				if !exist2 && exist {
					if !exist2 {
						// 读出来
						temp, e := util.ReadfromFile(filedir)
						// 出错不管
						if e != nil {
							return
						}
						// 写，出错不管
						util.SaveToFile(hashfiledir, temp)
						return
					}
					// spider.Log().Infof("image file %s exist", filedir)
					return
				}
				// hash存在
				if exist2 {
					return
				}
			}

			// 补充img url
			if strings.HasPrefix(imgurl, "//") {
				imgurl = "http:" + imgurl
			}

			// 抓取开始
			client.SetUrl(imgurl).SetRefer(index)
			data, e = client.Go()
			if e != nil {
				spider.Log().Errorf("抓取图片：%s 出错:%s", imgurl, e.Error())
				return
			}

			spider.Log().Infof("抓取图片：%s 成功", imgurl)

			// 大本营保存
			if saveroot {
				e = util.SaveToFile(filedir, data)
				if e != nil {
					spider.Log().Errorf("图片保存： %s 出错:%s", filedir, e.Error())
				} else {
					spider.Log().Infof("图片保存： %s", filedir)
				}
			}

			if savehash {
				e = util.SaveToFile(hashfiledir, data)
				if e != nil {
					spider.Log().Errorf("图片保存： %s 出错:%s", hashfiledir, e.Error())
				} else {
					spider.Log().Infof("图片保存： %s", hashfiledir)
				}
			}

		})
	}
}

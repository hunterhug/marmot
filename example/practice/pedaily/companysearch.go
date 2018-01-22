// 
// 	Copyright 2017 by marmot author: gdccmcm14@live.com.
// 	Licensed under the Apache License, Version 2.0 (the "License");
// 	you may not use this file except in compliance with the License.
// 	You may obtain a copy of the License at
// 		http://www.apache.org/licenses/LICENSE-2.0
// 	Unless required by applicable law or agreed to in writing, software
// 	distributed under the License is distributed on an "AS IS" BASIS,
// 	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// 	See the License for the specific language governing permissions and
// 	limitations under the License
//

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hunterhug/marmot/expert"
	"github.com/hunterhug/marmot/miner"
	"github.com/hunterhug/parrot/util"
)

var client *miner.Worker
var dir = "./data/company/raw"
var dirdetail = "./data/company/detailraw"
var dirresult = "./data/company/result"

func main() {
	welcome()
	initx()
	mainx()
	//fmt.Printf("%#v", detail("http://zdb.pedaily.cn/company/show587/"))
}
func welcome() {
	fmt.Println(`
************************************************************

		投资界关键字查找公司信息

		1.输入关键字
		首先翻页后抓取详情页信息保存在data文件夹中

		2.查看结果
		查看data/company/result中csv文件


		/*
		go build *.go，然后点击exe运行或go run *.go
		*/
************************************************************
`)
}
func initx() {
	var e error = nil
	client, e = miner.NewWorker(nil)
	if e != nil {
		panic(e.Error())
	}

	util.MakeDir(dir)
	util.MakeDir(dirresult)
	util.MakeDir(dirdetail)
}
func mainx() {
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please input kyword: ")
		//fmt.Scanln(&keyword)
		keyword, err := inputReader.ReadString('\n')
		keyword = strings.Replace(keyword, "\n", "", -1)
		keyword = trip(strings.Replace(keyword, "\r", "", -1))
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		e := util.MakeDir(dir + "/" + keyword)
		if e != nil {
			fmt.Println(e.Error())
		}
		//keyword := "投资管理股份有限公司"
		result, numt, e := featchcompany(keyword)
		if e != nil {
			fmt.Println(e.Error())
			continue
		}
		if len(result) == 0 {
			fmt.Println("empty")
			continue
		} else {
			num := int(math.Ceil(float64(numt) / 20.0))
			for i := 2; i <= num; i++ {
				temp, _, e := featchcompany(keyword + "/" + util.IS(i))
				if e != nil {
					fmt.Println(e.Error())
				} else {
					result = append(result, temp...)
				}
			}
			fmt.Printf("total comepany:%d\n", numt)
			txt := []string{}
			txt = append(txt, "公司名,英语名称,简称,详情URL,投资URL,资本类型,机构性质,注册地点,机构总部,投资阶段,成立时间,官网,简介,联系方式")
			for _, k := range result {
				detailr := detail(k["href"])

				stemp := k["title"] + "," + detailr["english"] + "," + k["abbr"] + "," + k["href"] + "," + k["hreft"] + "," + detailr["zitype"] + "," + detailr["jx"]
				stemp = stemp + "," + detailr["rloca"] + "," + detailr["tloca"] + "," + detailr["tstage"] + "," + detailr["date"] + "," + detailr["website"] + "," + detailr["desc"] + "," + detailr["contact"]
				fmt.Println(stemp)
				txt = append(txt, stemp)
			}

			e := util.SaveToFile(dirresult+"/"+keyword+".csv", []byte(strings.Join(txt, "\n")))
			if e != nil {
				fmt.Println(e.Error())
			}
		}

		fmt.Println("------------")
	}
}

func detail(url string) map[string]string {
	returnmap := map[string]string{
		"english": "", //英文名称
		"zitype":  "", //资本类型
		"jx":      "", //机构性质
		"rloca":   "", //注册地点
		"tloca":   "", //机构总部
		"tstage":  "", //投资阶段
		"date":    "", //成立时间
		"website": "", //官方网站
		"desc":    "", //简介
		"contact": "", //联系方式
	}
	hashmd := util.Md5(url)
	keep := dirdetail + "/" + hashmd + ".html"
	body := []byte("")
	var e error = nil
	if util.FileExist(keep) {
		body, e = util.ReadfromFile(keep)
	} else {
		client.Url = url
		body, e = client.Get()
	}
	if e != nil {
		return returnmap
	}
	util.SaveToFile(keep, body)
	doc, e := expert.QueryBytes(body)
	if e != nil {
		return returnmap
	}
	returnmap["english"] = trip(doc.Find(".info h2").Text())
	returnmap["contact"] = strings.Replace(trip(doc.Find("#contact").Text()), "\n", "<br/>", -1)
	returnmap["desc"] = strings.Replace(trip(doc.Find("#desc").Text()), "\n", "<br/>", -1)
	returnmap["website"] = trip(doc.Find("li.link a").Text())
	info := ""
	doc.Find(".info ul li").Each(func(i int, node *goquery.Selection) {
		temp := node.Text()
		temp = trip(strings.Replace(temp, "\n", "", -1))
		temp = strings.Replace(temp, "&nbsp;", "", -1)
		temp = strings.Replace(temp, "　", "", -1)
		info = info + "\n" + temp
	})

	dudu := tripemptyl(strings.Split(info, "\n"))
	for _, r := range dudu {
		rr := strings.Split(r, "：")
		dd := ""
		if len(rr) == 2 {
			dd = rr[1]
		} else {
			continue
		}
		if strings.Contains(r, "资本类型") {
			returnmap["zitype"] = dd
		} else if strings.Contains(r, "机构性质") {
			returnmap["jx"] = dd
		} else if strings.Contains(r, "注册地点") {
			returnmap["rloca"] = dd
		} else if strings.Contains(r, "成立时间") {
			returnmap["date"] = dd
		} else if strings.Contains(r, "机构总部") {
			returnmap["tloca"] = dd
		} else if strings.Contains(r, "投资阶段") {
			returnmap["tstage"] = dd
		} else {
		}
	}
	return returnmap
}

func trip(s string) string {
	return strings.TrimSpace(strings.Replace(s, ",", "", -1))
}

func tripemptyl(dudu []string) []string {
	returnlist := []string{}
	for _, r := range dudu {
		if trip(r) != "" {
			returnlist = append(returnlist, trip(r))
		}
	}
	return returnlist
}
func featchcompany(keyword string) ([]map[string]string, int, error) {
	returnmap := []map[string]string{}
	url := "http://zdb.pedaily.cn/company/w" + keyword
	rootdir := strings.Split(keyword, "/")[0]
	hashmd := util.Md5(url)
	keep := dir + "/" + rootdir + "/" + hashmd + ".html"
	if util.FileExist(keep) {
		dudu, _ := util.ReadfromFile(keep)
		return parsecompany(dudu)
	}
	fmt.Printf("featch:%s\n", url)
	client.Url = url
	body, err := client.Get()
	if err != nil {
		return returnmap, 0, err
	}
	e := util.SaveToFile(keep, body)
	if e != nil {
		fmt.Println(url + ":" + e.Error())
	}
	return parsecompany(body)
}

func parsecompany(body []byte) ([]map[string]string, int, error) {
	returnmap := []map[string]string{}
	d, e := expert.QueryBytes(body)
	if e != nil {
		return returnmap, 0, e
	}
	total := d.Find(".total").Text()
	num, e := util.SI(total)
	if e != nil {
		return returnmap, 0, nil
	}
	d.Find(".company-list li").Each(func(i int, node *goquery.Selection) {
		temp := map[string]string{}
		content := node.Find(".txt a.f16")
		abbr := content.Next().Text()
		title := content.Text()
		href, ok := content.Attr("href")
		if !ok {
			return
		} else {
			href = "http://zdb.pedaily.cn" + href
		}
		//location := node.Find(".txt .location").Text()
		//desc := strings.Replace(node.Find(".desc").Text(), ",", "", -1)
		//desc = strings.Replace(desc, "\n", "", -1)
		//desc = strings.TrimSpace(desc)
		//desc = strings.Replace(desc, "\r", "", -1)
		temp["title"] = title
		temp["abbr"] = abbr
		temp["href"] = href
		hreft := strings.Split(href, "show")
		if len(hreft) == 0 {
			return
		}
		temp["hreft"] = "http://zdb.pedaily.cn/company/" + hreft[len(hreft)-1] + "vc/"
		//temp["location"] = location
		//temp["desc"] = desc
		//fmt.Printf("%s,%s,%s,%s,%s\n", title, href, temp["hreft"], abbr, location)
		returnmap = append(returnmap, temp)
	})

	return returnmap, num, nil
}

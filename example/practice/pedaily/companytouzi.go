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
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hunterhug/marmot/expert"
	"github.com/hunterhug/marmot/miner"
	"github.com/hunterhug/parrot/util"
)

var tclient *miner.Worker
var tresult = "./data/companyt/result"
var traw = "./data/companyt/raw"

func main() {
	dudu()
	inittouzi()
	tmain()
	//b, _ := util.ReadfromFile(tresult + "/3392.html")
	//l := parsetouzi(b)
	//fmt.Printf("%#v", l)
}

func parset(body []byte) ([]string, string) {
	returnlist := []string{}
	doc, _ := expert.QueryBytes(body)
	total := doc.Find(".total").Text()
	doc.Find("#inv-list li").Each(func(i int, node *goquery.Selection) {
		href, ok := node.Find("dt.view a").Attr("href")
		if !ok {
			return
		}
		href = "http://zdb.pedaily.cn" + href
		fmt.Printf("Inv: %s:%s\n", node.Find(".company a").Text(), href)
		returnlist = append(returnlist, href)
	})
	return returnlist, total
}
func inittouzi() {
	var e error = nil
	tclient, e = miner.NewWorker(nil)
	if e != nil {
		panic(e.Error())
	}

	util.MakeDir(tresult)
	util.MakeDir(traw)
}
func tmain() {
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please input url: ")
		//fmt.Scanln(&keyword)
		keyword, err := inputReader.ReadString('\n')
		keyword = strings.Replace(keyword, "\n", "", -1)
		keyword = trip(strings.Replace(keyword, "\r", "", -1))
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		r, _ := regexp.Compile(`[\d]+`)
		mark := r.FindString(keyword)
		if mark == "" {
			fmt.Println("找不到")
			continue
		}
		urls := []string{}
		loop := []string{}
		loop = append(loop, "y-2004")
		tyear, _ := util.SI(util.TodayString(1))
		for i := 2014; i <= tyear; i++ {
			loop = append(loop, "y"+util.IS(i))
		}

		result := []map[string]string{}
		for _, pp := range loop {
			url := "http://zdb.pedaily.cn/company/" + mark + "/vc/" + pp
			body, e := fetchpage(url)
			if e != nil {
				fmt.Println(e.Error())
				continue
			} else {
				fmt.Println("fetch " + url)
			}
			//e = util.SaveToFile(tresult+"/"+mark+".html", body)
			//if e != nil {
			//	fmt.Println(e.Error())
			//}
			l, t := parset(body)
			fmt.Printf("%s:total:%s\n", pp, t)
			total, e := util.SI(t)
			if e != nil {
				fmt.Println(e.Error())
				continue
			}
			if total == 0 {
				fmt.Println("empty")
				continue
			}
			urls = append(urls, l...)
			page := int(math.Ceil(float64(total) / 20.0))
			for i := 2; i <= page; i++ {
				url := "http://zdb.pedaily.cn/company/" + mark + "/vc/" + pp + "/" + util.IS(i)
				body, e = fetchpage(url)
				if e != nil {
					fmt.Println(e.Error())
					continue
				} else {
					fmt.Println("fetch " + url)
				}
				temp, _ := parset(body)
				urls = append(urls, temp...)
			}
		}
		if len(urls) == 0 {
			fmt.Println("empty")
			continue
		}
		//fmt.Printf("%#v\n", urls)
		for _, url := range urls {
			body := []byte("")
			var e error = nil
			keep := traw + "/" + util.Md5(url) + ".html"
			if util.FileExist(keep) {
				body, e = util.ReadfromFile(keep)
			} else {
				body, e = fetchpage(url)
			}
			if e != nil {
				fmt.Println(e.Error())
				continue
			} else {
				fmt.Println("fetch " + url)
			}
			util.SaveToFile(keep, []byte(body))
			dududu := parsetouzi(body)
			dududu["url"] = url
			result = append(result, dududu)
		}

		if len(result) == 0 {
			fmt.Println("empty")
			continue
		}
		s := []string{"页面,事件名称,融资方,投资方,金额,融资时间,轮次,所属行业,简介"}
		for _, jinhan := range result {
			s = append(s, jinhan["url"]+","+jinhan["name"]+","+jinhan["rf"]+","+jinhan["tf"]+","+jinhan["money"]+","+jinhan["date"]+","+jinhan["times"]+","+jinhan["han"]+","+jinhan["desc"])
		}

		util.SaveToFile(tresult+"/"+mark+".csv", []byte(strings.Join(s, "\n")))
	}
}

func parsetouzi(body []byte) map[string]string {
	returnmap := map[string]string{
		"name":  "",
		"rf":    "",
		"tf":    "",
		"money": "",
		"date":  "",
		"times": "",
		"han":   "",
		"desc":  "",
	}
	doc, _ := expert.QueryBytes(body)
	returnmap["name"] = doc.Find(".info h1").Text()
	returnmap["desc"] = strings.Replace(trip(doc.Find("#desc").Text()), "\n", "<br/>", -1)

	info := ""
	doc.Find(".info ul li").Each(func(i int, node *goquery.Selection) {
		temp := node.Text()
		temp = trip(strings.Replace(temp, "\n", "", -1))
		temp = strings.Replace(temp, "&nbsp;", "", -1)
		temp = strings.Replace(temp, "　", "", -1)
		info = info + "\n" + temp
	})

	//fmt.Println(info)
	dudu := tripemptyl(strings.Split(info, "\n"))
	//fmt.Printf("%#v\n", dudu)
	for _, r := range dudu {
		rr := strings.Split(r, "：")
		dd := ""
		if len(rr) == 2 {
			dd = strings.Replace(rr[1], " ", "", -1)
		} else {
			continue
		}
		if strings.Contains(r, "融") && strings.Contains(r, "方") {
			returnmap["rf"] = dd
		} else if strings.Contains(r, "投") && strings.Contains(r, "方") {
			returnmap["tf"] = dd
		} else if strings.Contains(r, "金") && strings.Contains(r, "额") {
			returnmap["money"] = dd
		} else if strings.Contains(r, "融资时间") {
			returnmap["date"] = dd
		} else if strings.Contains(r, "轮") && strings.Contains(r, "次") {
			returnmap["times"] = dd
		} else if strings.Contains(r, "所属行业") {
			returnmap["han"] = dd
		} else {
		}
	}
	return returnmap
}
func fetchpage(url string) ([]byte, error) {
	tclient.Url = url
	return tclient.Get()
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
func dudu() {
	fmt.Println(`
************************************************************

		投资界根据公司查找投资案例

		1.输入URL或者输入数字587等
		如：http://zdb.pedaily.cn/company/587/vc/

		2.查看结果
		查看data/company/tresult中csv文件

		/*
		go build *.go，然后点击exe运行或go run *.go
		*/
************************************************************
`)
}

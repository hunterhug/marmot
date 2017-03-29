/*
Copyright 2017 hunterhug/一只尼玛.
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

package spider

import (
	"github.com/hunterhug/GoSpider/util"
	"math/rand"
	"strings"
	"sync"
)

var (
	// 爬虫池子
	Pool = &_Spider{brower: make(map[string]*Spider)}
	Ua   = map[int]string{}
)

type _Spider struct {
	mux    sync.RWMutex
	brower map[string]*Spider
}

func (sb *_Spider) Get(name string) (b *Spider, ok bool) {
	sb.mux.RLock()
	b, ok = sb.brower[name]
	sb.mux.RUnlock()
	return
}

func (sb *_Spider) Set(name string, b *Spider) {
	sb.mux.Lock()
	sb.brower[name] = b
	sb.mux.Unlock()
	return
}

func (sb *_Spider) Delete(name string) {
	sb.mux.Lock()
	delete(sb.brower, name)
	sb.mux.Unlock()
	return
}

// Ua初始化
func UaInit() {
	Ua[0] = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71 Safari/537.36"
	temp, err := util.ReadfromFile(util.CurDir() + "/config/ua.txt")
	if err != nil {
		//panic(err.Error())
	} else {
		uas := strings.Split(string(temp), "\n")
		for i, ua := range uas {
			Ua[i] = strings.TrimSpace(strings.Replace(ua, "\r", "", -1))
		}
	}

}

// 返回随机Ua
func RandomUa() string {
	length := len(Ua)
	if length == 0 {
		return ""
	}
	return Ua[rand.Intn(length-1)]
}

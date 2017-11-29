/*
Copyright 2017 by GoSpider author. Email: gdccmcm14@live.com
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
	"github.com/hunterhug/GoTool/util"
	"math/rand"
	"strings"
)

var Ua = map[int]string{}

// User-Agent init
func UaInit() {
	Ua[0] = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71 Safari/537.36"

	// this *.txt maybe not found if you exec binary, so we just fill several ua
	temp, err := util.ReadfromFile(util.CurDir() + "/config/ua.txt")

	if err != nil {
		Ua[1] = "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; de-de) AppleWebKit/125.5.5 (KHTML, like Gecko) Safari/125.12"
		Ua[2] = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:46.0) Gecko/20100101 Firefox/46.0"
	} else {
		uas := strings.Split(string(temp), "\n")
		for i, ua := range uas {
			Ua[i] = strings.TrimSpace(strings.Replace(ua, "\r", "", -1))
		}
	}

}

// Reback random User-Agent
func RandomUa() string {
	length := len(Ua)
	if length == 0 {
		return ""
	}
	return Ua[rand.Intn(length-1)]
}

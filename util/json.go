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

package util

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// 将字符串转成适合JSON格式的，比如中文转为\\u
func StringsToJson(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}

	return jsons
}

// 将JSON中的/u unicode乱码转回来，笨方法，弃用
func JsonEncode(raw string) string {
	raw = strings.Replace(raw, "\"", "\\u\"", -1)
	sUnicodev := strings.Split(raw, "\\u")
	leng := len(sUnicodev)
	if leng <= 1 {
		return raw
	}
	var context string
	for _, v := range sUnicodev {
		if len(v) <= 1 {
			context += fmt.Sprintf("%s", v)
			continue
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			context += fmt.Sprintf("%s", v)
		} else {
			context += fmt.Sprintf("%c", temp)
		}
	}
	return context
}

// 最好的方法
func JsonBack(s []byte) ([]byte, error) {
	temp := new(interface{})
	json.Unmarshal(s, temp)
	resut, err := json.Marshal(temp)
	return resut, err
}

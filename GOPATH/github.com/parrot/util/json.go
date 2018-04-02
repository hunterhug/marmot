/*
	版权所有，侵权必究
	署名-非商业性使用-禁止演绎 4.0 国际
	警告： 以下的代码版权归属hunterhug，请不要传播或修改代码
	你可以在教育用途下使用该代码，但是禁止公司或个人用于商业用途(在未授权情况下不得用于盈利)
	商业授权请联系邮箱：gdccmcm14@live.com QQ:459527502

	All right reserved
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
	For more information on commercial licensing please contact hunterhug.
	Ask for commercial licensing please contact Mail:gdccmcm14@live.com Or QQ:459527502

	2017.7 by hunterhug
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

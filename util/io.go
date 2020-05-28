/*
	All right reserved https://github.com/hunterhug/marmot at 2016-2020
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
*/
package util

import (
	"fmt"
	"strings"
)

func Input(say, defaults string) string {
	fmt.Println(say)
	var str string
	fmt.Scanln(&str)
	if strings.TrimSpace(str) == "" {
		if strings.TrimSpace(defaults) != "" {
			return defaults
		} else {
			fmt.Println("can not empty")
			return Input(say, defaults)
		}
	}
	//fmt.Println("--" + str + "--")
	return str
}

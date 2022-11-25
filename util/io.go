package util

import (
	"fmt"
	"strings"
)

func Input(say, defaults string) string {
	fmt.Println(say)

	var str string
	_, _ = fmt.Scanln(&str)

	if strings.TrimSpace(str) == "" {
		if strings.TrimSpace(defaults) != "" {
			return defaults
		} else {
			fmt.Println("can not empty")
			return Input(say, defaults)
		}
	}

	return str
}

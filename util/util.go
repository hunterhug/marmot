package util

import (
	"errors"
	"strconv"
	"strings"
)

func SI(s string) (i int, e error) {
	i, e = strconv.Atoi(s)
	return
}

func IS(i int) string {
	return strconv.Itoa(i)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func DivideStringList(files []string, num int) (map[int][]string, error) {
	length := len(files)
	split := map[int][]string{}
	if num <= 0 {
		return split, errors.New("num must not negative")
	}

	if num > length {
		num = length
	}

	process := length / num
	for i := 0; i < num; i++ {
		split[i] = append(split[i], files[i*process:(i+1)*process]...)
	}

	remain := files[num*process:]
	for i := 0; i < len(remain); i++ {
		split[i%num] = append(split[i%num], remain[i])
	}

	return split, nil
}

func InArray(sa []string, a string) bool {
	for _, v := range sa {
		if a == v {
			return true
		}
	}
	return false
}

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}

	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}

	if start > rl {
		start = rl
	}

	if end < 0 {
		end = 0
	}

	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

package util

import (
	"time"
)

func Sleep(waitTime int) {
	time.Sleep(time.Duration(waitTime) * time.Second)
}

func Second(times int) time.Duration {
	return time.Duration(times) * time.Second
}

func TodayString(level int) string {
	formats := "20060102150405"
	switch level {
	case 1:
		formats = "2006"
	case 2:
		formats = "200601"
	case 3:
		formats = "20060102"
	case 4:
		formats = "2006010215"
	case 5:
		formats = "200601021504"
	default:

	}
	return time.Now().Format(formats)
}

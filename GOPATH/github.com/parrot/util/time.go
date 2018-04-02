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
	"time"
)

// sleep
func Sleep(waittime int) {
	time.Sleep(time.Duration(waittime) * time.Second)
}

// time
func Second(times int) time.Duration {
	return time.Duration(times) * time.Second
}

// get secord times
// 172606056
func GetSecordTimes() int64 {
	return time.Now().Unix()
}

//201611112113
func GetSecord2DateTimes(secord int64) string {
	tm := time.Unix(secord, 0)
	return tm.Format("20060102150405")

}

func GetDateTimes2Secord(datestring string) int64 {
	tm2, _ := time.Parse("20060102150405", datestring)
	return tm2.Unix()

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

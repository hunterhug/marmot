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
package miner

import (
	"testing"
)

func TestWorker(t *testing.T) {
	// Global log record
	SetLogLevel("debug")

	// GLOBAL TIMEOUT
	SetGlobalTimeout(3)

	// a new spider without proxy
	// NewWorker(nil)
	worker, err := NewWorker(nil)

	//proxy := "http://smart:smart2016@104.128.121.46:808"
	//worker, err := NewWorker(proxy)

	if err != nil {
		panic(err)
	}
	// method can be get and post
	worker.SetMethod(GET)
	// wait times,can zero
	worker.SetWaitTime(1)
	// which url fetch
	worker.SetUrl("http://www.cjhug.me")

	//worker.SetUa(spider.RandomUa())

	// go!fetch url --||
	body, err := worker.Go()
	if err != nil {
		Log().Error(err.Error())
	} else {
		// bytes get!
		// fmt.Printf("%s", string(body))
	}

	// if filesize small than 500KB
	err = TooSortSizes(body, 500)
	Log().Error(err.Error())
}

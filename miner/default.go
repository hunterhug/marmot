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
*
*/

package miner

import (
	"net/http"
	"net/url"
)

// Global Worker
var DefaultWorker *Worker

func init() {
	UaInit()

	// New a Worker
	worker := new(Worker)
	worker.Header = http.Header{}
	worker.Data = url.Values{}
	worker.BData = []byte{}
	worker.Client = Client

	// Global Worker!
	DefaultWorker = worker

}

// This make effect only your Worker exec serial! Attention!
// Change Your Raw data To string
func ToString() string {
	return DefaultWorker.ToString()
}

// This make effect only your Worker exec serial! Attention!
// Change Your JSON'like Raw data to string
func JsonToString() (string, error) {
	return DefaultWorker.JsonToString()
}

func Get() (body []byte, e error) {
	return DefaultWorker.Get()
}

func Delete() (body []byte, e error) {
	return DefaultWorker.Delete()
}

func Go() (body []byte, e error) {
	return DefaultWorker.Go()
}

func GoByMethod(method string) (body []byte, e error) {
	return DefaultWorker.SetMethod(method).Go()
}

func OtherGo(method, contenttype string) (body []byte, e error) {
	return DefaultWorker.OtherGo(method, contenttype)
}

func Post() (body []byte, e error) {
	return DefaultWorker.Post()
}

func PostJSON() (body []byte, e error) {
	return DefaultWorker.PostJSON()
}

func PostFILE() (body []byte, e error) {
	return DefaultWorker.PostFILE()
}

func PostXML() (body []byte, e error) {
	return DefaultWorker.PostXML()
}

func Put() (body []byte, e error) {
	return DefaultWorker.Put()
}
func PutJSON() (body []byte, e error) {
	return DefaultWorker.PutJSON()
}

func PutFILE() (body []byte, e error) {
	return DefaultWorker.PutFILE()
}

func PutXML() (body []byte, e error) {
	return DefaultWorker.PutXML()
}

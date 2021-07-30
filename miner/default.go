/*
	All right reserved https://github.com/hunterhug/marmot at 2016-2021
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
*/
package miner

import "net/http"

// DefaultWorker Global Worker
var DefaultWorker *Worker

func init() {
	UaInit()

	// New a Worker
	worker := new(Worker)
	worker.Request = newRequest()
	worker.Header = http.Header{}
	worker.Response = new(Response)
	worker.Client = Client

	// Global Worker!
	DefaultWorker = worker
}

// ToString This make effect only your Worker exec serial! Attention!
// Change Your Raw data To string
func ToString() string {
	return DefaultWorker.ToString()
}

// JsonToString This make effect only your Worker exec serial! Attention!
// Change Your JSON like Raw data to string
func JsonToString() (string, error) {
	return DefaultWorker.JsonToString()
}

func Clone() *Worker {
	return DefaultWorker.Clone()
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

func OtherGo(method, contentType string) (body []byte, e error) {
	return DefaultWorker.OtherGo(method, contentType)
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

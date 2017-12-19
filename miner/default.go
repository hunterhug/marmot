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

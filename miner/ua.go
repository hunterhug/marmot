package miner

import (
	"math/rand"
	"strings"

	"github.com/hunterhug/parrot/util"
)

// Global User-Agent provide
var Ua = map[int]string{}

// User-Agent init
func UaInit() {
	Ua = map[int]string{
		0: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; de-de) AppleWebKit/125.5.5 (KHTML, like Gecko) Safari/125.12",
		1: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:46.0) Gecko/20100101 Firefox/46.0",
		2: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71 Safari/537.36",
		3: "Opera/9.60 (Macintosh; Intel Mac OS X; U; en) Presto/2.1.1",
	}
	// this *.txt maybe not found if you exec binary, so we just fill several ua
	temp, err := util.ReadfromFile(util.CurDir() + "/config/ua.txt")

	if err == nil {
		uas := strings.Split(string(temp), "\n")
		for i, ua := range uas {
			Ua[i] = strings.TrimSpace(strings.Replace(ua, "\r", "", -1))
		}
	}

}

// Reback random User-Agent
func RandomUa() string {
	length := len(Ua)
	if length == 0 {
		return ""
	}
	return Ua[rand.Intn(length-1)]
}

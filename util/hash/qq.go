package hash

import (
	"bitbucket.org/xtalpi/vm-srv/src/config"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/hunterhug/GoSpider/spider"
	"math/rand"
	urll "net/url"
	"sort"
	"strings"
	"time"
)

const (
	qqurl = ".api.qcloud.com/v2/index.php"
)

func InitCred(c map[string]config.QQ) {

}
func joinbody(body map[string]interface{}) string {

	t := []string{}
	for k := range body {
		t = append(t, k)
	}
	sort.Strings(t)
	s := []string{}
	for _, v := range t {
		s = append(s, fmt.Sprintf("%s=%v", strings.Replace(v, "_", ".", -1), body[v]))
	}
	return strings.Join(s, "&")
}

func Sign(secretId, secretKey string, apitype string, body map[string]interface{}) ([]byte, error) {
	if body["SecretId"] == nil {
		body["SecretId"] = secretId
	}
	if body["Timestamp"] == nil {
		body["Timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	}
	if body["Nonce"] == nil {
		rand.Seed(time.Now().UnixNano())
		body["Nonce"] = fmt.Sprintf("%v", rand.Int())
	}
	/*
		if body["Region"] == nil {
			body["Region"] = "gz"
		}
		if body["Action"] == nil {
			body["Action"] = "DescribeInstances"
		}
	*/

	if body["Version"] == nil {
		body["Version"] = "2017-03-12"
	}

	temp1 := joinbody(body)
	//fmt.Println(temp1)
	temp2 := "POST" + apitype + qqurl + "?" + temp1

	//fmt.Println(temp2)

	al := fmt.Sprintf("%v", body["SignatureMethod"])

	signstr := ""
	if al == "HmacSHA256" {
		signstr = ComputeHmac256(temp2, secretKey)
	} else {
		signstr = ComputeHmac1(temp2, secretKey)
	}

	api := spider.NewAPI()
	api.SetUrl("https://" + apitype + qqurl)

	body["Signature"] = signstr

	pform := urll.Values{}
	for k, v := range body {
		pform.Add(k, fmt.Sprintf("%v", v))
	}

	fmt.Printf("ASK %#v\n", body)
	api.SetForm(pform)
	result, err := api.Post()

	debug := true
	if debug {
		var a interface{}
		json.Unmarshal(result, &a)
		jsonOut, _ := json.MarshalIndent(a, "", "  ")
		return jsonOut, err

	}
	return result, err
}

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func ComputeHmac1(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

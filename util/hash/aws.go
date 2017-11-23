package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"sort"
	"strings"
	"time"
)

// 认证的配置
type AwsAuth struct {
	AwsID      string
	AwsKey     string
	AwsRegion  string
	AwsService string
}


/*

			req.ParseForm()
			//fmt.Println(AwsConfig, req.URL.Path, req.Method, req.URL.Host, req.Form)
			amzdate, authorization_header := AwsAuthSignature(AwsConfig, UriEncode(req.URL.Path, true), req.Method, req.URL.Host, req.Form, buf)
			req.Header.Set("X-Amz-Date", amzdate)
			req.Header.Set("Authorization", authorization_header)
*/
func AwsAuthSignature(auth AwsAuth, uri, method, host string, query url.Values, data []byte) (amzdate, authorization_header string) {

	// 基本配置
	access_key := auth.AwsID
	secret_key := auth.AwsKey
	region := auth.AwsRegion
	service := auth.AwsService

	request_parameters := ""

	// 查询字符串排序
	if query != nil {
		temp := []string{}
		for k, _ := range query {
			temp = append(temp, k)
		}
		sort.Strings(temp)
		temp1 := []string{}
		for _, v := range temp {
			temp1 = append(temp1, UriEncode(v, false)+"="+UriEncode(query.Get(v), false))
		}
		request_parameters = strings.Join(temp1, "&")
	}

	// 现在时间
	now := time.Now().UTC()
	amzdate = now.Format("20060102T150405Z")
	datestamp := now.Format("20060102")

	canonical_uri := uri
	canonical_querystring := request_parameters
	canonical_headers := "host:" + host + "\n" + "x-amz-date:" + amzdate + "\n"
	signed_headers := "host;x-amz-date"

	payload_hash := hex.EncodeToString(getSha256Code(""))
	if data != nil {
		payload_hash = hex.EncodeToString(getSha256Code(string(data)))
	}

	canonical_request := method + "\n" + canonical_uri + "\n" + canonical_querystring + "\n" + canonical_headers + "\n" + signed_headers + "\n" + payload_hash
	//fmt.Printf("%q\n", canonical_request)
	algorithm := "AWS4-HMAC-SHA256"

	credential_scope := datestamp + "/" + region + "/" + service + "/" + "aws4_request"
	string_to_sign := algorithm + "\n" + amzdate + "\n" + credential_scope + "\n" + hex.EncodeToString(getSha256Code(canonical_request))
	signing_key := GetSignatureKey(secret_key, datestamp, region, service)
	signature := hex.EncodeToString(SignAws(signing_key, []byte(string_to_sign)))
	authorization_header = algorithm + " " + "Credential=" + access_key + "/" + credential_scope + ", " + "SignedHeaders=" + signed_headers + ", " + "Signature=" + signature

	return amzdate, authorization_header
}

func SignAws(key, msg []byte) []byte {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(msg))
	return h.Sum(nil)
}

func GetSignatureKey(key, dateStamp, regionName, serviceName string) []byte {
	kDate := SignAws([]byte("AWS4"+key), []byte(dateStamp))
	kRegion := SignAws(kDate, []byte(regionName))
	kService := SignAws(kRegion, []byte(serviceName))
	kSigning := SignAws(kService, []byte("aws4_request"))
	return kSigning
}

func getSha256Code(s string) []byte {
	h := sha256.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}

func UriEncode(src string, encodeSlash bool) string {
	// application/x-www-form-urlencoded will have +
	back := url.QueryEscape(src)
	// all change but + must replace, RFC3986
	temp := strings.Replace(back, "+", "%20", -1)

	// uri / must be keep
	if encodeSlash {
		return strings.Replace(temp, "%2F", "/", -1)
	}
	return temp
}

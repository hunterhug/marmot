/*
	All right reserved https://github.com/hunterhug/marmot at 2016-2021
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
*/
package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"net/url"
	"strings"
)

// ComputeHmac256 HMAC with the SHA256
func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// StrToMd5 create md5 string
func StrToMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

func Md5(str string) string {
	return StrToMd5(str)
}

func Base64E(body string) string {
	str := []byte(body)
	data := base64.StdEncoding.EncodeToString(str)
	return data
}

func Base64D(body string) string {
	data, err := base64.StdEncoding.DecodeString(body)
	if err != nil {
		return ""
	}
	s := fmt.Sprintf("%q", data)
	s = strings.Replace(s, "\"", "", -1)
	return s
}

func UrlE(s string) string {
	return url.QueryEscape(s)
}

func UrlD(s string) string {
	s, e := url.QueryUnescape(s)
	if e != nil {
		return e.Error()
	} else {
		return s
	}
}

func Md5FS(src io.Reader) string {
	h := md5.New()
	if err := CopyFF(src, h); err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return fmt.Sprintf("%x", h.Sum([]byte("hunterhug")))
}

/*
	All right reserved https://github.com/hunterhug/marmot at 2016-2021
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
*/
package miner

import "net/http"

const (
	VERSION = "1.0.9"

	// GET HTTP method
	GET      = "GET"
	POST     = "POST"
	POSTJSON = "POSTJSON"
	POSTXML  = "POSTXML"
	POSTFILE = "POSTFILE"
	PUT      = "PUT"
	PUTJSON  = "PUTJSON"
	PUTXML   = "PUTXML"
	PUTFILE  = "PUTFILE"
	DELETE   = "DELETE"
	OTHER    = "OTHER" // this stand for you can use other method this lib not own.

	// HTTPFORMContentType HTTP content type
	HTTPFORMContentType = "application/x-www-form-urlencoded"
	HTTPJSONContentType = "application/json"
	HTTPXMLContentType  = "text/xml"
	HTTPFILEContentType = "multipart/form-data"

	// CRITICAL Log mark
	CRITICAL = "CRITICAL"
	ERROR    = "ERROR"
	WARNING  = "WARNING"
	NOTICE   = "NOTICE"
	INFO     = "INFO"
	DEBUG    = "DEBUG"
)

var (
	// Browser User-Agent, Our default Http ua header!
	ourLoveUa = "Marmot+" + VERSION + "+github:hunterhug"

	DefaultHeader = map[string][]string{
		"User-Agent": {
			ourLoveUa,
		},
	}

	// DefaultTimeOut http get and post No timeout
	DefaultTimeOut = 0
)

// SetGlobalTimeout Set global timeout, it can only by this way!
func SetGlobalTimeout(num int) {
	DefaultTimeOut = num
}

// MergeCookie Merge Cookie, not use
func MergeCookie(before []*http.Cookie, after []*http.Cookie) []*http.Cookie {
	cs := make(map[string]*http.Cookie)

	for _, b := range before {
		cs[b.Name] = b
	}

	for _, a := range after {
		if a.Value != "" {
			cs[a.Name] = a
		}
	}

	res := make([]*http.Cookie, 0, len(cs))

	for _, q := range cs {
		res = append(res, q)

	}

	return res

}

// CloneHeader Clone a header, If not exist Ua, Set our Ua!
func CloneHeader(h map[string][]string) map[string][]string {
	if h == nil || len(h) == 0 {
		h = DefaultHeader
		return h
		//return map[string][]string{}
	}

	if len(h["User-Agent"]) == 0 {
		h["User-Agent"] = []string{ourLoveUa}
	}
	return CopyM(h)
}

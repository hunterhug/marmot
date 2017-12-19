package miner

import "net/http"

const (
	// Default wait time
	WaitTime = 5

	// HTTP method
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

	// HTTP content type
	HTTPFORMContentType = "application/x-www-form-urlencoded"
	HTTPJSONContentType = "application/json"
	HTTPXMLContentType  = "text/xml"
	HTTPFILEContentType = "multipart/form-data"

	// Log mark
	CRITICAL = "CRITICAL"
	ERROR    = "ERROR"
	WARNING  = "WARNING"
	NOTICE   = "NOTICE"
	INFO     = "INFO"
	DEBUG    = "DEBUG"
)

var (
	// Browser User-Agent, Our default Http ua header!
	ourloveUa = "Marmot+hunterhug"

	DefaultHeader = map[string][]string{
		"User-Agent": {
			ourloveUa,
		},
	}

	// DefaultTimeOut,http get and post No timeout
	DefaultTimeOut = 0
)

// Set global timeout, it can only by this way!
func SetGlobalTimeout(num int) {
	DefaultTimeOut = num
}

// Merge Cookie, not use
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

// Clone a header, If not exist Ua, Set our Ua!
func CloneHeader(h map[string][]string) map[string][]string {
	if h == nil || len(h) == 0 {
		h = DefaultHeader
		return h
		//return map[string][]string{}
	}

	if len(h["User-Agent"]) == 0 {
		h["User-Agent"] = []string{ourloveUa}
	}
	return CopyM(h)
}

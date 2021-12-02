package goxios

import (
	"fmt"
	httppkg "net/http"
	urlpkg "net/url"
	"time"
)

const (
	// Method
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	PATCH  = "PATCH"
	DELETE = "DELETE"
	//
	ContentTypeJSON = "application/json"
	ContentTypeText = "application/x-www-form-urlencoded"
	UserAgentKey    = "User-Agent"
	UserAgentValue  = "Mozilla/5.0 (Linux; Android 8.0.0; Pixel 2 XL Build/OPD1.170816.004) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Mobile Safari/537.36"
)

type RequestConfig struct {
	Method    string
	Url       string
	Params    urlpkg.Values
	Header    httppkg.Header
	Data      Reader
	Timeout   time.Duration
	Transport *httppkg.Transport //socks5
	Cookies   []*httppkg.Cookie
}

func (r *RequestConfig) params() error {
	url := r.Url
	params := r.Params
	parse, err := urlpkg.Parse(url)
	if err != nil {
		return err
	}
	values, err := urlpkg.ParseQuery(parse.RawQuery)
	if err != nil {
		return err
	}
	for k, v := range params {
		values[k] = append(values[k], v...)
	}
	parse.RawQuery = values.Encode()
	r.Url = parse.String()
	return nil
}

type Header map[string]interface{}

func NewHeader(data Header) httppkg.Header {
	dast := make(httppkg.Header)
	for k, v := range data {
		dast.Add(k, fmt.Sprintf("%v", v))
	}
	return dast
}

type Params map[string]interface{}

func NewParams(data Params) urlpkg.Values {
	dast := make(urlpkg.Values)
	for k, v := range data {
		dast.Add(k, fmt.Sprintf("%v", v))
	}
	return dast
}

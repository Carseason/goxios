package goxios

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

//default
const (
	ContentTypeJSON = "application/json"
	ContentTypeText = "application/x-www-form-urlencoded"
	UserAgent       = "Mozilla/5.0 (Linux; Android 8.0.0; Pixel 2 XL Build/OPD1.170816.004) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Mobile Safari/537.36"
)

type Config struct {
	Url     string
	Method  string
	Params  map[string]string
	Headers Headers
	Data    io.Reader
}

type Headers struct {
	ContentType string
	Cookie      string
	UserAgent   string
	Host        string
	Referer     string
	Data        map[string]string
}

type Goxios struct {
	res           *http.Request
	url           *url.URL
	contentType   string
	err           error
	body          []byte
	headers       http.Header
	cookies       []*http.Cookie
	statusCode    int
	contentLength int64
	proto         string
}

//set URL params
func (a *Goxios) paramsURL(uri string, data map[string]string) string {
	if len(data) > 0 {
		params := make(url.Values)
		for k, v := range data {
			params.Add(k, v)
		}
		if strings.Contains(uri, "?") {
			uri = uri + "&" + params.Encode()
		} else {
			uri = uri + "?" + params.Encode()
		}
	}
	return uri
}

//setheader
func (a *Goxios) setHeaders(header Headers) *Goxios {
	a.res.Header.Set("Cookie", header.Cookie)
	a.res.Header.Set("User-Agent", header.UserAgent)
	a.res.Header.Set("Content-Type", header.ContentType)
	for k, v := range header.Data {
		a.res.Header.Add(k, v)
	}
	return a
}

//data
func (a *Goxios) bodyReader(e string, data interface{}) (io.Reader, error) {
	//"application/json"
	if e == "" || strings.Contains(e, "json") {
		switch value := data.(type) {
		case string:
			return bytes.NewBuffer([]byte(value)), nil
		case []byte:
			return bytes.NewBuffer(value), nil
		case io.Reader:
			return value, nil
		default:
			bs, err := json.Marshal(data)
			return bytes.NewBuffer(bs), err
		}

	} else {
		//"application/x-www-form-urlencoded"
		switch value := data.(type) {
		case url.Values:
			return strings.NewReader(value.Encode()), nil
		case map[string]interface{}:
			body := make(url.Values)
			for k, v := range value {
				body.Set(k, fmt.Sprintf("%v", v))
			}
			return strings.NewReader(body.Encode()), nil
		case string:
			return strings.NewReader(value), nil
		default:
			switch reflect.TypeOf(value).Kind() {
			case reflect.Struct,
				reflect.Map,
				reflect.Slice,
				reflect.Array:
				bs, err := json.Marshal(value)
				return bytes.NewBuffer(bs), err
			default:
				return strings.NewReader(""), nil
			}
		}
	}
}

//do
func (a *Goxios) do() *Goxios {
	res := a.res
	resp, err := new(http.Client).Do(res)
	if err != nil {
		a.err = err
		return a
	}
	defer resp.Body.Close()
	a.url = resp.Request.URL
	a.statusCode = resp.StatusCode
	a.headers = resp.Header
	a.cookies = resp.Cookies()
	a.proto = resp.Proto
	a.contentLength = resp.ContentLength
	a.body, a.err = ioutil.ReadAll(resp.Body)
	return a
}

//success
func (a *Goxios) Then(f func(res Goxios)) *Goxios {
	goxios := *a
	if goxios.err == nil {
		f(goxios)
	}
	return a
}

//failing
func (a *Goxios) Catch(f func(err error)) *Goxios {
	goxios := *a
	if goxios.err != nil {
		f(goxios.err)
	}
	return a
}

//result
func (a *Goxios) Body() []byte {
	return a.body
}
func (a *Goxios) Content() string {
	return string(a.body)
}
func (a *Goxios) Cookie() []*http.Cookie {
	return a.cookies
}
func (a *Goxios) StatusCode() int {
	return a.statusCode
}
func (a *Goxios) Headers() http.Header {
	return a.headers
}
func (a *Goxios) ContentLength() int64 {
	return a.contentLength
}
func (a *Goxios) URL() *url.URL {
	return a.url
}
func (a *Goxios) Proto() string {
	return a.proto
}

//GET
func Get(c Config) (goxios *Goxios) {
	c.Method = "GET"
	return Request(c)
}

//POST
func Post(c Config) (goxios *Goxios) {
	c.Method = "POST"
	return Request(c)
}

//PUT
func Put(c Config) (goxios *Goxios) {
	c.Method = "PUT"
	return Request(c)
}

//DELETE
func Delete(c Config) (goxios *Goxios) {
	c.Method = "DELETE"
	return Request(c)
}

// request
func Request(c Config) (goxios *Goxios) {
	goxios = &Goxios{}
	method := strings.ToUpper(c.Method)
	uri := goxios.paramsURL(c.Url, c.Params)
	goxios.res, goxios.err = http.NewRequest(method, uri, nil)
	if goxios.err == nil {
		goxios.setHeaders(c.Headers)
		goxios.do()
	}
	return
}

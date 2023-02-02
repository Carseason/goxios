package goxios

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	url "net/url"
	"strings"
	"time"

	"golang.org/x/net/proxy"
)

type GoxiosConfig struct {
	Cookies       []*http.Cookie
	Header        http.Header
	Transport     *http.Transport //socks5
	Timeout       time.Duration
	ClearRedirect bool //不允许重定向
}
type RequestConfig struct {
	GoxiosConfig
	Params url.Values
	Body   io.Reader
}

func NewGoxiosConfig() *GoxiosConfig {
	return &GoxiosConfig{}
}
func NewRequestConfig() *RequestConfig {
	return &RequestConfig{}
}
func (g *GoxiosConfig) SetUserAget(ua string) *GoxiosConfig {
	g.AddHeader(http.Header{
		"User-Agent": []string{
			ua,
		},
	})
	return g
}
func (g *GoxiosConfig) SetDefaultUserAget() *GoxiosConfig {
	g.SetUserAget(UserAgentValue)
	return g
}
func (g *GoxiosConfig) SetTimeout(timeout time.Duration) *GoxiosConfig {
	g.Timeout = timeout
	return g
}
func (g *GoxiosConfig) SetTransport(transport *http.Transport) *GoxiosConfig {
	g.Transport = transport
	return g
}
func (g *GoxiosConfig) AddHeader(header http.Header) *GoxiosConfig {
	headers := make(http.Header)
	for k, v := range g.Header {
		for i := range v {
			headers.Add(k, v[i])
		}
	}
	for k, v := range header {
		for i := range v {
			headers.Add(k, v[i])
		}
	}
	g.Header = headers
	return g
}
func (g *GoxiosConfig) AddCookies(cookies []*http.Cookie) *GoxiosConfig {
	var value []*http.Cookie
	value = append(value, g.Cookies...)
	value = append(value, cookies...)
	g.Cookies = value
	return g
}
func NewHeader(values map[string]string) http.Header {
	header := make(http.Header)
	for k, v := range values {
		header.Add(k, v)
	}
	return header
}
func NewParams(values map[string]string) url.Values {
	params := make(url.Values)
	for k, v := range values {
		params.Add(k, v)
	}
	return params
}

func NewStringReader(values string) io.Reader {
	return strings.NewReader(values)
}
func NewFileReader(file *multipart.FileHeader, fieldname string, data map[string]string) (io.Reader, error) {
	files, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer files.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	defer writer.Close()
	for v, k := range data {
		writer.WriteField(v, k)
	}
	part, err := writer.CreateFormFile(fieldname, file.Filename)
	if err != nil {
		return nil, err
	}
	io.Copy(part, files)
	return body, nil
}

// "tcp" ":3908" , "user" ,"password"
func NewSOCKS5(network, addr, user, password string) (*http.Transport, error) {
	dialer, err := proxy.SOCKS5(network, addr, &proxy.Auth{
		User:     user,
		Password: password,
	}, proxy.Direct)
	if err != nil {
		return nil, err
	}
	tr := &http.Transport{Dial: dialer.Dial}
	return tr, nil
}

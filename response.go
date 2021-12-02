package goxios

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

type Response struct {
	Status           string
	StatusCode       int
	Proto            string
	ProtoMajor       int
	ProtoMinor       int
	Header           http.Header
	ContentLength    int64
	TransferEncoding []string
	Close            bool
	Uncompressed     bool
	Trailer          http.Header
	Request          *http.Request
	TLS              *tls.ConnectionState
	Body             []byte
}

func NewResponse(resp *http.Response) (Response, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}
	return Response{
		Status:           resp.Status,
		StatusCode:       resp.StatusCode,
		Proto:            resp.Proto,
		ProtoMajor:       resp.ProtoMajor,
		ProtoMinor:       resp.ProtoMajor,
		Header:           resp.Header,
		ContentLength:    resp.ContentLength,
		TransferEncoding: resp.TransferEncoding,
		Close:            resp.Close,
		Uncompressed:     resp.Uncompressed,
		Trailer:          resp.Trailer,
		Request:          resp.Request,
		TLS:              resp.TLS,
		Body:             body,
	}, nil
}
func (r *Response) Title() string {
	if r.Body == nil {
		return ""
	}
	data := regexp.MustCompile(`<title[^>]{0,}>([^>]+)</title[^>]{0,}>`).FindSubmatch(r.Body)
	if len(data) == 0 {
		return ""
	}
	return string(data[1])
}
func (r *Response) Content() string {
	return string(r.Body)
}
func (r *Response) Decoder(dast interface{}) error {
	return json.Unmarshal(r.Body, dast)
}
func (r *Response) Cookies() []*http.Cookie {
	if r.Request == nil {
		return nil
	}
	return r.Request.Cookies()
}
func (r *Response) Cookie(name string) (*http.Cookie, error) {
	if r.Request == nil {
		return nil, errors.New("request is nil")
	}
	return r.Request.Cookie(name)
}
func (r *Response) URL() *url.URL {
	if r.Request == nil {
		return new(url.URL)
	}
	return r.Request.URL
}

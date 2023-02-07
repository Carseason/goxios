package goxios

import (
	"encoding/json"
	"io"
	"net/http"
)

type Response struct {
	body     []byte
	resp     *http.Response
	Redirect Redirect
	*Node
}

func newResponse(resp *http.Response, redirect Redirect) (*Response, error) {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if by, err := DecodeEncoding(body); err == nil {
		body = by
	}
	response := &Response{
		body:     body,
		Redirect: redirect,
		resp:     resp,
		Node:     newNode(body),
	}
	return response, err
}
func (r *Response) Body() []byte {
	return r.body
}
func (r *Response) Content() string {
	return string(r.Body())
}
func (r *Response) Cookies() []*http.Cookie {
	return r.resp.Cookies()
}
func (r *Response) StatusCode() int {
	return r.resp.StatusCode
}

// 是否重定向
func (r *Response) IsRedirect() bool {
	if r.Redirect.IsRedirect() {
		return true
	}
	return isRedirect(r.StatusCode())
}
func (r *Response) JSON(obj interface{}) error {
	return json.Unmarshal(r.Body(), obj)
}
func (r *Response) Title() string {
	title, _ := r.QueryText("title")
	return title
}

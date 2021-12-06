package goxios

import (
	"net/http"
	"strings"
)

func Do(option RequestConfig) (dast *Request) {
	option.Method = strings.ToUpper(option.Method)
	switch option.Method {
	case GET, POST, PATCH, PUT, DELETE:
	default:
		if option.Method == "" {
			option.Method = GET
		}
	}
	dast = new(Request)
	if err := option.params(); err != nil {
		dast.err = err
		return
	}
	//
	req, err := http.NewRequest(option.Method, option.Url, option.Data)
	if err != nil {
		dast.err = err
		return
	}
	// cookie
	if option.Cookies != nil {
		for _, k := range option.Cookies {
			req.AddCookie(k)
		}
	}
	//
	dast.request = req
	// header
	dast.request.Header = option.Header
	client := &http.Client{}
	// socks5
	if option.Transport != nil {
		client.Transport = option.Transport
	}
	// timeout
	if option.Timeout > 0 {
		client.Timeout = option.Timeout
	}
	// do
	resp, err := client.Do(dast.request)
	if err != nil {
		dast.err = err
		return
	}
	response, err := NewResponse(resp)
	if err != nil {
		dast.err = err
	} else {
		dast.response = response
	}
	return
}

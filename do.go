package goxios

import (
	"errors"
	"net/http"
)

func Do(configs ...GoxiosConfig) *Goxios {
	g := &Goxios{}
	if len(configs) > 0 {
		g.SetConfig(configs[0])
	}
	return g
}

type Goxios struct {
	config GoxiosConfig
}

func (g *Goxios) SetConfig(config GoxiosConfig) *Goxios {
	g.config = config
	return g
}
func (g *Goxios) newRequest(url, method string, config RequestConfig) (*Response, error) {
	uri := NewURL(url, config.Params)
	if err := uri.Error(); err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, uri.String(), config.Body)
	if err != nil {
		return nil, err
	}
	header := make(http.Header)
	if g.config.Header != nil {
		for k, v := range g.config.Header {
			for i := range v {
				header.Add(k, v[i])
			}
		}
	}
	if config.Header != nil {
		for k, v := range config.Header {
			for i := range v {
				header.Add(k, v[i])
			}
		}
	}
	req.Header = header
	// 写入cookie
	var cookies []*http.Cookie
	if len(g.config.Cookies) > 0 {
		cookies = append(cookies, g.config.Cookies...)
	}
	if len(config.Cookies) > 0 {
		cookies = append(cookies, config.Cookies...)
	}
	for i := range cookies {
		req.AddCookie(cookies[i])
	}
	client := &http.Client{}
	// 是否重定向
	var redirect Redirect
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if val := len(via); val > 0 {
			redirect.redirectCount = val
			if val == 1 {
				redirect.isRedirect = true
				if req.Response != nil {
					redirect.code = req.Response.StatusCode
				}
			}

		}
		// 禁止重定向
		if g.config.ClearRedirect || config.ClearRedirect {
			return http.ErrUseLastResponse
		}
		// 超过10次后
		if len(via) >= 10 {
			return errors.New("stopped after 10 redirects")
		}
		return nil
	}
	// socks5
	if g.config.Transport != nil {
		client.Transport = g.config.Transport
	}
	// timeout
	if config.Timeout > 0 {
		client.Timeout = config.Timeout
	} else if g.config.Timeout > 0 {
		client.Timeout = g.config.Timeout
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		if len(resp.Cookies()) > 0 {
			g.config.AddCookies(resp.Cookies())
		}
		if resp.Header != nil {
			g.config.AddHeader(resp.Header)
		}
	}
	return newResponse(resp, redirect)
}
func (g *Goxios) getRequestConfig(configs ...RequestConfig) RequestConfig {
	if len(configs) > 0 {
		return configs[0]
	}
	return RequestConfig{}
}
func (g *Goxios) GET(url string, configs ...RequestConfig) *Result {
	method := GET_
	config := g.getRequestConfig(configs...)
	return newResult(g.newRequest(url, method.String(), config))
}
func (g *Goxios) POST(url string, configs ...RequestConfig) *Result {
	method := POST_
	config := g.getRequestConfig(configs...)
	return newResult(g.newRequest(url, method.String(), config))
}
func (g *Goxios) PUT(url string, configs ...RequestConfig) *Result {
	method := PUT_
	config := g.getRequestConfig(configs...)
	return newResult(g.newRequest(url, method.String(), config))
}
func (g *Goxios) PATCH(url string, configs ...RequestConfig) *Result {
	method := PATCH_
	config := g.getRequestConfig(configs...)
	return newResult(g.newRequest(url, method.String(), config))
}
func (g *Goxios) DELETE(url string, configs ...RequestConfig) *Result {
	method := DELETE_
	config := g.getRequestConfig(configs...)
	return newResult(g.newRequest(url, method.String(), config))
}

func GET(url string, configs ...RequestConfig) *Result {
	return Do().GET(url, configs...)
}
func POST(url string, configs ...RequestConfig) *Result {
	return Do().POST(url, configs...)
}
func PUT(url string, configs ...RequestConfig) *Result {
	return Do().PUT(url, configs...)
}
func PATCH(url string, configs ...RequestConfig) *Result {
	return Do().PATCH(url, configs...)
}
func DELETE(url string, configs ...RequestConfig) *Result {
	return Do().DELETE(url, configs...)
}

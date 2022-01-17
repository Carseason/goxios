package goxios

import (
	"net/http"
)

type Request struct {
	request  *http.Request
	response Response
	err      error
}
type Finally struct {
	Response
	Error error
}

//success
func (g *Request) Then(f func(responseData Response)) *Request {
	if g.err == nil {
		f(g.response)
	}
	return g
}

//failing
func (g *Request) Catch(f func(responseError error)) *Request {
	if g.err != nil {
		f(g.err)
	}
	return g
}

// result
func (g *Request) Finally(f ...func(Response, error)) (Response, error) {
	for i := range f {
		f[i](g.response, g.err)
	}
	return g.response, g.err
}

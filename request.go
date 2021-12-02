package goxios

import (
	"net/http"
)

type Request struct {
	request  *http.Request
	response Response
	err      error
}

//success
func (g *Request) Then(f func(responseData Response)) *Request {
	f(g.response)
	return g
}

//failing
func (g *Request) Catch(f func(responseError error)) *Request {
	f(g.err)
	return g
}
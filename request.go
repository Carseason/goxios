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
	f(g.response)
	return g
}

//failing
func (g *Request) Catch(f func(responseError error)) *Request {
	f(g.err)
	return g
}

// result
func (g *Request) Finally() Finally {
	return Finally{
		Response: g.response,
		Error:    g.err,
	}
}

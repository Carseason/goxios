package goxios

type Result struct {
	Response *Response
	Error    error
}

func newResult(resp *Response, err error) *Result {
	return &Result{
		Response: resp,
		Error:    err,
	}
}

// success
func (g Result) Then(f func(resp Response)) Result {
	if g.Error == nil {
		f(*g.Response)
	}
	return g
}

// failing
func (g Result) Catch(f func(respError error)) Result {
	if g.Error != nil {
		f(g.Error)
	}
	return g
}

// result
func (g Result) Finally(fs ...func(Response, error)) (Response, error) {
	for i := range fs {
		fs[i](*g.Response, g.Error)
	}
	return *g.Response, nil
}

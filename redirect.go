package goxios

type Redirect struct {
	code          int
	isRedirect    bool
	redirectCount int //重定向次数
}

func (r Redirect) StatusCode() int {
	return r.code
}
func (r Redirect) IsRedirect() bool {
	return r.isRedirect
}
func isRedirect(code int) bool {
	switch code {
	case 301: // (Moved Permanently)
		return true
	case 302: // (Found)
		return true
	case 303: // (See Other)
		return true
	case 307: // (Temporary Redirect)
		return true
	case 308: // (Permanent Redirect)
		return true
	}
	return false
}

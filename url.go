package goxios

import "net/url"

type URL struct {
	host   string
	params url.Values
	err    error
}

func (u *URL) String() string {
	uri := u.host
	params := u.params
	parse, err := url.Parse(uri)
	if err != nil {
		u.err = err
		return uri
	}
	values, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		u.err = err
		return uri
	}
	for k, v := range params {
		values[k] = append(values[k], v...)
	}
	parse.RawQuery = values.Encode()
	return parse.String()
}
func (u *URL) Error() error {
	return u.err
}

func NewURL(uri string, params url.Values) *URL {
	return &URL{
		host:   uri,
		params: params,
	}
}

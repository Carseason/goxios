package goxios

import (
	"net/http"

	"golang.org/x/net/proxy"
)

// "tcp" ":3908" , "user" ,"password"
func SOCKS5(network, addr, user, password string) (*http.Transport, error) {
	dialer, err := proxy.SOCKS5(network, addr, &proxy.Auth{
		User:     user,
		Password: password,
	}, proxy.Direct)
	if err != nil {
		return nil, err
	}
	tr := &http.Transport{Dial: dialer.Dial}
	return tr, nil
}

package goxios

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestDo(t *testing.T) {
	config := NewGoxiosConfig().AddHeader(NewHeader(map[string]string{
		"user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
	}))
	Do(*config).GET("https://www.baidu.com/").
		Catch(func(respError error) {
			fmt.Println(respError)
		}).
		Then(func(resp *Response) {
			// fmt.Println(resp.StatusCode(), resp.IsRedirect(), resp.Title())
			fmt.Println(resp.Title())
			value, err := resp.QueryText(".weather_p > .con")
			fmt.Println(strings.TrimSpace(value), err)
			value, err = resp.QueryAttr(`input[name="formhash"]`, "value")
			fmt.Println(strings.TrimSpace(value), err)
			// fmt.Println(resp.Content())
		})
	t.Fail()
}

func TestGoxios(t *testing.T) {
	header := make(http.Header)
	request := Do(GoxiosConfig{
		Header: header,
	})
	var resp *Result
	var url = "https://www.baidu.com/"
	switch strings.ToUpper("GET") {
	case string(POST_):
		var body io.Reader = nil
		resp = request.POST(url, RequestConfig{
			Body: body,
		})
	case string(PUT_):
		var body io.Reader = nil
		resp = request.PUT(url, RequestConfig{
			Body: body,
		})
	case string(PATCH_):
		var body io.Reader = nil

		resp = request.PATCH(url, RequestConfig{
			Body: body,
		})
	case string(DELETE_):
		resp = request.DELETE(url)
	default:
		resp = request.GET(url)
	}
	res, err := resp.Finally()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(res.Content())
	t.Fail()
}

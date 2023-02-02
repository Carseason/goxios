package goxios

import (
	"fmt"
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
		Then(func(resp Response) {
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

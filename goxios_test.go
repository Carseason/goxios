package goxios

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	Do(RequestConfig{
		Method: GET,
		Url:    "https://www.baidu.com",
		Params: NewParams(Params{
			"a": 1,
			"b": 2,
		}),
		Header: NewHeader(Header{
			UserAgentKey: UserAgentValue,
		}),
	}).Then(func(responseData Response) {
		fmt.Println(responseData.Title())
	}).Catch(func(responseError error) {
		t.Error(responseError)
	})
	t.Fail()
}

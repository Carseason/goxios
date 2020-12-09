package goxios

import "testing"

func TestGet(t *testing.T) {
	Get(Config{
		Url: "https://www.baidu.com",
		Headers: Headers{
			UserAgent: UserAgent,
			Cookie:    "",
		},
	}).Then(func(res Goxios) {
		t.Log(string(res.Body()))

	}).Catch(func(err error) {
		t.Error(err)
	})
	t.Fail()

}

func TestPost(t *testing.T) {
	Post(Config{
		Url: "https://www.baidu.com",
		Headers: Headers{
			UserAgent:   UserAgent,
			Cookie:      "",
			ContentType: ContentTypeJSON,
		},
		Data: nil,
	}).Then(func(res Goxios) {
		t.Log(res.StatusCode())

	}).Catch(func(err error) {
		t.Error(err)
	})
	t.Fail()

}
func TestPut(t *testing.T) {
	Put(Config{
		Url: "https://www.baidu.com",
		Headers: Headers{
			UserAgent:   UserAgent,
			Cookie:      "",
			ContentType: ContentTypeJSON,
		},
		Data: nil,
	}).Then(func(res Goxios) {
		t.Log(res.StatusCode())
	}).Catch(func(err error) {
		t.Error(err)
	})
	t.Fail()
}
func TestDelete(t *testing.T) {
	Delete(Config{
		Url: "https://www.baidu.com",
	}).Then(func(res Goxios) {
		t.Log(res.StatusCode())
	}).Catch(func(err error) {
		t.Error(err)
	})
	t.Fail()
}

package goxios

import (
	"bytes"
	"errors"
	"io"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	return io.ReadAll(reader)
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	return io.ReadAll(reader)
}

// 将对应格式文本转换成utf-8
func DecodeEncoding(body []byte) ([]byte, error) {
	determineEncoding, _, _ := charset.DetermineEncoding(body, "")
	if determineEncoding == nil {
		return nil, errors.New("determineEncoding.NewDecoder is empty")
	}
	return io.ReadAll(transform.NewReader(bytes.NewBuffer(body), determineEncoding.NewDecoder()))
}

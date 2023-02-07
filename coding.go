package goxios

import (
	"bytes"
	"errors"
	"io"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/charmap"
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
	determineEncoding, name, _ := charset.DetermineEncoding(body, "")
	if determineEncoding == nil {
		return nil, errors.New("determineEncoding.NewDecoder is empty")
	}
	var err error
	switch name {
	case "gbk":
		// body, _, err = transform.Bytes(determineEncoding.NewDecoder(), body)
		// return body, err
		return io.ReadAll(transform.NewReader(bytes.NewBuffer(body), determineEncoding.NewDecoder()))
	case "windows-1252":
		// body, _, err = transform.Bytes(charmap.Windows1252.NewEncoder(), body)
		// return body, err
		return io.ReadAll(transform.NewReader(bytes.NewBuffer(body), charmap.Windows1252.NewEncoder()))
	case "utf-8":
	}
	return body, err

}

package goxios

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	urlpkg "net/url"
	"strings"
)

type Reader io.Reader

func NewGoxiosReader(data interface{}) (Reader, error) {
	switch value := data.(type) {
	case urlpkg.Values: //表单
		return strings.NewReader(value.Encode()), nil
	case string: //json
		return strings.NewReader(value), nil
	case map[string]interface{}:
		body := make(urlpkg.Values)
		for k, v := range value {
			body.Set(k, fmt.Sprintf("%v", v))
		}
		return strings.NewReader(body.Encode()), nil
	default:
		if bs, err := json.Marshal(value); err == nil {
			return bytes.NewBuffer(bs), nil
		}
	}
	return nil, nil
}
func NewGoxiosFileReader(file *multipart.FileHeader, fieldname string, data map[string]string) (Reader, error) {
	files, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer files.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	defer writer.Close()
	for v, k := range data {
		writer.WriteField(v, k)
	}
	part, err := writer.CreateFormFile(fieldname, file.Filename)
	if err != nil {
		return nil, err
	}
	io.Copy(part, files)
	return body, nil
}
func NewGoxiosFormReader(values urlpkg.Values) Reader {
	return strings.NewReader(values.Encode())
}
func NewGoxiosStringToReader(value string) Reader {
	return strings.NewReader(value)
}

package utils

import (
	"bytes"
	"errors"
	"fmt"
	"goSMSBoom/log"
	"io"
	"net/http"
	"reflect"
)

type HttpDos struct {
	url     string
	body    []byte
	headers map[string]string
}

func NewHttpDos(url string, body []byte, headers map[string]string) *HttpDos {
	return &HttpDos{url: url, body: body, headers: headers}
}

// Get get请求
func (h *HttpDos) Get() ([]byte, error) {
	// 创建一个 GET 请求
	req, err := http.NewRequest("GET", h.url, nil)
	if err != nil {
		return nil, err
	}

	// 添加自定义请求头11
	for key, value := range h.headers {
		req.Header.Set(key, value)
	}

	//  发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应体0.0
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// MontageURL 拼接url
func MontageURL(oldUrl string, data interface{}) (string, error) {
	var suf = oldUrl + "?"
	// 使用反射遍历 data
	v := reflect.ValueOf(data).Elem()
	if v.Kind() != reflect.Struct {
		log.Println("data is not a struct")
		return "", errors.New("data is not a struct")
	}
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		fieldName := v.Type().Field(i).Name
		tag := v.Type().Field(i).Tag.Get("json")
		if tag != "" {
			suf += fmt.Sprintf("%v=%v&", tag, f.Interface())
		} else {
			suf += fmt.Sprintf("%v=%v&", fieldName, f.Interface())
		}
	}

	return suf[:len(suf)-1], nil // 去除最后一个多余的 "&" 符号
}

// Post 发送 HTTP POST 请求并返回响应体和可能的错误
func (h *HttpDos) Post() ([]byte, error) {
	// 创建一个请求
	req, err := http.NewRequest("POST", h.url, bytes.NewBuffer(h.body))
	if err != nil {
		return nil, err
	}

	// 添加自定义请求头
	for key, value := range h.headers {
		req.Header.Set(key, value)
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应体
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func (h *HttpDos) GetUrl() string {
	return h.url
}

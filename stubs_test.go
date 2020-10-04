package zaleycash_sdk

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

func BuildStubConfig() *Config {
	cfg := &Config{
		Uri:       "https://github.com",
		PublicKey: "PublicKey",
		SecretKey: "SecretKey",
	}
	return cfg
}

func LoadStubResponseData(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func BuildStubResponseFromString(statusCode int, json string) *http.Response {
	body := ioutil.NopCloser(strings.NewReader(json))
	return &http.Response{Body: body, StatusCode: statusCode}
}

func BuildStubResponseFromFile(statusCode int, path string) *http.Response {
	data, _ := LoadStubResponseData(path)
	body := ioutil.NopCloser(bytes.NewReader(data))
	return &http.Response{Body: body, StatusCode: statusCode}
}
package zaleycash

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

func BuildStubConfig() *Config {
	return &Config{
		Uri:       "https://github.com",
		PublicKey: "PublicKey",
		SecretKey: "SecretKey",
	}
}

func BuildStubToken() *Token {
	return &Token{
		AccessToken: "AccessToken",
		ExpiresAt:   100,
	}
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

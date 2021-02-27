package zaleycash

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func Test_Auth_TokenIsNotExpiredSuccess(t *testing.T) {
	expired := time.Now().Unix() + 1000
	token := Token{ExpiresAt: int(expired)}
	assert.True(t, token.IsNotExpired())
}

func Test_Auth_TokenIsNotExpiredFail(t *testing.T) {
	expired := time.Now().Unix() - 1000
	token := Token{ExpiresAt: int(expired)}
	assert.False(t, token.IsNotExpired())
}

func Test_Auth_TokenIsValidSuccess(t *testing.T) {
	expired := time.Now().Unix() + 1000
	token := Token{ExpiresAt: int(expired), AccessToken: "qwerty"}
	assert.True(t, token.IsValid())
}

func Test_Auth_TokenIsValidFailExpired(t *testing.T) {
	expired := time.Now().Unix() - 1000
	token := Token{ExpiresAt: int(expired), AccessToken: "qwerty"}
	assert.False(t, token.IsValid())
}

func Test_Auth_TokenIsValidFailEmpty(t *testing.T) {
	expired := time.Now().Unix() + 1000
	token := Token{ExpiresAt: int(expired)}
	assert.False(t, token.IsValid())
}

func Test_Auth_NewAuthFromConfig(t *testing.T) {
	auth := NewAuthFromConfig(BuildStubConfig(), nil)
	assert.NotEmpty(t, auth)
}

func Test_Auth_NewAuthFromCredentials(t *testing.T) {
	auth := NewAuthFromCredentials("secret-key", "public-key", nil)
	assert.NotEmpty(t, auth)
}

func Test_Auth_HandleTokenStructSuccess(t *testing.T) {
	rsp := BuildStubResponseFromFile(http.StatusOK, "stubs/data/auth/token.success.json")
	auth := NewAuthFromConfig(BuildStubConfig(), nil)
	token, _ := auth.HandleTokenStruct(&Response{raw: rsp})
	assert.Equal(t, "foo", token.AccessToken)
	assert.Equal(t, 1601644827, token.ExpiresAt)
}

func Test_Auth_HandleTokenStructBadRequest(t *testing.T) {
	rsp := BuildStubResponseFromFile(http.StatusBadRequest, "stubs/data/auth/token.invalid.json")
	auth := NewAuthFromConfig(BuildStubConfig(), nil)
	_, err := auth.HandleTokenStruct(&Response{raw: rsp})
	assert.Error(t, err)
	assert.Equal(t, "Auth@GetTokenStruct Unknown access token", err.Error())
}

func Test_Auth_HandleTokenStructUnmarshalError(t *testing.T) {
	rsp := BuildStubResponseFromFile(http.StatusOK, "stubs/data/auth/token.invalid.json")
	auth := NewAuthFromConfig(BuildStubConfig(), nil)
	_, err := auth.HandleTokenStruct(&Response{raw: rsp})
	assert.Error(t, err)
	assert.Equal(t, "Auth@GetTokenStruct unmarshal token: ResultPayloadType@Unmarshal parse json error: unexpected end of JSON input", err.Error())
}

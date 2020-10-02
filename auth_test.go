package zaleycash_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_ZaleyCash_SDK_Auth_TokenIsNotExpiredSuccess(t *testing.T) {
	expired := time.Now().Unix() + 1000
	token := Token{ExpiresAt: int(expired)}
	assert.True(t, token.IsNotExpired())
}

func Test_ZaleyCash_SDK_Auth_TokenIsNotExpiredFail(t *testing.T) {
	expired := time.Now().Unix() - 1000
	token := Token{ExpiresAt: int(expired)}
	assert.False(t, token.IsNotExpired())
}

func Test_ZaleyCash_SDK_Auth_TokenIsValidSuccess(t *testing.T) {
	expired := time.Now().Unix() + 1000
	token := Token{ExpiresAt: int(expired), AccessToken: "qwerty"}
	assert.True(t, token.IsValid())
}

func Test_ZaleyCash_SDK_Auth_TokenIsValidFailExpired(t *testing.T) {
	expired := time.Now().Unix() - 1000
	token := Token{ExpiresAt: int(expired), AccessToken: "qwerty"}
	assert.False(t, token.IsValid())
}

func Test_ZaleyCash_SDK_Auth_TokenIsValidFailEmpty(t *testing.T) {
	expired := time.Now().Unix() + 1000
	token := Token{ExpiresAt: int(expired)}
	assert.False(t, token.IsValid())
}

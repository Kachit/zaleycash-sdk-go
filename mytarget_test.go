package zaleycash_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Resource_MyTarget_TokenIsNotExpiredSuccess(t *testing.T) {
	expired := time.Now().Unix() + 1000
	token := MyTargetToken{ExpiresIn: float64(expired)}
	assert.True(t, token.IsNotExpired())
}

func Test_Resource_MyTarget_TokenIsNotExpiredFail(t *testing.T) {
	expired := time.Now().Unix() - 1000
	token := MyTargetToken{ExpiresIn: float64(expired)}
	assert.False(t, token.IsNotExpired())
}

func Test_Resource_MyTarget_TokenIsValidSuccess(t *testing.T) {
	expired := time.Now().Unix() + 1000
	token := MyTargetToken{ExpiresIn: float64(expired), AccessToken: "qwerty"}
	assert.True(t, token.IsValid())
}

func Test_Resource_MyTarget_TokenIsValidFailExpired(t *testing.T) {
	expired := time.Now().Unix() - 1000
	token := MyTargetToken{ExpiresIn: float64(expired), AccessToken: "qwerty"}
	assert.False(t, token.IsValid())
}

func Test_Resource_MyTarget__TokenIsValidFailEmpty(t *testing.T) {
	expired := time.Now().Unix() + 1000
	token := MyTargetToken{ExpiresIn: float64(expired)}
	assert.False(t, token.IsValid())
}

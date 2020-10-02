package zaleycash_sdk

import (
	"net/http"
	"time"
)

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int    `json:"expires_at"`
}

func (t *Token) IsValid() bool {
	return t.IsNotExpired() && t.AccessToken != ""
}

func (t *Token) IsNotExpired() bool {
	ts := time.Now().Unix()
	return int64(t.ExpiresAt) > ts
}

type Auth struct {
	*ResourceAbstract
}

func NewAuthFromConfig(config *Config, cl *http.Client) *Auth {
	if cl == nil {
		cl = &http.Client{}
	}
	transport := NewHttpTransport(config, cl)
	return &Auth{ResourceAbstract: NewResourceAbstract(transport)}
}

func NewAuthFromCredentials(secretKey string, publicKey string, cl *http.Client) *Auth {
	if cl == nil {
		cl = &http.Client{}
	}
	config := NewConfig(secretKey, publicKey)
	transport := NewHttpTransport(config, cl)
	return &Auth{ResourceAbstract: NewResourceAbstract(transport)}
}

func (a *Auth) GetToken() (*Response, error) {
	return a.Post("api/v2/token", nil, nil)
}

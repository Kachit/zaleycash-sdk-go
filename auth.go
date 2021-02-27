package zaleycash

import (
	"errors"
	"fmt"
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

/**
 * @unmarshal Token
 */
func (a *Auth) GetToken() (*Response, error) {
	return a.Post("api/v2/token", nil, nil)
}

func (a *Auth) GetTokenStruct() (*Token, error) {
	response, err := a.GetToken()
	if err != nil {
		return nil, fmt.Errorf("Auth@GetTokenStruct token request error: %v", err)
	}
	return a.HandleTokenStruct(response)
}

func (a *Auth) HandleTokenStruct(response *Response) (*Token, error) {
	if !response.IsSuccess() {
		respError, err := response.GetError()
		if err != nil {
			return nil, fmt.Errorf("Auth@GetTokenStruct parse error result: %v", err)
		}
		return nil, errors.New("Auth@GetTokenStruct " + respError.Message)
	}
	var token Token
	err := response.Unmarshal(&token)
	if err != nil {
		return nil, fmt.Errorf("Auth@GetTokenStruct unmarshal token: %v", err)
	}
	return &token, nil
}

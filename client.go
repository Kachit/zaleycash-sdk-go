package zaleycash_sdk

import "net/http"

type Client struct {
	transport *Transport
}

func NewClientFromConfig(config *Config, token *Token, cl *http.Client) *Client {
	if cl == nil {
		cl = &http.Client{}
	}
	transport := NewHttpTransportAuthenticated(config, token, cl)
	return &Client{transport}
}

func (c *Client) MyTarget() *MyTarget {
	return &MyTarget{ResourceAbstract: NewResourceAbstract(c.transport)}
}

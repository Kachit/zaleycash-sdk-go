package zaleycash

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

func (c *Client) MyTarget() *MyTargetResource {
	return &MyTargetResource{ResourceAbstract: NewResourceAbstract(c.transport)}
}

func (c *Client) Users() *UsersResource {
	return &UsersResource{ResourceAbstract: NewResourceAbstract(c.transport)}
}

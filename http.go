package zaleycash

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type RequestBuilderInterface interface {
	buildUri(path string, query map[string]interface{}) (uri *url.URL, err error)
	buildHeaders() http.Header
	buildBody(data map[string]interface{}) (io.Reader, error)
	isValidToken() bool
}

type RequestBuilder struct {
	cfg *Config
}

func (rb *RequestBuilder) isValidToken() bool {
	return true
}

func (rb *RequestBuilder) buildUri(path string, query map[string]interface{}) (uri *url.URL, err error) {
	u, err := url.Parse(rb.cfg.Uri)
	if err != nil {
		return nil, fmt.Errorf("RequestBuilder@buildUri parse: %v", err)
	}
	u.Path = "/" + path
	if query != nil {
		q := u.Query()
		for k, v := range query {
			q.Set(k, fmt.Sprintf("%v", v))
		}
		u.RawQuery = q.Encode()
	}
	return u, err
}

func (rb *RequestBuilder) buildHeaders() http.Header {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+rb.cfg.SecretKey)
	return headers
}

func (rb *RequestBuilder) buildBody(data map[string]interface{}) (io.Reader, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("RequestBuilder@buildBody json convert: %v", err)
	}
	return bytes.NewBuffer(b), nil
}

type RequestBuilderAuthenticated struct {
	*RequestBuilder
	token *Token
}

func (rb *RequestBuilderAuthenticated) isValidToken() bool {
	return rb.token.IsValid()
}

func (rb *RequestBuilderAuthenticated) buildHeaders() http.Header {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+rb.token.AccessToken)
	return headers
}

func NewHttpTransport(config *Config, h *http.Client) *Transport {
	if h == nil {
		h = &http.Client{}
	}
	rb := &RequestBuilder{cfg: config}
	return &Transport{http: h, rb: rb}
}

func NewHttpTransportAuthenticated(config *Config, token *Token, h *http.Client) *Transport {
	if h == nil {
		h = &http.Client{}
	}
	rb := &RequestBuilderAuthenticated{RequestBuilder: &RequestBuilder{cfg: config}, token: token}
	return &Transport{http: h, rb: rb}
}

type Transport struct {
	http *http.Client
	rb   RequestBuilderInterface
}

func (t *Transport) Request(method string, path string, query map[string]interface{}, body map[string]interface{}) (resp *http.Response, err error) {
	if !t.rb.isValidToken() {
		return nil, fmt.Errorf("transport@request invalid token: %v", err)
	}
	//build uri
	uri, err := t.rb.buildUri(path, query)
	if err != nil {
		return nil, fmt.Errorf("transport@request build uri: %v", err)
	}
	//build body
	bodyReader, err := t.rb.buildBody(body)
	if err != nil {
		return nil, fmt.Errorf("transport@request build request body: %v", err)
	}
	//build request
	req, err := http.NewRequest(strings.ToUpper(method), uri.String(), bodyReader)
	if err != nil {
		return nil, fmt.Errorf("transport@request new request error: %v", err)
	}
	//build headers
	req.Header = t.rb.buildHeaders()
	return t.http.Do(req)
}

func (t *Transport) Get(path string, query map[string]interface{}) (resp *http.Response, err error) {
	return t.Request("GET", path, query, nil)
}

func (t *Transport) Post(path string, body map[string]interface{}, query map[string]interface{}) (resp *http.Response, err error) {
	return t.Request("POST", path, query, body)
}

type Response struct {
	raw    *http.Response
	result *Result
}

func (r *Response) IsSuccess() bool {
	return r.raw.StatusCode < http.StatusMultipleChoices
}

func (r *Response) GetRawResponse() *http.Response {
	return r.raw
}

func (r *Response) GetRawBody() string {
	body, _ := r.ReadBody()
	return string(body)
}

func (r *Response) Unmarshal(v interface{}) error {
	data, err := r.ReadBody()
	if err != nil {
		return fmt.Errorf("Response@Unmarshal read body: %v", err)
	}
	return r.UnmarshalData(data, &v)
}

func (r *Response) GetError() (*ErrorResult, error) {
	var result ErrorResult
	err := r.UnmarshalError(&result)
	if err != nil {
		return nil, fmt.Errorf("Response@GetError Unmarshal: %v", err)
	}
	return &result, nil
}

func (r *Response) UnmarshalError(errorResult *ErrorResult) error {
	data, err := r.ReadBody()
	if err != nil {
		return fmt.Errorf("Response@UnmarshalError read body: %v", err)
	}

	err = json.Unmarshal(data, &errorResult)
	if err != nil {
		return fmt.Errorf("Response@UnmarshalError Unmarshal: %v", err)
	}
	return nil
}

func (r *Response) UnmarshalData(data []byte, v interface{}) error {
	err := json.Unmarshal(data, &r.result)
	if err != nil {
		return fmt.Errorf("Response@Unmarshal parse json error: %v", err)
	}
	err = r.result.Payload.Unmarshal(&v)
	if err != nil {
		return err
	}
	return nil
}

func (r *Response) ReadBody() ([]byte, error) {
	defer r.raw.Body.Close()
	return ioutil.ReadAll(r.raw.Body)
}

func NewResponse(raw *http.Response) *Response {
	return &Response{raw: raw}
}

type Result struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Payload ResultPayloadType `json:"response"`
}

type ResultPayloadType struct {
	Payload string
}

func (rp *ResultPayloadType) UnmarshalJSON(data []byte) error {
	rp.Payload = string(data)
	return nil
}

func (rp *ResultPayloadType) Unmarshal(v interface{}) error {
	err := json.Unmarshal([]byte(rp.Payload), &v)
	if err != nil {
		return fmt.Errorf("ResultPayloadType@Unmarshal parse json error: %v", err)
	}
	return nil
}

type ErrorResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

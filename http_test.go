package zaleycash_sdk

import (
	"encoding/json"
	"github.com/jarcoal/httpmock"
	"net/http"
	"time"

	//"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	//"net/http"
	"testing"
)

func Test_HTTP_RequestBuilder_BuildUriWithoutQueryParams(t *testing.T) {
	cfg := BuildStubConfig()
	builder := RequestBuilder{cfg: cfg}
	uri, err := builder.buildUri("qwerty", nil)
	assert.NotEmpty(t, uri)
	assert.Equal(t, "https://github.com/qwerty", uri.String())
	assert.Nil(t, err)
}

func Test_HTTP_RequestBuilder_BuildUriWithQueryParams(t *testing.T) {
	cfg := BuildStubConfig()
	builder := RequestBuilder{cfg: cfg}

	data := make(map[string]interface{})
	data["foo"] = "bar"
	data["bar"] = "baz"

	uri, err := builder.buildUri("qwerty", data)
	assert.NotEmpty(t, uri)
	assert.Equal(t, "https://github.com/qwerty?bar=baz&foo=bar", uri.String())
	assert.Nil(t, err)
}

func Test_HTTP_RequestBuilder_BuildHeaders(t *testing.T) {
	cfg := BuildStubConfig()
	builder := RequestBuilder{cfg: cfg}

	headers := builder.buildHeaders()
	assert.NotEmpty(t, headers)
	assert.Equal(t, "application/json", headers.Get("Content-Type"))
	assert.Equal(t, "Bearer "+cfg.SecretKey, headers.Get("Authorization"))
}

func Test_HTTP_RequestBuilder_BuildBody(t *testing.T) {
	cfg := BuildStubConfig()
	builder := RequestBuilder{cfg: cfg}

	data := make(map[string]interface{})
	data["foo"] = "bar"
	data["bar"] = "baz"

	body, _ := builder.buildBody(data)
	assert.NotEmpty(t, body)
}

func Test_HTTP_RequestBuilder_IsValidToken(t *testing.T) {
	cfg := BuildStubConfig()
	builder := RequestBuilder{cfg: cfg}
	assert.True(t, builder.isValidToken())
}

func Test_HTTP_RequestBuilderAuthenticated_IsValidTokenSuccess(t *testing.T) {
	cfg := BuildStubConfig()
	token := BuildStubToken()
	token.ExpiresAt = int(time.Now().Unix() + 1000)
	builder := RequestBuilderAuthenticated{RequestBuilder: &RequestBuilder{cfg: cfg}, token: token}
	assert.True(t, builder.isValidToken())
}

func Test_HTTP_RequestBuilderAuthenticated_IsValidTokenExpired(t *testing.T) {
	cfg := BuildStubConfig()
	token := BuildStubToken()
	token.ExpiresAt = int(time.Now().Unix() - 1000)
	builder := RequestBuilderAuthenticated{RequestBuilder: &RequestBuilder{cfg: cfg}, token: token}
	assert.False(t, builder.isValidToken())
}

func Test_HTTP_RequestBuilderAuthenticated_IsValidTokenEmpty(t *testing.T) {
	cfg := BuildStubConfig()
	token := BuildStubToken()
	token.AccessToken = ""
	builder := RequestBuilderAuthenticated{RequestBuilder: &RequestBuilder{cfg: cfg}, token: token}
	assert.False(t, builder.isValidToken())
}

func Test_HTTP_NewHttpTransport(t *testing.T) {
	cfg := BuildStubConfig()
	transport := NewHttpTransport(cfg, nil)
	assert.NotEmpty(t, transport)
}

func Test_HTTP_Transport_RequestSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := NewHttpTransport(cfg, nil)

	body, _ := LoadStubResponseData("stubs/data/auth/token.success.json")

	httpmock.RegisterResponder("GET", cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, _ := transport.Request("GET", "foo", nil, nil)
	assert.NotEmpty(t, resp)
}

func Test_HTTP_Transport_RequestGET(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := NewHttpTransport(cfg, nil)

	body, _ := LoadStubResponseData("stubs/data/auth/token.success.json")

	httpmock.RegisterResponder("GET", cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, _ := transport.Get("foo", nil)
	assert.NotEmpty(t, resp)
}

func Test_HTTP_Transport_RequestPOST(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := NewHttpTransport(cfg, nil)

	body, _ := LoadStubResponseData("stubs/data/auth/token.success.json")

	httpmock.RegisterResponder("POST", cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, _ := transport.Post("foo", nil, nil)
	assert.NotEmpty(t, resp)
}

func Test_HTTP_Response_IsSuccessTrue(t *testing.T) {
	response := &Response{raw: BuildStubResponseFromFile(http.StatusOK, "stubs/data/auth/token.success.json")}
	assert.True(t, response.IsSuccess())
}

func Test_HTTP_Response_IsSuccessFalse(t *testing.T) {
	response := &Response{raw: BuildStubResponseFromFile(http.StatusBadRequest, "stubs/data/auth/token.success.json")}
	assert.False(t, response.IsSuccess())
}

func Test_HTTP_Response_GetRawResponse(t *testing.T) {
	rsp := BuildStubResponseFromFile(http.StatusOK, "stubs/data/auth/token.success.json")
	response := &Response{raw: rsp}
	raw := response.GetRawResponse()
	assert.NotEmpty(t, raw)
	assert.Equal(t, http.StatusOK, raw.StatusCode)
}

func Test_HTTP_Response_Unmarshal(t *testing.T) {
	rsp := BuildStubResponseFromFile(http.StatusBadRequest, "stubs/data/auth/token.success.json")
	response := &Response{raw: rsp}
	var result Token
	_ = response.Unmarshal(&result)
	assert.Equal(t, "foo", result.AccessToken)
	assert.Equal(t, 1601644827, result.ExpiresAt)
}

func Test_HTTP_Response_GetError(t *testing.T) {
	rsp := BuildStubResponseFromFile(http.StatusBadRequest, "stubs/data/auth/token.invalid.json")
	response := &Response{raw: rsp}
	result, _ := response.GetError()
	assert.Equal(t, "Unknown access token", result.Message)
	assert.Equal(t, 1, result.Code)
}

func Test_HTTP_Response_UnmarshalError(t *testing.T) {
	rsp := BuildStubResponseFromFile(http.StatusBadRequest, "stubs/data/auth/token.invalid.json")
	response := &Response{raw: rsp}
	var result ErrorResult
	_ = response.UnmarshalError(&result)
	assert.Equal(t, "Unknown access token", result.Message)
	assert.Equal(t, 1, result.Code)
}

func Test_HTTP_Response_GetRawBody(t *testing.T) {
	data, _ := LoadStubResponseData("stubs/data/auth/token.success.json")
	rsp := BuildStubResponseFromFile(http.StatusBadRequest, "stubs/data/auth/token.success.json")
	response := &Response{raw: rsp}
	assert.Equal(t, string(data), response.GetRawBody())
}

func Test_HTTP_NewResponse(t *testing.T) {
	rsp := BuildStubResponseFromFile(http.StatusOK, "stubs/data/auth/token.success.json")
	response := NewResponse(rsp)
	assert.NotEmpty(t, response)
}

func Test_HTTP_ResultPayloadType_UnmarshalJSON(t *testing.T) {
	data, _ := json.Marshal(BuildStubToken())
	payload := &ResultPayloadType{}
	_ = payload.UnmarshalJSON(data)
	assert.Equal(t, `{"access_token":"AccessToken","expires_at":100}`, payload.Payload)
}

func Test_HTTP_ResultPayloadType_Unmarshal(t *testing.T) {
	token := &Token{}
	payload := &ResultPayloadType{}
	payload.Payload = `{"access_token":"AccessToken","expires_at":100}`
	_ = payload.Unmarshal(&token)
	assert.Equal(t, "AccessToken", token.AccessToken)
	assert.Equal(t, 100, token.ExpiresAt)
}

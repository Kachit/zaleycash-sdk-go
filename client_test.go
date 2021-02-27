package zaleycash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Client_NewClientFromConfig(t *testing.T) {
	client := NewClientFromConfig(BuildStubConfig(), nil, nil)
	assert.NotEmpty(t, client)
}

func Test_Client_GetMyTargetResource(t *testing.T) {
	client := NewClientFromConfig(BuildStubConfig(), nil, nil)
	result := client.MyTarget()
	assert.NotEmpty(t, result)
}

func Test_Client_GetUsersResource(t *testing.T) {
	client := NewClientFromConfig(BuildStubConfig(), nil, nil)
	result := client.Users()
	assert.NotEmpty(t, result)
}

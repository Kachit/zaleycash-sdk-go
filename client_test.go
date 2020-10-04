package zaleycash_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Client_NewClientFromConfig(t *testing.T) {
	client := NewClientFromConfig(BuildStubConfig(), nil, nil)
	assert.NotEmpty(t, client)
}

func Test_Client_GetMyTargetAPI(t *testing.T) {
	client := NewClientFromConfig(BuildStubConfig(), nil, nil)
	result := client.MyTarget()
	assert.NotEmpty(t, result)
}

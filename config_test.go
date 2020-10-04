package zaleycash_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Config_NewConfig(t *testing.T) {
	cfg := NewConfig("secret-key", "public-key")
	assert.NotEmpty(t, cfg)
}

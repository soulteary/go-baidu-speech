package oauth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBaiduOauth(t *testing.T) {
	o := NewBaiduOauth("api_key", "secret_key", nil)
	assert.NotNil(t, o)
}

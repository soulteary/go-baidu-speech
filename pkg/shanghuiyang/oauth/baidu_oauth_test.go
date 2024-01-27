package oauth_test

import (
	"testing"

	"github.com/shanghuiyang/oauth"
	"github.com/stretchr/testify/assert"
)

func TestNewBaiduOauth(t *testing.T) {
	o := oauth.NewBaiduOauth("api_key", "secret_key", nil)
	assert.NotNil(t, o)
}

package oauth_test

import (
	"testing"

	"github.com/shanghuiyang/oauth"
	"github.com/stretchr/testify/assert"
)

func TestCacheImp(t *testing.T) {
	c := oauth.NewCacheImp()
	assert.NotNil(t, c)

	token := "a token"
	err := c.Put(token, 60)
	assert.NoError(t, err)

	vaild := c.IsValid()
	assert.True(t, vaild)

	newToken, err := c.Get()
	assert.NoError(t, err)
	assert.Equal(t, token, newToken)
}

/*
CacheImp is used to cache tokens.
*/

package oauth

import (
	"time"
)

// CacheImp implements Cache interface
type CacheImp struct {
	token    string
	expireAt time.Time
}

// NewCacheImp ...
func NewCacheImp() *CacheImp {
	return &CacheImp{
		token:    "",
		expireAt: time.Now(),
	}
}

// GetToken ...
func (c *CacheImp) Get() (string, error) {
	return c.token, nil
}

// SetToken ...
func (c *CacheImp) Put(token string, expiresInSec int64) error {
	c.token = token
	c.expireAt = time.Now().Add(time.Second * time.Duration(expiresInSec))
	return nil
}

// IsValid ...
func (c *CacheImp) IsValid() bool {
	sub := time.Until(c.expireAt)
	return sub.Seconds() > 10
}

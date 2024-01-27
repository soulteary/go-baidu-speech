/*
oauth is used to get token from web services. Usually, we use a cache to manage the tokens to avoid to request token from web services frequently.
*/

package oauth

type Oauth interface {
	Token() (string, error)
}

type Cache interface {
	Get() (string, error)
	Put(token string, expiresIn int64) error
	IsValid() bool
}

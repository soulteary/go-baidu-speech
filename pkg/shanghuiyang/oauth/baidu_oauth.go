/*
BaiduOauth implememt Oauth interface using baidu oauth.
*/

package oauth

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	baiduTokenURL = "https://openapi.baidu.com/oauth/2.0/token"
	grantType     = "client_credentials"
)

// BaiduOauth implememts Oauth interface
type BaiduOauth struct {
	apiKey    string
	secretKey string
	cache     Cache
}

// BaiduToken ...
type BaiduToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// NewBaiduOauth ...
func NewBaiduOauth(apiKey, secretKey string, cache Cache) *BaiduOauth {
	return &BaiduOauth{
		apiKey:    apiKey,
		secretKey: secretKey,
		cache:     cache,
	}
}

// Token ...
func (o *BaiduOauth) Token() (string, error) {
	if o.cache != nil && o.cache.IsValid() {
		return o.cache.Get()
	}

	formData := url.Values{
		"grant_type":    {grantType},
		"client_id":     {o.apiKey},
		"client_secret": {o.secretKey},
	}

	resp, err := http.PostForm(baiduTokenURL, formData)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if bytes.Contains(body, []byte("error")) {
		return "", errors.New("failed to get access token")
	}

	token := BaiduToken{}
	err = json.Unmarshal(body, &token)
	if err != nil {
		return "", err
	}
	if o.cache != nil {
		if err := o.cache.Put(token.AccessToken, token.ExpiresIn); err != nil {
			return "", err
		}
	}
	return token.AccessToken, nil
}

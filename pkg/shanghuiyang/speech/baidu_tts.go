package speech

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/shanghuiyang/oauth"
)

const (
	baiduTtsURL = "http://tsn.baidu.com/text2audio"
)

// TTS ...
type BaiduTTS struct {
	auth oauth.Oauth
}

type baiduTtsResponse struct {
	ErrNo  int    `json:"err_no"`
	ErrMsg string `json:"err_msg"`
	SN     string `json:"sn"`
	Idx    int    `json:"idx"`
}

// NewBaiduTTS ...
func NewBaiduTTS(auth oauth.Oauth) *BaiduTTS {
	return &BaiduTTS{
		auth: auth,
	}
}

// ToSpeech ...
func (t *BaiduTTS) ToSpeech(text string) ([]byte, error) {
	token, err := t.auth.Token()
	if err != nil {
		return nil, err
	}

	formData := url.Values{
		"tex":  {text},
		"lan":  {"zh"},
		"tok":  {token},
		"ctp":  {"1"},
		"cuid": {"go-speech"},
		"aue":  {"6"}, // wav
	}
	resp, err := http.PostForm(baiduTtsURL, formData)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	contentType := resp.Header.Get("Content-type")
	if contentType == "audio/wav" {
		return body, nil
	}

	var errResp baiduTtsResponse
	if err := json.Unmarshal(body, &errResp); err != nil {
		return nil, err
	}
	return nil, fmt.Errorf("error: %v, %v", errResp.ErrNo, errResp.ErrMsg)
}

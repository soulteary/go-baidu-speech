package speech

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/shanghuiyang/oauth"
)

const (
	baiduAsrURL    = "http://vop.baidu.com/server_api"
	baiduAsrProURL = "https://vop.baidu.com/pro_api"
)

// BaiduASR ...
type BaiduASR struct {
	auth oauth.Oauth
}

type baiduAsrResponse struct {
	ErrNo  int      `json:"err_no"`
	ErrMsg string   `json:"err_msg"`
	SN     string   `json:"sn"`
	Result []string `json:"result"`
}

type baiduAsrRequest struct {
	Format  string `json:"format"`
	Rate    int    `json:"rate"`
	Channel int    `json:"channel"`
	Token   string `json:"token"`
	Cuid    string `json:"cuid"`
	Len     int    `json:"len"`
	Speech  string `json:"speech"`
	// DevPid  int    `json:"dev_pid"` // for pro api
}

// NewBaiduASR ...
func NewBaiduASR(auth oauth.Oauth) *BaiduASR {
	return &BaiduASR{
		auth: auth,
	}
}

// ToText ...
func (a *BaiduASR) ToText(speech []byte, format Format) (string, error) {
	token, err := a.auth.Token()
	if err != nil {
		return "", err
	}

	req := &baiduAsrRequest{
		Format:  string(format),
		Rate:    16000,
		Channel: 1,
		Cuid:    "asr",
		Token:   token,
		Len:     len(speech),
		Speech:  base64.StdEncoding.EncodeToString(speech),
		// DevPid:  80001,
	}

	reqData, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", baiduAsrURL, bytes.NewReader(reqData))
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Content-Length", fmt.Sprintf("%d", len(reqData)))
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var r baiduAsrResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return "", err
	}
	if r.ErrNo > 0 {
		return "", fmt.Errorf("error: %v, %v", r.ErrNo, r.ErrMsg)
	}
	text := strings.TrimRight(r.Result[0], "。")
	return text, nil
}

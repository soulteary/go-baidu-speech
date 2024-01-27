package server

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var ttsAuth = Auth(os.Getenv("BAIDU_TTS_API_KEY"), os.Getenv("BAIDU_TTS_SECRET_KEY"))

func TTS(c *gin.Context) {
	text := strings.TrimSpace(c.PostForm("text"))
	if text == "" {
		c.String(http.StatusBadRequest, "未提供有效文本内容")
		return
	}

	success, file := BaiduTTS(ttsAuth, text)
	if !success {
		c.String(http.StatusBadRequest, "TTS API 调用失败")
		return
	}
	c.String(http.StatusOK, file)
}

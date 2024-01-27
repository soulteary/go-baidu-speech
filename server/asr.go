package server

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

var asrAuth = Auth(os.Getenv("BAIDU_ASR_API_KEY"), os.Getenv("BAIDU_ASR_SECRET_KEY"))

func ASR(c *gin.Context) {
	body, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "获取上传文件出错: %s", err.Error())
		return
	}

	file := fmt.Sprintf("%s.wav", GetUnixStr())
	fullPath := path.Join("./public", file)

	err = c.SaveUploadedFile(body, fullPath)
	if err != nil {
		c.String(http.StatusInternalServerError, "保存文件失败: %s", err.Error())
		return
	}

	success, text := BaiduASR(asrAuth, fullPath)
	if !success {
		c.String(http.StatusInternalServerError, "识别声音到文本失败: %s", err.Error())
		return
	}

	c.String(http.StatusOK, text)
}

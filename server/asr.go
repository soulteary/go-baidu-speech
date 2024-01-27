package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var asrAuth = Auth(os.Getenv("BAIDU_ASR_API_KEY"), os.Getenv("BAIDU_ASR_SECRET_KEY"))

func ASR(c *gin.Context) {
	success, text := BaiduASR(asrAuth, "file")
	if !success {
		log.Printf("failed to recognize the speech")
		return
	}
	log.Println("ASR 识别内容：", text)
}

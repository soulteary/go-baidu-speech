package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shanghuiyang/oauth"
	"github.com/shanghuiyang/speech"
)

func Auth(accessKey, secretKey string) *oauth.BaiduOauth {
	return oauth.NewBaiduOauth(accessKey, secretKey, oauth.NewCacheImp())
}

func TTS(auth *oauth.BaiduOauth, text string) (bool, string) {
	tts := speech.NewBaiduTTS(auth)
	data, err := tts.ToSpeech(text)
	if err != nil {
		log.Printf("failed to convert text to speech, error: %v", err)
		return false, ""
	}

	file := fmt.Sprintf("%s.wav", "test")
	if err := os.WriteFile(file, data, 0644); err != nil {
		log.Printf("failed to save %s, error: %v", file, err)
		return false, ""
	}
	return true, file
}

func ASR(auth *oauth.BaiduOauth, file string) (bool, string) {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Printf("failed to read wav/sample.wav, error: %v\n", err)
		return false, ""
	}

	asr := speech.NewBaiduASR(auth)
	text, err := asr.ToText(data, speech.Wav)
	if err != nil {
		log.Printf("failed to recognize the speech, error: %v", err)
		return false, ""
	}
	return true, text
}

func main() {
	text := "阳光彩虹小白马"
	ttsAuth := Auth(os.Getenv("BAIDU_TTS_API_KEY"), os.Getenv("BAIDU_TTS_SECRET_KEY"))
	success, file := TTS(ttsAuth, text)
	if !success {
		log.Printf("failed to convert text to speech")
		return
	}
	log.Println("TTS 文件保存地址：", file)

	asrAuth := Auth(os.Getenv("BAIDU_ASR_API_KEY"), os.Getenv("BAIDU_ASR_SECRET_KEY"))
	success, text = ASR(asrAuth, file)
	if !success {
		log.Printf("failed to recognize the speech")
		return
	}
	log.Println("ASR 识别内容：", text)
}

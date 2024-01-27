package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/shanghuiyang/oauth"
	"github.com/shanghuiyang/speech"
)

func main() {
	text := "阳光彩虹小白马"
	AK := os.Getenv("BAIDU_SPEECH_API_KEY")
	SK := os.Getenv("BAIDU_SPEECH_SECRET_KEY")

	auth := oauth.NewBaiduOauth(AK, SK, oauth.NewCacheImp())
	tts := speech.NewBaiduTTS(auth)
	data, err := tts.ToSpeech(text)
	if err != nil {
		log.Printf("failed to convert text to speech, error: %v", err)
		os.Exit(1)
	}

	if err := ioutil.WriteFile("test.wav", data, 0644); err != nil {
		log.Printf("failed to save test.wav, error: %v", err)
		os.Exit(1)
	}

	fmt.Println("success: test.wav")
}

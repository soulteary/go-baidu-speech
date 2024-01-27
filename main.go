package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/shanghuiyang/oauth"
	"github.com/shanghuiyang/speech"
)

const (
	baiduSpeechAPIKey    = "your_baidu_api_key"
	baiduSpeechSecretKey = "your_baidu_speech_secret_key"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("error: invalid args")
		fmt.Println(`usage: tts "a test"`)
		os.Exit(1)
	}
	text := os.Args[1]

	auth := oauth.NewBaiduOauth(baiduSpeechAPIKey, baiduSpeechSecretKey, oauth.NewCacheImp())
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

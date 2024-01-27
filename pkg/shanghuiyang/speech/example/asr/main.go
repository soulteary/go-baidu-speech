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
		fmt.Println("usage: asr test.wav")
		os.Exit(1)
	}
	speechFile := os.Args[1]

	data, err := ioutil.ReadFile(speechFile)
	if err != nil {
		log.Printf("failed to read %v, error: %v", speechFile, err)
		os.Exit(1)
	}

	auth := oauth.NewBaiduOauth(baiduSpeechAPIKey, baiduSpeechSecretKey, oauth.NewCacheImp())
	asr := speech.NewBaiduASR(auth)
	text, err := asr.ToText(data, speech.Wav)
	if err != nil {
		log.Printf("failed to recognize the speech, error: %v", err)
		os.Exit(1)
	}
	fmt.Println(text)
}

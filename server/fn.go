package server

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shanghuiyang/oauth"
	"github.com/shanghuiyang/speech"
)

func Auth(accessKey, secretKey string) *oauth.BaiduOauth {
	return oauth.NewBaiduOauth(accessKey, secretKey, oauth.NewCacheImp())
}

func BaiduTTS(auth *oauth.BaiduOauth, text string) (bool, string) {
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

func BaiduASR(auth *oauth.BaiduOauth, file string) (bool, string) {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Printf("failed to read %s, error: %v\n", file, err)
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

func GetPort() int {
	defaultPort := 8080
	portStr := os.Getenv("PORT")

	if portStr == "" {
		log.Printf("The PORT environment variable is empty, using the default port: %d\n", defaultPort)
		return defaultPort
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Printf("The PORT environment variable is not a valid integer, using the default port: %d\n", defaultPort)
		return defaultPort
	}

	if port < 1 || port > 65535 {
		log.Printf("The PORT environment variable is not a valid port number, using the default port: %d\n", defaultPort)
		return defaultPort
	}

	log.Printf("Using the port: %d\n", port)
	return port
}

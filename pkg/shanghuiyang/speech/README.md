# speech
speech implements ASR and TTS using APIs provided by web services like baidu and google.

## Usage
### [ASR](/example/asr)
convert speech to text
```go
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
	data, err := ioutil.ReadFile("wav/sample.wav")
	if err != nil {
		log.Printf("failed to read wav/sample.wav, error: %v\n", err)
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
```

### [TTS](/example/tts)
convert text to speech
```go
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
	auth := oauth.NewBaiduOauth(baiduSpeechAPIKey, baiduSpeechSecretKey, oauth.NewCacheImp())
	tts := speech.NewBaiduTTS(auth)
	data, err := tts.ToSpeech("中国北京")
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
```

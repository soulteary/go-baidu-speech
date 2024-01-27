package speech

const (
	Amr Format = "amr"
	Wav Format = "wav"
	M4a Format = "m4a"
	Pcm Format = "pcm"
)

// Format defines the format of speech files.
type Format string

// ASR ...
type ASR interface {
	ToText(speech []byte, format Format) (string, error)
}

// TTS ...
type TTS interface {
	ToSpeech(text string) ([]byte, error)
}

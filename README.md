# Go Baidu Speech

API Service of Baidu ASR/TTS.

## TTS

![](.github/screenshots/tts.jpg)

```bash
curl -X POST -d "text=阳光彩虹小白马" http://localhost:8080/tts
1706342489.wav
```

## ASR

```bash
curl -X POST -F "file=@/Users/soulteary/Downloads/1706341924.wav" http://localhost:8080/asr
阳光彩虹小白马
```

## Credit

- pkg: https://github.com/shanghuiyang/oauth / https://github.com/shanghuiyang/speech

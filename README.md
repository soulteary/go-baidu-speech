# Go Baidu Speech

API Service of Baidu ASR/TTS.

## TTS

![](.github/screenshots/tts.jpg)

```bash
curl -X POST -d "text=阳光彩虹小白马" http://localhost:8080/tts
1706342489.wav
```

Python example: [tts](./example/python/tts.py)

## ASR

```bash
curl -X POST -F "file=@/Users/soulteary/Downloads/1706341924.wav" http://localhost:8080/asr
阳光彩虹小白马
```

Python example: [asr](./example/python/asr.py)

## Docker

build it:

```bash
docker build -t soulteary/go-baidu-speech . -f docker/Dockerfile
```

play:

```bash
docker compose up -d
```

## Credit

- pkg: https://github.com/shanghuiyang/oauth / https://github.com/shanghuiyang/speech

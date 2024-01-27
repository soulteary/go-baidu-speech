module github.com/soulteary/go-baidu-speech

go 1.21

replace github.com/shanghuiyang/oauth => ./pkg/shanghuiyang/oauth

replace github.com/shanghuiyang/speech => ./pkg/shanghuiyang/speech

require (
	github.com/shanghuiyang/oauth v0.0.0-20210815163430-0403e002fe02
	github.com/shanghuiyang/speech v0.0.0-00010101000000-000000000000
)

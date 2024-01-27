package main

import (
	"os"
	"strings"

	"github.com/soulteary/go-baidu-speech/server"
)

func main() {
	var debug bool
	if strings.ToLower(strings.TrimSpace(os.Getenv("DEBUG"))) == "on" {
		debug = true
	} else {
		debug = false
	}
	server.Launch(debug)
}

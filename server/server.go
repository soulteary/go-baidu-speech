package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	static "github.com/soulteary/gin-static"
)

func Launch(debugMode bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if debugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	os.MkdirAll("./public", os.ModePerm)

	route := gin.New()
	route.Use(gin.Recovery())
	route.Use(gzip.Gzip(gzip.DefaultCompression))

	route.Use(static.Serve("/", static.LocalFile("./public", true)))
	route.GET("/ping", Ping)
	route.POST("/tts", TTS)
	route.POST("/asr", ASR)

	srv := &http.Server{
		Addr:              "0.0.0.0:" + strconv.Itoa(GetPort()),
		Handler:           route,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Program start error: %s\n", err)
		}
	}()
	log.Println("github.com/soulteary/go-baidu-speech has started 🚀")

	<-ctx.Done()

	stop()
	log.Println("The program is closing, if you want to end it immediately, please press `CTRL+C`")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Program was forced to close: %s\n", err)
	}
}

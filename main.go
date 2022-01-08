package main

import (
	"os"
	"log"
	"time"
	"net/http"

	"github.com/scmn-dev/get-latest/router"
)

func main() {
	logger := log.New(os.Stdout, "[sm-get-latest] ", 0)
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		MaxHeaderBytes: 10, // 10 MB
		Addr:           ":" + port,
		WriteTimeout:   time.Second * 24,
		ReadTimeout:    time.Second * 24,
		IdleTimeout:    time.Second * 60,
		Handler:        router.New(),
	}

	logger.Printf("📡 Server listening on %s", port)

	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}

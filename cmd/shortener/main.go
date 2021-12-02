package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/miramariadev/go-musthave-shortener-tpl/internal/app"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	storage := app.NewMemoryStorage()
	service := app.NewURLShortenerService(storage)
	handler := app.NewURLShortenerHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HandleShortURL)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}

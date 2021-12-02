package main

import (
	"log"
	"net/http"

	"github.com/miramariadev/go-musthave-shortener-tpl/internal/app"
)

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

//go build  github.com/miramariadev/go-musthave-shortener-tpl/cmd/shortener

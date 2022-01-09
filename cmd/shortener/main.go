package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/miramariadev/go-musthave-shortener-tpl/internal/app"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	storage := app.NewMemoryStorage()
	service := app.NewURLShortenerService(storage)
	handler := app.NewURLShortenerHandler(service)

	r := chi.NewRouter()

	r.Get("/{urlID}", handler.HandleShortURL)
	r.Post("/", handler.HandleShortURL)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Fatal(server.ListenAndServe())
}

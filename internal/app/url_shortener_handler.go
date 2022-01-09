package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
	"net/url"
)

const host = "http://localhost:8080/"

type URLService interface {
	CreateShortURL(url string) string
	GetLongURLByID(id string) (string, error)
}

type URLShortenerHandler struct {
	service URLService
}

func NewURLShortenerHandler(service URLService) *URLShortenerHandler {
	return &URLShortenerHandler{
		service: service,
	}
}

func (c *URLShortenerHandler) HandleShortURL(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		defer r.Body.Close()
		requestData, _ := io.ReadAll(r.Body)
		longURL := string(requestData)

		_, err := url.ParseRequestURI(longURL)
		if err != nil {
			http.Error(w, "StatusBadRequest", http.StatusBadRequest)
			return
		}

		idURL := c.service.CreateShortURL(longURL)
		responseURL := host + idURL
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, responseURL)

		return

	case "GET":
		id := chi.URLParam(r, "urlID")
		if id == "" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}

		longURL, err := c.service.GetLongURLByID(id)
		log.Println(id)
		if err != nil {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Location", longURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
		fmt.Fprint(w, "")
		return

	default:
		http.Error(w, "Bad method", http.StatusBadRequest)
		return
	}
}

//curl -X POST -H "Content-Type: text/plain" -d  'http://rp21sh.yandex/avshz7qrl/h4ululwnp7bow' http://localhost:8080
//curl -X GET http://localhost:8080/56430

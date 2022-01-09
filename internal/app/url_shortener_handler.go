package app

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const maxInt = 99999

type URLShortenerHandler struct {
	service *URLShortenerService
}

func NewURLShortenerHandler(service *URLShortenerService) *URLShortenerHandler {
	return &URLShortenerHandler{
		service: service,
	}
}

func (c *URLShortenerHandler) HandleShortURL(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		requestData, _ := io.ReadAll(r.Body)
		longURL := string(requestData)

		_, err := url.ParseRequestURI(longURL)
		if err != nil {
			http.Error(w, "StatusBadRequest", http.StatusBadRequest)
			return
		}

		idURL := c.service.CreateShortURL(longURL)
		responseURL := "http://localhost:8080/" + idURL
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(201)
		fmt.Fprint(w, responseURL)

		return

	case "GET":
		data := strings.Split(r.URL.String(), "/")
		id := data[len(data)-1]

		longURL, err := c.service.GetLongURLByID(id)
		if err != nil {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Location", longURL)
		w.WriteHeader(307)
		fmt.Fprint(w, "")
		return

	default:
		http.Error(w, "Bad method", http.StatusBadRequest)
		return
	}
}

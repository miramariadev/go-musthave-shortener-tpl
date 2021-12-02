package app

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
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

		newURL := c.service.CreateShortURL(longURL)

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(201)
		fmt.Fprint(w, newURL)

		return

	case "GET":
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "StatusBadRequest", http.StatusBadRequest)
			return
		}

		/*data := strings.Split(r.URL.String(), "/")
		id = data[len(data) - 1]*/

		//id это символы после слеша

		longURL, err := c.service.GetLongURLByID(id)
		if err != nil {
			http.NotFoundHandler()
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Location", longURL)
		w.WriteHeader(307)
		return

	default:
		http.Error(w, "Bad method", http.StatusBadRequest)
		return
	}
}

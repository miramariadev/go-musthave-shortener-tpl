package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleURLShortener)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}

func handleURLShortener(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		ur := r.FormValue("url")

		_, err := url.ParseRequestURI(ur)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)

		}
		newUrl := ur + "/1"

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(201)

		w.Write([]byte(newUrl))

	case "GET":
		id := r.URL.Query().Get("id")

		fullUrl := "http://localhost:8080"

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Location", fullUrl)
		w.WriteHeader(307)

		w.Write([]byte(id))

	default:
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
}

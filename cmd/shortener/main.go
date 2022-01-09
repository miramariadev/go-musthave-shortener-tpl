package main

import (
	"fmt"
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

		var str string
		for key, values := range r.Form {
			for _, value := range values {
				fmt.Println(key, value)
				str = str + key + value
			}
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, str, http.StatusBadRequest)
			return
		}

		ur := r.FormValue("url")

		newURL := ur + "/1"

		if ur == "" {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(201)
			fmt.Fprintln(w, "http://localhost:8080/1")

			return
		}

		_, err := url.ParseRequestURI(ur)
		if err != nil {
			http.Error(w, "ParseRequestURI", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(201)

		fmt.Fprintln(w, newURL)
		return

	case "GET":
		id := r.URL.Query().Get("id")

		fullURL := "http://localhost:8080"

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Location", fullURL)
		w.WriteHeader(307)

		w.Write([]byte(id))
		return

	default:
		http.Error(w, "Bad method", http.StatusBadRequest)
		return
	}
}

// вместо url возвращается 400
// вместо 400 возвращается 201

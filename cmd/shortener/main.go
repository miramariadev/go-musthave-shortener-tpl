package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleUrlShortener)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}

func handleUrlShortener(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		url := r.FormValue("url")
		if url == "" {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		newUrl := url + "/1"

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(201)

		resp := make(map[string]string)
		resp["url"] = newUrl
		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)

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

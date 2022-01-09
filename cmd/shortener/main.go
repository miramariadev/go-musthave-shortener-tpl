package main

import (
	"encoding/json"
	"fmt"
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
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		url := r.FormValue("url")
		if url == "" {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(200)
			fmt.Fprintln(w, "No")
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(201)
		resp := make(map[string]string)
		resp["message"] = "Unauthorized"
		jsonResp, _ := json.Marshal(url + "/1")
		w.Write(jsonResp)

	case "GET":
		url := r.URL.Query().Get("url")
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(201)
		resp := make(map[string]string)
		resp["message"] = "Unauthorized"
		jsonResp, _ := json.Marshal(url + "/2")
		w.Write(jsonResp)

	default:
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

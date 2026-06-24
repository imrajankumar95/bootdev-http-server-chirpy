package main

import (
	"net/http"
	"log"
)


func main() {

	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("."))))

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		
		w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(http.StatusText(http.StatusOK)))
	})

	srv := &http.Server{

		Handler : mux,
		Addr : ":8080",
	}
	
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}

	
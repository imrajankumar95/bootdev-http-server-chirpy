package main

import (
	"net/http"
	"log"
)


func main() {

	mux := http.NewServeMux()

	srv := &http.Server{

		Handler : mux,
		Addr : ":8080",
	}
	
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("HTTP server ListenAndServe: %v", err)
	}
}

	
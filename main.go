package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/imad-almansi/backend-test-golang/pkg/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/facts", handlers.HandleFacts)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

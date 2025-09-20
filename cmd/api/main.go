package main

import (
	"fmt"
	"log"
	"net/http"

	"duck/internal/api"
	"duck/internal/store"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Create a new in memory store to act as our data store

	// Make a server by calling api.NewServer with our db

	// Create a new chi router/mux by calling chi.NewRouter()
	mux := chi.NewRouter()

	// Call the RegisterHandler function with chi router to register our handler
	srv.RegisterHandler(mux)

	// Our server that will listen on our port and use our mux to handle requests
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Listening on http://localhost:8080")

	// Run the server
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("http server error: %s", err))
	}
}


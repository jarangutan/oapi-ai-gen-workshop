package main

import (
	"log"
	"net/http"

	"duck/internal/api"
	"duck/internal/store"
)

func main() {
	// This is dependency injection. InMemoryStore implements api.DuckStore
	// implicitly, so we can pass it straight into NewServer.
	// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection
	duckStore := store.NewInMemoryStore()
	apiServer := api.NewServer(duckStore)

	// apiServer.Handler gives net/http our generated routes and middleware.
	// https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go
	server := &http.Server{
		Addr:    ":8080",
		Handler: apiServer.Handler(),
	}

	log.Println("Listening on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

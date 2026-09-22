package main

import (
	"log"
	"net/http"

	"duck/internal/api"
	"duck/internal/store"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("duck.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	duckStore := store.NewSQLiteStore(db)
	if err := duckStore.Migrate(); err != nil {
		log.Fatal(err)
	}

	apiServer := api.NewServer(duckStore)
	server := &http.Server{
		Addr:    ":8080",
		Handler: apiServer.Handler(),
	}

	log.Println("Listening on http://localhost:8080 with SQLite")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

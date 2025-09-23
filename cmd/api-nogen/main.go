package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Duck - https://go.dev/wiki/Well-known-struct-tags
type Duck struct {
	Color string `json:"color"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Size  string `json:"size"`
}

// This is a small Demo of what it would look like to build the API without codegen
func main() {
	// A Mux is an HTTP Multiplexer.
	// It matches the URL of each incoming request against a list
	// of registered patterns and calls the handler for the pattern
	// that most closely matches the URL.
	// See: https://pkg.go.dev/net/http#ServeMux
	mux := http.NewServeMux()

	// We create a handler that will be called when someone calls GET /duck
	mux.HandleFunc("GET /duck", func(w http.ResponseWriter, r *http.Request) {
		// Set the content type so our caller knows what we are responding with
		w.Header().Set("Content-Type", "application/json")
		// Return a JSON response back
		encoder := json.NewEncoder(w)
		encoder.Encode(Duck{
			ID:    1,
			Name:  "Donna",
			Color: "pink",
			Size:  "medium",
		})
	})

	// This is the actual HTTP server that will handle traffic on the
	// port we set using the handlers we made
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("listening on http://localhost:8080")

	// Start up the server
	if err := server.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("http server error: %s", err))
	}
}

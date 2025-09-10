package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type Duck struct {
	Color string `json:"color"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Size  string `json:"size"`
}

// This is a small Demo of what it would look like to build the API without codegen
func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Port for test HTTP server")
	flag.Parse()

	// A Mux is an HTTP Multiplexer. See: https://pkg.go.dev/net/http#ServeMux
	mux := http.NewServeMux()

	// We create a handler that will act when someone calls GET /duck
	mux.HandleFunc("GET /duck", func(w http.ResponseWriter, r *http.Request) {
		// Set the content type so our caller knows what we are responding with
		w.Header().Set("Content-Type", "application/json")
		// Return a JSON response back
		json.NewEncoder(w).Encode(Duck{
			ID:    1,
			Name:  "Donna",
			Color: "pink",
			Size:  "medium",
		})
	})

	// This is the actual HTTP server that will handle traffic on the
	// port we set using the handlers we made
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
		// Recommended timeouts from
		// https://blog.cloudflare.com/exposing-go-on-the-internet/
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	slog.Info(fmt.Sprintf("Server listening on :%d", port))

	// Start up the server
	if err := server.ListenAndServe(); err != nil {
		slog.Error("Server failed to start", "error", err)
	}
}

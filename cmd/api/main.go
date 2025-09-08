package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"duck/internal/server"
	"duck/internal/store"

	"github.com/go-chi/chi/v5"
)

func main() {
	port := flag.Int("port", 8080, "Port for test HTTP server")
	flag.Parse()

	db := store.NewInMemoryStore()
	srv := server.NewServer(db)

	strictHandler := server.NewStrictHandler(srv, nil)

	r := chi.NewRouter()
	r.Use(server.WithSwaggerValidate())
	server.HandlerFromMux(strictHandler, r)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", *port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)

	// Run graceful shutdown in a separate goroutine
	go gracefulShutdown(server, done)

	fmt.Printf("Listening on http://localhost:%d\n", *port)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	// Wait for the graceful shutdown to complete
	<-done
	log.Println("Graceful shutdown complete.")
}

// gracefulShutdown listens for the a kill signal and an optional interrupt from the user
// see https://victoriametrics.com/blog/go-graceful-shutdown/
func gracefulShutdown(apiServer *http.Server, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")
	stop() // Allow Ctrl+C to force shutdown

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}

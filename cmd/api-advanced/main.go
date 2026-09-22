package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"duck/internal/api"
	"duck/internal/store"
)

func main() {
	port := flag.Int("port", 8080, "HTTP server port")
	flag.Parse()

	// This pattern is called dependency injection. The in-memory store implements
	// api.DuckStore implicitly, so we can swap it for another store or a fake in tests.
	// I usually prefer testing with the real thing when I can: https://testcontainers.com/
	duckStore := store.NewInMemoryStore()
	apiServer := api.NewServer(duckStore)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: apiServer.Handler(),
		// These timeouts help protect a public HTTP server from slow or broken clients.
		// See: https://blog.cloudflare.com/exposing-go-on-the-internet/
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  2 * time.Minute,
	}

	// Shutdown waits for an OS signal in another goroutine while ListenAndServe
	// handles traffic down below.
	shutdownComplete := make(chan struct{})
	go gracefulShutdown(server, shutdownComplete)

	log.Printf("Listening on http://localhost:%d", *port)
	// Shutdown makes ListenAndServe return ErrServerClosed, which isn't actually an error here.
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}

	<-shutdownComplete
	log.Println("Graceful shutdown complete.")
}

// gracefulShutdown waits for an interrupt and gives the server time to finish
// any requests it is already handling.
// See: https://victoriametrics.com/blog/go-graceful-shutdown/
func gracefulShutdown(server *http.Server, done chan<- struct{}) {
	defer close(done)

	// NotifyContext catches Ctrl+C or a termination signal and cancels ctx
	// instead of immediately killing the process.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()

	log.Println("Shutting down gracefully, press Ctrl+C again to force it")
	// Restore the normal signal behavior so a second Ctrl+C exits right away.
	stop()

	// Give in-flight requests five seconds to finish before giving up.
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("Server forced to shut down: %v", err)
	}
}

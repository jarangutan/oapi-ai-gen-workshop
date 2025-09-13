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

	"duck/internal/api"
	"duck/internal/store"

	"github.com/go-chi/chi/v5"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Port for test HTTP server")
	flag.Parse()

	// This pattern here is called ["Dependency Injection"](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection), you'll see a lot of that in Go.
	// The in memory store implements the DuckStore interface from the Server implicitly so it can be passed to NewServer.
	// With this, we can feed NewServer a fake implementation in tests if we want to isolate behavior.
	// But usually it's best to test with the [real things](https://testcontainers.com/) in Go.
	db := store.NewInMemoryStore()
	srv := api.NewServer(db)

	// Chi is a lightweight router (mux) that works with the built in standard library http handlers.
	// It's got a lot of nice features and it's compatible with any tooling that works off the standard library like middlewares.
	// To learn more on handlers and muxes, see: https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go
	//
	// NOTE! Chi ain't the fastest but don't look at benchmarks when picking your tools.
	// The fastests libraries make a lot of tradeofs and most of the time, the speed increase isn't worth the pain.
	r := chi.NewRouter()
	srv.RegisterHandler(r)

	// Our server that will listen on our port and use our mux to handle requests
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
		// Recommended timeouts from
		// https://blog.cloudflare.com/exposing-go-on-the-internet/
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)

	// Run graceful shutdown in a separate goroutine
	go gracefulShutdown(server, done)

	log.Printf("Listening on http://localhost:%d\n", port)
	// Run the server
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	// Wait for the graceful shutdown to complete
	<-done
	log.Println("Graceful shutdown complete.")
}

// gracefulShutdown listens for the a kill signal and an optional interrupt from the user.
// see https://victoriametrics.com/blog/go-graceful-shutdown/
func gracefulShutdown(apiServer *http.Server, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	// This is done by catching those signals and instead of performing the default exit action,
	// we send a Done to the context to let the server know to start shutting down.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")
	// Allow Ctrl+C to force shutdown by reseting the default behavior of those syscalls back to their original
	// behavior i.e. exiting the program.
	stop()

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}

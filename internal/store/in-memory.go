package store

import (
	"context"
	"sync"

	"duck/internal/api"
)

// InMemoryStore will store our rubber ducks in memory until the server shuts down
// We will need a map, an index we can increment as a primary key,
// and a Read Write mutex for concurrent access
type InMemoryStore struct{}

// Same as NewServer, we need to create a function that returns us a new InMemoryStore
func NewInMemoryStore() *InMemoryStore {
	return nil
}

// We will create methods on our InMemoryStore struct that immplements the 
// DuckStore interface from our api.
// 
// Why not declare the interface here?
// From: https://go.dev/doc/effective_go#interfaces
// > Interfaces in Go provide a way to specify the behavior of an object: if something can do this, then it can be used *here*
// The server.go file needs a DuckStore and anything that satisfies that DuckStore interface can be used by the server

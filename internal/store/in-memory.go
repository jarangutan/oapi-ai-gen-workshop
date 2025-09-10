// Package store provides different stores that implement the DuckStore interface
//
// NOTE! Unless you are building something stupid simple like this or a Proof of concept (POC),
// NEVER COUPLE YOUR API TYPES TO YOUR DATABASE TYPES! We are breaking that rule for demonstrations only.
//
// Your database types seldom look like your API types in real apps so keep them separate even if they look very similar.
// It will make it easier if you ever have to change your API to GRPC or have to release a v2 while keeping backwards compatibility.
// A good [go proverb](https://go-proverbs.github.io/) to remember is "A little copying is better than a little dependency."
//
// If you're curious about this pattern of using interfaces to separate your store implementations,
// look up the ["repository pattern"](https://threedots.tech/post/repository-pattern-in-go/).
// You'll see folks call the package repo, storage, db, database, etc. Pick a name you like :-)
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

// We will create methods on our InMemoryStore struct that implements the 
// DuckStore interface from our api.
// 
// Why not declare the interface here?
// From: https://go.dev/doc/effective_go#interfaces
// > Interfaces in Go provide a way to specify the behavior of an object: if something can do this, then it can be used *here*
// The server.go file needs a DuckStore and anything that satisfies that DuckStore interface can be used by the server

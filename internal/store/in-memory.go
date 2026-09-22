// Package store provides different ways to store our rubber ducks.
//
// NOTE! This in-memory store uses our generated API types directly to keep the
// workshop short. Your database types usually won't look exactly like your API
// types, and keeping them separate makes API changes or a v2 much less painful.
// A good Go proverb to remember is "A little copying is better than a little dependency."
// https://go-proverbs.github.io/
//
// If you're curious, this is usually called the repository pattern:
// https://threedots.tech/post/repository-pattern-in-go/
// Folks call this package repo, storage, db, database, and plenty more. Pick a name you like :-)
package store

import (
	"cmp"
	"context"
	"slices"
	"sync"

	"duck/internal/api"
)

// InMemoryStore keeps our rubber ducks in memory until the server shuts down.
// We need a map, an index we can increment for IDs, and a mutex because HTTP
// handlers can use the store at the same time.
type InMemoryStore struct {
	ducks map[int]api.RubberDuck
	index int
	mu    sync.RWMutex // https://gobyexample.com/mutexes
}

// NewInMemoryStore gives us an empty in-memory store.
func NewInMemoryStore() *InMemoryStore {
	// sync.RWMutex's zero value is ready to use, so only the map needs initialization.
	return &InMemoryStore{ducks: make(map[int]api.RubberDuck)}
}

// InMemoryStore implements the api.DuckStore interface, so why not declare the
// interface here?
//
// Effective Go puts it nicely:
// "Interfaces in Go provide a way to specify the behavior of an object:
// if something can do this, then it can be used here."
// https://go.dev/doc/effective_go#interfaces
//
// As long as InMemoryStore satisfies DuckStore, the server can use it. Since
// Server is what needs a DuckStore, that's where we declare the interface.

// ListDucks returns every duck in ID order.
func (s *InMemoryStore) ListDucks(_ context.Context) ([]api.RubberDuck, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	ducks := make([]api.RubberDuck, 0, len(s.ducks))
	for _, duck := range s.ducks {
		ducks = append(ducks, duck)
	}

	slices.SortFunc(ducks, func(a, b api.RubberDuck) int {
		return cmp.Compare(a.ID, b.ID)
	})

	return ducks, nil
}

// CreateDuck assigns an ID and stores duck.
func (s *InMemoryStore) CreateDuck(_ context.Context, duck api.DuckRequest) (api.RubberDuck, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.index++
	created := api.RubberDuck{
		ID:    s.index,
		Color: duck.Color,
		Name:  duck.Name,
		// DuckRequestSize and RubberDuckSize are two different named string types.
		// See: https://go.dev/blog/constants#string-constants
		Size: api.RubberDuckSize(duck.Size),
	}
	s.ducks[created.ID] = created

	return created, nil
}

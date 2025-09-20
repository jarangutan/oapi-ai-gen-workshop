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
	"cmp"
	"context"
	"slices"
	"sync"

	"duck/internal/api"
)

// InMemoryStore will store our rubber ducks in memory until the server shuts down
// We will need a map, an index we can increment as a primary key,
// and a Read Write mutex for concurrent access
type InMemoryStore struct {
	ducks map[uint]api.RubberDuck
	index uint
	mu    sync.RWMutex // https://gobyexample.com/mutexes
}

// Same as NewServer, we need to create a function that returns us a new InMemoryStore
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		ducks: make(map[uint]api.RubberDuck),
		mu:    sync.RWMutex{},
	}
}

// We will fill in methods on our InMemoryStore struct that implements the 
// DuckStore interface from our api.
// 
// Why not declare the interface here?
// From: https://go.dev/doc/effective_go#interfaces
// > Interfaces in Go provide a way to specify the behavior of an object: if something can do this, then it can be used *here*
// The server.go file needs a DuckStore and anything that satisfies that DuckStore interface can be used by the server

func (i *InMemoryStore) GetDucks(ctx context.Context) ([]api.RubberDuck, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	d := make([]api.RubberDuck, 0, len(i.ducks))
	for _, v := range i.ducks {
		d = append(d, v)
	}

	slices.SortFunc(d, func(a,b api.RubberDuck) int {
		return cmp.Compare(a.Id, b.Id)
	})

	return d, nil
}

func (i *InMemoryStore) CreateDuck(ctx context.Context, duck api.NewRubberDuck) (api.RubberDuck, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.index++
	id := i.index
	i.ducks[id] = api.RubberDuck{
		Id:    int(id),
		Color: duck.Color,
		Name:  duck.Name,
		Size:  api.RubberDuckSize(duck.Size), // see https://go.dev/blog/constants#string-constants
	}

	return i.ducks[id], nil
}

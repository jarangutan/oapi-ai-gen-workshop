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

// InMemoryStore stores rubber ducks in memory until the server shuts down
// Ideally, you would create domain types and have store import them
type InMemoryStore struct {
	ducks map[uint]api.RubberDuck
	index uint
	mu    sync.RWMutex // https://gobyexample.com/mutexes
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		ducks: make(map[uint]api.RubberDuck),
		mu:    sync.RWMutex{},
	}
}

func (i *InMemoryStore) GetDucks(ctx context.Context) ([]api.RubberDuck, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	d := make([]api.RubberDuck, 0, len(i.ducks))
	for _, v := range i.ducks {
		d = append(d, v)
	}
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

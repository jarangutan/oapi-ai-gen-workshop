package store

import (
	"context"
	"sync"

	"duck/internal/server"
)

// InMemoryStore stores rubber ducks in memory until the server shuts down
// Ideally, you would create domain types and have store import them
type InMemoryStore struct {
	ducks map[int]server.RubberDuck
	index int
	mu    sync.RWMutex
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		ducks: make(map[int]server.RubberDuck),
		mu:    sync.RWMutex{},
	}
}

func (i *InMemoryStore) GetDucks(ctx context.Context) ([]server.RubberDuck, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	d := []server.RubberDuck{}
	for _, v := range i.ducks {
		d = append(d, v)
	}
	return d, nil
}

func (i *InMemoryStore) CreateDuck(ctx context.Context, duck server.NewRubberDuck) (server.RubberDuck, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.index++
	id := i.index
	i.ducks[id] = server.RubberDuck{
		Id:    id,
		Color: duck.Color,
		Name:  duck.Name,
		Size:  server.RubberDuckSize(duck.Size), // see https://go.dev/blog/constants#string-constants
	}

	return i.ducks[id], nil
}

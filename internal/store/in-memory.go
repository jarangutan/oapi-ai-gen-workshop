package store

import (
	"context"
	"sync"

	"duck/internal/api"
)

// InMemoryStore stores rubber ducks in memory until the server shuts down
// Ideally, you would create domain types and have store import them
type InMemoryStore struct {
	ducks map[int]api.RubberDuck
	index int
	mu    sync.RWMutex // https://gobyexample.com/mutexes
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		ducks: make(map[int]api.RubberDuck),
		mu:    sync.RWMutex{},
	}
}

func (i *InMemoryStore) GetDucks(ctx context.Context) ([]api.RubberDuck, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	d := []api.RubberDuck{}
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
		Id:    id,
		Color: duck.Color,
		Name:  duck.Name,
		Size:  api.RubberDuckSize(duck.Size), // see https://go.dev/blog/constants#string-constants
	}

	return i.ducks[id], nil
}

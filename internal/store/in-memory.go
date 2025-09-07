package store

import (
	"context"
	"sync"

	"duck/internal/server"
)

// InMemory stores rubber ducks in memory until the server shuts down
// Ideally, you would create domain types and have store import them
type InMemory struct {
	db map[int]server.RubberDuck
	mu sync.RWMutex
}

func (i *InMemory) GetDucks(ctx context.Context) ([]server.RubberDuck, error) {
	return nil, nil
}

func (i *InMemory) CreateDuck(ctx context.Context, duck server.NewRubberDuck) error {
	return nil
}

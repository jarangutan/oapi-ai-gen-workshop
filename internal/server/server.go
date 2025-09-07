package server

import (
	"context"
	"fmt"
)

var _ StrictServerInterface = (*Server)(nil)

type DuckStore interface {
	GetDucks(ctx context.Context) ([]RubberDuck, error)
	CreateDuck(ctx context.Context, duck NewRubberDuck) error
}

type Server struct {
	duckStore DuckStore
}

func NewServer(ds DuckStore) *Server {
	server := &Server{
		duckStore: ds,
	}

	return server
}

func (s *Server) GetDucks(ctx context.Context, request GetDucksRequestObject) (GetDucksResponseObject, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (s *Server) CreateDuck(ctx context.Context, request CreateDuckRequestObject) (CreateDuckResponseObject, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

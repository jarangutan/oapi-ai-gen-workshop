package server

import (
	"context"
	"fmt"
)

var _ StrictServerInterface = (*Server)(nil)

type DuckStore interface {
	GetDucks(ctx context.Context) ([]RubberDuck, error)
	CreateDuck(ctx context.Context, duck NewRubberDuck) (RubberDuck, error)
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
	ducks, err := s.duckStore.GetDucks(ctx)
	if err != nil {
		return GetDucks500JSONResponse{
			Code:    500,
			Message: fmt.Sprintf("failed to get ducks: %s", err),
		}, nil
	}
	return GetDucks200JSONResponse(ducks), nil
}

func (s *Server) CreateDuck(ctx context.Context, request CreateDuckRequestObject) (CreateDuckResponseObject, error) {
	body := *request.Body

	duck, err := s.duckStore.CreateDuck(ctx, body)
	if err != nil {
		return CreateDuck500JSONResponse{
			Code:    500,
			Message: fmt.Sprintf("failed to get ducks: %s", err),
		}, nil
	}

	return CreateDuck201JSONResponse(duck), nil
}

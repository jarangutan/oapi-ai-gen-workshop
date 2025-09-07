package server

import (
	"context"

	"duck/internal/server/oapigen"
	"duck/internal/service"
)

var _ oapigen.StrictServerInterface = (*Server)(nil)

type DuckService interface {
	GetDuck(id int) service.RubberDuck
	AddDuck(duck service.RubberDuck) error
}

type Server struct {
	duckService DuckService
}

func NewServer(ds DuckService) *Server {
	server := &Server{
		duckService: ds,
	}

	return server
}

func (s *Server) GetDucks(ctx context.Context, request oapigen.GetDucksRequestObject) (oapigen.GetDucksResponseObject, error) {
	return nil, nil
}

func (s *Server) CreateDuck(ctx context.Context, request oapigen.CreateDuckRequestObject) (oapigen.CreateDuckResponseObject, error) {
	return nil, nil
}

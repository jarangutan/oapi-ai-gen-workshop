package api

import (
	"context"
	"fmt"

	"github.com/go-chi/chi/v5"
)

// Guarantees our Server adheres to the StrictServerInterface
// see https://dev.to/kittipat1413/checking-if-a-type-satisfies-an-interface-in-go-432n
var _ StrictServerInterface = (*Server)(nil)

// DuckStore interface describes the methods needed by the Server
//
// Go idiom is to declare interfaces where they are used. If you repeat it a lot,
//
//	then consider exporting it or moving it to a neutral package
//
// See standard library io -> net -> net/http for an idea of exporting interfaces in practice
type DuckStore interface {
	GetDucks(ctx context.Context) ([]RubberDuck, error)
	CreateDuck(ctx context.Context, duck NewRubberDuck) (RubberDuck, error)
}

// Server holds our handlers
type Server struct {
	duckStore DuckStore
}

// NewServer will load create a new Server struct loaded with the duck store
func NewServer(ds DuckStore) *Server {
	server := &Server{
		duckStore: ds,
	}

	return server
}

// RegisterHandler takes a mux and registers the server handlers onto it
// The swagger validator is specific to this server so we load it here
func (s *Server) RegisterHandler(r *chi.Mux) {
	strictHandler := NewStrictHandler(s, nil)
	r.Use(withSwaggerValidate())

	HandlerFromMux(strictHandler, r)
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

package api

import (
	"context"
	"net/http"
)

// Guarantees our Server still implements the generated StrictServerInterface.
// See: https://dev.to/kittipat1413/checking-if-a-type-satisfies-an-interface-in-go-432n
var _ StrictServerInterface = (*Server)(nil)

// DuckStore describes the methods Server needs from a store.
// Go interfaces usually live where they are used instead of where they are implemented.
type DuckStore interface {
	ListDucks(ctx context.Context) ([]RubberDuck, error)
	CreateDuck(ctx context.Context, duck DuckRequest) (RubberDuck, error)
}

// Server holds our handlers and duckStore.
type Server struct {
	duckStore DuckStore
}

// NewServer loads our duck store into a new Server.
func NewServer(duckStore DuckStore) *Server {
	return &Server{duckStore: duckStore}
}

// Handler loads Server into the generated routes and adds our OpenAPI validator.
// This is also where API-specific middleware can go.
func (s *Server) Handler() http.Handler {
	strictHandler := NewStrictHandler(s, nil)
	handler := HandlerFromMux(strictHandler, http.NewServeMux())
	return withOpenAPIValidation()(handler)
}

// oapi-codegen's strict server gives us typed requests and only lets us return
// responses defined for that operation in our OpenAPI spec. NewStrictHandler
// wraps these methods in the usual net/http handlers. If you're curious, look
// for StrictServerInterface and strictHandler in server.gen.go.

// GetDucks returns all the ducks in our store.
func (s *Server) GetDucks(ctx context.Context, _ GetDucksRequestObject) (GetDucksResponseObject, error) {
	ducks, err := s.duckStore.ListDucks(ctx)
	if err != nil {
		return GetDucks500JSONResponse{
			Code:    http.StatusInternalServerError,
			Message: "Unable to retrieve ducks.",
		}, nil
	}

	return GetDucks200JSONResponse{Ducks: ducks}, nil
}

// CreateDuck stores and returns a new duck.
func (s *Server) CreateDuck(ctx context.Context, request CreateDuckRequestObject) (CreateDuckResponseObject, error) {
	duck, err := s.duckStore.CreateDuck(ctx, *request.Body)
	if err != nil {
		return CreateDuck500JSONResponse{
			Code:    http.StatusInternalServerError,
			Message: "Unable to store duck.",
		}, nil
	}

	return CreateDuck201JSONResponse(duck), nil
}

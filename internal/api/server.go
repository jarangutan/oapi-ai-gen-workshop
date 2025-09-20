package api

import (
	"context"

	"github.com/go-chi/chi/v5"
)

// We'll get to this later ;-)
var _ StrictServerInterface = (*Server)(nil)

// Define an interface that describes what
// is needed by the server to interact with the database

// Make a struct called Server that holds our duckStore

// Make a NewServer func to create a new Server struct loaded
// with our duck store and return a pointer to this new server

// This one is a freebie! 
// RegisterHandler takes a mux and registers the server handlers onto it
// The swagger (OpenAPI) validator is specific to this api so we load it here
func (s *Server) RegisterHandler(mux *chi.Mux) {
	strictHandler := NewStrictHandler(s, nil)
	mux.Use(withSwaggerValidate())

	HandlerFromMux(strictHandler, mux)
}

// Implement the strict handlers made by oapi codegen

// GetDucks

// CreateDuck
func (s *Server) CreateDuck(ctx context.Context, request CreateDuckRequestObject) (CreateDuckResponseObject, error) {
	return nil, nil
}

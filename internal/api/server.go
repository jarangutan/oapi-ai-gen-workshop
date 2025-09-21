package api

import (
	"context"
	"fmt"

	"github.com/go-chi/chi/v5"
)

// Guarantees our Server adheres to the StrictServerInterface
// see https://dev.to/kittipat1413/checking-if-a-type-satisfies-an-interface-in-go-432n
var _ StrictServerInterface = (*Server)(nil)

// DuckStore interface describes the methods needed by the Server for a store
// Go idiom is to declare interfaces where they are used. If you repeat it a lot,
// then consider exporting it or moving it to a neutral package
// See standard library io -> net -> net/http for an idea of exporting interfaces in practice
type DuckStore interface {
	GetDucks(ctx context.Context) ([]RubberDuck, error)
	CreateDuck(ctx context.Context, duck NewRubberDuck) (RubberDuck, error)
}

// Server holds our handlers and our duckStore
type Server struct {
	duckStore DuckStore
}

// NewServer will create a new Server struct loaded with a duck store that implements the DuckStore interfacce
func NewServer(ds DuckStore) *Server {
	server := &Server{
		duckStore: ds,
	}

	return server
}

// RegisterHandler takes a mux and registers the server handlers onto it
// The swagger (OpenAPI) validator is specific to this api so we load it here
// You can also add more handler specific middlewares here!
func (s *Server) RegisterHandler(mux *chi.Mux) {
	strictHandler := NewStrictHandler(s, nil)
	mux.Use(withSwaggerValidate())

	HandlerFromMux(strictHandler, mux)
}

// GetDucks here implements the interface method `GetDucks` from the StrictServerInterface
// This looks very different from your usual Go handler which looks like:
// `GetDucks(w http.ResponseWriter, r *http.Request)`
//
// In actuality, the handler still looks like the one above! But Oapi Codegen has changed the internals
// of this handler function to wrap the GetDucks from your server passed in through the NewStrictHandler function.
// By doing so, it can take care of giving you the request context as ctx, loading params and unmarshalling the request
// body into the GetDucksRequestObject, and building a strict interface around the response so you can only
// return types that this method expects you to return.
//
// If you are curious, go to server.gen.go and do a search for "siw *ServerInterfaceWrapper" to see the wrapper in action
//
// If you want less magic, you can swap to the regular go handler interface also in server.gen.go (search for "type ServerInterface interface")
func (s *Server) GetDucks(ctx context.Context, request GetDucksRequestObject) (GetDucksResponseObject, error) {
	ducks, err := s.duckStore.GetDucks(ctx)
	if err != nil {
		return GetDucks500JSONResponse{
			Code:    500,
			Message: fmt.Sprintf("failed to get ducks: %s", err), // in prod, don't ever let a user get your internal erros :^)
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

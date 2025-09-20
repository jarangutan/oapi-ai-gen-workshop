package api

// This file is a freebie! We'll walk through it together

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	middleware "github.com/oapi-codegen/nethttp-middleware"
)

// withSwaggerValidate will prevent bad requests that don't conform to our OpenAPI schema
// from hitting our handlers
func withSwaggerValidate() func(http.Handler) http.Handler {
	swagger, err := GetSwagger()
	if err != nil {
		// This should never error
		panic("there was an error getting the swagger")
	}

	// Clear out the servers array in the swagger spec. It is recommended to do this so that it skips validating
	// that server names match.
	swagger.Servers = nil

	// Registering our own ErrorHandler to conform to our Error schema
	f := middleware.OapiRequestValidatorWithOptions(swagger, &middleware.Options{
		ErrorHandlerWithOpts: func(ctx context.Context, err error, w http.ResponseWriter, r *http.Request, opts middleware.ErrorHandlerOpts) {
			w.WriteHeader(opts.StatusCode)
			// Encode can fail but we control what is encoded so ignoring the error :-)
			_ = json.NewEncoder(w).Encode(
				Error{
					Code:    400,
					Message: fmt.Sprint("bad request:", err.Error()),
				},
			)
		},
	})
	// Middleware pattern: https://www.alexedwards.net/blog/making-and-using-middleware
	return func(next http.Handler) http.Handler {
		return f(next)
	}
}

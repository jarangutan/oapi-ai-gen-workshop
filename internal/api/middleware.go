package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	middleware "github.com/oapi-codegen/nethttp-middleware"
)

// This file is a freebie! We'll walk through it together

// withOpenAPIValidation prevents requests that don't match our OpenAPI spec
// from hitting our handlers.
// Middleware pattern: https://www.alexedwards.net/blog/making-and-using-middleware
func withOpenAPIValidation() func(http.Handler) http.Handler {
	spec, err := GetSpec()
	if err != nil {
		panic(fmt.Errorf("load OpenAPI spec: %w", err))
	}

	// Clear out the servers array so the validator doesn't require requests to use
	// the exact server URL from our spec. Tests and deployments may use another host.
	spec.Servers = nil

	return middleware.OapiRequestValidatorWithOptions(spec, &middleware.Options{
		// Register our own ErrorHandler so validation errors match our Error schema.
		ErrorHandlerWithOpts: func(_ context.Context, err error, w http.ResponseWriter, _ *http.Request, opts middleware.ErrorHandlerOpts) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(opts.StatusCode)
			// Encode can fail but we control what is encoded so ignoring the error :-)
			_ = json.NewEncoder(w).Encode(Error{
				Code:    opts.StatusCode,
				Message: fmt.Sprintf("bad request: %s", err),
			})
		},
	})
}

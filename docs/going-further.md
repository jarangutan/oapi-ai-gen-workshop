# Going further

Want to keep playing with the project after the workshop?

The advanced examples below are on the `main` branch.

## More Go

- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)

## Advanced HTTP server

Run:

```bash
go run cmd/api-advanced/main.go
```

This version adds a configurable port, timeouts, and graceful shutdown.

- [Go HTTP server timeouts](https://blog.cloudflare.com/exposing-go-on-the-internet/)
- [Graceful shutdown in Go](https://victoriametrics.com/blog/go-graceful-shutdown/)

## SQLite

Run:

```bash
go run cmd/api-sqlite/main.go
```

This swaps the in-memory store for SQLite and shows how to keep database types
separate from API types.

The example uses `go-sqlite3`, so you will need CGO and a C compiler.

## Testing with real dependencies

The `DuckStore` interface lets us use a fake store in a test. When possible, I
prefer testing with the real thing.

- [Testcontainers](https://testcontainers.com/)

## What did oapi-codegen make?

The strict handler signatures look different from the usual
`func(http.ResponseWriter, *http.Request)`, but the generated code wraps them
in regular `net/http` handlers.

- [An introduction to handlers and servemuxes in Go](https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go)
- [Making and using middleware](https://www.alexedwards.net/blog/making-and-using-middleware)

## Code-first APIs

On production APIs, it's a lot easier to work out the spec first. Other teams
can review it, and you can build mocks and clients before committing to an
implementation.

You can use code-first libraries like [Swaggo](https://github.com/swaggo/swag)
or [Huma](https://huma.rocks/) for smaller projects. The larger the project gets,
the more painful it becomes to shift the API design, especially between versions.

## Code generation

Do not edit `*.gen.go` by hand. Change the OpenAPI spec or generator config and
run `go generate ./...`.

- [`oapi-codegen`](https://github.com/oapi-codegen/oapi-codegen)
- [Go's `generate` command](https://go.dev/blog/generate)

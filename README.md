# Build an API in Go using Copilot and CodeGen

This is a workshop built for ShellHacks 2026.

- [Workshop Notes](./docs/workshop/notes.md)

## Setup

> We recommend using Go 1.26 or higher

Clone this repo and start in the base branch:

```bash
git clone https://github.com/jarangutan/oapi-ai-gen-workshop.git
cd oapi-ai-gen-workshop
git checkout base
go mod tidy
```

Generate the API and make sure everything builds:

```bash
go generate ./...
go mod tidy
go build ./...
```

## Run sample

```bash
# On Mac/Linux, `make run` does the same thing.
go run cmd/api/main.go
```

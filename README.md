# Build an API in Go using Copilot and CodeGen

This is a workshop built for ShellHacks 2025

- [Workshop Notes](./docs/workshop/notes.md)

## Setup

> We recommend using Go 1.24.2 or higher

Clone this repo, start in the base branch, and install dependencies

```bash
git clone https://github.com/jarangutan/oapi-ai-gen-workshop.git
cd oapi-ai-gen-workshop
git checkout base
go mod tidy
```

## Run sample

```bash
go mod tidy

## on Mac/Linux, you can use the make commands found in Makefile
## make run
go run cmd/api/main.go
```

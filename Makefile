# Simple Makefile for a Go project

# Build the application
all: build test

build:
	@echo "Building..."
	@CGO_ENABLED=1 GOOS=linux go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Run generators
gen:
	@echo "Generating"
	@go generate ./... 

# Live Reload
watch:
	@if command -v air > /dev/null; then \
            echo "Watching...";\
            air; \
						make clean; \
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@v1.62.0; \
                echo "Watching...";\
                air; \
								make clean; \
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

# Presentation
presi-watch:
	@npx @marp-team/marp-cli@latest --html --template bespoke -w docs/workshop/slides.md

.PHONY: all build run test clean watch presi-watch

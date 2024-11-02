# Makefile for Go Library

# Variables
PKG=./...

# Run tests
test:
	go test $(PKG) -v

# Run lint check (ensure golint is installed: go install golang.org/x/lint/golint@latest)
lint:
	golangci-lint run $(PKG)

# Format the code
fmt:
	go fmt $(PKG)

# Tidy up go.mod and go.sum
tidy:
	go mod tidy

# Clean up any generated or temporary files (modify as needed)
clean:
	rm -rf ./bin ./coverage.out

# Generate documentation (optional)
docs:
	godoc -http=:6060

# Build the library (e.g., for benchmarking or testing compiled output)
build:
	go build $(PKG)

# Phony targets to prevent conflicts with file names
.PHONY: test lint fmt tidy clean docs build

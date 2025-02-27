BIN := bin

.PHONY: help
help: ## Show help message.
	@printf "Usage:\n"
	@printf "  make <target>\n\n"
	@printf "Targets:\n"
	@perl -nle'print $& if m{^[a-zA-Z0-9_-]+:.*?## .*$$}' $(MAKEFILE_LIST) | \
		sort | \
		awk 'BEGIN {FS = ":.*?## "}; \
		{printf "  %-18s %s\n", $$1, $$2}'

.PHONY: test
test: ## Run tests.
	go test -v ./...

.PHONY: lint
lint: ## Run linter.
	golangci-lint run ./...

.PHONY: clean
clean: ## Clean up any generated or temporary files (modify as needed)
	rm -rf ./bin ./coverage.out

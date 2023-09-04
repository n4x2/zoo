.PHONY: servedoc test lint it

servedoc:
	@golds ./...

test:
	@go test ./... \
		-coverprofile=coverage.out \
		-race \
		-v

lint:
	@golangci-lint run ./...

it: test lint
	@echo "Tests and linting completed."

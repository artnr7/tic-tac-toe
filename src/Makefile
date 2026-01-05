.PHONY: build

build:
	@go build -v ./cmd/t03

.PHONY: test

test:
	@go test -v -race -timeout 30s ./...

.DEFAULT_GOAL: build


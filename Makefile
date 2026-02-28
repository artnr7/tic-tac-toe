.PHONY: build r run

build:
	@go build -o main cmd/tictactoe/main.go

r:
	@./main

run:
	@go run cmd/tiktok/main.go

br:
	@go build -o main cmd/tictactoe/main.go
	@./main

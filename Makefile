.PHONY: run
run:
	go run ./cmd/bot/main.go -bot "$(shell pwd)/config/bots/telegram.json"

build:
	go build -o bin/bot cmd/bot/main.go

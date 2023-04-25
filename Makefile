.PHONY: run

build:
	go build -o ./.bin/nfc-cash cmd/nfc-cash/main.go

run: build
	./.bin/nfc-cash

.DEFAULT_GOAL := run 
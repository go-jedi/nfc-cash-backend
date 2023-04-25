.PHONY: run

build:
	go build -o ./.bin/meet-site cmd/meet-site/main.go

run: build
	./.bin/meet-site

.DEFAULT_GOAL := run
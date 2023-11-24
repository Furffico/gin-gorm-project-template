GO=go

.PHONY: run build dev default;

default: build;

dev:
	$(GO) run ./internal/app/main.go --config ./config/dev.toml

run:
	$(GO) run ./internal/app/main.go

build:
	$(GO) build -o ./dist/server.exe -ldflags='-w -s' ./internal/app/main.go
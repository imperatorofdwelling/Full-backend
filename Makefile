run: build
	./bin/app

build: wire
	go build -o bin/app cmd/app/main.go
wire:
	wire ./internal/di
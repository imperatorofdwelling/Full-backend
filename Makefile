run: build
	./bin/app
build: swag
	go build -o bin/app cmd/app/main.go
swag: wire
	swag init --md ./docs --parseInternal  --parseDependency --parseDepth 2 -g cmd/app/main.go
wire:
	wire ./internal/di
migration-create:
	migrate create -ext sql -dir .\cmd\migrator\migrations -seq $(filter-out $@,$(MAKECMDGOALS))
migrate-up:
	go run cmd/migrator/main.go up
migrate-down:
	go run cmd/migrator/main.go down
docker:
	docker compose -p iod up --build -d
test:
	go test ./internal/api/handler/...
mock:
	go generate ./internal/domain/interfaces/...
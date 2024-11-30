run:
	./bin/app
build:
	go build -o bin/app cmd/app/main.go
swag:
	swag init --md ./docs --parseInternal  --parseDependency --parseDepth 2 -g cmd/app/main.go
wire:
	wire ./internal/di
migration-create:
	migrate create -ext sql -dir .\cmd\migrator\migrations -seq $(filter-out $@,$(MAKECMDGOALS))
migrate-up:
	go run cmd/migrator/main.go up
migrate-down:
	go run cmd/migrator/main.go down
docker-local:
	docker compose -f ./local.docker-compose.yml -p iod up --build -d
docker-dev:
	docker compose -f ./dev.docker-compose.yml -p iod up --build -d
test:
	go test ./internal/api/handler/...
mock:
	go generate ./internal/domain/interfaces/...
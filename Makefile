run-local:
	./bin/app -env='local'
run-dev:
	./bin/app -env='dev'
run-prod:
	./bin/app -env='prod'
build:
	go build -o bin/app cmd/app/main.go
swag:
	swag init --exclude docker,nginx,assets,pkg --md ./docs --parseInternal --parseDependency --parseDepth 2 -g cmd/app/main.go
wire:
	wire ./internal/di
migration-create:
	migrate create -ext sql -dir .\cmd\migrator\migrations -seq $(filter-out $@,$(MAKECMDGOALS))
migrate-up:
	go run cmd/migrator/main.go up
migrate-down:
	go run cmd/migrator/main.go down
docker-local: wire swag
	docker compose --env-file ./.env.local -f ./local.docker-compose.yml -p iod up --build -d
docker-dev:
	docker compose --env-file ./.env.dev -f ./dev.docker-compose.yml -p iod up --build -d
test:
	go test ./internal/api/handler/...
mock:
	go generate ./internal/domain/interfaces/...
run: build
	./bin/app
build: swag
	go build -o bin/app cmd/app/main.go
swag: wire
	swag init --md ./docs --parseInternal  --parseDependency --parseDepth 2 -g cmd/app/main.go
wire:
	wire ./internal/di
bjiake-wire-swag:
	swag init -g cmd/app/main.go
	google-wire ./internal/di
migration-create:
	migrate create -ext sql -dir .\cmd\migrator\migrations -seq $(filter-out $@,$(MAKECMDGOALS))
migrate-up:
	go run cmd/migrator/main.go up
migrate-down:
	go run cmd/migrator/main.go down
docker:
	docker-compose up --build
test:
	go test -cover ./internal/api/handler/...

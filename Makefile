run: build
	./bin/app
build: wire
	go build -o bin/app cmd/app/main.go
wire: swag
	wire ./internal/di
swag:
	swag init -g cmd/app/main.go
migration-create:
	migrate create -ext sql -dir .\cmd\migrator\migrations -seq $(filter-out $@,$(MAKECMDGOALS))
migrate-up:
	go run cmd/migrator/main.go up
migrate-down:
	go run cmd/migrator/main.go down
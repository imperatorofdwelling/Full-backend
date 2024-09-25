FROM golang:1.23.1

RUN go install github.com/google/wire/cmd/wire@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/migrator ./cmd/migrator/main.go
RUN ./bin/migrator up

RUN make build

ENTRYPOINT ["/app/bin/app"]

EXPOSE 8080
FROM golang:alpine

RUN apk add --no-cache bash

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/migrator ./cmd/migrator/main.go
RUN ./bin/migrator

RUN go build -o ./bin/app ./cmd/app/main.go

ENTRYPOINT ["/app/bin/app"]

EXPOSE 8080
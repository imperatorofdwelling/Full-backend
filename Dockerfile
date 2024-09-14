FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/app ./cmd/app/main.go

ENTRYPOINT ["/app/bin/app"]

EXPOSE 8080
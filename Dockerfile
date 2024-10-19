FROM golang:1.23.2

RUN go install github.com/google/wire/cmd/wire@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

ENV PATH="/go/bin:${PATH}"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod tidy

COPY . .

RUN make build

ENTRYPOINT ["/app/bin/app"]

EXPOSE 8080
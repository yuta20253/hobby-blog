FROM golang:1.26

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

CMD ["go", "run", "./cmd/server"]

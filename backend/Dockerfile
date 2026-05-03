FROM golang:1.26

WORKDIR /app

COPY . .

RUN go mod tidy

CMD ["go", "run", "./cmd/server"]

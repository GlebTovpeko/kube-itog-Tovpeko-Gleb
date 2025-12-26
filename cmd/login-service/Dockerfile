FROM golang:1.24.9-alpine3.22 AS builder

WORKDIR /app

COPY ./cmd/login-service .
COPY go.mod .

RUN go mod tidy

RUN go build -o login-service /app/main.go

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/login-service .

CMD [ "/app/login-service"]
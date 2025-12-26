FROM golang:1.24.9-alpine3.22 AS builder

WORKDIR /app

COPY ./cmd/root-service .
COPY go.mod .

RUN go mod tidy

RUN go build -o root-service /app/main.go

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/root-service .

CMD [ "/app/root-service"]
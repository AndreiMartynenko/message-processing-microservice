FROM golang:1.21.7-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main /app/main

ENV DB_HOST=db
ENV DB_PORT=5432
ENV DB_USER=m-user
ENV DB_PASSWORD=m-password
ENV DB_NAME=messages
ENV KAFKA_BROKER=kafka:9092

CMD ["./main"]


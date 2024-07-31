#FROM golang:1.21.7-alpine AS builder
#
#COPY . /github.com/AndreiMartynenko/message-processing-microservice/source/
#WORKDIR /github.com/AndreiMartynenko/message-processing-microservice/source/
#
#RUN go mod download
#RUN go build -o ./bin/crud_server cmd/server/main.go
#
#FROM alpine:latest
#
#WORKDIR /root/
#
#COPY --from=builder /github.com/AndreiMartynenko/message-processing-microservice/bin/crud_server .
#
#
#CMD ["./crud_server"]


FROM golang:1.21.7-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

#RUN go build -o ./bin/consumer cmd/kafka/consumer/main.go
#RUN go build -o ./bin/producer cmd/kafka/producer/main.go
RUN go build -o ./main.go

FROM alpine:latest



#COPY --from=builder /app/bin/consumer .
#COPY --from=builder /app/bin/producer .

COPY --from=builder /app/main /app/main

ENV DB_HOST=db
ENV DB_PORT=5432
ENV DB_USER=m-user
ENV DB_PASSWORD=m-password
ENV DB_NAME=messages
ENV KAFKA_BROKER=kafka:9092

WORKDIR /app

#ENTRYPOINT ["/main"]

CMD ["./main"]


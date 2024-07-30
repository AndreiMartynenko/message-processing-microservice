# Message processing microservice

## Review

This microservice is written in Go and performs the following tasks:
- Receives messages via HTTP API.
- Saves them in PostgreSQL.
- Sends them to Kafka for further processing.
- Marks processed messages.
- Provides an API for obtaining statistics on processed messages.

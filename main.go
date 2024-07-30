package main

import (
	"github.com/AndreiMartynenko/message-processing-microservice/src"
	"log"
	"net/http"
)

func main() {
	// Initialize Kafka and PostgreSQL clients
	src.InitializeKafka()
	src.InitializePostgres()

	// Set up HTTP routes
	http.HandleFunc("/messages", src.MessageHandler)
	http.HandleFunc("/stats", src.StatsHandler)

	// Start the server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

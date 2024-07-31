package main

import (
	"database/sql"
	"github.com/AndreiMartynenko/message-processing-microservice/cmd/kafka"
	"github.com/AndreiMartynenko/message-processing-microservice/postgres"
	"github.com/AndreiMartynenko/message-processing-microservice/src"
	"github.com/IBM/sarama"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var db *sql.DB
var kafkaProducer sarama.SyncProducer

func main() {
	var err error
	db, err = postgres.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	kafkaProducer, err = kafka.NewSyncProducer([]string{"localhost:9092"})
	if err != nil {
		log.Fatalf("Failed to connect to Kafka: %v", err)
	}
	defer kafkaProducer.Close()

	router := mux.NewRouter()

	// Define the API routes
	router.HandleFunc("/messages", src.HandlePostMessage).Methods("POST")
	router.HandleFunc("/statistics", src.HandleGetStatistics).Methods("GET")

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

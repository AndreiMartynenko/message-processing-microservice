//package main
//
//import (
//	"github.com/AndreiMartynenko/message-processing-microservice/src"
//	"log"
//	"net/http"
//)
//
//func main() {
//	// Initialize Kafka and PostgreSQL clients
//	src.InitializeKafka()
//	src.InitializePostgres()
//
//	// Set up HTTP routes
//	http.HandleFunc("/messages", src.MessageHandler)
//	http.HandleFunc("/stats", src.StatsHandler)
//
//	// Start the server
//	log.Println("Starting server on :8080")
//	log.Fatal(http.ListenAndServe(":8080", nil))
//}

package main

import (
	"log"
	"net/http"
	"github.com/AndreiMartynenko/message-processing-microservice/src"

	"github.com/gorilla/mux"
)

func main() {
	var err error
	db, err = initDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	kafkaProducer, err = initKafka()
	if err != nil {
		log.Fatalf("Failed to connect to Kafka: %v", err)
	}
	defer kafkaProducer.Close()

	router := mux.NewRouter()

	// Define the API routes
	router.HandleFunc("/messages", src.handlePostMessage).Methods("POST")
	router.HandleFunc("/statistics", src.handleGetStatistics).Methods("GET")

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

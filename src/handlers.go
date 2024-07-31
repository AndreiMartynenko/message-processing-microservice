package src

import (
	"encoding/json"
	"github.com/AndreiMartynenko/message-processing-microservice/cmd/kafka"
	"github.com/AndreiMartynenko/message-processing-microservice/postgres"
	"net/http"
)

func HandlePostMessage(w http.ResponseWriter, r *http.Request) {
	var msg struct {
		Content string `json:"content"`
	}

	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := postgres.db.Exec(postgres.InsertMessage, msg.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := kafka.SendMessageToKafka("test-topic", msg.Content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func HandleGetStatistics(w http.ResponseWriter, r *http.Request) {
	rows, err := postgres.db.Query(postgres.GetStatistics)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	stats := make(map[string]int)
	for rows.Next() {
		var status string
		var count int
		if err := rows.Scan(&status, &count); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		stats[status] = count
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

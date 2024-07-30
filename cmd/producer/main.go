package main

import (
	"github.com/IBM/sarama"
	"github.com/brianvoe/gofakeit"
	"log"
)

const (
	brokerAddress = "localhost:9092"
	topicName     = "test-topic"
)

func main() {
	producer, err := newSyncProducer([]string{brokerAddress})
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("Failed to close producer: %s", err)
		}
	}()

	msg := &sarama.ProducerMessage{
		Topic: topicName,
		Value: sarama.StringEncoder(gofakeit.StreetName()),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to send message: %v\n", err.Error())
		return
	}

	log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
}

// brokerList []string for fault tolerance and load balancing
func newSyncProducer(brokerList []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

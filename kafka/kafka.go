package kafka

import (
	"context"
	"log"
	"github.com/segmentio/kafka-go"
	"booking-api/config"
)

func Produce(topic string, message []byte) {
	w := &kafka.Writer{
		Addr:     kafka.TCP(config.KafkaBroker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Value: message,
		},
	)
	if err != nil {
		log.Fatalf("Failed to produce message: %v", err)
	}
}

func Consume(topic string) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{config.KafkaBroker},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}
		log.Printf("Received message: %s", string(m.Value))
		// Process the message (e.g., update database, notify users, etc.)
	}
}
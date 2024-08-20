package worker

import (
    "context"
    "log"
    "encoding/json"

    "github.com/segmentio/kafka-go"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "booking-api/models"
    "booking-api/config"
)

func StartWorker() {
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers:   []string{config.KafkaBroker},
        Topic:     "bookings",
        Partition: 0,
        MinBytes:  10e3, // 10KB
        MaxBytes:  10e6, // 10MB
    })

    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoURI))
    if err != nil {
        log.Fatal(err)
    }
    appointmentCollection := client.Database("clinic").Collection("appointments")

    for {
        m, err := r.ReadMessage(context.Background())
        if err != nil {
            log.Printf("Error reading message: %v", err)
            continue
        }

        var appointment models.Appointment
        err = json.Unmarshal(m.Value, &appointment)
        if err != nil {
            log.Printf("Failed to unmarshal message: %v", err)
            continue
        }

        _, err = appointmentCollection.InsertOne(context.Background(), appointment)
        if err != nil {
            log.Printf("Failed to insert appointment: %v", err)
            continue
        }

        log.Printf("Successfully booked appointment: %s", appointment.ID.Hex())
    }
}
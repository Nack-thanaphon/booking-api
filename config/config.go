package config

import (
	"os"
)

var (
	MongoURI   = os.Getenv("MONGODB_URI")
	KafkaBroker = os.Getenv("KAFKA_BROKER")
)
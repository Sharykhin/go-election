package mongodb

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(connectionUrl string) *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionUrl))
	if err != nil {
		log.Fatalf("Failed to create mongodb client: %v", err)
	}

	return client
}

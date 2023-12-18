// mongo.go
package handlers

import (
    "context"
    "log"
    "os"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func init() {
    mongoURL := os.Getenv("MONGO_URL")
    if mongoURL == "" {
        log.Fatal("MONGO_URL is not set in the environment")
    }

    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURL))
    if err != nil {
        log.Fatal(err)
    }

    mongoClient = client
}

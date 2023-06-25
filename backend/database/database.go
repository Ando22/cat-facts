package database

import (
	"context"
	"log"
	"time"

	"user-message/backend/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Database
)

// InitDatabase initializes the MongoDB database connection
func InitDatabase() {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.DatabaseURL))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db = client.Database(config.DatabaseName)

	log.Println("Connected to database")
}

// GetDB returns the MongoDB database instance
func GetDB() *mongo.Database {
	return db
}

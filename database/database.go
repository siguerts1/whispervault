// database/database.go
package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var secretsCollection *mongo.Collection

// InitDB initializes the MongoDB connection
func InitDB() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	secretsCollection = client.Database("whispervault").Collection("secrets")
	return nil
}

// GetSecretsCollection returns the MongoDB collection for secrets
func GetSecretsCollection() *mongo.Collection {
	return secretsCollection
}


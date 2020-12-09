package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database describe a database connection
type Database struct {
	databaseName string
	client       *mongo.Client
	Handle       *mongo.Database
}

// NewDatabase returns a new instance of Database.
func NewDatabase(connectURI string, databaseName string, username string, password string) (*Database, error) {
	db := &Database{
		client:       nil,
		databaseName: databaseName,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(connectURI)
	// clientOptions.SetAuth(options.Credential{
	// 	Username:    username,
	// 	Password:    password,
	// 	PasswordSet: true,
	// })

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	db.client = client
	db.Handle = client.Database(databaseName)

	return db, nil
}

const (
	// Timeout operations after N seconds
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s"
)

// GetConnection - Retrieves a client to the DocumentDB
func GetConnection() (*mongo.Client, context.Context, context.CancelFunc) {

	devURI := "mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&ssl=false"
	// devURI := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(devURI))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	return client, ctx, cancel
}

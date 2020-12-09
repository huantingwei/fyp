package util

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
	Handle *mongo.Database
}

func NewDatabase() (Database, context.Context) {
	dbClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://jeff:CYyTXZahfIMkoBhk@cluster0.ngxmf.mongodb.net/Cluster0?retryWrites=true&w=majority"))
	//dbClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = dbClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to MongoDB...")
	}

	db := Database{
		Client: dbClient,
		Handle: dbClient.Database("fyp"),
	}
	fmt.Printf("db: %v\n", db)
	return db, ctx
}

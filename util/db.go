package util

import (
	"context"
	"log"
	"time"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct{
	Client *mongo.Client
	Handle *mongo.Database
}

func NewDatabase() (Database, context.Context){

	dbClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://jeff:CYyTXZahfIMkoBhk@cluster0.ngxmf.mongodb.net/Cluster0?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = dbClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Println("Connected to MongoDB...");
	}

	db := Database{
		Client: dbClient,
		Handle: dbClient.Database("fyp"),
	}

	return db, ctx;
}
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
	// dbClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://jeff:CYyTXZahfIMkoBhk@cluster0.ngxmf.mongodb.net/Cluster0?retryWrites=true&w=majority"))
	// ting's atlas
	// dbClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://admin:admin@fyp-demo-shard-00-00.uqci5.mongodb.net:27017,fyp-demo-shard-00-01.uqci5.mongodb.net:27017,fyp-demo-shard-00-02.uqci5.mongodb.net:27017/fyp?authSource=admin&retryWrites=true&w=majority"))
	// ting's windows mongo
	// dbClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"))
	// ting's ubuntu mongo
	// dbClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"))

	// docker
	dbClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"))

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
	return db, ctx
}

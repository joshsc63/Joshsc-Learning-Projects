package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort = "80"
	rpcPort = "5001"
	mongoURL = "mongodb://mongo:27017" // defined in docker-compose
	gRpcPort = "50001"
)

var client *mongo.Client

type Config struct {
	
}

func main () {
	// connect to mongoDB
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}
	
	client = mongoClient
	
	// create context to disconnect. Needed by Mongo
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	
	//close connection
	defer func () {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	} ()
}

func connectToMongo() (*mongo.Client, err) {
	// create connection options
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})
	
	// connect
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting:", err)
		return nil, err
	}
	
	return c, nil
}
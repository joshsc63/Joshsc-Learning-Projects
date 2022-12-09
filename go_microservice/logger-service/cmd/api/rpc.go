package main

import (
	"context"
	"log"
	"log-service/data"
	"time"
)

// no members
type RPCServer struct{}

// type of payload to receive from RPC
type RPCPayload struct {
	Name string
	Data string
}

// sending resp back as *string
func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	//write to mongo
	collection := client.Database("logs").Collection("logs")

	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Println(err)
		return err
	}

	// send msg back w/ resp
	*resp = "Processed payload via RPC:" + payload.Name
	return nil
}

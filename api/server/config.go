package server

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetSession Databases atlas mongodb...
func GetSession() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://alejodb:Alejandro-10@comicsmarvel-pum8a.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//defer client.Disconnect(ctx)
	return client
}

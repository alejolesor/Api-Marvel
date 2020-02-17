package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ErrorResponse ...
type ErrorResponse struct {
	StatusCode   int    `json:"statuscode"`
	ErrorMessage string `json:"message"`
}

//ConnectedDB Databases atlas mongodb...
func ConnectedDB() *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb+srv://alejodb:Alejandro-10@comicsmarvel-pum8a.mongodb.net/test?retryWrites=true&w=majority")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conectado mongo db")

	collection := client.Database("ApiMarvel").Collection("comics")

	return collection
}

//GetError ...
func GetError(err error) {
	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}
	message, _ := json.Marshal(response)
	fmt.Println(message)
}

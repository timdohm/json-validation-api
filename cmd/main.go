package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/timdohm/json-validation-api/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("tasker").Collection("tasks")

}

func main() {
	l := log.New(os.Stdout, "json-validation-api", log.LstdFlags)
	dio := handlers.NewDataIO(l, collection)
	r := mux.NewRouter()

	// POST /schema/SCHMAID
	// GET /schema/SCHEMAID
	// POST /validate/SCHMAID

	gr := r.Methods(http.MethodGet).Subrouter()
	gr.HandleFunc("/schema/{id}", dio.GetSchema)

	pr := r.Methods(http.MethodPost).Subrouter()
	pr.HandleFunc("/schema/{id}", dio.PostSchema)
	pr.HandleFunc("/validate/{id}", dio.ValidateSchema)

}

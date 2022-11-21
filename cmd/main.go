package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/timdohm/json-validation-api/database"
	"github.com/timdohm/json-validation-api/handlers"
)

func main() {
	l := log.New(os.Stdout, "json-validation-api", log.LstdFlags)
	client, ctx := database.InitDB()
	collection := client.Database("api").Collection("schema")
	defer client.Disconnect(ctx)
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
	pr.Use(dio.VerifyJSON)

	server := http.Server{
		Addr:         ":4060",
		Handler:      r,
		ErrorLog:     l,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		l.Println("Starting the server on port 4060")
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(ctx)
	l.Println("Shutting down server")
	os.Exit(0)
}

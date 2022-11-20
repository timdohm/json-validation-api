package handlers

import (
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type DataIO struct {
	l  *log.Logger
	db *mongo.Collection
}

func NewDataIO(l *log.Logger, db *mongo.Collection) *DataIO {
	return &DataIO{l, db}
}

func (db *DataIO) GetSchema(rw http.ResponseWriter, r *http.Request) {

}

func (db *DataIO) PostSchema(rw http.ResponseWriter, r *http.Request) {

}

func (db *DataIO) ValidateSchema(rw http.ResponseWriter, r *http.Request) {

}

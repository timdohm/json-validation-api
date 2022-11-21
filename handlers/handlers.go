package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/timdohm/json-validation-api/database"
	"github.com/timdohm/json-validation-api/middleware"
	"github.com/timdohm/json-validation-api/validate"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataIO struct {
	l  *log.Logger
	db *mongo.Collection
}

func NewDataIO(l *log.Logger, db *mongo.Collection) *DataIO {
	return &DataIO{l, db}
}

func (dio *DataIO) GetSchema(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// get json
	schema, err := database.GetSchema(dio.db, key)

	if err != nil {
		dio.l.Printf("Error retrieving schema: %s\n", err)
		////////////// Inform user
		return
	}

	//////////////// finish return and verification
	rw.Write([]byte(schema))

}

func (dio *DataIO) PostSchema(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	body := r.Context().Value(middleware.KeyBody{}).([]byte)

	err := database.PutSchema(dio.db, key, []byte(body))

	if err != nil {
		dio.l.Printf("Error adding schema to database: %s\n", err)
		////////////Inform user
		return

	}

	////// Inform user

}

func (dio *DataIO) ValidateSchema(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	schema, err := database.GetSchema(dio.db, key)
	body := r.Context().Value(middleware.KeyBody{}).([]byte)

	if err != nil {
		dio.l.Printf("Error retrieving schema for validation: %s\n", err)
		////////////////Inform user
		return

	}

	if err = validate.ValidateAgainstSchema(body, schema); err != nil {
		/////decision based on err type
	}

	///////write positive result

}

package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"github.com/timdohm/json-validation-api/database"
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
		dio.l.Printf("Error retrieving schema: %s\n", err.Error())
		message := "Unable to retrieve schema at this id."
		response := validate.NewReponseWithMessage("downloadSchema", key, "error", message)
		rw.WriteHeader(http.StatusNotFound)
		rw.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(response)
		if err != nil {
			dio.l.Fatalf("Error in JSON marshall: %s\n", err.Error())
		}
		rw.Write(jsonResp)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(schema)

}

func (dio *DataIO) PostSchema(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	body := r.Context().Value(KeyBody{}).([]byte)

	err := database.PutSchema(dio.db, key, []byte(body))

	if err != nil {
		dio.l.Printf("Error adding schema to database: %s\n", err.Error())
		message := "Schema already exists at this id."
		response := validate.NewReponseWithMessage("uploadSchema", key, "error", message)
		rw.WriteHeader(http.StatusConflict)
		rw.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(response)
		if err != nil {
			dio.l.Fatalf("Error in JSON marshall: %s\n", err.Error())
		}
		rw.Write(jsonResp)
		return
	}

	// Inform client
	response := validate.NewReponse("uploadSchema", key, "success")
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(response)
	if err != nil {
		dio.l.Fatalf("Error in JSON marshall: %s\n", err.Error())
	}
	rw.Write(jsonResp)

}

func (dio *DataIO) ValidateSchema(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	schema, err := database.GetSchema(dio.db, key)
	body := r.Context().Value(KeyBody{}).([]byte)

	if err != nil {
		dio.l.Printf("Error retrieving schema for validation: %s\n", err)
		message := "Unable to retrieve schema at this id."
		response := validate.NewReponseWithMessage("validateDocument", key, "error", message)
		rw.WriteHeader(http.StatusNotFound)
		rw.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(response)
		if err != nil {
			dio.l.Fatalf("Error in JSON marshall: %s\n", err.Error())
		}
		rw.Write(jsonResp)
		return

	}

	if err = validate.ValidateAgainstSchema(body, schema); err != nil {
		// decision based on err type

		switch err.(type) {
		case *jsonschema.ValidationError:
			// validation failure
			message := err.Error()
			response := validate.NewReponseWithMessage("validateDocument", key, "error", message)
			rw.WriteHeader(http.StatusOK)
			rw.Header().Set("Content-Type", "application/json")
			jsonResp, err := json.Marshal(response)
			if err != nil {
				dio.l.Fatalf("Error in JSON marshall: %s\n", err.Error())
			}
			rw.Write(jsonResp)

		default:
			dio.l.Fatalf("Fatal error encountered while validating against schema: %s\n", err.Error())
		}
		return
	}

	// write successful result
	response := validate.NewReponse("validateDocument", key, "success")
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(response)
	if err != nil {
		dio.l.Fatalf("Error in JSON marshall: %s\n", err.Error())
	}
	rw.Write(jsonResp)
}

package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/timdohm/json-validation-api/validate"
)

type KeyBody struct{}

func (dio *DataIO) VerifyJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		// Verify valid JSON
		body, err := io.ReadAll(r.Body)
		if err != nil {
			dio.l.Fatalf("Unable to read request body: %s\n", err.Error())

		}

		if !json.Valid(body) {
			response := validate.NewReponseWithMessage("uploadSchema", vars["id"], "error", "Invalid JSON")
			rw.WriteHeader(http.StatusBadRequest)
			rw.Header().Set("Content-Type", "application/json")
			jsonResp, err := json.Marshal(response)
			if err != nil {
				dio.l.Fatalf("Error in JSON marshall: %s\n", err.Error())
			}
			rw.Write(jsonResp)
			return
		}

		// pass the valid json to the context
		ctx := context.WithValue(r.Context(), KeyBody{}, body)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}

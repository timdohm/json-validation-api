package middleware

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type KeyBody struct{}

func VerifyJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		// Verify valid JSON
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// report error

			return
		}

		if !json.Valid(body) {
			// report error
			return
		}

		// pass the valid json to the context
		ctx := context.WithValue(r.Context(), KeyBody{}, body)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}

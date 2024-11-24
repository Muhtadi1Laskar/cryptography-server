package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HandleEmptyJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read the request body
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusBadRequest)
			return
		}

		fmt.Println(string(reqBody))

		// Check if body is empty or only whitespace
		if len(reqBody) == 0 || string(reqBody) == "{}" {
			http.Error(w, "Request body cannot be empty", http.StatusBadRequest)
			return
		}

		// Validate JSON format
		var body map[string]interface{}
		if err := json.Unmarshal(reqBody, &body); err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		// Reassign the request body for the next handler
		r.Body = io.NopCloser(bytes.NewReader(reqBody))

		// Pass control to the next handler
		next.ServeHTTP(w, r)
	})
}


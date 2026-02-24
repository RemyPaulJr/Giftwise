package main

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func checkError(w http.ResponseWriter, status int, message string) {
	// Set HTTP Header's content-type to json
	w.Header().Set("Content-Type", "application/json")
	// Write the Header Status Code as the status parameter
	w.WriteHeader(status)

	pER := ErrorResponse{
		Message: message,
	}

	json.NewEncoder(w).Encode(pER)

}

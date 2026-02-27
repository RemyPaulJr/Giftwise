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

	pErrorResponse := ErrorResponse{
		Message: message,
	}

	json.NewEncoder(w).Encode(pErrorResponse)

}

func writeJSON(w http.ResponseWriter, status int, data any) {
	// Set HTTP Header's content-type to json
	w.Header().Set("Content-Type", "application/json")
	// Write the Header Status Code as the status parameter
	w.WriteHeader(status)
	// Encode data that is passed to JSON
	json.NewEncoder(w).Encode(data)
}

package utils

import (
	"encoding/json"
	"net/http"
)

// APIError represents an API error response
type APIError struct {
	Message string `json:"message"`
}

// RespondWithError sends an error response to the client
func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(APIError{Message: message})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

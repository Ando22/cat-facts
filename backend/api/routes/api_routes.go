package routes

import (
	"user-message/backend/api/handlers"

	"github.com/gorilla/mux"
)

// RegisterAPIRoutes registers API routes
func RegisterAPIRoutes(router *mux.Router) {
	router.HandleFunc("/api/messages", handlers.GetMessages).Methods("GET")
	router.HandleFunc("/api/messages", handlers.CreateMessage).Methods("POST")
}

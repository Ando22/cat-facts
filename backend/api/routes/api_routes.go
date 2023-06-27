package routes

import (
	"github.com/Ando22/user-message/backend/api/handlers"
	"github.com/gorilla/mux"
)

// RegisterAPIRoutes registers API routes
func RegisterAPIRoutes(router *mux.Router) {
	router.HandleFunc("/api/facts", handlers.GetFacts).Methods("GET")
}

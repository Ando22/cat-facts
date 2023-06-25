package main

import (
	"log"
	"net/http"

	"user-message/backend/api/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	routes.RegisterAPIRoutes(router)

	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := corsOptions.Handler(router)

	log.Println("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}

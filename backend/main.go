package main

import (
	"log"
	"net/http"

	"github.com/Ando22/user-message/backend/api/routes"
	"github.com/Ando22/user-message/backend/database"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	log.Println("Initialising db")
	// database.InitDatabase()
	err := database.InitDatabase()
	if err != nil {
		panic(err.Error())
	}

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

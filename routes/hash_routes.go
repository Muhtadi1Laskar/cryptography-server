package routes

import (
	"cryptographyServer/handlers"
	"cryptographyServer/middlewares"

	"github.com/gorilla/mux"
)

func HashRoutes(router *mux.Router) {
	router.HandleFunc("/hash-list", handlers.ShowHashList).Methods("GET")
	router.Use(middlewares.HandleEmptyJSON)
	router.HandleFunc("/hash", handlers.HashData).Methods("POST")
}
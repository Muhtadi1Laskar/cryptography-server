package routes

import (
	"cryptographyServer/handlers"
	"cryptographyServer/middlewares"

	"github.com/gorilla/mux"
)

func HashRoutes(router *mux.Router) {
	hashRouter := router.PathPrefix("/hash").Subrouter()
	hashRouter.Use(middlewares.HandleEmptyJSON)
	hashRouter.HandleFunc("", handlers.HashData).Methods("POST")

	router.HandleFunc("/hash-list", handlers.ShowHashList).Methods("GET")
}
package routes

import (
	"cryptographyServer/handlers"

	"github.com/gorilla/mux"
)

func HashRoutes(router *mux.Router) {
	router.HandleFunc("/hash", handlers.HashData).Methods("POST")
	router.HandleFunc("/hash-list", handlers.ShowHashList).Methods("GET")
}
package routes

import (
	"cryptographyServer/handlers"

	"github.com/gorilla/mux"
)

func RsaRoutes(router *mux.Router) {
	router.HandleFunc("/get-keys", handlers.GenerateKeys).Methods("GET")
}
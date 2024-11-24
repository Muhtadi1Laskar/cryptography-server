package routes

import (
	"cryptographyServer/handlers"

	"github.com/gorilla/mux"
)

func RsaRoutes(router *mux.Router) {
	router.HandleFunc("rsa/get-keys", handlers.GenerateKeys).Methods("GET")
}
package routes

import (
	"cryptographyServer/handlers"

	"github.com/gorilla/mux"
)

func HMACRoutes(router *mux.Router) {
	router.HandleFunc("/hmac/create-signature", handlers.SignMessage).Methods("POST")
}
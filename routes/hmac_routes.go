package routes

import (
	"cryptographyServer/handlers"

	"github.com/gorilla/mux"
)

func HMACRoutes(router *mux.Router) {
	router.HandleFunc("/hmac/create-signature", handlers.SignMessage).Methods("POST")
	router.HandleFunc("/hmac/verify-signature", handlers.VerifyMessage).Methods("POST")
}
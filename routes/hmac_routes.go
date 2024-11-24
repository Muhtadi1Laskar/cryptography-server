package routes

import (
	"cryptographyServer/handlers"
	"cryptographyServer/middlewares"

	"github.com/gorilla/mux"
)

func HMACRoutes(router *mux.Router) {
	router.Use(middlewares.HandleEmptyJSON)
	router.HandleFunc("/hmac/create-signature", handlers.SignMessage).Methods("POST")
	router.HandleFunc("/hmac/verify-signature", handlers.VerifyMessage).Methods("POST")
}
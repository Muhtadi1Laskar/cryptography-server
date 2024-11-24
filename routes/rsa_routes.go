package routes

import (
	"cryptographyServer/handlers"
	"cryptographyServer/middlewares"

	"github.com/gorilla/mux"
)

func RsaRoutes(router *mux.Router) {
	router.Use(middlewares.HandleEmptyJSON)
	router.HandleFunc("/rsa/get-keys", handlers.GenerateKeys).Methods("GET")
	router.HandleFunc("/rsa/encrypt", handlers.RSAEncryptMessage).Methods("POST")
	router.HandleFunc("/rsa/decrypt", handlers.RSADecryptMessage).Methods("POST")
}
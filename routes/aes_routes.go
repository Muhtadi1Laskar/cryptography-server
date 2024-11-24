package routes

import (
	"cryptographyServer/handlers"
	"cryptographyServer/middlewares"

	"github.com/gorilla/mux"
)

func AesCipherRoutes(router *mux.Router) {
	router.Use(middlewares.HandleEmptyJSON)
	router.HandleFunc("/aes-cipher/encrypt", handlers.AesEncrypt).Methods("POST")
	router.HandleFunc("/aes-cipher/decrypt", handlers.AesDecrypt).Methods("POST")
}
package routes

import (
	"cryptographyServer/handlers"

	"github.com/gorilla/mux"
)

func AesCipherRoutes(router *mux.Router) {
	router.HandleFunc("/aes-cipher/encrypt", handlers.AesEncrypt).Methods("POST")
	router.HandleFunc("/aes-cipher/decrypt", handlers.AesDecrypt).Methods("POST")
}
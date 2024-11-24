package routes

import (
	"cryptographyServer/handlers"
	"cryptographyServer/middlewares"

	"github.com/gorilla/mux"
)

func AesCipherRoutes(router *mux.Router) {
	hashRouter := router.PathPrefix("/aes-cipher").Subrouter()
	hashRouter.Use(middlewares.HandleEmptyJSON)
	router.HandleFunc("/encrypt", handlers.AesEncrypt).Methods("POST")
	router.HandleFunc("/decrypt", handlers.AesDecrypt).Methods("POST")
}
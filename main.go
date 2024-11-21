package main

import (
	"cryptographyServer/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	routes.HashRoutes(router)
	routes.AesCipherRoutes(router)

	log.Println("Server running on http://localhost:500")
	http.ListenAndServe(":5000", router)
}

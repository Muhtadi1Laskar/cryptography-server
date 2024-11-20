package main

import (
	"cryptographyServer/routes"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter()
	routes.HashRoutes(router)
	routes.AesCipherRoutes(router)

	log.Println("Server running on http://localhost:500")
	http.ListenAndServe(":5000", router)
}

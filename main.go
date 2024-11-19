package main

import (
	"cryptographyServer/hashs"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func hashData (w http.ResponseWriter, r *http.Request) {
	h, _ := hashs.Hash("Hello World", "sha256")

	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(h)
}


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hashData", hashData).Methods("GET")

	fmt.Printf("Server running on port 5000\n")
	http.ListenAndServe(":5000", router)
}

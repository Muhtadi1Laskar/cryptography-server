package main

import (
	"cryptographyServer/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


// type CipherRequest struct {
// 	Cipher string `json:"cipher"`
// 	Type   string `json:"type"`
// 	Key    string `json:"key"`
// 	Data   string `json:"Data"`
// }


// func aesEncrypt(w http.ResponseWriter, r *http.Request) {
// 	reqBody, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, "Unable to read the request body", http.StatusBadRequest)
// 		return
// 	}
// 	defer r.Body.Close()

// 	var requestData CipherRequest
// 	var responseBody ResponseData
// 	if err := json.Unmarshal(reqBody, &requestData); err != nil {
// 		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
// 	}

// 	cipherText, err := authenticatedEncryption.Encrypt(requestData.Data, requestData.Key)
// 	if err != nil {
// 		responseBody = buildResponse(err.Error(), "Failed")
// 	} else {
// 		responseBody = buildResponse(cipherText, requestData.Type + " successful")
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(responseBody)
// }


func main() {
	router := mux.NewRouter()
	routes.HashRoutes(router)

	log.Println("Server running on http://localhost:500")
	http.ListenAndServe(":5000", router)
}

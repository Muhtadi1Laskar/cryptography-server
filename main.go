package main

import (
	"cryptographyServer/ciphers"
	"cryptographyServer/hashs"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type RequestData struct {
	Hash string `json:"hash"`
	Data string `json:"data"`
}

type ResponseData struct {
	Data   string `json:"Data"`
	Status string `json:"status"`
}

type HashList struct {
	List []string `json:"list"`
}

type CipherRequest struct {
	Cipher string `json:"cipher"`
	Type   string `json:"type"`
	Key    string `json:"key"`
	Data   string `json:"Data"`
}

func hashData(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var requestData RequestData
	var responseBody ResponseData
	if err := json.Unmarshal(reqBody, &requestData); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	hashedData, err := hashs.Hash(requestData.Data, requestData.Hash)
	if err != nil {
		responseBody = buildResponse(err.Error(), "Failed")
	} else {
		responseBody = buildResponse(hashedData, "Successfully hashed the data")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)
}

func showHashList(w http.ResponseWriter, r *http.Request) {
	var list []string = hashs.GetHashList()
	var response HashList = HashList{
		List: list,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func aesEncrypt(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read the request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var requestData CipherRequest
	var responseBody ResponseData
	if err := json.Unmarshal(reqBody, &requestData); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
	}

	cipherText, err := authenticatedEncryption.Encrypt(requestData.Data, requestData.Key)
	if err != nil {
		responseBody = buildResponse(err.Error(), "Failed")
	} else {
		responseBody = buildResponse(cipherText, requestData.Type + " successful")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)
}

func buildResponse(response string, status string) ResponseData {
	return ResponseData{
		Data: response,
		Status: status,
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hashData", hashData).Methods("POST")
	router.HandleFunc("/cipher/aes", aesEncrypt).Methods("POST")
	router.HandleFunc("/getHashes", showHashList).Methods("GET")

	fmt.Printf("Server running on port 5000\n")
	http.ListenAndServe(":5000", router)
}

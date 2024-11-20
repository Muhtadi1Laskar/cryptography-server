package main

import (
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
	HashedData string `json:"hashedData"`
	Status     string `json:"status"`
}

type HashList struct {
	List []string `json:"list"`
}

func hashData(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var requestData RequestData
	var response ResponseData
	if err := json.Unmarshal(reqBody, &requestData); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	hashedData, err := hashs.Hash(requestData.Data, requestData.Hash)
	if err != nil {
		response = ResponseData{
			HashedData: err.Error(),
			Status:     "Failed",
		}
	} else {
		response = ResponseData{
			HashedData: hashedData,
			Status:     "Success",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func showHashList(w http.ResponseWriter, r *http.Request) {
	var list []string = hashs.GetHashList()
	var response HashList = HashList{
		List: list,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hashData", hashData).Methods("POST")
	router.HandleFunc("/getHashes", showHashList).Methods("GET")

	fmt.Printf("Server running on port 5000\n")
	http.ListenAndServe(":5000", router)
}

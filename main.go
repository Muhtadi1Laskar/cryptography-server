package main

import (
	"cryptographyServer/hashs"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/gorilla/mux"
)

type RequestData struct {
	Hash string `json:"hash"`
	Data string `json:"data"`
}

type ResponseData struct {
	HashedData string `json:"hashedData"`
	Status string `json:"status"`
}

func hashData (w http.ResponseWriter, r *http.Request) {
	reqBody, err:= io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var requestData RequestData
	if err := json.Unmarshal(reqBody, &requestData); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return 
	}

	hashedData, _ := hashs.Hash(requestData.Data, requestData.Hash);

	response := ResponseData{
		HashedData: hashedData,
		Status: "Success",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hashData", hashData).Methods("POST")

	fmt.Printf("Server running on port 5000\n")
	http.ListenAndServe(":5000", router)
}

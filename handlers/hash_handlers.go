package handlers

import (
	"io"
	"net/http"
	"encoding/json"
	"cryptographyServer/hashs"
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

func HashData(w http.ResponseWriter, r *http.Request) {
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

func ShowHashList(w http.ResponseWriter, r *http.Request) {
	var list []string = hashs.GetHashList()
	var response HashList = HashList{
		List: list,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func buildResponse(response string, status string) ResponseData {
	return ResponseData{
		Data: response,
		Status: status,
	}
}

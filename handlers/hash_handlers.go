package handlers

import (
	"cryptographyServer/hashs"
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

func HashData(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestData
	if err := readRequestBody(r, &requestBody); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	hashedData, err := hashs.Hash(requestBody.Data, requestBody.Hash)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responseBody := ResponseData{
		Data:   hashedData,
		Status: "Successfully hashed the data",
	}
	writeJSONResponse(w, http.StatusOK, responseBody)
}

func ShowHashList(w http.ResponseWriter, r *http.Request) {
	var list []string = hashs.GetHashList()
	var response HashList = HashList{
		List: list,
	}
	writeJSONResponse(w, http.StatusOK, response)
}

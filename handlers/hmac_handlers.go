package handlers

import (
	"cryptographyServer/hashs"
	"net/http"
)

type HMACRequest struct {
	Message string `json:"message"`
	Key string `json:"key"`
}

type HMACResponse struct {
	Hash string `json:"hash"`
}

func SignMessage(w http.ResponseWriter, r *http.Request) {
	var requestBody HMACRequest
	if err := readRequestBody(r, &requestBody); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	hash, err := hashs.CreateSignature(requestBody.Message, requestBody.Key)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responseBody := HMACResponse{
		Hash: hash,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}
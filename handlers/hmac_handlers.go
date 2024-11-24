package handlers

import (
	"cryptographyServer/hashs"
	"net/http"
)

type HMACRequest struct {
	Message string `json:"message" validate:"required"`
	Key string `json:"key" validate:"required"`
}

type VerifyRequest struct {
	Message string `json:"message" validate:"required"`
	Key string `json:"key" validate:"required"`
	Hash string `json:"hash" validate:"required"`
}

type VerifyResponse struct {
	IsAltered bool `json:"isaltered"`
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

func VerifyMessage(w http.ResponseWriter, r *http.Request) {
	var requestBody VerifyRequest
	if err := readRequestBody(r, &requestBody); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	isVerified, err := hashs.VerifySignature(requestBody.Message, requestBody.Key, requestBody.Hash)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return 
	}

	responseBody := VerifyResponse{
		IsAltered: isVerified,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}
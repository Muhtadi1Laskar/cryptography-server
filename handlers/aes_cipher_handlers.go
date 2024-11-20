package handlers

import (
	authenticatedEncryption "cryptographyServer/ciphers"
	"encoding/json"
	"io"
	"net/http"
)

type EncryptRequest struct {
	Key string `json:"key"`
	Data string `json:"data"`
}

type EncryptResponse struct {
	Cipher string `json:"cipher"`
	Nonce string `json:"nonce"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func AesEncrypt(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read the request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var requestBody EncryptRequest
	if err := json.Unmarshal(reqBody, &requestBody); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
	}

	cipherText, nonce, err := authenticatedEncryption.Encrypt(requestBody.Data, requestBody.Key)
	if err != nil {
		responseBody := ErrorResponse{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseBody)
		return
	}

	responseBody := EncryptResponse{
		Cipher: cipherText,
		Nonce: nonce,
		Message: "Successfully encrypted the message",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)
}
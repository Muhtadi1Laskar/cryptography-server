package handlers

import (
	authenticatedEncryption "cryptographyServer/ciphers"
	"net/http"
)

type EncryptRequest struct {
	Key  string `json:"key"`
	Data string `json:"data"`
}

type EncryptResponse struct {
	Cipher  string `json:"cipher"`
	Nonce   string `json:"nonce"`
	Message string `json:"message"`
}

type DecryptRequest struct {
	Cipher string `json:"cipher"`
	Nonce  string `json:"nonce"`
	Key    string `json:"key"`
}

type DecryptResponse struct {
	Data    string `json:"data"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func AesEncrypt(w http.ResponseWriter, r *http.Request) {
	var requestBody EncryptRequest
	if err := readRequestBody(r, &requestBody); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	cipherText, nonce, err := authenticatedEncryption.Encrypt(requestBody.Data, requestBody.Key)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	responseBody := EncryptResponse{
		Cipher:  cipherText,
		Nonce:   nonce,
		Message: "Successfully encrypted the message",
	}
	writeJSONResponse(w, http.StatusOK, responseBody)
}

func AesDecrypt(w http.ResponseWriter, r *http.Request) {
	var requestBody DecryptRequest
	if err := readRequestBody(r, &requestBody); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	plainText, err := authenticatedEncryption.Decrypt(requestBody.Cipher, requestBody.Nonce, requestBody.Key)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	responseBody := DecryptResponse{
		Data:    plainText,
		Message: "Successfully decrypted the message",
	}
	writeJSONResponse(w, http.StatusOK, responseBody)
}

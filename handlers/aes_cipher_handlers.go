package handlers

import (
	authenticatedEncryption "cryptographyServer/ciphers"
	"net/http"
)

type EncryptRequest struct {
	Key  string `json:"key" validate:"required"`
	Data string `json:"data" validate:"required"`
}

type EncryptResponse struct {
	Cipher  string `json:"cipher"`
	Nonce   string `json:"nonce"`
	Message string `json:"message"`
}

type DecryptRequest struct {
	Cipher string `json:"cipher" validate:"required"`
	Nonce  string `json:"nonce" validate:"required"`
	Key    string `json:"key" validate:"required"`
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
		writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	cipherText, nonce, err := authenticatedEncryption.Encrypt(requestBody.Data, requestBody.Key)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
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
		writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	plainText, err := authenticatedEncryption.Decrypt(requestBody.Cipher, requestBody.Nonce, requestBody.Key)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responseBody := DecryptResponse{
		Data:    plainText,
		Message: "Successfully decrypted the message",
	}
	writeJSONResponse(w, http.StatusOK, responseBody)
}

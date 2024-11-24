package handlers

import (
	"cryptographyServer/ciphers"
	"net/http"
	"crypto/rsa"
)

type KeyResponseBody struct {
	PrivateKey *rsa.PrivateKey `json:"privateKey"`
	PublicKey *rsa.PublicKey `json:"publicKey"` 
}

type RsaEncryptRequest struct {
	PlainText string `json:"plaintext"`
	PublicKey *rsa.PublicKey `json:"publickey"`
}

type RsaEncryptResponse struct {
	CipherText string `json:"ciphertext"`
}

func GenerateKeys(w http.ResponseWriter, r *http.Request) {
	privateKey, publicKey, err := ciphers.GenerateKeys()
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	var response KeyResponseBody = KeyResponseBody{
		PrivateKey: privateKey,
		PublicKey: publicKey,
	}
	writeJSONResponse(w, http.StatusOK, response)
}

func RSAEncryptMessage(w http.ResponseWriter, r *http.Request) {
	var requestBody RsaEncryptRequest
	if err := readRequestBody(r, &requestBody); err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	cipherText, err := ciphers.EncryptRSA(requestBody.PublicKey, requestBody.PlainText)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	responseBody := RsaEncryptResponse{
		CipherText: cipherText,
	}
	writeJSONResponse(w, http.StatusOK, responseBody)
}
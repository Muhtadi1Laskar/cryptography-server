package handlers

import (
	"cryptographyServer/ciphers"
	"net/http"
)

type KeyResponseBody struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string  `json:"publicKey"`
}

type RsaEncryptRequest struct {
	PlainText string         `json:"plaintext"`
	PublicKey string `json:"publickey"`
}

type RsaDecryptRequest struct {
	CipherText string          `json:"ciphertext"`
	PrivateKey string `json:"privatekey"`
}

type RsaEncryptResponse struct {
	CipherText string `json:"ciphertext"`
}

type RsaDecryptResponse struct {
	PlaintText string `json:"plaintext"`
}

func GenerateKeys(w http.ResponseWriter, r *http.Request) {
	privateKey, publicKey, err := ciphers.GenerateKeys()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	privateKeyPEM := ciphers.PrivateKeyToPEM(privateKey)
	publicKeyPEM := ciphers.PublicKeyToPEM(publicKey)

	var response KeyResponseBody = KeyResponseBody{
		PrivateKey: privateKeyPEM,
		PublicKey:  publicKeyPEM,
	}

	writeJSONResponse(w, http.StatusOK, response)
}

func RSAEncryptMessage(w http.ResponseWriter, r *http.Request) {
	var requestBody RsaEncryptRequest
	if err := readRequestBody(r, &requestBody); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	publicKey, err := ciphers.PEMToPublicKey(requestBody.PublicKey)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	cipherText, err := ciphers.EncryptRSA(publicKey, requestBody.PlainText)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	responseBody := RsaEncryptResponse{
		CipherText: cipherText,
	}
	writeJSONResponse(w, http.StatusOK, responseBody)
}

func RSADecryptMessage(w http.ResponseWriter, r *http.Request) {
	var requestBody RsaDecryptRequest
	if err := readRequestBody(r, &requestBody); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	privateKey, err := ciphers.PEMToPrivateKey(requestBody.PrivateKey)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	plainText, err := ciphers.DecryptRSA(privateKey, requestBody.CipherText)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responseBody := RsaDecryptResponse{
		PlaintText: plainText,
	}
	writeJSONResponse(w, http.StatusOK, responseBody)
}

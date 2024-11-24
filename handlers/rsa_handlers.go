package handlers

import (
	"cryptographyServer/ciphers"
	"fmt"
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
		writeJSONResponse(w, http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	privateKeyPEM := ciphers.PrivateKeyToPEM(privateKey)
	publicKeyPEM := ciphers.PublicKeyToPEM(publicKey)

	var response KeyResponseBody = KeyResponseBody{
		PrivateKey: privateKeyPEM,
		PublicKey:  publicKeyPEM,
	}
	fmt.Println("Private Key: ", privateKey)
	fmt.Println()
	writeJSONResponse(w, http.StatusOK, response)
}

func RSAEncryptMessage(w http.ResponseWriter, r *http.Request) {
	var requestBody RsaEncryptRequest
	if err := readRequestBody(r, &requestBody); err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		})
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
		writeJSONResponse(w, http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		})
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
		writeJSONResponse(w, http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	privateKey, err := ciphers.PEMToPrivateKey(requestBody.PrivateKey)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	plainText, err := ciphers.DecryptRSA(privateKey, requestBody.CipherText)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	fmt.Println("Plain Text: ", plainText)

	responseBody := RsaDecryptResponse{
		PlaintText: plainText,
	}
	writeJSONResponse(w, http.StatusOK, responseBody)
}

package handlers

import (
	"crypto/rsa"
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
	PublicKey *rsa.PublicKey `json:"publickey"`
}

type RsaDecryptRequest struct {
	CipherText string          `json:"ciphertext"`
	PrivateKey *rsa.PrivateKey `json:"privatekey"`
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

	cipherText, err := ciphers.EncryptRSA(requestBody.PublicKey, requestBody.PlainText)
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

	fmt.Println(requestBody.PrivateKey)
	fmt.Println()
	fmt.Println(requestBody.CipherText)

	plainText, err := ciphers.DecryptRSA(requestBody.PrivateKey, requestBody.CipherText)
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

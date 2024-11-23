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
package ciphers

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"crypto/sha256"
	"fmt"
)

func GenerateKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, fmt.Errorf("error generating private key: %v", err)
	}
	return privateKey, &privateKey.PublicKey, nil
}

func EncryptRSA(publicKey *rsa.PublicKey, plainText string) (string, error) {
	hash := sha256.New()

	cipherText, err := rsa.EncryptOAEP(
		hash,             
		rand.Reader,       
		publicKey,        
		[]byte(plainText), 
		nil, 
	)
	if err != nil {
		return "", fmt.Errorf("errof encrypting data: %v", err)
	}

	return base64.StdEncoding.EncodeToString(cipherText), nil
}


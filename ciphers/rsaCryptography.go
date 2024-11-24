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

func DecryptRSA(privateKey *rsa.PrivateKey, cipherText string) (string, error) {
	cipherBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", fmt.Errorf("error decoding cipherText: %v", err)
	}
	hash := sha256.New()

	plaintText, err := rsa.DecryptOAEP(
		hash,
		rand.Reader, 
		privateKey,
		cipherBytes,
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("error decrypting data: %v", err)
	}

	return string(plaintText), nil
}


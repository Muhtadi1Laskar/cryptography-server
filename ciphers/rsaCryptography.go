package ciphers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
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
		return "", fmt.Errorf("error decoding ciphertext: %v", err)
	}
	hash := sha256.New()

	plainText, err := rsa.DecryptOAEP(
		hash,          
		rand.Reader,   
		privateKey,     
		cipherBytes,    
		nil,            
	)
	if err != nil {
		return "", fmt.Errorf("error decrypting data: %v", err)
	}

	return string(plainText), nil
}

func PrivateKeyToPEM(privateKey *rsa.PrivateKey) string {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	return string(pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}))
}
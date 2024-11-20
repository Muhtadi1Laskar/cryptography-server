package authenticatedEncryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func Encrypt(data, cipherKey string) (string, error) {
	key, _ := hex.DecodeString(cipherKey)
	plainText := []byte(data)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("%s", err.Error())
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("%s", err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("%s", err.Error()) 
	}

	cipherText := aesgcm.Seal(nil, nonce, plainText, nil)

	return hex.EncodeToString(cipherText), nil
}

func Decrypt(data, nonceHex, cipherKey string) (string, error) {
	key, err := hex.DecodeString(cipherKey)
	if err != nil {
		return "", fmt.Errorf("invalid cipher key: %v", err)
	}
	cipherText, err := hex.DecodeString(data)
	if err != nil {
		return "", fmt.Errorf("invalid cipher text: %v", err)
	}
	nonce, err := hex.DecodeString(nonceHex)
	if err != nil {
		return "", fmt.Errorf("invalid Nonce: %v", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %v", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %v", err)
	}

	plaintText, err := aesgcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", fmt.Errorf("decryption failed: %v", err)
	}

	return string(plaintText), nil
}
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
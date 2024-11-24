package hashs

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func CreateSignature(message, key string) (string, error) {
	keyBytes := []byte(key)
	mac := hmac.New(sha256.New, keyBytes)

	if _, err := mac.Write([]byte(message)); err != nil {
		return "", fmt.Errorf("failed to write to HMAC: %v", err)
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}

func VerifySignature(message, key, hash string) (bool, error) {
	expectedSignature, err := hex.DecodeString(hash)
	if err != nil {
		return false, fmt.Errorf("invalid hash format: %v", err)
	}
	mac := hmac.New(sha256.New, []byte(key))

	if _, err := mac.Write([]byte(message)); err != nil {
		return false, fmt.Errorf("failed to write to HMAC: %v", err)
	}

	return hmac.Equal(expectedSignature, mac.Sum(nil)), nil
}
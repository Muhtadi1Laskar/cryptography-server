package hashs

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func CreateSignatur(message, key string) (string, error) {
	keyBytes := []byte(key)
	mac := hmac.New(sha256.New, keyBytes)

	if _, err := mac.Write([]byte(message)); err != nil {
		return "", fmt.Errorf("failed to write to HMAC: %v", err)
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}
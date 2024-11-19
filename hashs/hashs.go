package hashs

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
)

type UnaryFunc func(int) int

func Sha256(param string) string {
	hash := sha256.New()
	hash.Write([]byte(param))

	hashedBytes := hash.Sum(nil)
	encodedStr := hex.EncodeToString(hashedBytes)

	return encodedStr
}

func Md5(param string) string {
	hash := md5.New()
	hash.Write([]byte(param))

	hashedBytes := hash.Sum(nil)
	encodedStr := hex.EncodeToString(hashedBytes)

	return encodedStr
}

func Hash(param, hashName string) (string, error) {
	hashes := map[string]func() hash.Hash {
		"md5": md5.New,
		"sha256": sha256.New,
		"sha1": sha1.New,
		"sha224": sha256.New224,
		"sha384": sha512.New384,
	}
	hashFunc, exists := hashes[hashName]
	if !exists {
		return "", fmt.Errorf("Unsupported hash: %s", hashName)
	}

	hash := hashFunc()
	hash.Write([]byte(param))

	hashedBytes := hash.Sum(nil)
	encodedStr := hex.EncodeToString((hashedBytes))

	return encodedStr, nil
}
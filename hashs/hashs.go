package hashs

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

// func Hash(param, hashName string) (string, error) {
// 	hashes := map[string]func() hash.Hash {
// 		"md5": md5.New,
// 		"sha256": sha256.New,
// 		"sha1": sha1.New,
// 		"sha224": sha256.New224,
// 		"sha384": sha512.New384,
// 		"sha512_224": sha512.New512_224,
// 		"sha512": sha512.New,
// 		"ripemd160": ripemd160.New,
// 		"md4": md4.New,
// 		"sha3_224": sha3.New224,
// 		"sha3_256": sha3.New256,
// 		"sha3_384": sha3.New384,
// 		"sha3_512": sha3.New512,
// 		"sha512_256": sha512.New512_256,
// 		"blake2s_128": func() hash.Hash {
// 			h, _ := blake2s.New128(nil)
// 			return h
// 		},
// 	}
// 	hashFunc, exists := hashes[hashName]
// 	if !exists {
// 		return "", fmt.Errorf("Unsupported hash: %s", hashName)
// 	}

// 	hash := hashFunc()
// 	hash.Write([]byte(param))

// 	hashedBytes := hash.Sum(nil)
// 	encodedStr := hex.EncodeToString((hashedBytes))

// 	return encodedStr, nil
// }

func Hash(param string, hashName string) (string, error) {
	hashes := map[string]func() (hash.Hash, error){
		"md5": func() (hash.Hash, error) { return md5.New(), nil },
		"sha256": func() (hash.Hash, error) { return sha256.New(), nil },
		"sha1": func() (hash.Hash, error) { return sha1.New(), nil },
		"sha224": func() (hash.Hash, error) { return sha256.New224(), nil },
		"sha384": func() (hash.Hash, error) { return sha512.New384(), nil },
		"sha512_224": func() (hash.Hash, error) { return sha512.New512_224(), nil },
		"sha512": func() (hash.Hash, error) { return sha512.New(), nil },
		"ripemd160": func() (hash.Hash, error) { return ripemd160.New(), nil },
		"md4": func() (hash.Hash, error) { return md4.New(), nil },
		"sha3_224": func() (hash.Hash, error) { return sha3.New224(), nil },
		"sha3_256": func() (hash.Hash, error) { return sha3.New256(), nil },
		"sha3_384": func() (hash.Hash, error) { return sha3.New384(), nil },
		"sha3_512": func() (hash.Hash, error) { return sha3.New512(), nil },
		"sha512_256": func() (hash.Hash, error) { return sha512.New512_256(), nil },
		"blake2s_256": func() (hash.Hash, error) {
			return blake2s.New256(nil) 
		},
		"blake2b_256": func() (hash.Hash, error) {
			return blake2b.New256(nil)
		},
		"blake2b_384": func() (hash.Hash, error) {
			return blake2b.New384(nil)
		},
	}

	hashFunc, exists := hashes[hashName]
	if !exists {
		return "", fmt.Errorf("unsupported hash algorithm: %s", hashName)
	}

	hasher, err := hashFunc()
	if err != nil {
		return "", fmt.Errorf("failed to create hash instance: %v", err)
	}

	hasher.Write([]byte(param))

	hashedBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashedBytes), nil
}
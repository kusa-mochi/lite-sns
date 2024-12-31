package auth_utils

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

func GenerateHashString() string {
	seed := time.Now().GoString()
	return GetHashStringFrom(seed)
}

func GetHashStringFrom(str string) string {
	binary := sha256.Sum256([]byte(str))
	hash := hex.EncodeToString(binary[:])
	return hash
}

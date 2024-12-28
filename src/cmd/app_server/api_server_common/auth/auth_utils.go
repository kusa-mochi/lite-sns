package auth_utils

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

func GenerateHashString() string {
	seed := time.Now().GoString()
	binary := sha256.Sum256([]byte(seed))
	hash := hex.EncodeToString(binary[:])
	return hash
}

package api_server_common

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

func GenerateToken() string {
	seed := time.Now().GoString()
	binary := sha256.Sum256([]byte(seed))
	hash := hex.EncodeToString(binary[:])
	return hash
}

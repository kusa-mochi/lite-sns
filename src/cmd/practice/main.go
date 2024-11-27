package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"time"
)

func main() {
	hash := GenerateToken()
	log.Println("token:", hash)
}

func GenerateToken() string {
	seed := time.Now().GoString()
	log.Println(seed)
	binary := sha256.Sum256([]byte(seed))
	hash := hex.EncodeToString(binary[:])
	return hash
}

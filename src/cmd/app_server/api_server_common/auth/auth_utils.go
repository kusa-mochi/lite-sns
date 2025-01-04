package auth_utils

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

func ValidateTokenMiddleware(c *gin.Context) {
	handlers := c.HandlerNames()
	nHandlers := len(handlers)
	handlerName := handlers[nHandlers-1]
	handlerName = strings.Split(handlerName, "(*ApiServer).")[1]
	handlerName = strings.Split(handlerName, "-")[0]
	log.Println("handler:", handlerName)

	c.Next()
}

package main

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("app server started")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		ping_res, err := http.Post(
			"http://localhost:18081/token",
			"application/json",
			bytes.NewBuffer([]byte("{}")),
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "failed to get response from the auth server " + err.Error(),
			})
			return
		}

		body, err := io.ReadAll(ping_res.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "failed to get response boy from PING response " + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": string(body),
		})
	})
	r.Run(":18080")
}

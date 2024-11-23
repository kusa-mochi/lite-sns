package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	var (
		ip   = flag.String("ip", "localhost", "IP address of the app server")
		port = flag.Int("port", 10081, "port number of the app server")
	)
	const (
		apiPathPrefix string = "/lite-sns/api/v1"
	)

	r := gin.Default()

	r.POST(fmt.Sprintf("%s/signup_request", apiPathPrefix), func(c *gin.Context) {})
	r.POST(fmt.Sprintf("%s/signup", apiPathPrefix), func(c *gin.Context) {})
	r.GET(fmt.Sprintf("%s/mail_addr_auth", apiPathPrefix), func(c *gin.Context) {})
	r.POST(fmt.Sprintf("%s/signin", apiPathPrefix), func(c *gin.Context) {})

	r.Run(fmt.Sprintf(":%d", *port))
	// log.Println("app server started")
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	ping_res, err := http.Post(
	// 		"http://localhost:18081/token",
	// 		"application/json",
	// 		bytes.NewBuffer([]byte("{}")),
	// 	)
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "failed to get response from the auth server " + err.Error(),
	// 		})
	// 		return
	// 	}

	// 	body, err := io.ReadAll(ping_res.Body)
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "failed to get response boy from PING response " + err.Error(),
	// 		})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": string(body),
	// 	})
	// })
	// r.Run(":18080")
}

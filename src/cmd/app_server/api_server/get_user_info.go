package api_server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *ApiServer) GetUserInfo(c *gin.Context) {
	log.Println("server GetUserInfo start")

	// TODO

	c.JSON(http.StatusOK, gin.H{
		"message": "GetUserInfo fin",
	})
}

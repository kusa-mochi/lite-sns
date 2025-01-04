package api_server

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (s *ApiServer) GetUserInfo(c *gin.Context) {
	log.Println("server GetUserInfo start")
}

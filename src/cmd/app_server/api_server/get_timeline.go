package api_server

import (
	"lite-sns/m/src/cmd/app_server/commands"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *ApiServer) GetTimeline(c *gin.Context) {
	log.Println("server GetTimeline start")

	var userIdStr string = c.GetHeader("X-User-Id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Printf("failed to convert the user ID string (%s) to int | %s", userIdStr, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	if userId < 0 {
		log.Printf("invalid user ID (ID=%v)", userId)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	resCh := make(chan *commands.GetTimelineRes)
	s.commandCh <- &commands.GetTimelineCommand{
		UserId:              userId,
		CurrentOldestPostId: 0,
		ResCh:               resCh,
	}
	result := <-resCh
	if result.Error != nil {
		log.Printf("failed to get timeline | %s", result.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "GetTimeline fin",
		"timeline": result.Timeline,
	})
}

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
	currentOldestPostIdStr := c.Query("current_oldest_post_id")
	currentOldestPostId, err := strconv.Atoi(currentOldestPostIdStr)
	if err != nil {
		log.Printf("failed to convert the current oldest post ID string (%s) to int | %s", currentOldestPostIdStr, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	var userIdStr string = c.GetHeader("X-User-Id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Printf("failed to convert the user ID string (%s) to int | %s", userIdStr, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	if userId < 1 {
		log.Printf("invalid user ID (ID=%v)", userId)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	resCh := make(chan *commands.GetTimelineRes)
	s.commandCh <- &commands.GetTimelineCommand{
		UserId:              userId,
		CurrentOldestPostId: currentOldestPostId,
		ResCh:               resCh,
	}
	result := <-resCh
	if result.Error != nil {
		log.Printf("failed to get timeline | %s", result.Error.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "GetTimeline fin",
		"timeline": result.Timeline,
	})
}

package api_server

import (
	"lite-sns/m/src/cmd/app_server/commands"
	"net/http"

	"github.com/gin-gonic/gin"
)

// メールアドレスとパスワードによるサインイン処理
func (s *ApiServer) Signin(c *gin.Context) {
	resCh := make(chan string)
	s.commandCh <- &commands.SigninCommand{
		ResCh: resCh,
	}
	result := <-resCh

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

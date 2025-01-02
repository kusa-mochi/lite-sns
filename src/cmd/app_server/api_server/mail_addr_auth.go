package api_server

import (
	"lite-sns/m/src/cmd/app_server/commands"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ユーザーアカウント本登録処理
// 認証用メールのリンクにアクセスされた場合の処理を想定したAPI
func (s *ApiServer) MailAddrAuth(c *gin.Context) {
	log.Println("server mailaddrauth start")

	tokenString := c.Query("t")
	// パラメータ t が取得できなかった場合、
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "bad request",
		})
		return
	}

	resCh := make(chan *commands.MailAddrAuthRes)
	s.commandCh <- &commands.MailAddrAuthCommand{
		TokenString: tokenString,
		ResCh:       resCh,
	}
	result := <-resCh
	if result.Error != nil {
		switch result.Error.Error() {
		case "invalid access token":
			c.JSON(http.StatusBadRequest, gin.H{
				"error": result.Error.Error(),
			})
		case "internal server error":
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error(),
			})
		default:
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error": result.Error.Error(),
			})
		}
		return
	}

	c.Redirect(http.StatusMovedPermanently, result.RedirectTo)
}

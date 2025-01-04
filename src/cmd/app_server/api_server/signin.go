package api_server

import (
	auth_utils "lite-sns/m/src/cmd/app_server/api_server_common/auth"
	"lite-sns/m/src/cmd/app_server/commands"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// メールアドレスとパスワードによるサインイン処理
func (s *ApiServer) Signin(c *gin.Context) {
	var (
		emailAddr string = c.PostForm("EmailAddr")
		password  string = c.PostForm("Password")
	)

	log.Println("email addr:", emailAddr)
	log.Println("password:", password)

	// 受信したデータのバリデーション
	// スループット確保のためバリデーションのみこのゴルーチンで処理する。

	// eメールアドレス のバリデーション
	err := s.validateEmailAddress(emailAddr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "invalid signin data",
		})
		return
	}

	// パスワード のバリデーション
	// パスワードハッシュではなくパスワードをバリデーションする。
	err = s.validatePassword(password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "invalid signin data",
		})
		return
	}

	// パスワードハッシュの取得
	passwordHash := auth_utils.GetHashStringFrom(password)

	resCh := make(chan *commands.SigninRes)
	s.commandCh <- &commands.SigninCommand{
		MailAddr:     emailAddr,
		PasswordHash: passwordHash,
		ResCh:        resCh,
	}
	result := <-resCh

	if result.Error != nil {
		errorMessage := result.Error.Error()
		switch errorMessage {
		case "invalid signin data":
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errorMessage,
			})
		case "internal server error":
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": errorMessage,
			})
		default:
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error": errorMessage,
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":  result.Message,
		"token":   result.TokenString,
		"user_id": result.UserId,
	})
}

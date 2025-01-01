package api_server

import (
	"fmt"
	"lite-sns/m/src/cmd/app_server/commands"
	"lite-sns/m/src/cmd/app_server/interfaces"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	r             *gin.Engine
	apiPathPrefix string
	port          int
	commandCh     chan<- interfaces.ApiServerCommandInterface
}

func NewApiServer(
	apiPathPrefix string,
	port int,
	commandCh chan<- interfaces.ApiServerCommandInterface,
) *ApiServer {
	s := &ApiServer{
		r:             gin.Default(),
		apiPathPrefix: apiPathPrefix,
		port:          port,
		commandCh:     commandCh,
	}

	s.r.POST(fmt.Sprintf("%s/signup_request", apiPathPrefix), s.SignupRequest)
	s.r.POST(fmt.Sprintf("%s/signup", apiPathPrefix), s.Signup)
	s.r.GET(fmt.Sprintf("%s/mail_addr_auth", apiPathPrefix), s.MailAddrAuth)
	s.r.POST(fmt.Sprintf("%s/signin", apiPathPrefix), s.Signin)

	return s
}

func (s *ApiServer) Run() {
	s.r.Run(fmt.Sprintf(":%d", s.port)) // エラーが発生しない限りここで処理がブロックされる。
}

// ユーザーアカウント登録リクエストを受け付けるAPI
func (s *ApiServer) SignupRequest(c *gin.Context) {
	resCh := make(chan string)
	s.commandCh <- &commands.SignupRequestCommand{
		ResCh: resCh,
	}
	result := <-resCh

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

// ユーザーアカウント本登録処理
// 認証用メールのリンクにアクセスされた場合の処理を想定したAPI
func (s *ApiServer) MailAddrAuth(c *gin.Context) {
	log.Println("server mailaddrauth start")

	tokenString := c.Query("t")
	// パラメータ t が取得できなかった場合、
	if tokenString == "" {
		c.JSON(http.StatusForbidden, gin.H{
			"result": "forbidden",
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
		case "server error":
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

package main

import (
	"fmt"
	"lite-sns/m/src/cmd/app_server/commands"

	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	r             *gin.Engine
	apiPathPrefix string
	port          int
	commandCh     chan<- ApiServerCommandInterface
}

func NewApiServer(
	apiPathPrefix string,
	port int,
	commandCh chan<- ApiServerCommandInterface,
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
	s.commandCh <- &commands.SignupRequestCommand{}
}

// ユーザーアカウント仮登録処理
func (s *ApiServer) Signup(c *gin.Context) {
	s.commandCh <- &commands.SignupCommand{}
}

// ユーザーアカウント本登録処理
// 認証用メールのリンクにアクセスされた場合の処理を想定したAPI
func (s *ApiServer) MailAddrAuth(c *gin.Context) {
	s.commandCh <- &commands.MailAddrAuthCommand{}
}

// メールアドレスとパスワードによるサインイン処理
func (s *ApiServer) Signin(c *gin.Context) {
	s.commandCh <- &commands.SigninCommand{}
}

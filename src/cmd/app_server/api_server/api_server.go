package api_server

import (
	"fmt"
	"lite-sns/m/src/cmd/app_server/interfaces"
	"lite-sns/m/src/cmd/app_server/server_configs"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	r         *gin.Engine
	configs   *server_configs.ServerConfigs
	commandCh chan<- interfaces.ApiServerCommandInterface
}

func NewApiServer(
	configs *server_configs.ServerConfigs,
	commandCh chan<- interfaces.ApiServerCommandInterface,
) *ApiServer {
	s := &ApiServer{
		r:         gin.Default(),
		configs:   configs,
		commandCh: commandCh,
	}

	// すべてのAPIに適用する設定・ミドルウェア
	s.r.Use(
		cors.New(
			cors.Config{
				AllowOrigins: []string{fmt.Sprintf("http://%s:%v", configs.Frontend.Ip, configs.Frontend.Port)}, // TODO: TLS対応
				AllowMethods: []string{"GET", "POST"},
				AllowHeaders: []string{"Origin"},
				MaxAge:       24 * time.Hour,
			},
		),
	)

	log.Println("configured CORS")

	// publicなhandler
	publicGroup := s.r.Group(fmt.Sprintf("%s/public", configs.App.ApiPrefix))
	{
		publicGroup.POST("/signup", s.Signup)
		publicGroup.GET("/mail_addr_auth", s.MailAddrAuth)
		publicGroup.POST("/signin", s.Signin)
	}

	// 適切なアクセストークンの使用でアクセス可能なhandler
	authUserGroup := s.r.Group(fmt.Sprintf("%s/auth_user", configs.App.ApiPrefix))
	authUserGroup.Use(
		s.ValidateTokenMiddleware,
	)
	{
		authUserGroup.POST("/get_user_info", s.GetUserInfo)
	}

	log.Println("gin callbacks is ready")

	return s
}

func (s *ApiServer) Run() {
	log.Println("app server is now listening...")
	s.r.Run(fmt.Sprintf(":%d", s.configs.App.Port)) // エラーが発生しない限りここで処理がブロックされる。
}

func (s *ApiServer) ValidateTokenMiddleware(c *gin.Context) {
	// TODO: HTTPヘッダからアクセストークンを取得する。

	// TODO: HTTPヘッダからユーザーIDを取得する。

	// TODO: ユーザーIDに対応する秘密鍵をDBから取得する。
	// TODO: 処理の原子性確保・順序保証のためDB処理はcommandsパッケージで行う。

	// TODO: アクセストークンを検証する。

	c.Next()
}

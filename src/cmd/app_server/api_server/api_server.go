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
	s.r.Use(cors.New(cors.Config{
		AllowOrigins: []string{fmt.Sprintf("http://%s:%v", configs.Frontend.Ip, configs.Frontend.Port)}, // TODO: TLS対応
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Origin"},
		MaxAge:       24 * time.Hour,
	}))

	log.Println("configured CORS")

	s.r.POST(fmt.Sprintf("%s/signup", configs.App.ApiPrefix), s.Signup)
	s.r.GET(fmt.Sprintf("%s/mail_addr_auth", configs.App.ApiPrefix), s.MailAddrAuth)
	s.r.POST(fmt.Sprintf("%s/signin", configs.App.ApiPrefix), s.Signin)

	log.Println("gin callbacks is ready")

	return s
}

func (s *ApiServer) Run() {
	log.Println("app server is now listening...")
	s.r.Run(fmt.Sprintf(":%d", s.configs.App.Port)) // エラーが発生しない限りここで処理がブロックされる。
}

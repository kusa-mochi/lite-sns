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

type RestMethod int

const (
	RestMethod_GET = iota
	RestMethod_POST
)

type ApiMetaData struct {
	Method       RestMethod
	CallbackFunc func(*gin.Context)
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

	// data for API definitions
	var callbacks map[string]ApiMetaData = map[string]ApiMetaData{
		"signup": {
			Method:       RestMethod_POST,
			CallbackFunc: s.Signup,
		},
		"mail_addr_auth": {
			Method:       RestMethod_GET,
			CallbackFunc: s.MailAddrAuth,
		},
		"signin": {
			Method:       RestMethod_POST,
			CallbackFunc: s.Signin,
		},
	}

	// set API callbacks
	for apiName, apiMetaData := range callbacks {
		apiPath := fmt.Sprintf("%s/%s", configs.App.ApiPrefix, apiName)
		switch apiMetaData.Method {
		case RestMethod_GET:
			s.r.GET(apiPath, apiMetaData.CallbackFunc)
		case RestMethod_POST:
			s.r.POST(apiPath, apiMetaData.CallbackFunc)
		}
	}

	log.Println("gin callbacks is ready")

	return s
}

func (s *ApiServer) Run() {
	log.Println("app server is now listening...")
	s.r.Run(fmt.Sprintf(":%d", s.configs.App.Port)) // エラーが発生しない限りここで処理がブロックされる。
}

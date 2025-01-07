package api_server

import (
	"fmt"
	"lite-sns/m/src/cmd/app_server/commands"
	"lite-sns/m/src/cmd/app_server/interfaces"
	"lite-sns/m/src/cmd/app_server/server_configs"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
				AllowHeaders: []string{"Origin", "Authorization", "X-User-Id"},
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
	// auth_userグループのAPIに適用する設定・ミドルウェア
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
	// HTTPヘッダからアクセストークンを取得する。
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		log.Println("no access token")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}

	// HTTPヘッダからユーザーIDを取得する。
	userIdString := c.GetHeader("X-User-Id")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		log.Println("failed to convert userIdString to int |", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}

	// ユーザーIDに対応する秘密鍵をDBから取得する。
	// 処理は通常のコマンド処理と同様にメインゴルーチンで行う。
	resCh := make(chan *commands.GetUserSecretKeyRes)
	s.commandCh <- &commands.GetUserSecretKeyCommand{
		UserId: userId,
		ResCh:  resCh,
	}
	res := <-resCh
	if res.Error != nil {
		log.Println("failed to get a secret key corresponding to the user ID |", res.Error.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}

	// アクセストークンを検証する。
	_, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("")
		}

		return []byte(res.SecretKey), nil
	})
	if err != nil {
		log.Println("failed to parse a token |", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}

	c.Next()
}

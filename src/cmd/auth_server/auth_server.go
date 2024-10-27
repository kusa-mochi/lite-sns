package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	TOKEN_LIFETIME time.Duration = 30 * time.Minute
)

type AuthServer struct {
	router      *gin.Engine
	accessToken string
	options     *AuthServerOptions
}

type AuthServerOptions struct {
	Addr      string `json:"auth_server_addr"`
	SecretKey string `json:"secret_key"`
}

func NewAuthServer(options *AuthServerOptions) *AuthServer {
	fmt.Println("*** options ***")
	fmt.Println(*options)
	return &AuthServer{
		router:      nil,
		accessToken: "",
		options:     options,
	}
}

func (s *AuthServer) Run() error {
	s.router = gin.Default()
	s.router.POST("/token", s.Token)
	log.Println("running auth server on ", s.options.Addr)
	return s.router.Run(s.options.Addr)
}

func (s *AuthServer) Token(c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(TOKEN_LIFETIME).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.options.SecretKey))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to generate a new token",
		})
		return
	}

	s.accessToken = tokenString

	c.JSON(http.StatusOK, gin.H{
		"token": fmt.Sprintf("Bearer %s", tokenString),
	})
}

package main

import (
	"fmt"
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
	Addr      string `json:"addr"`
	SecretKey string `json:"secret_key"`
}

func NewAuthServer(options *AuthServerOptions) *AuthServer {
	return &AuthServer{
		router:      nil,
		accessToken: "",
		options:     options,
	}
}

func (s *AuthServer) Run() error {
	s.router = gin.Default()
	s.router.POST("/token", s.Token)
	return s.router.Run(s.options.Addr)
}

func (s *AuthServer) Token(c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(TOKEN_LIFETIME).Unix(),
	})

	tokenString, err := token.SignedString(s.options.SecretKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to generate new token",
		})
	}

	s.accessToken = tokenString

	c.JSON(http.StatusOK, gin.H{
		"token": fmt.Sprintf("Bearer %s", tokenString),
	})
}

package commands

import (
	"fmt"
	"lite-sns/m/src/cmd/app_server/server_configs"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

type MailAddrAuthCommand struct {
	TokenString string
	ResCh       chan<- string
}

func (c *MailAddrAuthCommand) Exec(configs *server_configs.ServerConfigs) {
	log.Println("mail addr auth exec")

	tokenString := c.TokenString

	// TODO: search and read a secret key from DB.
	const secretKey string = "secretkey"

	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		c.ResCh <- err.Error()
		return
	}

	c.ResCh <- "mail addr auth fin"
}

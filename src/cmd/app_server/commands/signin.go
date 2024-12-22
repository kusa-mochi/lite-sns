package commands

import (
	"lite-sns/m/src/cmd/app_server/server_configs"
	"log"
)

type SigninCommand struct {
	ResCh chan<- string
}

func (c *SigninCommand) Exec(configs *server_configs.ServerConfigs) {
	log.Println("signin exec")
	c.ResCh <- "signin fin"
}

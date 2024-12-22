package commands

import (
	"lite-sns/m/src/cmd/app_server/server_configs"
	"log"
)

type MailAddrAuthCommand struct {
	ResCh chan<- string
}

func (c *MailAddrAuthCommand) Exec(configs *server_configs.ServerConfigs) {
	log.Println("mail addr auth exec")
	c.ResCh <- "mail addr auth fin"
}

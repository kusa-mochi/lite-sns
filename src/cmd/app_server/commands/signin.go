package commands

import (
	"database/sql"
	"lite-sns/m/src/cmd/app_server/server_configs"
	"log"
)

type SigninCommand struct {
	MailAddr     string
	PasswordHash string
	ResCh        chan<- string
}

func (c *SigninCommand) Exec(configs *server_configs.ServerConfigs, db *sql.DB) {
	log.Println("signin exec")
	c.ResCh <- "signin fin"
}

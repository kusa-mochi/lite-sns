package commands

import (
	"database/sql"
	"lite-sns/m/src/cmd/app_server/server_configs"
	"log"
)

type SignupRequestCommand struct {
	ResCh chan<- string
}

func (c *SignupRequestCommand) Exec(configs *server_configs.ServerConfigs, db *sql.DB) {
	log.Println("signup request exec")
	c.ResCh <- "signup request fin"
}

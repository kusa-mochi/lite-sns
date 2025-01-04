package commands

import (
	"database/sql"
	"lite-sns/m/src/cmd/app_server/server_configs"
)

type GetUserSecretKeyCommand struct {
	UserId string
	ResCh  chan<- *GetUserSecretKeyRes
}

type GetUserSecretKeyRes struct {
	SecretKey string
	Error     error
}

func (c *GetUserSecretKeyCommand) Exec(configs *server_configs.ServerConfigs, db *sql.DB) {
}

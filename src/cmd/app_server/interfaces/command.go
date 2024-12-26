package interfaces

import (
	"database/sql"
	"lite-sns/m/src/cmd/app_server/server_configs"
)

type ApiServerCommandInterface interface {
	Exec(configs *server_configs.ServerConfigs, db *sql.DB)
}

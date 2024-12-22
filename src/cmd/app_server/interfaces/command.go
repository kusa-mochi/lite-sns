package interfaces

import "lite-sns/m/src/cmd/app_server/server_configs"

type ApiServerCommandInterface interface {
	Exec(configs *server_configs.ServerConfigs)
}

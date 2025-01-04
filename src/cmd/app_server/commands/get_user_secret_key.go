package commands

import (
	"database/sql"
	"fmt"
	db_utils "lite-sns/m/src/cmd/app_server/api_server_common/db"
	"lite-sns/m/src/cmd/app_server/server_configs"
	"log"
)

type GetUserSecretKeyCommand struct {
	UserId int
	ResCh  chan<- *GetUserSecretKeyRes
}

type GetUserSecretKeyRes struct {
	SecretKey string
	Error     error
}

func (c *GetUserSecretKeyCommand) Exec(configs *server_configs.ServerConfigs, db *sql.DB) {
	selectData, err := db_utils.SelectFrom(
		db,
		[]string{"access_token_secret_key"},
		"sns_user",
		"WHERE id = $1",
		c.UserId,
	)
	if err != nil {
		log.Println("failed to get a secret key from DB |", err.Error())
		c.ResCh <- &GetUserSecretKeyRes{
			SecretKey: "",
			Error:     fmt.Errorf("internal server error"),
		}
		return
	}

	userInfo := selectData[0]
	secretKey := userInfo["access_token_secret_key"].(string)

	c.ResCh <- &GetUserSecretKeyRes{
		SecretKey: secretKey,
		Error:     nil,
	}
}

package commands

import (
	"database/sql"
	"fmt"
	db_utils "lite-sns/m/src/cmd/app_server/api_server_common/db"
	"lite-sns/m/src/cmd/app_server/server_configs"
	"log"
)

// アクセストークン検証を行うミドルウェアでの使用を想定したコマンド。
type GetUserSecretKeyCommand struct {
	UserId int
	ResCh  chan<- *GetUserSecretKeyRes
}

type GetUserSecretKeyRes struct {
	SecretKey string // 重要：この情報はクライアントに渡したりクライアントからアクセスできる状態にはしないこと。
	Error     error
}

func (c *GetUserSecretKeyCommand) Exec(configs *server_configs.ServerConfigs, db *sql.DB) {
	log.Println("get user secret key exec")

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
	if len(selectData) == 0 {
		log.Println("secret key not found on DB")
		c.ResCh <- &GetUserSecretKeyRes{
			SecretKey: "",
			Error:     fmt.Errorf("bad request"),
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

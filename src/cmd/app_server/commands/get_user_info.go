package commands

import (
	"database/sql"
	"fmt"
	db_utils "lite-sns/m/src/cmd/app_server/api_server_common/db"
	"lite-sns/m/src/cmd/app_server/server_configs"
	"log"
)

type GetUserInfoCommand struct {
	UserId int
	ResCh  chan<- *GetUserInfoRes
}

type GetUserInfoRes struct {
	Username            string
	IconType            string
	IconBackgroundColor string
	Error               error
}

func (c *GetUserInfoCommand) Exec(configs *server_configs.ServerConfigs, db *sql.DB) {
	log.Println("get user info exec")

	// UserIdに対応するユーザー情報をDBから取得する。
	selectData, err := db_utils.SelectFrom(
		db,
		[]string{"name", "icon_type", "icon_background_color"},
		"sns_user",
		"WHERE id = $1",
		c.UserId,
	)
	if err != nil {
		// 何もせずコマンド終了。
		log.Printf("failed to get a user info corresponding to the user ID (ID=%v) from DB | %s", c.UserId, err.Error())
		c.ResCh <- &GetUserInfoRes{
			Username:            "",
			IconType:            "",
			IconBackgroundColor: "",
			Error:               fmt.Errorf("bad request"),
		}
		return
	}

	userInfo := selectData[0]

	c.ResCh <- &GetUserInfoRes{
		Username:            userInfo["name"].(string),
		IconType:            userInfo["icon_type"].(string),
		IconBackgroundColor: userInfo["icon_background_color"].(string),
		Error:               nil,
	}
}

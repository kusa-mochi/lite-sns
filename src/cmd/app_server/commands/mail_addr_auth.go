package commands

import (
	"database/sql"
	"fmt"
	db_utils "lite-sns/m/src/cmd/app_server/api_server_common/db"
	"lite-sns/m/src/cmd/app_server/server_configs"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

type MailAddrAuthCommand struct {
	TokenString string
	ResCh       chan<- string
}

func (c *MailAddrAuthCommand) Exec(configs *server_configs.ServerConfigs, db *sql.DB) {
	log.Println("mail addr auth exec")

	tokenString := c.TokenString

	// DBから、サインアップ用トークンに対応する秘密鍵を取得する。
	keys, err := db_utils.SelectFrom(
		db,
		[]string{"secret_key"},
		"signup_access_token",
		"where access_token = $1",
		tokenString,
	)
	if err != nil {
		// 何もせずコマンド終了。
		log.Println("failed to query SELECT command |", err.Error())
		return
	}
	// トークン検証に必要な秘密鍵が見つからなかった場合、
	if len(keys) == 0 {
		// 何もせずコマンド終了。
		log.Println("there is no secret key related to an access token")
		return
	}
	secretKey := keys[0].(string)

	// DBからアクセストークンとそれに対応する秘密鍵を削除する。
	err = db_utils.DeleteFrom(
		db,
		"signup_access_token",
		"where access_token = $1",
		tokenString,
	)
	if err != nil {
		// 何もせずコマンド終了
		log.Println("failed to execute DELETE command |", err.Error())
		return
	}

	// DBから取得した秘密鍵を用いてトークンを検証する。
	_, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		c.ResCh <- err.Error()
		return
	}

	// コマンドの正常終了
	c.ResCh <- "mail addr auth fin"
}

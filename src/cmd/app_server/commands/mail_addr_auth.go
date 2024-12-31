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
	ResCh       chan<- *MailAddrAuthRes
	Error       error
}

type MailAddrAuthRes struct {
	Message string
	Error   error
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
		log.Println("failed to find a secret key for an access token |", err.Error())
		c.ResCh <- &MailAddrAuthRes{
			Message: "",
			Error:   fmt.Errorf("invalid access token"),
		}
		return
	}
	// トークン検証に必要な秘密鍵が見つからなかった場合、
	if len(keys) == 0 {
		// 何もせずコマンド終了。
		log.Println("there is no secret key related to the signup access token")
		c.ResCh <- &MailAddrAuthRes{
			Message: "",
			Error:   fmt.Errorf("invalid access token"),
		}
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
		log.Println("failed to delete a signup token |", err.Error())
		c.ResCh <- &MailAddrAuthRes{
			Message: "",
			Error:   fmt.Errorf("server error"),
		}
		return
	}

	// DBから取得した秘密鍵を用いてトークンを検証する。
	_, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			log.Printf("unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("")
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		log.Println("failed to parse a token |", err.Error())
		c.ResCh <- &MailAddrAuthRes{
			Message: "",
			Error:   fmt.Errorf("invalid access token"),
		}
		return
	}

	// コマンドの正常終了
	c.ResCh <- &MailAddrAuthRes{
		Message: "mail addr auth fin",
		Error:   nil,
	}
}

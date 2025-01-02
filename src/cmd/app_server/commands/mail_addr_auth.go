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
}

type MailAddrAuthRes struct {
	Message    string
	RedirectTo string
	Error      error
}

func (c *MailAddrAuthCommand) Exec(configs *server_configs.ServerConfigs, db *sql.DB) {
	log.Println("mail addr auth exec")

	tokenString := c.TokenString

	cols := []string{"email_address", "nickname", "password_hash", "secret_key"}

	// DBから、サインアップ用トークンに対応する秘密鍵を取得する。
	selectData, err := db_utils.SelectFrom(
		db,
		cols,
		"signup_access_token",
		"where access_token = $1",
		tokenString,
	)
	if err != nil {
		// 何もせずコマンド終了。
		log.Println("failed to find a secret key for an access token |", err.Error())
		c.ResCh <- &MailAddrAuthRes{
			Message:    "",
			RedirectTo: "",
			Error:      fmt.Errorf("invalid access token"),
		}
		return
	}

	// トークン検証に必要な秘密鍵が見つからなかった場合、
	if len(selectData) == 0 {
		// 何もせずコマンド終了。
		log.Println("there is no secret key related to the signup access token")
		c.ResCh <- &MailAddrAuthRes{
			Message:    "",
			RedirectTo: "",
			Error:      fmt.Errorf("invalid access token"),
		}
		return
	}

	signupData := selectData[0]

	secretKey := signupData["secret_key"].(string)

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
			Message:    "",
			RedirectTo: "",
			Error:      fmt.Errorf("internal server error"),
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
			Message:    "",
			RedirectTo: "",
			Error:      fmt.Errorf("invalid access token"),
		}
		return
	}

	// ユーザーアカウントをDBに新規登録する。
	db_utils.InsertInto(
		db,
		"sns_user",
		db_utils.KeyValuePair{
			Key:   "name",
			Value: signupData["nickname"],
		},
		db_utils.KeyValuePair{
			Key:   "icon_type",
			Value: "IconType_Default",
		},
		db_utils.KeyValuePair{
			Key:   "icon_background_color",
			Value: "F0F0F0",
		},
		db_utils.KeyValuePair{
			Key:   "email_address",
			Value: signupData["email_address"],
		},
		db_utils.KeyValuePair{
			Key:   "password_hash",
			Value: signupData["password_hash"],
		},
		db_utils.KeyValuePair{
			Key:   "access_token_secret_key",
			Value: "",
		},
	)

	// コマンドの正常終了
	c.ResCh <- &MailAddrAuthRes{
		Message:    "mail addr auth fin",
		RedirectTo: fmt.Sprintf("http://%s:%v/mail_addr_auth", configs.Frontend.Ip, configs.Frontend.Port),
		Error:      nil,
	}
}

package commands

import (
	"database/sql"
	"fmt"
	auth_utils "lite-sns/m/src/cmd/app_server/api_server_common/auth"
	db_utils "lite-sns/m/src/cmd/app_server/api_server_common/db"
	"lite-sns/m/src/cmd/app_server/server_configs"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type SigninCommand struct {
	MailAddr     string
	PasswordHash string
	ResCh        chan<- *SigninRes
}

type SigninRes struct {
	Message     string
	TokenString string
	UserId      int64
	Error       error
}

func (c *SigninCommand) Exec(configs *server_configs.ServerConfigs, db *sql.DB) {
	log.Println("signin exec")

	emailAddress := c.MailAddr
	passwordHash := c.PasswordHash

	log.Println("email address:", emailAddress)

	// DBのユーザーテーブルについて、メールアドレスが一致するレコードを検索・取得する。
	selectData, err := db_utils.SelectFrom(
		db,
		[]string{"id", "name", "icon_type", "icon_background_color", "password_hash"},
		"sns_user",
		"WHERE email_address = $1",
		emailAddress,
	)
	if err != nil {
		// 何もせずコマンド終了。
		log.Println("failed to find a user data |", err.Error())
		c.ResCh <- &SigninRes{
			Message:     "",
			TokenString: "",
			UserId:      -1,
			Error:       fmt.Errorf("invalid signin data"),
		}
		return
	}

	// レコードが存在しない場合
	if len(selectData) == 0 {
		// 何もせずコマンド終了。
		log.Printf("there is no user data corresponding to the email address <%s>", emailAddress)
		c.ResCh <- &SigninRes{
			Message:     "",
			TokenString: "",
			UserId:      -1,
			Error:       fmt.Errorf("invalid signin data"),
		}
		return
	}

	// レコードからパスワードハッシュを取得する。
	signinData := selectData[0]
	passwordHashInDB := signinData["password_hash"].(string)

	// signin API に渡されたパスワードハッシュと一致しない場合
	if passwordHash != passwordHashInDB {
		log.Printf("invalid password")
		c.ResCh <- &SigninRes{
			Message:     "",
			TokenString: "",
			UserId:      -1,
			Error:       fmt.Errorf("invalid signin data"),
		}
		return
	}

	//// 以下、認証に成功した前提の処理

	// このユーザー専用の秘密鍵を生成する。（ユーザーには共有せず、サーバー上のみで秘匿するデータ）
	secretKey := auth_utils.GenerateHashString()

	// 秘密鍵を用いて有効期限付のアクセストークンを発行する。
	const tokenLifetime time.Duration = 7 * 24 * time.Hour
	expirationDatetime := time.Now().Add(tokenLifetime).Unix()
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"exp": expirationDatetime,
		},
	)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Println("failed to generate a signin-token string")
		c.ResCh <- &SigninRes{
			Message:     "",
			TokenString: "",
			UserId:      -1,
			Error:       fmt.Errorf("internal server error"),
		}
		return
	}

	// 秘密鍵をDBに保存する。
	log.Println("saving secret key:", secretKey)
	rowCnt, err := db_utils.PrepareAndExec(
		db,
		"UPDATE sns_user SET access_token_secret_key = $1 WHERE email_address = $2",
		secretKey,
		emailAddress,
	)
	if err != nil {
		log.Println("failed to update secret key")
		c.ResCh <- &SigninRes{
			Message:     "",
			TokenString: "",
			UserId:      -1,
			Error:       fmt.Errorf("internal server error"),
		}
		return
	}
	log.Printf("ID = <not supported>, affected = %d", rowCnt)

	// コマンド正常終了
	c.ResCh <- &SigninRes{
		Message:     "signin fin",
		TokenString: tokenString,
		UserId:      signinData["id"].(int64),
		Error:       nil,
	}
}

package commands

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	auth_utils "lite-sns/m/src/cmd/app_server/api_server_common/auth"
	db_utils "lite-sns/m/src/cmd/app_server/api_server_common/db"
	"lite-sns/m/src/cmd/app_server/server_configs"
	"log"
	"net/mail"
	"net/smtp"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type SignupCommand struct {
	EmailAddr    string
	Nickname     string
	PasswordHash string
	ResCh        chan<- string
}

func (c *SignupCommand) sendAuthMail(configs *server_configs.SmtpConfig, toAddr string, subject string, body string) {
	// Setup headers
	from := mail.Address{Name: "", Address: configs.Username}
	to := mail.Address{Name: "", Address: toAddr}
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	auth := smtp.PlainAuth("", configs.Username, configs.Password, configs.Hostname)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         configs.Hostname,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%v", configs.Hostname, configs.Port), tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	client, err := smtp.NewClient(conn, configs.Hostname)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = client.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = client.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := client.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	client.Quit()
}

func (c *SignupCommand) Exec(configs *server_configs.ServerConfigs, db *sql.DB) {
	var (
		emailAddress string = c.EmailAddr
		nickname     string = c.Nickname
		passwordHash string = c.PasswordHash
		subj         string = "lite-sns email test"
	)

	// このサインアップ処理でのみ有効な秘密鍵を生成する。
	secretKey := auth_utils.GenerateHashString()

	// サインアップ用アクセストークンを発行する。
	// アクセストークンには有効期限が設定されている。
	const tokenLifetime time.Duration = 10 * time.Minute
	var expirationDatetime int64 = time.Now().Add(tokenLifetime).Unix()
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"exp": expirationDatetime,
		},
	)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Println("failed to generate a signup-token string")
		c.ResCh <- "internal server error"
		return
	}
	log.Println("token string:", tokenString)

	// DBに、サインアップ用アクセストークンと秘密鍵を登録する。
	rowCnt, err := db_utils.InsertInto(
		db,
		"signup_access_token",
		db_utils.KeyValuePair{
			Key:   "access_token",
			Value: tokenString,
		},
		db_utils.KeyValuePair{
			Key:   "email_address",
			Value: emailAddress,
		},
		db_utils.KeyValuePair{
			Key:   "nickname",
			Value: nickname,
		},
		db_utils.KeyValuePair{
			Key:   "password_hash",
			Value: passwordHash,
		},
		db_utils.KeyValuePair{
			Key:   "secret_key",
			Value: secretKey,
		},
		db_utils.KeyValuePair{
			Key:   "expiration_datetime",
			Value: expirationDatetime,
		},
	)
	if err != nil {
		log.Println("failed to insert a data into signup_access_token")
		c.ResCh <- "server internal error"
		return
	}
	log.Printf("ID = <not supported>, affected = %d", rowCnt)

	// 認証用メールの本文を生成する。
	body := fmt.Sprintf("access to the following link:\nhttp://%s:%v%s/mail_addr_auth?t=%s", configs.App.Ip, configs.App.Port, configs.App.ApiPrefix, tokenString)

	// 認証用メールを送信する。
	c.sendAuthMail(&configs.Smtp, c.EmailAddr, subj, body)

	c.ResCh <- "signup fin"
}

package commands

import (
	"crypto/tls"
	"fmt"
	"lite-sns/m/src/cmd/app_server/api_server_common"
	"lite-sns/m/src/cmd/app_server/server_configs"
	"log"
	"net/mail"
	"net/smtp"
)

type SignupCommand struct {
	EmailAddr string
	ResCh     chan<- string
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

func (c *SignupCommand) Exec(configs *server_configs.ServerConfigs) {
	log.Println("email addr:", c.EmailAddr)

	var (
		subj string = "lite-sns email test"
		body string = "This is a test email from lite-sns site.\nThis is a second line :)"
	)

	// このサインアップ処理でのみ有効な秘密鍵を生成する。
	secretKey := api_server_common.GenerateToken()

	// サインアップ用アクセストークンを発行する。

	// DBに、サインアップ用アクセストークンと秘密鍵を登録する。

	// 認証用メールの本文を生成する。

	// 認証用メールを送信する。
	c.sendAuthMail(&configs.Smtp, c.EmailAddr, subj, body)

	c.ResCh <- "signup fin"
}

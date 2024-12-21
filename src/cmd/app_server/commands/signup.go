package commands

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
)

type SignupCommand struct {
	EmailAddr string
	ResCh     chan<- string
}

func (c *SignupCommand) Exec() {
	log.Println("email addr:", c.EmailAddr)

	const (
		servername string = "<smtp_server_host>:<port>"
		fromAddr   string = "<email_addr>"
		password   string = "<password>"
	)

	var (
		toAddr string = c.EmailAddr
		subj   string = "lite-sns email test"
		body   string = "This is a test email from lite-sns site.\nThis is a second line :)"
	)

	// Setup headers
	from := mail.Address{Name: "", Address: fromAddr}
	to := mail.Address{Name: "", Address: toAddr}
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", fromAddr, password, host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	client, err := smtp.NewClient(conn, host)
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

	c.ResCh <- "signup fin"
}

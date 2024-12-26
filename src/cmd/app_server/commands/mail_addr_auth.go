package commands

import (
	"database/sql"
	"fmt"
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

	rows, err := db.Query("select secret_key from signup_access_token where access_token = $1", tokenString)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	log.Println("query done")

	var secretKey string = ""

	for rows.Next() {
		err := rows.Scan(&secretKey)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("get secretKey from DB:", secretKey)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

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

	c.ResCh <- "mail addr auth fin"
}

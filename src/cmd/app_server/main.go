package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	var (
		// ip   = flag.String("ip", "localhost", "IP address of the app server")
		port   = flag.Int("port", 10081, "port number of the app server")
		dbPort = flag.Int("db_port", 5432, "port number of the db server")
	)
	const (
		apiPathPrefix string = "/lite-sns/api/v1"
	)

	db, err := sql.Open("postgres", fmt.Sprintf("host=lite-sns-db port=%v user=user password=postgres dbname=lite_sns_db sslmode=disable", *dbPort))
	// db, err := sql.Open("postgres", fmt.Sprintf("postgres://user:postgres@lite-sns-db:%v/lite_sns_db", *dbPort))
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	log.Println("DB connected")

	// practice select
	rows, err := db.Query("select id, name from users where id = $1", 2)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	log.Println("query done")

	for rows.Next() {
		var (
			id   int
			name string
		)
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("get data:", id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	r.POST(fmt.Sprintf("%s/signup_request", apiPathPrefix), func(c *gin.Context) {})
	r.POST(fmt.Sprintf("%s/signup", apiPathPrefix), func(c *gin.Context) {})
	r.GET(fmt.Sprintf("%s/mail_addr_auth", apiPathPrefix), func(c *gin.Context) {})
	r.POST(fmt.Sprintf("%s/signin", apiPathPrefix), func(c *gin.Context) {})

	r.Run(fmt.Sprintf(":%d", *port))
	// log.Println("app server started")
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	ping_res, err := http.Post(
	// 		"http://localhost:18081/token",
	// 		"application/json",
	// 		bytes.NewBuffer([]byte("{}")),
	// 	)
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "failed to get response from the auth server " + err.Error(),
	// 		})
	// 		return
	// 	}

	// 	body, err := io.ReadAll(ping_res.Body)
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "failed to get response boy from PING response " + err.Error(),
	// 		})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": string(body),
	// 	})
	// })
	// r.Run(":18080")
}

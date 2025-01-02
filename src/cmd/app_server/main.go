package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"lite-sns/m/src/cmd/app_server/api_server"
	"lite-sns/m/src/cmd/app_server/interfaces"
	"lite-sns/m/src/cmd/app_server/server_configs"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// read a config file
	configFile, err := os.Open("conf/release/app_server.json")
	if err != nil {
		log.Fatalln(err)
	}
	configBytes, err := io.ReadAll(configFile)
	if err != nil {
		log.Fatalln(err)
	}
	var serverConfigs server_configs.ServerConfigs
	json.Unmarshal(configBytes, &serverConfigs)

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%v user=%s password=%s dbname=%s sslmode=disable",
			serverConfigs.Db.Hostname,
			serverConfigs.Db.Port,
			serverConfigs.Db.Username,
			serverConfigs.Db.Password,
			serverConfigs.Db.Dbname,
		),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	log.Println("DB connected")

	// // practice select
	// rows, err := db.Query("select id, name from users where id = $1", 2)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer rows.Close()

	// log.Println("query done")

	// for rows.Next() {
	// 	var (
	// 		id   int
	// 		name string
	// 	)
	// 	err := rows.Scan(&id, &name)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	log.Println("get data:", id, name)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// // practice insert
	// stmt, err := db.Prepare("insert into users(name) values($1)")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// res, err := stmt.Exec("doraemon")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// // lastId, err := res.LastInsertId()
	// // if err != nil {
	// // 	log.Fatalln(err)
	// // }
	// rowCnt, err := res.RowsAffected()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Printf("ID = <not supported>, affected = %d\n", rowCnt)

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

	apiServerCommandCh := make(chan interfaces.ApiServerCommandInterface)
	apiServer := api_server.NewApiServer(
		serverConfigs.App.ApiPrefix,
		serverConfigs.App.Port,
		serverConfigs.Frontend.Ip,
		serverConfigs.Frontend.Port,
		apiServerCommandCh,
	)

	// APIリクエスト受付開始
	go apiServer.Run()

	for {
		select {
		case cmd := <-apiServerCommandCh: // APIリクエストはこのチャネルで受け取り、シングルスレッドで処理する。
			cmd.Exec(&serverConfigs, db)
		}
	}
}

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

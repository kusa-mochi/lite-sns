package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func main() {
	log.Println("auth server started")

	configFile, err := os.Open("auth_server_config.json")
	if err != nil {
		log.Fatal("failed to open the config file")
	}
	defer configFile.Close()

	configContents, err := io.ReadAll(configFile)
	if err != nil {
		log.Fatal("failed to read the config file")
	}

	var config AuthServerOptions
	json.Unmarshal(configContents, &config)

	server := NewAuthServer(&config)
	server.Run()
}

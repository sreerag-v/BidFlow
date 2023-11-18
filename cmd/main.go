package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/sreerag_v/BidFlow/pkg/config"
	"github.com/sreerag_v/BidFlow/pkg/di"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading the env file")
	}

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}

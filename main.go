package main

import (
	"log"

	"github.com/gogoalish/doodocs-test/internal/api"
	"github.com/gogoalish/doodocs-test/utils"
)

func main() {
	config, err := utils.LoadConfig("config/config.json")
	if err != nil {
		log.Fatal(err)
	}
	err = utils.ParseEnv("config/credentials.env")
	if err != nil {
		log.Fatal(err)
	}
	server := api.NewServer(config)
	log.Fatal(server.Start())
}

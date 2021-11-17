package main

import (
	"github.com/denys-semenkov/go-clean-architecture/config"
	"github.com/denys-semenkov/go-clean-architecture/server"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.NewApp()

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}

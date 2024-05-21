package main

import (
	"golang-hexagon-example/internal/app"
	"golang-hexagon-example/internal/app/infrastructure/config"
	"log"
)

func main() {
	cfg, err := config.NewAppConfig("./configs/config.json")
	if err != nil {
		log.Fatalln("Unable to read configuration file", err)
	}

	application := app.NewApp(cfg)
	application.Run()
}

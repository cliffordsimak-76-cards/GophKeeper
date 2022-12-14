package main

import (
	"log"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("error running server: ", err)
	}

	if err = app.Run(cfg); err != nil {
		log.Fatal("error running server", err)
	}
}

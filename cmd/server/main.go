package main

import (
	"context"
	"log"

	_ "github.com/jackc/pgx/stdlib"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("error running server: ", err)
	}

	if err = app.Run(context.Background(), cfg); err != nil {
		log.Fatal("error running server: ", err)
	}
}

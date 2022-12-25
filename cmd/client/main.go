package main

import (
	"context"
	"log"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/client"
)

func main() {
	cfg, err := client.NewConfig()
	if err != nil {
		log.Fatal("error running client: ", err)
	}

	if err = client.Run(context.Background(), cfg); err != nil {
		log.Fatal("error running client: ", err)
	}
}

package main

import (
	"context"
	"log"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/client"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
)

func main() {
	printBuildData()

	cfg, err := client.NewConfig()
	if err != nil {
		log.Fatal("error running client: ", err)
	}

	if err = client.Run(context.Background(), cfg); err != nil {
		log.Fatal("error running client: ", err)
	}
}

func printBuildData() {
	log.Printf("Build version: %v\n", buildVersion)
	log.Printf("Build date: %v\n", buildDate)
}

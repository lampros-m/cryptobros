package main

import (
	"log"
	"os"
	"repositories/cryptobros/internal/config"
	"repositories/cryptobros/internal/services"
)

var (
	DebugMode = true
)

func main() {
	config := config.NewConfig(DebugMode)
	coinsCreator := services.NewCoinsCreator(config)

	err := coinsCreator.CreateTodayCoins()
	if err != nil {
		log.Println("error create today coins: ", err)
		os.Exit(1)
	}

	log.Println("Coins created successfully")
}

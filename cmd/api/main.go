package main

import (
	"fmt"
	"log"

	"github.com/levi9111/goshort/internal/config"
	"github.com/levi9111/goshort/internal/database"
)

func main() {
	// Initialize our config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to MogoDB
	mongoClient, err := database.ConnectMongo(cfg.MongoURI)
	if err !=nil {
		log.Fatalf("Failed to connect to MongoDB: %v",err)
		log.Println(mongoClient)
	}

	// Let's print out a value to prove it loaded from .env
	fmt.Printf("Config loaded successfully! Mongo URI is: %s\n", cfg.Port)
}
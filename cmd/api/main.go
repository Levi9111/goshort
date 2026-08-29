package main

import (
	"fmt"
	"log"

	// Replace 'github.com/yourusername/goshort' with whatever you put in go.mod
	"github.com/levi9111/goshort/internal/config"
)

func main() {
	// Initialize our config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Let's print out a value to prove it loaded from .env
	fmt.Printf("Config loaded successfully! Mongo URI is: %s\n", cfg.MongoURI)
}
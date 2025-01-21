package main

import (
	"go-backend-demo/config"
	"log"
)

func init() {
	log.Println("Init start")
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	log.Printf("config load success: %v\n", config)
	log.Println("Init end")
}

func main() {
	log.Printf("Main thread start...\n")
}

package main

import (
	"log"
	"stock-sentiment-cli/cmd"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	cmd.Execute()
}

package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadLocalEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

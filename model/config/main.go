package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	SrvHost    string
	SrvPort    string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

func Get() Config {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	return Config{
		SrvHost:    os.Getenv("ADDRESS"),
		SrvPort:    os.Getenv("PORT"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
	}
}

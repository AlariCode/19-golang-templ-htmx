package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file")
		return
	}
	log.Println(".env file loaded")
}

type DatabaseConfig struct {
	url string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		url: os.Getenv("DATABASE_URL"),
	}
}

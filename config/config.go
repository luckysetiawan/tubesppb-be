package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig(key string) string {
	if os.Getenv("DEPLOY_STATUS") == "PRODUCTION" {
		return os.Getenv(key)
	}

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Cannot load given key: " + key)
	}

	return os.Getenv(key)
}

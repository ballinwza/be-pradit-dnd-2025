package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadingEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env")
	}

	return
}

func GetMongoURI() string {
	return os.Getenv("MONGO_URI")
}
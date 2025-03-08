package utils

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		return
	}
	log.Println("Environment variables loaded successfully")
}

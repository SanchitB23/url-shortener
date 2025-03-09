package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnvs() error {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" || goEnv == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file")
			return err
		}
		fmt.Println("Environment variables loaded successfully")
	}
	return nil
}

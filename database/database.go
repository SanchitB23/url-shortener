package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
	"url-shortener/logger"
)

var DB *sql.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		logger.Log.Panic("DATABASE_URL environment variable not set")
	}
	DB, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Log.Panic("Failed to connect to the database", err)
	}
	err = DB.Ping()
	if err != nil {
		logger.Log.Panic("Failed to ping the database", err)
	}
	logger.Log.Info("Successfully connected to the database")
}

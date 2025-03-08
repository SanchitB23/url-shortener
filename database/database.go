package database

import (
	"database/sql"
	"fmt"
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
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		logger.Log.Panic("Failed to connect to the database", err)
	}
	err = DB.Ping()
	if err != nil {
		logger.Log.Panic("Failed to ping the database", err)
	}
	logger.Log.Info("Successfully connected to the database")
	createTables()
}

func createTables() {
	tableName := os.Getenv("URL_TABLE_NAME")
	createURLsTable := fmt.Sprintf(`
				CREATE TABLE IF NOT EXISTS %s (
					id SERIAL PRIMARY KEY,
					long_url TEXT NOT NULL,
					short_url TEXT NOT NULL,
					user_id INT NOT NULL
				)
			`, tableName)
	_, err := DB.Exec(createURLsTable)
	if err != nil {
		logger.Log.Panic("Failed to create urls table", err)
	} else {
		logger.Log.Info("Successfully created urls table")
	}
}

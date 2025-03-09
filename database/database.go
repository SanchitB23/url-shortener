package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"url-shortener/config"
)

var DB *sql.DB
var urlTableName string
var userTableName string

func InitPostgresDB() error {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		config.Log.Panic("DATABASE_URL environment variable not set")
		return errors.New("DATABASE_URL environment variable not set")
	}
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		config.Log.Panic("Failed to connect to the database", err)
		return err
	}
	err = DB.Ping()
	if err != nil {
		config.Log.Panic("Failed to ping the database", err)
		return err
	}
	config.Log.Info("Successfully connected to the database")
	urlTableName = os.Getenv("URL_TABLE_NAME")
	userTableName = os.Getenv("USER_TABLE_NAME")

	err = createTables()
	if err != nil {
		return err
	}
	return nil
}

func createTables() error {
	createURLsTable := fmt.Sprintf(`
				CREATE TABLE IF NOT EXISTS %s (
					id SERIAL PRIMARY KEY,
					long_url TEXT NOT NULL,
					short_url TEXT NOT NULL,
					user_id INT NOT NULL
				)
			`, GetUrlTableName())
	_, err := DB.Exec(createURLsTable)
	if err != nil {
		config.Log.Panic("Failed to create urls table", err)
		return err
	}
	config.Log.Info("Successfully created urls table")
	return nil
}

func GetUrlTableName() string {
	return urlTableName
}
func GetUserTableName() string {
	return userTableName
}

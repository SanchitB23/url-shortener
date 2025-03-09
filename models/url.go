package models

import (
	"database/sql"
	"errors"
	"fmt"
	"url-shortener/config"
	"url-shortener/database"
	"url-shortener/utils"
)

type URL struct {
	ID          int64  `json:"id"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url" binding:"required"`
	CreatedAt   string `json:"created_at"`
	UserID      int64  `json:"user_id"`
}

func checkIfOriginalURLExists(originalUrl string, userId int64) (bool, error) {
	query := fmt.Sprintf(`SELECT id FROM %s WHERE original_url = $1 AND user_id = $2`, database.GetUrlTableName())
	row := database.DB.QueryRow(query, originalUrl, userId)
	var id int64
	err := row.Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			config.Log.Debug("Long URL does not exist", err)
			return false, nil
		}
		config.Log.Debugf("Error in checking if long URL exists: %v", err)
		return false, err
	}
	return true, nil
}

func (u *URL) Shorten() error {
	// Logic to save the URL goes here
	exists, err := checkIfOriginalURLExists(u.OriginalURL, u.UserID)
	if err != nil {
		config.Log.Debugf("Error in checking if long URL exists: %v", err)
		return err
	}
	if exists {
		config.Log.Debug("URL already exists", u.OriginalURL, u.UserID)
		return errors.New("URL already exists for the user")
	}

	u.ShortURL = utils.CreateShortUUID()

	query := fmt.Sprintf(`INSERT INTO %s (original_url, short_url, user_id) VALUES ($1, $2, $3)`, database.GetUrlTableName())
	smt, err := database.DB.Prepare(query)
	defer func(smt *sql.Stmt) {
		err := smt.Close()
		if err != nil {
			config.Log.Debug("Error closing statement", err)
		}
	}(smt)
	if err != nil {
		config.Log.Debug("Error preparing insert statement", err)
		return err
	}
	_, err = smt.Exec(u.OriginalURL, u.ShortURL, u.UserID)
	if err != nil {
		config.Log.Debug("Error inserting URL", err)
		return err
	}
	config.Log.Debug("Shortened URL saved in Db", u.ShortURL)
	return nil
}

func (u *URL) GetOriginalURL() error {
	query := fmt.Sprintf(`SELECT original_url FROM %s WHERE short_url = $1 AND user_id = $2`, database.GetUrlTableName())
	row := database.DB.QueryRow(query, u.ShortURL, u.UserID)
	return row.Scan(&u.OriginalURL)
}

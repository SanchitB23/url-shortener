package models

import (
	"database/sql"
	"errors"
	"fmt"
	"url-shortener/config"
	"url-shortener/database"
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

func (u *URL) Shorten(userId int64) (string, error) {
	// Logic to save the URL goes here
	exists, err := checkIfOriginalURLExists(u.OriginalURL, userId)
	if err != nil {
		config.Log.Debugf("Error in checking if long URL exists: %v", err)
		return "", err
	}
	if exists {
		config.Log.Debug("URL already exists", u.ShortURL, userId)
		return "", errors.New("URL already exists for the user")
	}

	//todo
	return "", nil
}

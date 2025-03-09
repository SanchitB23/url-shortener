package models

type URL struct {
	ID          int64  `json:"id"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url" binding:"required"`
	DateCreated string `json:"date_created"`
	UserID      int64  `json:"user_id"`
}

func (u *URL) Shorten() (string, error) {
	// Logic to save the URL goes here
	return "", nil
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"url-shortener/config"
	_ "url-shortener/docs" // This line is necessary for go-swagger to find your docs
	"url-shortener/middlewares"
	"url-shortener/models"
)

// ShortenURL @Summary Shorten URL
// @Description Shortens a given URL
// @Tags url
// @Accept json
// @Produce json
// @Param url body string true "URL to shorten" example({original_url: "https://example.com"})
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /shorten [post]
func ShortenURL(context *gin.Context) {

	userID := middlewares.GetCurrentUser(context)
	var urlObj models.URL

	if err := context.ShouldBindJSON(&urlObj); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		config.Log.Debug("Error binding JSON: ", err, context.JSON)
		return
	}
	config.Log.Debugf("Host URL: %v", context.Request.Host)
	shortenedURL, err := urlObj.Shorten(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to shorten URL"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"shortenedURL": shortenedURL})
}

// RedirectToOriginalURL @Summary Redirect to original URL
// @Description Redirects to the original URL based on the shortened URL
// @Tags url
// @Param shortURL path string true "Shortened URL"
// @Success 301
// @Router /{shortURL} [get]
func RedirectToOriginalURL(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, gin.H{"message": "Not Implemented"})
}

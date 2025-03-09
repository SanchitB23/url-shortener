package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"url-shortener/config"
	_ "url-shortener/docs" // This line is necessary for go-swagger to find your docs
	"url-shortener/middlewares"
	"url-shortener/models"
	"url-shortener/utils"
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

	var urlObj models.URL
	urlObj.UserID = middlewares.GetCurrentUser(context)

	if err := context.ShouldBindJSON(&urlObj); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		config.Log.Debug("Error binding JSON: ", err, context.JSON)
		return
	}

	if !utils.IsValidURL(urlObj.OriginalURL) {
		config.Log.Debug("Invalid URL: ", urlObj.OriginalURL)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL", "message": "Please provide a valid URL"})
		return
	}

	config.Log.Debugf("Host URL: %v", context.Request.Host)
	err := urlObj.Shorten()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to shorten URL"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"shortenedURL": utils.GetHttpProtocol(context) + `://` + context.Request.Host + `/` + urlObj.ShortURL})
}

// RedirectToOriginalURL @Summary Redirect to original URL
// @Description Redirects to the original URL based on the shortened URL
// @Tags url
// @Param shortURL path string true "Shortened URL"
// @Success 301
// @Router /{shortURL} [get]
func RedirectToOriginalURL(context *gin.Context) {
	var urlObj models.URL
	urlObj.ShortURL = context.Request.URL.Path[1:]
	urlObj.UserID = middlewares.GetCurrentUser(context)

	err := urlObj.GetOriginalURL()
	if err != nil {
		config.Log.Debug("Error getting original URL: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to get original URL"})
		return
	}
	context.Redirect(http.StatusMovedPermanently, urlObj.OriginalURL)
}
